package users

import (
	"net/http"

	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("user_model", userModelValidator.userModel)
	serialiser := UserSerialiser{c}
	c.JSON(http.StatusCreated, gin.H{"user": serialiser.Response()})
}

func UserRetrieve(c *gin.Context) {
	serialiser := UserSerialiser{c}
	c.JSON(http.StatusOK, gin.H{"user": serialiser.Response()})
}
