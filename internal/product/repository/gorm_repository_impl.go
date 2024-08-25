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

func (r *productRepository) AddColor(ctx context.Context, productId uint, colors []*entity.Color) error {
	product := &entity.Product{}

	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	// mencari product dengan id dari argument function
	if err := tx.Where("id = ?", productId).First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return fmt.Errorf("product with ID %d not found. Please check product ID and try again", productId)
		}

		return err
	}

	// mencoba untuk menyimpan data color ke product
	if err := tx.Model(product).Association("Colors").Append(&colors); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetAll(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product
	result := r.db.WithContext(ctx).Preload("Colors").Preload("Sizes").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *productRepository) AddSize(ctx context.Context, productId uint, sizes []*entity.Size) error {
	product := &entity.Product{}

	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := tx.Where("id = ?", productId).First(product).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed find product with id %d", productId)
		}
		return err
	}

	if err := tx.Model(product).Association("Sizes").Append(sizes); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add color for product with id %d", productId)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit change")
	}

	tx.Commit()

	return nil
}

func (r *productRepository) UpdatePath(ctx context.Context,productId uint,path string) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := tx.Model(&entity.Product{}).Where("id = ?",productId).Update("image",path).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("cannot save path to the database : %v",err)
	}

	if err := tx.Commit().Error ;err != nil{
		tx.Rollback()
		return fmt.Errorf("failed to commit change : %v",err)
	}

	return nil
}