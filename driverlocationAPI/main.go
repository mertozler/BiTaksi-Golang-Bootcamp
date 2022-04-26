package main

import (
	"driverlocationapi-bitaksi/configs"
	"driverlocationapi-bitaksi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

// @title           Driver Location API
// @version         1.0
// @description     The API provides finding the nearest driver location with a given GeoJSON point and radius.

// @contact.name   Mert Özler
// @contact.url    http://github.com/mertozler
// @contact.email  meozler@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:7000
// @BasePath  /

func main() {
	app := fiber.New()
	app.Use(logger.New())
	configs.ConnectDB()
	err := configs.CreateDataIfisNotExist()
	if err != nil {
		log.Fatalln("Error while importing initial data")
	}
	routes.DriverRoute(app)
	routes.SwaggerRoute(app)
	app.Listen(":7000")
}
