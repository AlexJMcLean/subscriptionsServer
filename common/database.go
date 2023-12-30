package common

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

const (
	HOST = "database"
	PORT = 5432
)

func Init(username, password, database string) (*gorm.DB, error) {
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	HOST, PORT, username, password, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	DB = db
	return DB, nil
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
