package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"matchingAPI-bitaksi/models"
	"matchingAPI-bitaksi/responses"
	"net/http"
	"strings"
)

var validate = validator.New()

// FindDriver godoc
// @Summary      Find The Driver
// @Description  The endpoint that allows searching with a GeoJSON point to find a driver if it matches the given criteria. Otherwise, the service should respond with a 404 - Not Found
// @Tags         Matching
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Authentication header: (You need to add Bearer to the beginning of the jwt key. For example: Bearer apiKey)"
// @Param        rider  body      models.RiderRequest  true  "Rider data"
// @Success      200  {object}   responses.RiderResponse
// @Failure      400  {object}   responses.RiderResponse
// @Failure      404  {object}   responses.RiderResponse
// @Failure      401  {object}   responses.RiderResponse
// @Router       /findDriver [post]
func FindDriver(c *fiber.Ctx) error {
	var rider models.RiderRequest
	checkAuthStatus := false
	if err := c.BodyParser(&rider); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	requestToken := c.GetReqHeaders()["Authorization"]
	if requestToken == "" || requestToken == "Bearer" {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "jwt token not found so authorization failed"}})
	}
	// Check the probability of user sending wrong token
	if len(strings.Fields(requestToken)) == 1 {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "jwt token not found so authorization failed"}})
	}
	requestTokenClean := strings.Fields(requestToken)[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(requestTokenClean, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<true>"), nil
	})
	for key, val := range claims {
		if key == "authenticated" && val == true {
			checkAuthStatus = true //this means user is authenticated
		} else {
			checkAuthStatus = false //this means user is unauthorized
		}
	}
	if checkAuthStatus == false {
		return c.Status(http.StatusUnauthorized).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Unauthorized user"}})
	}
	if validationErr := validate.Struct(&rider); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	riderRequestData := models.RiderRequest{
		Id:          primitive.NewObjectID(),
		Type:        rider.Type,
		Coordinates: rider.Coordinates,
		Radius:      rider.Radius,
	}
	riderRequestDataConvertedJson, err := json.Marshal(riderRequestData)
	req, err := http.NewRequest("POST", "http://localhost:7000/findDriver", bytes.NewBuffer(riderRequestDataConvertedJson))
	req.Header.Set("Apikey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs")
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}
	body, err := ioutil.ReadAll(resp.Body)
	var result responses.RiderResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.RiderResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}
	if result.Status == http.StatusNotFound {
		return c.Status(http.StatusNotFound).JSON(
			responses.RiderResponse{Status: http.StatusNotFound, Message: "404 - Not Found", Data: &fiber.Map{"Could not match any rider according to rider information": result.Status}},
		)
	}

	if result.Status == http.StatusBadRequest {
		return c.Status(http.StatusBadRequest).JSON(
			responses.RiderResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: &fiber.Map{"Could not match any rider according to rider information": result.Data}},
		)
	}
	return c.Status(http.StatusOK).JSON(responses.RiderResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Matching Data": result.Data}})
}
