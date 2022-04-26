package routes

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestSwaggerRoute(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		args args
	}{
		{"swagger route test", args{app: fiber.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SwaggerRoute(tt.args.app)
		})
	}
}
