package http

import (
	"github.com/Zeo-dev3/fashionmart/internal/product"
	"github.com/gofiber/fiber/v2"
)

func MapProductRoutes(app *fiber.App, productHandler product.Handler, authMW fiber.Handler, adminMw fiber.Handler) {
	group := app.Group("/product")

	group.Get("/", productHandler.GetAllProducts())
	group.Post("/", authMW, adminMw, productHandler.AddProduct())
	group.Get("/:id", productHandler.GetProductById())
	group.Post("/:id/colors", authMW, adminMw, productHandler.AddProductColor())
	group.Post("/:id/sizes", authMW, adminMw, productHandler.AddProductSize())
}
