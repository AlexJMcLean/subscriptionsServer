package products

import (
	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

type ProductModelValidator struct {
	Product struct {
		Reference string `form:"reference" json:"reference" binding:"required,alphanum,min=4,max=255"`
		Name string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
		Price float64 `form:"price" json:"price" binding:"required"`
	}
	productModel ProductModel `json:"-"`
}

func (productModelValidator *ProductModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, productModelValidator)
	if err != nil {
		return err
	}
	productModelValidator.productModel.Name = productModelValidator.Product.Name
	productModelValidator.productModel.Price = productModelValidator.Product.Price
	productModelValidator.productModel.Reference = productModelValidator.Product.Reference

	return nil
}

func NewProductModelValidator() ProductModelValidator {
 productModelValidator := ProductModelValidator{}
 return productModelValidator
}