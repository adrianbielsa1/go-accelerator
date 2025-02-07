package products

import (
	"context"
	"log"

	"github.com/adrianbiesa1/go-accelerator/internal/repositories/products"
	"github.com/adrianbiesa1/go-accelerator/internal/utils/slices"
)

type Service interface {
	GetProducts(ctx context.Context) ([]Product, error)
	CreateProduct(ctx context.Context, product Product) (Product, error)
	UpdateProductByCode(ctx context.Context, code string, productUpdate ProductUpdate) (Product, error)
	DeleteProductByCode(ctx context.Context, code string) error
}

type defaultService struct {
	productsRepository products.Repository
	loggerService      *log.Logger
}

func NewService(
	productsRepository products.Repository,
	loggerService *log.Logger,
) *defaultService {
	return &defaultService{
		productsRepository: productsRepository,
		loggerService:      loggerService,
	}
}

func (service *defaultService) GetProducts(ctx context.Context) ([]Product, error) {
	products, err := service.productsRepository.GetProducts(ctx)

	if err != nil {
		return nil, err
	}

	return slices.Map(products, service.mapProductToService), nil
}

func (service *defaultService) CreateProduct(ctx context.Context, product Product) (Product, error) {
	repositoryProduct, err := service.productsRepository.CreateProduct(
		ctx, service.mapProductToRepository(product),
	)

	if err != nil {
		return Product{}, err
	}

	return service.mapProductToService(repositoryProduct), nil
}

func (service *defaultService) UpdateProductByCode(ctx context.Context, code string, productUpdate ProductUpdate) (Product, error) {
	repositoryProduct, err := service.productsRepository.UpdateProductByCode(
		ctx, code, service.mapProductUpdateToRepository(productUpdate),
	)

	if err != nil {
		return Product{}, err
	}

	return service.mapProductToService(repositoryProduct), nil
}

func (service *defaultService) DeleteProductByCode(ctx context.Context, code string) error {
	return service.productsRepository.DeleteProductByCode(ctx, code)
}

func (service *defaultService) mapProductToService(product products.Product) Product {
	return Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (service *defaultService) mapProductToRepository(product Product) products.Product {
	return products.Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (service *defaultService) mapProductUpdateToRepository(productUpdate ProductUpdate) products.ProductUpdate {
	return products.ProductUpdate{
		Code:        productUpdate.Code,
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
		Price:       productUpdate.Price,
	}
}
