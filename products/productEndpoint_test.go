package products

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var test_db *gorm.DB

// Creates n users and adds them into the test_db
// Param n => num of users to add to db
func ProductModelMocker(n int64) []ProductModel {
	var offset int64
	test_db.Model(&ProductModel{}).Count(&offset)
	var ret []ProductModel
	for i := offset + 1; i <= offset+n; i++ {
		productModel := ProductModel{
			Name: fmt.Sprintf("Product%v", i),
			Price: 1.99,
			Reference: fmt.Sprintf("PRODUCT%v", i),
		}
		test_db.Create(&productModel)
		ret = append(ret, productModel)
	}
	return ret
}

// resests DB with i prducts
// @Param i = number of products to populate DB with
func resetDBWithMock(i int64) {
	common.TestDBFree(test_db)
	test_db = common.TestDBInit()
	AutoMigrateProduct()
	ProductModelMocker(i)
}

var CreateProducts = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	{
		func(req *http.Request) {
			resetDBWithMock(1)
		},
		"/product/",
		"POST",
		`{"product":{"name":"TestProduct","price":2.99,"reference":"TESTPRODUCT"}}`,
		http.StatusCreated,
		`{"product":{"name":"TestProduct","price":2.99,"reference":"TESTPRODUCT"}}`,
		"valid data and should return StatusCreated",
	},

	{
		func(req *http.Request) {
			resetDBWithMock(2)
		},
		"/product/",
		"GET",
		``,
		http.StatusOK,
		`{"products":[{"name":"Product1","price":1.99,"reference":"PRODUCT1"},{"name":"Product2","price":1.99,"reference":"PRODUCT2"}]}`,
		"valid data returned and should return StatusOK",
	},
}

func TestProductEndpoint(t *testing.T) {
	asserts := assert.New(t)
	
	r := gin.New()
	ProductsRegister(r.Group("/product"))
	
	for _, testData := range CreateProducts {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)
		
		testData.init(req)
		defer common.TestDBFree(test_db)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)

		asserts.Equal(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
	}
}