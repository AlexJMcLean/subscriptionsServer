package products

import "github.com/gin-gonic/gin"

type ProductResponse struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Reference string `json:"reference"` 
}

type ProductSerialiser struct {
	C *gin.Context
	ProductModel
}

func (serialiser *ProductSerialiser) Response() ProductResponse {
	product := ProductResponse{ 
		Name: serialiser.Name,
		Price: serialiser.Price,
		Reference: serialiser.Reference,
	}
	return product
}

type ProductsSerialiser struct {
	C *gin.Context
	Products []ProductModel
}

func (s ProductsSerialiser) Response() []ProductResponse {
	response := []ProductResponse{}

	for _, product := range s.Products {
		serialiser := ProductSerialiser{s.C, product}
		response = append(response, serialiser.Response())
	}

	return response
}