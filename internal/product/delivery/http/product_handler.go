package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/internal/models"
	"github.com/Zeo-dev3/fashionmart/internal/product"
	"github.com/Zeo-dev3/fashionmart/pkg/utils"
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

		if err := utils.ValidateRequest(c, &productReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad request, cannot parse json :" + err.Error(),
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
				"product_name": createdProduct.Name,
			},
		})
	}
}

func (h *productHandler) GetAllProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		products, err := h.productUC.GetAllProducts(ctx)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(http.StatusOK).JSON(products)
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

func (h *productHandler) AddProductColor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		idParam := c.Params("id")
		productId, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid product ID",
			})
		}

		var colors []*entity.Color
		if err := c.BodyParser(&colors); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if err := h.productUC.AddProductColor(ctx, uint(productId), colors); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "colors added successfully",
		})
	}
}

func (h *productHandler) AddProductSize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		idParam := c.Params("id")

		idUint64, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Invalid ID format: %s", idParam),
			})
		}

		var sizes []*entity.Size
		if err := c.BodyParser(&sizes); err != nil {
			log.Println(sizes)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid request body",
			})
		}

		productId := uint(idUint64)

		if err := h.productUC.AddProductSize(ctx, productId, sizes); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "size added successfully",
		})
	}
}

func (h *productHandler) AddProductImage() fiber.Handler {
    return func(c *fiber.Ctx) error {
        ctx := c.Context()
        productId := c.Params("id")
        file, err := c.FormFile("image")

        if err != nil {
            return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
                "message": "failed to retrieve file",
            })
        }

        productUint64, err := strconv.ParseUint(productId, 10, 64)
        if err != nil {
            return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
                "message": "invalid product id",
            })
        }

        productUint := uint(productUint64)
        fileName := fmt.Sprintf("%d-%s", productUint, file.Filename)

        // Get the current working directory
        currentDir, err := os.Getwd()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "failed to get current working directory: " + err.Error(),
            })
        }

        // Combine current working directory with the upload path
        uploadDir := filepath.Join(currentDir, "uploads", "image", "product")

        // Create the directory if it doesn't exist
        if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "failed to create directory: " + err.Error(),
            })
        }

        destination := filepath.Join(uploadDir, fileName)

        err = c.SaveFile(file, destination)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "failed save product image to server :" + err.Error(),
            })
        }

        if err := h.productUC.AddProductImage(ctx, productUint, destination); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "failed add image path in database : "+err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "message": "successfully add product image to server",
        })
    }
}

func (h *productHandler) GetProductImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		id := c.Params("id")
		idUint64,err := strconv.ParseUint(id,10,64)

		if err != nil{
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"message":"invalid product id",
			})
		}

		productId := uint(idUint64)
		product,err := h.productUC.GetProductById(ctx,productId)

		if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": fmt.Sprintf("Product with ID %d not found", productId),
            })
        }

		imagePath := product.Image

		if imagePath == "" {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Image not found for the specified product",
            })
        }

		err = c.SendFile(imagePath)
		if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": fmt.Sprintf("Failed to send image file: %v", err),
            })
        }

		return nil
	}
}