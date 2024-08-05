package product

import "github.com/gofiber/fiber/v2"

type Handler interface {
	AddProduct() fiber.Handler
	GetProductById() fiber.Handler
	AddProductColor() fiber.Handler
	GetAllProducts() fiber.Handler
}
