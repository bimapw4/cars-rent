package routes

import (
	"car-rent/internal/handlers"
	"car-rent/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func CarsRouter(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/cars", middleware.Authentication, handler.Cars.Create)
	app.Put("/cars/:cars_id", middleware.Authentication, handler.Cars.Update)
	app.Get("/cars", middleware.Authentication, handler.Cars.List)
	app.Get("/cars/:cars_id", middleware.Authentication, handler.Cars.Detail)
	app.Delete("/cars/:cars_id", middleware.Authentication, handler.Cars.Delete)
	app.Put("/cars/activate/:cars_id", middleware.Authentication, handler.Cars.Activate)
	app.Put("/cars/deactivate/:cars_id", middleware.Authentication, handler.Cars.Deactivate)
	app.Get("/cars/schedulle/available", middleware.Authentication, handler.Cars.AvailableCars)
	app.Get("/cars/preview/:cars_id", middleware.Authentication, handler.Cars.Preview)
}
