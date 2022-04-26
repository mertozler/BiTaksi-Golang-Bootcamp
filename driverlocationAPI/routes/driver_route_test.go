package routes

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestDriverRoute(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		args args
	}{
		{"Driver Route Test", args{app: fiber.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DriverRoute(tt.args.app)
		})
	}
}
