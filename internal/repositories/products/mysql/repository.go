package mysql

import (
	"context"
	"fmt"

	"github.com/adrianbiesa1/go-accelerator/internal/repositories/products"
	"github.com/adrianbiesa1/go-accelerator/internal/utils/slices"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySQLRepository struct {
	gormDB *gorm.DB
}

type Configuration struct {
	Username string
	Password string
	Host     string
	Port     uint16
	Database string
}

func NewMySQLRepository(configuration Configuration) (*mySQLRepository, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configuration.Username, configuration.Password,
		configuration.Host, configuration.Port,
		configuration.Database,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &mySQLRepository{
		gormDB: database,
	}, nil
}

func (repository *mySQLRepository) GetProducts(ctx context.Context) ([]products.Product, error) {
	products := []Product{}
	transaction := repository.gormDB.Find(&products)

	if transaction.Error != nil {
		return nil, transaction.Error
	}

	return slices.Map(products, repository.mapProductToRepository), nil
}

func (repository *mySQLRepository) CreateProduct(ctx context.Context, product products.Product) (products.Product, error) {
	mappedProduct := repository.mapProductToMySQL(product)
	transaction := repository.gormDB.Create(&mappedProduct)

	if transaction.Error != nil {
		return products.Product{}, transaction.Error
	}

	return product, nil
}

func (repository *mySQLRepository) UpdateProductByCode(ctx context.Context, code string, productUpdate products.ProductUpdate) (products.Product, error) {
	product := Product{}

	if err := repository.gormDB.Transaction(func(tx *gorm.DB) error {
		if nestedTx := tx.
			Where(Product{Code: code}).
			First(&product); nestedTx.Error != nil {
			return nestedTx.Error
		}

		if productUpdate.Code != nil {
			product.Code = *productUpdate.Code
		}

		if productUpdate.Name != nil {
			product.Name = *productUpdate.Name
		}

		if productUpdate.Description != nil {
			product.Description = *productUpdate.Description
		}

		if productUpdate.Price != nil {
			product.Price = *productUpdate.Price
		}

		if nestedTx := tx.Save(&product); nestedTx.Error != nil {
			return nestedTx.Error
		}

		return nil
	}); err != nil {
		return products.Product{}, err
	}

	return repository.mapProductToRepository(product), nil
}

func (repository *mySQLRepository) DeleteProductByCode(ctx context.Context, code string) error {
	transaction := repository.gormDB.Where("code = ?", code).Delete(&Product{})

	if transaction.RowsAffected <= 0 {
		return fmt.Errorf("code `%s` not found", code)
	}

	return transaction.Error
}

func (repository *mySQLRepository) mapProductToRepository(product Product) products.Product {
	return products.Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (repository *mySQLRepository) mapProductToMySQL(product products.Product) Product {
	return Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}
