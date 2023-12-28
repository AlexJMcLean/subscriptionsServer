package products

import "github.com/gin-gonic/gin"

type ProductResponse struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Reference string `json:"reference"` 
}

type ProductSerialiser struct {
	c *gin.Context
}

func (serialiser *ProductSerialiser) Response() ProductResponse {
	productModel := serialiser.c.MustGet("product_model").(ProductModel)
	product := ProductResponse{ 
		Name: productModel.Name,
		Price: productModel.Price,
		Reference: productModel.Reference,
	}
	return product
}