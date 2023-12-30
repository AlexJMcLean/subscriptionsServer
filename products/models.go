package products

import "github.com/AlexJMcLean/subscriptions/common"

type ProductModel struct {
	ID uint `gorm:"primary_key"`
	Reference string `gorm:"column:reference;uniqueindex"`
	Name string `gorm:"column:name"`
	Price float64 `gorm:"column:price"`
}

func AutoMigrateProduct() {
	db := common.GetDB()
	db.AutoMigrate(&ProductModel{})
}

func SaveProduct(data *ProductModel) error {
	db := common.GetDB()
	err := db.Create(data).Error
	return err
}

func FindAllProducts() ([]ProductModel, error) {
	db := common.GetDB()
	var products []ProductModel
	
	tx := db.Begin()
	result := tx.Find(&products)

	err := result.Error

	return products, err
}