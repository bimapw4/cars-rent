package routes

import (
	"car-rent/internal/handlers"
	"car-rent/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func CarsRouter(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/cars", handler.Cars.Create)
	app.Put("/cars/:cars_id", handler.Cars.Update)
	app.Get("/cars", handler.Cars.List)
	app.Get("/cars/:cars_id", handler.Cars.Detail)
	app.Delete("/cars/:cars_id", handler.Cars.Delete)
	app.Put("/cars/activate/:cars_id", handler.Cars.Activate)
	app.Put("/cars/deactivate/:cars_id", handler.Cars.Deactivate)
}
