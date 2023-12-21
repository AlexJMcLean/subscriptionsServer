package users

import (
	"github.com/AlexJMcLean/subscriptions/common"
	"github.com/gin-gonic/gin"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, userValidator)
// Then, you can call model.Save() after the data is ready in DataModel
type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Email string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}


func (userValidator *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, userValidator)
	if err != nil {
		return err
	}
	userValidator.userModel.Username = userValidator.User.Username
	userValidator.userModel.Email = userValidator.User.Email

	passwordErr := userValidator.userModel.setPassword(userValidator.User.Password)
	if passwordErr != nil {
		return passwordErr
	}
	
	return nil
}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}