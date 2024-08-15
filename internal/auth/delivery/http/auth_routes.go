package http

import (
	"github.com/Zeo-dev3/fashionmart/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(route fiber.Router, h auth.Handler, authMW fiber.Handler, adminMW fiber.Handler) {
	route.Post("/auth/register", h.Register())
	route.Post("/auth/login", h.Login())
	route.Get("/auth/check", authMW, h.Check())
	route.Get("/auth/check/admin", authMW, adminMW, h.CheckAdmin())
	route.Post("/auth/update-role", authMW, h.UpdateRole())
}
