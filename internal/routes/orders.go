package routes

import (
	"car-rent/internal/handlers"
	"car-rent/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrdersRouter(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/order", middleware.Authentication, handler.Order.Create)
	app.Put("/order/:order_id", middleware.Authentication, handler.Order.Update)
	app.Get("/order", middleware.Authentication, handler.Order.List)
	app.Get("/order/:order_id", middleware.Authentication, handler.Order.Detail)
	app.Delete("/order/:order_id", middleware.Authentication, handler.Order.Delete)
	app.Put("/order/activate/:order_id", middleware.Authentication, handler.Order.Activate)
	app.Put("/order/deactivate/:order_id", middleware.Authentication, handler.Order.Deactivate)
	app.Get("/order/summary/trx", middleware.Authentication, handler.Order.Summary)
}
