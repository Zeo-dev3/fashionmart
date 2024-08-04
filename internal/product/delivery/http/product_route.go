package http

import (
	"github.com/Zeo-dev3/fashionmart/internal/product"
	"github.com/gofiber/fiber/v2"
)

func MapProductRoutes(app *fiber.App, productHandler product.Handler, authMW fiber.Handler, adminMw fiber.Handler) {
	group := app.Group("/product")

	group.Post("/", authMW, adminMw, productHandler.AddProduct())
	group.Get("/:id", productHandler.GetProductById())
}
