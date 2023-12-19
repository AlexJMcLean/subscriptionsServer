package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../gorm.db")
	if err != nil {
		fmt.Println("db err: (Init)", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
} 

func GetDB() *gorm.DB {
	return DB
}