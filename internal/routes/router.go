package routes

import (
	"car-rent/internal/handlers"
	"car-rent/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	// register route
	routes := []func(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication){
		CarsRouter,
	}

	for _, route := range routes {
		route(app, handler, m)
	}
}
