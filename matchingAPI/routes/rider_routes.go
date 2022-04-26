package routes

import (
	"github.com/gofiber/fiber/v2"
	"matchingAPI-bitaksi/controllers"
)

func RiderRoute(app *fiber.App) {
	app.Post("/findDriver", controllers.FindDriver)
}
