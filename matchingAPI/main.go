package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"matchingAPI-bitaksi/routes"
)

// @title           Matching API
// @version         1.0
// @description     Find the nearest driver to rider.

// @contact.name   Mert Özler
// @contact.url    http://github.com/mertozler
// @contact.email  meozler@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @in                          header
// @name                        Authorization
func main() {
	app := fiber.New()
	app.Use(logger.New())
	routes.RiderRoute(app)
	routes.SwaggerRoute(app)
	app.Listen(":8000")
}
