package main

import (
	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/AlexJMcLean/subscriptions/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	
	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))

	r.Run()
}