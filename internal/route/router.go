package router

import (
	"notification-service/internal/module/notification/handler"
	"notification-service/internal/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App, handlerTicket *handler.NotificationHandler, m *middleware.Middleware) *fiber.App {

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	return app

}
