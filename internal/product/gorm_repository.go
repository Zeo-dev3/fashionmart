package product

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
)

type Reporitory interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetById(ctx context.Context, id uint) (entity.Product, error)
}
