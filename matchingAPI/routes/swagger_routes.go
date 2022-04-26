package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "matchingAPI-bitaksi/docs"
)

func SwaggerRoute(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
}
