package users

import (
	"net/http"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", usersRegistration)
}


func usersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("my_user_model", userModelValidator.userModel)
	serialiser := userSerialiser{c}
	c.JSON(http.StatusCreated, gin.H{"user": serialiser.Response()})
}