package controllers

import (
	"context"
	"driverlocationapi-bitaksi/configs"
	"driverlocationapi-bitaksi/models"
	"driverlocationapi-bitaksi/responses"
	"driverlocationapi-bitaksi/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"sort"
	"time"
)

var driverCollection *mongo.Collection = configs.GetCollection(configs.DB, "driverLocations")
var validate = validator.New()

// CreateDriverLocation godoc
// @Summary      Create a Driver Location
// @Description  An endpoint for creating a driver location. It would support batch operations to handle the bulk update.
// @Tags         Driver-Create
// @Accept       json
// @Produce      json
// @Param        driver  body     models.RiderCreateAndUpdateRequest  true  "Driver data"
// @Success      200  {object}   responses.DriverResponse
// @Failure      400  {object}   responses.DriverResponse
// @Failure      404  {object}   responses.DriverResponse
// @Failure      401  {object}   responses.DriverResponse
// @Router       /createDriver [post]
func CreateDriverLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var driver []models.RiderCreateAndUpdateRequest
	defer cancel()
	//validate the request body
	if err := c.BodyParser(&driver); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "Data is not valid, please try again", Data: &fiber.Map{"data": err.Error()}})
	}
	for _, oneDriver := range driver {
		objID, err := primitive.ObjectIDFromHex(oneDriver.Id)
		if err != nil {
			objID = primitive.NewObjectID()
		}
		newDriver := models.DriverLocation{
			Id:       objID,
			Location: oneDriver.Location,
		}
		// The update variable determines which value will be updated or inserted.
		update := bson.M{
			"$set": newDriver,
		}
		// If there is a driver on that id, update it, otherwise insert it. This is what Upsert means.
		opts := options.Update().SetUpsert(true)
		result, err := driverCollection.UpdateOne(ctx, bson.M{"_id": newDriver.Id}, update, opts)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
			fmt.Println(result)
		}

	}
	return c.Status(http.StatusCreated).JSON(responses.DriverResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": "your driver locations succesfully updated or created"}})
}

// GetAllDriverLocation godoc
// @Summary      Get All Driver Location
// @Description  Returns the current location of all drives.
// @Tags         Drivers-Location-Get-All
// @Accept       json
// @Produce      json
// @Success      200  {object}   responses.DriverResponse
// @Failure      400  {object}   responses.DriverResponse
// @Failure      404  {object}   responses.DriverResponse
// @Failure      401  {object}   responses.DriverResponse
// @Router       /getAllDriversLocation [get]
func GetAllDriverLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	var driverLocations []models.DriverLocation
	defer cancel()
	// Query for all driver locations
	results, err := driverCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var driverLocation models.DriverLocation
		if err = results.Decode(&driverLocation); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		driverLocations = append(driverLocations, driverLocation)
	}
	return c.Status(http.StatusOK).JSON(
		responses.DriverResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": driverLocations}},
	)
}

// FindDriverForRider godoc
// @Summary      Find The Driver
// @Description  The endpoint that allows searching with a GeoJSON point to find a driver if it matches the given criteria. Otherwise, the service should respond with a 404 - Not Found
// @Tags         Driver-Find
// @Accept       json
// @Produce      json
// @Param        Apikey  header    string  true  "Apikey header: (You need to add Bearer to the beginning of the jwt key. For example: Bearer apiKey)"
// @Param        rider  body      models.RiderRequest  true  "Rider data"
// @Success      200  {object}   responses.DriverResponse
// @Failure      400  {object}   responses.DriverResponse
// @Failure      404  {object}   responses.DriverResponse
// @Failure      401  {object}   responses.DriverResponse
// @Router       /findDriver [post]
func FindDriverForRider(c *fiber.Ctx) error {
	checkAuthStatus := false
	ctx, cancel := context.WithCancel(context.Background())
	var riderRequest models.RiderRequest
	var DistanceofDriversFromRider []models.DistanceofDriversFromRider
	// Getting JWT data from Matching API
	secretKeyByMatchingAPI := c.GetReqHeaders()["Apikey"]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(secretKeyByMatchingAPI, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Apikey"), nil
	})
	// Claims checking
	for key, val := range claims {
		if key == "authenticated" && val == true {
			checkAuthStatus = true //this means request is authenticated
		} else {
			checkAuthStatus = false //this means request is unauthorized
		}
	}
	if checkAuthStatus == false {
		return c.Status(http.StatusUnauthorized).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Unauthorized request"}})
	}
	defer cancel()
	if err := c.BodyParser(&riderRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "Data is not valid, please try again", Data: &fiber.Map{"data": err.Error()}})
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&riderRequest); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	riderRequestData := models.RiderRequest{
		Id:          riderRequest.Id,
		Type:        riderRequest.Type,
		Coordinates: riderRequest.Coordinates,
		Radius:      riderRequest.Radius,
	}
	// Query.
	results, err := driverCollection.Find(ctx, bson.M{"location": bson.M{
		"$nearSphere": bson.M{
			"$geometry": bson.M{
				"type":        "Point",
				"coordinates": []float64{riderRequestData.Coordinates[0], riderRequestData.Coordinates[1]},
			},
			"$maxDistance": riderRequestData.Radius * 1000,
		},
	}})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.DriverResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//Reading DB data and calculate distance.
	defer results.Close(ctx)
	for results.Next(ctx) {
		var driverLocation models.DriverLocation
		if err = results.Decode(&driverLocation); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DriverResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		distance := utils.Haversine(driverLocation.Location.Coordinates[0], driverLocation.Location.Coordinates[1], riderRequestData.Coordinates[0], riderRequestData.Coordinates[1])
		if distance <= riderRequestData.Radius {
			distanceData := models.DistanceofDriversFromRider{
				DriverID: driverLocation.Id,
				Distance: distance,
			}
			DistanceofDriversFromRider = append(DistanceofDriversFromRider, distanceData)
		}
	}
	//Sorting Struct field distance. The distance values in the struct arrays are compared and the struct with the minimum distance value can be found at index 0 of the struct array.
	sort.SliceStable(DistanceofDriversFromRider, func(i, j int) bool {
		return DistanceofDriversFromRider[i].Distance < DistanceofDriversFromRider[j].Distance
	})
	//Get best distance from rider location.
	if len(DistanceofDriversFromRider) == 0 {
		return c.Status(http.StatusNotFound).JSON(
			responses.DriverResponse{Status: http.StatusNotFound, Message: "404 - Not Found", Data: &fiber.Map{"Could not match any rider according to rider information": DistanceofDriversFromRider}},
		)
	}
	return c.Status(http.StatusOK).JSON(
		responses.DriverResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Nearest driver data": DistanceofDriversFromRider[0]}},
	)
}
