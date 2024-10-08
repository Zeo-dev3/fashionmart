package product

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
)

type Reporitory interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetById(ctx context.Context, id uint) (entity.Product, error)
	AddColor(ctx context.Context, productId uint, colors []*entity.Color) error
	GetAll(ctx context.Context) ([]entity.Product, error)
	AddSize(ctx context.Context, productId uint, sizes []*entity.Size) error
	UpdatePath(ctx context.Context,productId uint,path string) error
}
