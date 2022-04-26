package routes

import (
	_ "driverlocationapi-bitaksi/docs"
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
}
