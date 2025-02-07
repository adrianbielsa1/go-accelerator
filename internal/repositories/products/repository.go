package products

import (
	"context"
)

type Repository interface {
	GetProducts(ctx context.Context) ([]Product, error)
	CreateProduct(ctx context.Context, product Product) (Product, error)
	UpdateProductByCode(ctx context.Context, code string, productUpdate ProductUpdate) (Product, error)
	DeleteProductByCode(ctx context.Context, code string) error
}
