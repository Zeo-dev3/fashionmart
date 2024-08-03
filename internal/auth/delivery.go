package auth

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Register() fiber.Handler
	Login() fiber.Handler
	Check() fiber.Handler
	UpdateRole() fiber.Handler
	CheckAdmin() fiber.Handler
}
