package products

import (
	"errors"
	"net/http"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

func ProductsRegister(router *gin.RouterGroup) {
	router.POST("/", CreateProduct)
	router.GET("/", GetProductList)
}

func CreateProduct (c *gin.Context) {
	productModelValidator := NewProductModelValidator()

	if err := productModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
	}

	if err:= SaveProduct(&productModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("product_model", productModelValidator.productModel)
	serialiser := ProductSerialiser{c, productModelValidator.productModel}
	c.JSON(http.StatusCreated, gin.H{"product": serialiser.Response()})
}

func GetProductList (c *gin.Context) {
	productModels, err := FindAllProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("no products found")))
		return
	}
	serialiser := ProductsSerialiser{c, productModels}
	c.JSON(http.StatusOK, gin.H{"products": serialiser.Response()})
}