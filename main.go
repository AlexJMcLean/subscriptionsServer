package main

import (
	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/AlexJMcLean/subscriptions/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {

	db := common.Init()
	Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))

	r.Run()
}
