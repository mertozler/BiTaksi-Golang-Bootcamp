package routes

import (
	"driverlocationapi-bitaksi/controllers"

	"github.com/gofiber/fiber/v2"
)

func DriverRoute(app *fiber.App) {
	app.Post("/createDriver", controllers.CreateDriverLocation)
	app.Get("/getAllDriversLocation", controllers.GetAllDriverLocation)
	app.Post("/findDriver", controllers.FindDriverForRider)
}
