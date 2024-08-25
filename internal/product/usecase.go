package product

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/internal/models"
)

type UseCase interface {
	AddProduct(ctx context.Context, productReq *models.ProductRequest) (models.ProductCreateResponse, error)
	GetProductById(ctx context.Context, id uint) (models.ProductDetailResponse, error)
	AddProductColor(ctx context.Context, productId uint, colors []*entity.Color) error
	GetAllProducts(ctx context.Context) ([]models.ProductDetailResponse, error)
	AddProductSize(ctx context.Context, productId uint, sizes []*entity.Size) error
	AddProductImage(ctx context.Context,productId uint,path string) error
}
