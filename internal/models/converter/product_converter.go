package converter

import (
	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/internal/models"
)

func ToProductEntity(product *models.ProductRequest) entity.Product {
	return entity.Product{
		Name:          product.Name,
		Description:   product.Description,
		CurrentPrice:  product.CurrentPrice,
		OriginalPrice: product.OriginalPrice,
		Currency:      product.Currency,
	}
}

func ToProductDetailResp(productEntity *entity.Product) models.ProductDetailResponse {
	colors := make([]models.Color, len(productEntity.Colors))
	for i, color := range colors {
		colors[i] = models.Color{
			ID:   color.ID,
			Name: color.Name,
			Hex:  color.Hex,
		}
	}

	sizes := make([]models.Size, len(productEntity.Sizes))
	for i, size := range sizes {
		sizes[i] = models.Size{
			ID:    size.ID,
			Value: size.Value,
		}
	}

	return models.ProductDetailResponse{
		ID:             productEntity.ID,
		Name:           productEntity.Name,
		Rating:         productEntity.Rating,
		ReviewsCount:   productEntity.ReviewsCount,
		ReviewersCount: productEntity.ReviewersCount,
		Description:    productEntity.Description,
		CurrentPrice:   productEntity.CurrentPrice,
		OriginalPrice:  productEntity.OriginalPrice,
		Currency:       productEntity.Currency,
		Colors:         colors,
		Sizes:          sizes,
		CreatedAt:      productEntity.CreatedAt,
		UpdatedAt:      productEntity.UpdatedAt,
	}
}
