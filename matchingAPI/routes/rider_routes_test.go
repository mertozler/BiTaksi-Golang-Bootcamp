package routes

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestRiderRoute(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		args args
	}{
		{"test rider route", args{app: fiber.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RiderRoute(tt.args.app)
		})
	}
}
