package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Type        string    `json:"type,omitempty" validate:"required"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

type DriverLocation struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Location Location           `json:"location,omitempty" validate:"required"`
}

type RiderRequest struct {
	Id          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty" validate:"required"`
	Coordinates []float64 `json:"coordinates,omitempty"`
	Radius      float64   `json:"radius,omitempty"`
}

type RiderCreateAndUpdateRequest struct {
	Id       string   `json:"id,omitempty"`
	Location Location `json:"location,omitempty" validate:"required"`
}

type DistanceofDriversFromRider struct {
	DriverID primitive.ObjectID `json:"id,omitempty"`
	Distance float64            `json:"distance(km),omitempty"`
}
