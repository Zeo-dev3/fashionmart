package http

import (
	"github.com/Zeo-dev3/fashionmart/internal/product"
	"github.com/gofiber/fiber/v2"
)

func MapProductRoutes(route fiber.Router, productHandler product.Handler, authMW fiber.Handler, adminMw fiber.Handler) {
	route.Get("/product", productHandler.GetAllProducts())
	route.Post("/product", authMW, adminMw, productHandler.AddProduct())
	route.Get("/product/:id", productHandler.GetProductById())
	route.Post("/product/:id/colors", authMW, adminMw, productHandler.AddProductColor())
	route.Post("/product/:id/sizes", authMW, adminMw, productHandler.AddProductSize())
	route.Post("/product/:id/image/upload",authMW,adminMw,productHandler.AddProductImage())
	route.Get("/product/:id/image",productHandler.GetProductImage())
}
