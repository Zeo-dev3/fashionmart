package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/Zeo-dev3/fashionmart/internal/product"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.Reporitory {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := tx.Create(product).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("cannot create product")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error while commit to db")
	}

	return nil
}

func (r *productRepository) GetById(ctx context.Context, id uint) (entity.Product, error) {
	var product entity.Product

	err := r.db.WithContext(ctx).Preload("Colors").Preload("Sizes").First(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, fmt.Errorf("product with ID %d not found", id)
		}

		return entity.Product{}, fmt.Errorf("cannot search product %w", err)
	}

	return product, nil
}
