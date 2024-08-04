package http

import (
	"fmt"
	"strconv"

	"github.com/Zeo-dev3/fashionmart/internal/models"
	"github.com/Zeo-dev3/fashionmart/internal/product"
	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productUC product.UseCase
}

func NewProductHandler(productUsecase product.UseCase) product.Handler {
	return &productHandler{
		productUC: productUsecase,
	}
}

func (h *productHandler) AddProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var productReq models.ProductRequest
		ctx := c.Context()

		if err := c.BodyParser(&productReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad request, cannot parse json",
			})
		}

		createdProduct, err := h.productUC.AddProduct(ctx, &productReq)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "succes added product",
			"product_data": fiber.Map{
				"product_id":   createdProduct.ID,
				"producr_name": createdProduct.Name,
			},
		})
	}
}

func (h *productHandler) GetProductById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		ctx := c.Context()

		idUint64, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Invalid ID format: %s", idParam),
			})
		}

		id := uint(idUint64)

		product, err := h.productUC.GetProductById(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"product_data": product,
		})
	}
}
