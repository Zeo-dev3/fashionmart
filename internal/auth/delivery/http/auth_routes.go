package http

import (
	"github.com/Zeo-dev3/fashionmart/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(app *fiber.App, h auth.Handler, authMW fiber.Handler, adminMW fiber.Handler) {
	group := app.Group("/api/auth")

	group.Post("/register", h.Register())
	group.Post("/login", h.Login())
	group.Get("/check", authMW, h.Check())
	group.Get("/check/admin", authMW, adminMW, h.CheckAdmin())
	group.Post("/update-role", authMW, h.UpdateRole())
}
