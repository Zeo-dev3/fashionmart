package usecase

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/models"
	"github.com/Zeo-dev3/fashionmart/internal/models/converter"
	"github.com/Zeo-dev3/fashionmart/internal/product"
)

type productUsecase struct {
	productRepo product.Reporitory
}

func NewProductUsecas(productRepo product.Reporitory) product.UseCase {
	return &productUsecase{productRepo: productRepo}
}

func (u *productUsecase) AddProduct(ctx context.Context, productReq *models.ProductRequest) (models.ProductCreateResponse, error) {
	productEntity := converter.ToProductEntity(productReq)

	err := u.productRepo.CreateProduct(ctx, &productEntity)
	if err != nil {
		return models.ProductCreateResponse{}, err
	}

	response := models.ProductCreateResponse{
		ID:   productEntity.ID,
		Name: productEntity.Name,
	}

	return response, nil
}

func (u *productUsecase) GetProductById(ctx context.Context, id uint) (models.ProductDetailResponse, error) {
	product, err := u.productRepo.GetById(ctx, id)
	if err != nil {
		return models.ProductDetailResponse{}, err
	}

	productResponse := converter.ToProductDetailResp(&product)

	return productResponse, nil
}
