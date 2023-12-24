package users

import (
	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserSerialiser struct {
	c *gin.Context
}

func (serialiser *UserSerialiser) Response() UserResponse {
	myUserModel := serialiser.c.MustGet("user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}
