package routes

import (
	"car-rent/internal/handlers"
	"car-rent/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrdersRouter(app *fiber.App, handler handlers.Handlers, middleware *middleware.Authentication) {
	app.Post("/order", handler.Order.Create)
	app.Put("/order/:order_id", handler.Order.Update)
	app.Get("/order", handler.Order.List)
	app.Get("/order/:order_id", handler.Order.Detail)
	app.Delete("/order/:order_id", handler.Order.Delete)
	app.Put("/order/activate/:order_id", handler.Order.Activate)
	app.Put("/order/deactivate/:order_id", handler.Order.Deactivate)
}
