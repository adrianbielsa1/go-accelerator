package main

import (
	"log"

	"github.com/adrianbiesa1/go-accelerator/internal/controllers/products"
	productsRepository "github.com/adrianbiesa1/go-accelerator/internal/repositories/products/mysql"
	productsService "github.com/adrianbiesa1/go-accelerator/internal/services/products"
	"github.com/gin-gonic/gin"
)

func main() {
	productsRepository, err := productsRepository.NewMySQLRepository(productsRepository.Configuration{
		Username: "root",
		Password: "examplerootpass",
		Database: "exampledatabase",
		Host:     "maria-db",
		Port:     3306,
	})

	if err != nil {
		panic(err)
	}

	loggerService := log.Default()
	productsService := productsService.NewService(productsRepository, loggerService)

	engine := gin.Default()

	if err := products.NewController(engine, productsService, loggerService); err != nil {
		panic(err)
	}

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
