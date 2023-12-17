package main

import (
	"github.com/AlexJMcLean/subscriptions/users"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	
	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))

	r.Run()
}