package product

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/models"
)

type UseCase interface {
	AddProduct(ctx context.Context, productReq *models.ProductRequest) (models.ProductCreateResponse, error)
	GetProductById(ctx context.Context, id uint) (models.ProductDetailResponse, error)
}
