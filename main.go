package main

import (
	"log"
	"os"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/AlexJMcLean/subscriptions/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {
	dbUser, dbPassword, dbName :=
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB")
	db, err := common.Init(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))

	r.Run()
}
