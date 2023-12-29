package common

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init)", err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open(sqlite.Open("../gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (testDBInit) ", err)
	}
	test_db.Logger.LogMode(logger.Info)
	DB = test_db
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	err := os.Remove("./../gorm.db")
	return err
}
