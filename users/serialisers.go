package users

import (
	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Token    string  `json:"token"`
}

type userSerialiser struct {
	c *gin.Context
}

func (serialiser *userSerialiser) Response() UserResponse {
	myUserModel := serialiser.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email: myUserModel.Email,
		Token: common.GenToken(myUserModel.ID),
	}
	return user
}