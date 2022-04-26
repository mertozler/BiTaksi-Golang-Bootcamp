package configs

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestCreateDataIfisNotExist(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Testing initial data importing", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDataIfisNotExist(); (err != nil) != tt.wantErr {
				t.Errorf("CreateDataIfisNotExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCollection(t *testing.T) {
	type args struct {
		client         *mongo.Client
		collectionName string
	}
	client := ConnectDB()
	tests := []struct {
		name string
		args args
		want *mongo.Collection
	}{
		{"driver location get collection test", args{client: DB, collectionName: "driverLocations"}, client.Database("driverlocation-api").Collection("driverLocations")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCollection(tt.args.client, tt.args.collectionName); !reflect.DeepEqual(got.Name(), tt.want.Name()) {
				t.Errorf("GetCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
