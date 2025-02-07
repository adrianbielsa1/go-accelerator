package products

import (
	"log"
	"net/http"

	"github.com/adrianbiesa1/go-accelerator/internal/services/products"
	"github.com/adrianbiesa1/go-accelerator/internal/utils/slices"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	productsService products.Service
	loggerService   *log.Logger
}

func NewController(
	engine *gin.Engine,
	productsService products.Service,
	loggerService *log.Logger,
) error {
	controller := &Controller{
		productsService: productsService,
		loggerService:   loggerService,
	}

	group := engine.Group("/v1")

	group.GET("/products", controller.GetProducts)
	group.POST("/products", controller.CreateProduct)
	group.PATCH("/products/:code", controller.UpdateProductByCode)
	group.DELETE("/products/:code", controller.DeleteProductByCode)

	return nil
}

func (controller *Controller) GetProducts(ctx *gin.Context) {
	products, err := controller.productsService.GetProducts(ctx)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		slices.Map(products, controller.mapProductToController),
	)
}

func (controller *Controller) CreateProduct(ctx *gin.Context) {
	product := Product{}

	if err := ctx.BindJSON(&product); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdProduct, err := controller.productsService.CreateProduct(
		ctx,
		controller.mapProductToService(product),
	)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		controller.mapProductToController(createdProduct),
	)
}

func (controller *Controller) UpdateProductByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	productUpdate := ProductUpdate{}

	if err := ctx.BindJSON(&productUpdate); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedProduct, err := controller.productsService.UpdateProductByCode(
		ctx,
		code,
		controller.mapProductUpdateToService(productUpdate),
	)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		controller.mapProductToController(updatedProduct),
	)
}

func (controller *Controller) DeleteProductByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	if err := controller.productsService.DeleteProductByCode(ctx, code); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (controller *Controller) mapProductToController(product products.Product) Product {
	return Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (controller *Controller) mapProductToService(product Product) products.Product {
	return products.Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func (controller *Controller) mapProductUpdateToService(productUpdate ProductUpdate) products.ProductUpdate {
	return products.ProductUpdate{
		Code:        productUpdate.Code,
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
		Price:       productUpdate.Price,
	}
}
