package usecase

import (
	"context"

	"github.com/Zeo-dev3/fashionmart/internal/entity"
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

func (u *productUsecase) AddProductColor(ctx context.Context, productId uint, colors []*entity.Color) error {
	err := u.productRepo.AddColor(ctx, productId, colors)
	if err != nil {
		return err
	}

	return nil
}

func (u *productUsecase) GetAllProducts(ctx context.Context) ([]models.ProductDetailResponse, error) {
	products, err := u.productRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToProductDetailResponses(products), nil
}

func (u *productUsecase) AddProductSize(ctx context.Context,productId uint,sizes []*entity.Size) error {
	err := u.productRepo.AddSize(ctx,productId,sizes)
	if err != nil {
		return err
	}
	return nil
}