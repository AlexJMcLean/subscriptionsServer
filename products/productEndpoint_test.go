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
			Name: fmt.Sprintf("Productr%v", i),
			Price: 1.99,
			Reference: fmt.Sprintf("PRODUCT%v", i),
		}
		test_db.Create(&productModel)
		ret = append(ret, productModel)
	}
	return ret
}

func resetDBWithMock() {
	common.TestDBFree(test_db)
	test_db = common.TestDBInit()
	AutoMigrateProduct()
	ProductModelMocker(3)
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
			resetDBWithMock()
		},
		"/product/",
		"POST",
		`{"product":{"name":"TestProduct","price":2.99,"reference":"TESTPRODUCT"}}`,
		http.StatusCreated,
		`{"product":{"name":"TestProduct","price":2.99,"reference":"TESTPRODUCT"}}`,
		"valid data and should return StatusCreated",
	},
}

func TestCreateProduct(t *testing.T) {
	asserts := assert.New(t)

	r := gin.New()
	ProductsRegister(r.Group("/product"))
	
	for _, testData := range CreateProducts {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
	}
}