package configs

import (
	"context"
	"driverlocationapi-bitaksi/models"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"path"
	"runtime"
	"strconv"

	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("driverlocation-api").Collection(collectionName)
	return collection
}

func CreateDataIfisNotExist() error {
	ctx, cancel := context.WithCancel(context.Background())
	var driverCollection *mongo.Collection = GetCollection(DB, "driverLocations")
	//var driverLocations []models.DriverLocation
	var checkDatabaseisNull models.DriverLocation
	defer cancel()
	err := driverCollection.FindOne(ctx, bson.M{}).Decode(&checkDatabaseisNull)
	if err != nil {
		fmt.Println("Database is empty, initial data is creating right now")
		// get path dynamically
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "..")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}
		pwd := path.Dir(filename)
		f, err := os.Open(pwd + "\\coordinates.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		csvReader := csv.NewReader(f)
		if err != nil {
			log.Fatal(err)
		}

		driverLocationCSVData, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		var oneDriverLocationModel models.DriverLocation
		var oneDriverLocation models.Location
		var allDriverLocation []models.DriverLocation
		fmt.Println("Converting CSV data to GeoJSON objects, please wait.")
		for _, each := range driverLocationCSVData {
			var location []float64
			if latitude, err := strconv.ParseFloat(each[0], 64); err == nil {
				location = append(location, latitude)
			}
			if longitude, err := strconv.ParseFloat(each[1], 64); err == nil {
				location = append(location, longitude)
			}
			if location != nil {
				oneDriverLocation.Type = "Point"
				oneDriverLocation.Coordinates = location
				oneDriverLocationModel.Id = primitive.NewObjectID()
				oneDriverLocationModel.Location = oneDriverLocation
				allDriverLocation = append(allDriverLocation, oneDriverLocationModel)
			}
		}
		var driverLocationDatas []interface{}
		for _, t := range allDriverLocation {
			driverLocationDatas = append(driverLocationDatas, t)
		}
		fmt.Println("Data is importing into database, please wait..")
		result, err := driverCollection.InsertMany(ctx, driverLocationDatas)
		if err != nil {
			fmt.Println(err)
			fmt.Println(result)
		}
		Indexes := mongo.IndexModel{
			Keys: bson.D{
				{"location", "2dsphere"},
			},
		}
		fmt.Println("Creating Index for searching GEOJson data")
		//creating index for searching geojson datas
		_, err = driverCollection.Indexes().CreateOne(context.Background(), Indexes)
		if err != nil {
			log.Fatal("Error while creating index")
		}
		fmt.Println("Initial data is ready.")
	}

	return nil
}
