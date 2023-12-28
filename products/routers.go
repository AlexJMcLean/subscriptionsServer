package products

import (
	"net/http"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

func ProductsRegister(router *gin.RouterGroup) {
	router.POST("/", CreateProduct)
	router.GET("/", GetProduct)
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
	serialiser := ProductSerialiser{c}
	c.JSON(http.StatusCreated, gin.H{"product": serialiser.Response()})
}

func GetProduct (c *gin.Context) {
	// TODO implement Get Product	
}