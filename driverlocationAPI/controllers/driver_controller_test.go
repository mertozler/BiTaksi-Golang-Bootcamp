package controllers

import (
	"bytes"
	"driverlocationapi-bitaksi/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetAllDriverLocation(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int
	}{

		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "http://localhost:7000/not-found",
			expectedCode: 404,
		},
		{
			description:  "getting all drivers location endpoint test",
			route:        "http://localhost:7000/getAllDriversLocation",
			expectedCode: 200,
		},
	}
	app := fiber.New()
	app.Get("/getAllDriversLocation", GetAllDriverLocation)
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.route, nil)
			resp, _ := app.Test(req, -1)
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})

	}
}

func TestCreateDriverLocation(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int
	}{
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "http://localhost:7000/not-found",
			expectedCode: 404,
		},
		{
			description:  "create a driver and bulk update endpoint test",
			route:        "http://localhost:7000/createDriver",
			expectedCode: 201,
		},
		{
			description:  "write expection test",
			route:        "http://localhost:7000/createDriver",
			expectedCode: 500,
		},
		{
			description:  "invalid body request test",
			route:        "http://localhost:7000/createDriver",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/createDriver", CreateDriverLocation)
	for _, test := range tests {
		if test.description == "write expection test" {
			t.Run(test.description, func(t *testing.T) {
				location := models.Location{
					Type:        "Hello",
					Coordinates: []float64{22.690080, 46.277659},
				}
				driver := models.RiderCreateAndUpdateRequest{
					Location: location,
				}
				var driverLocations []models.RiderCreateAndUpdateRequest
				driverLocations = append(driverLocations, driver)
				riderRequestDataConvertedJson, err := json.Marshal(driverLocations)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.description == "invalid body request test" {
			t.Run(test.description, func(t *testing.T) {
				testInvalidData := []float64{32, 25, 56, 22}
				riderRequestDataConvertedJson, err := json.Marshal(testInvalidData)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else {
			t.Run(test.description, func(t *testing.T) {
				location := models.Location{
					Type:        "Point",
					Coordinates: []float64{22.690080, 46.277659},
				}
				driver := models.RiderCreateAndUpdateRequest{
					Location: location,
				}
				var driverLocations []models.RiderCreateAndUpdateRequest
				driverLocations = append(driverLocations, driver)
				riderRequestDataConvertedJson, err := json.Marshal(driverLocations)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		}

	}
}

func TestFindDriverForRider(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int
	}{
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "http://localhost:7000/not-found",
			expectedCode: 404,
		},
		{
			description:  "find driver endpoint test",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 200,
		},
		{
			description:  "find driver bad request check",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "find driver validation check",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "find driver authorization check with false value",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 401,
		},
		{
			description:  "invalid body request test",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 400,
		},
		{
			description:  "404 not found test",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 404,
		},
		{
			description:  "validate error test",
			route:        "http://localhost:7000/findDriver",
			expectedCode: 400,
		},
	}
	app := fiber.New()
	app.Post("/findDriver", FindDriverForRider)
	riderGoodRequest := models.RiderRequest{
		Type:        "Point",
		Coordinates: []float64{40.881087, 29.1146},
		Radius:      3,
	}
	riderBadRequest := models.RiderRequest{
		Type:        "Point",
		Coordinates: []float64{400.881087, 29.1146},
		Radius:      3,
	}
	riderValidateCheckRequest := models.RiderRequest{
		Type:        "Point",
		Coordinates: []float64{40.881087, 29.1146},
	}
	riderRequestDataConvertedJson, err := json.Marshal(riderGoodRequest)
	riderBadRequestConvertedToJson, err := json.Marshal(riderBadRequest)
	riderValidateCheckConvertedToJson, err := json.Marshal(riderValidateCheckRequest)
	if err != nil {
		assert.Error(t, err)
	}
	for _, test := range tests {
		if test.description == "find driver authorization check with false value" {
			t.Run(test.description, func(t *testing.T) {
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemF0aW9uIjpmYWxzZX0.X2VwpiIBD46bYO1m2MqR9DiI5sHCNdHD4HUktcE_x5I")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.description == "validate error test" {
			t.Run(test.description, func(t *testing.T) {
				riderValidateTestData := models.RiderRequest{
					Coordinates: []float64{40.881087, 29.1146},
					Radius:      1,
				}
				riderValidateTestDataConvertedJson, err := json.Marshal(riderValidateTestData)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderValidateTestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.description == "404 not found test" {
			t.Run(test.description, func(t *testing.T) {
				rider404NotFoundTestData := models.RiderRequest{
					Type:        "Point",
					Coordinates: []float64{77.881087, 15.1146},
					Radius:      1,
				}
				rider404NotFoundTestDataConvertedJson, err := json.Marshal(rider404NotFoundTestData)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(rider404NotFoundTestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.description == "invalid body request test" {
			t.Run(test.description, func(t *testing.T) {
				invalidBodyRequestTest := []float64{32, 25, 56, 22}
				invalidBodyRequestJson, err := json.Marshal(invalidBodyRequestTest)
				if err != nil {
					assert.Error(t, err)
				}
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(invalidBodyRequestJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.expectedCode == 200 {
			t.Run(test.description, func(t *testing.T) {
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.expectedCode == 400 {
			t.Run(test.description, func(t *testing.T) {
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderBadRequestConvertedToJson))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.expectedCode == 404 {
			t.Run(test.description, func(t *testing.T) {
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderRequestDataConvertedJson))
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				req.Header.Set("Content-Type", "application/json")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})
		} else if test.description == "find driver validation check" {
			t.Run(test.description, func(t *testing.T) {
				req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(riderValidateCheckConvertedToJson))
				req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
				req.Header.Set("Content-Type", "application/json")
				resp, _ := app.Test(req, -1)
				assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
			})

		}
	}
}
