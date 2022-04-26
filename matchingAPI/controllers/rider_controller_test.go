package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"matchingAPI-bitaksi/models"
	"net/http/httptest"
	"testing"
)

func TestFindDriver(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int
	}{
		{
			description:  "Find driver",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 200,
		},
		{
			description:  "404 not found",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 404,
		},
		{
			description:  "Invalid request test",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "Null auth token test",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "Unauthorized User test",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 401,
		},
		{
			description:  "Sending wrong auth token test",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "Body parser test",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "Bad value test",
			route:        "http://localhost:8000/findDriver",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/findDriver", FindDriver)
	for _, tt := range tests {
		if tt.expectedCode == 200 {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						41.51087, 29.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		} else if tt.expectedCode == 404 {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						71.51087, 69.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		} else if tt.description == "Unauthorized User test" {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						41.51087, 29.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjpmYWxzZX0.GeyNcvlidGJOB2rmmwtxTNoXmVb-G-jQuN4Z_sJ2j8E")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		} else if tt.description == "Body parser test" {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						41.51087, 29.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})

		} else if tt.description == "Bad value test" {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						410.51087, 29.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		} else if tt.description == "Null auth token test" {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						71.51087, 69.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		} else if tt.description == "Sending wrong auth token test" {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Radius: 3,
					Type:   "Point",
					Coordinates: []float64{
						41.51087, 29.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "BiTaksi")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		} else if tt.expectedCode == 400 {
			t.Run(tt.description, func(t *testing.T) {
				rider := models.RiderRequest{
					Type: "Point",
					Coordinates: []float64{
						71.51087, 69.55146,
					},
				}
				riderRequestDataConvertedJson, err := json.Marshal(rider)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
			})
		}

	}

}
