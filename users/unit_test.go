package users

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var test_db *gorm.DB

func newUserModel() UserModel {
	return UserModel{
		ID: 2,
		Username: "abc123",
		Email: "test@test.com",
		PasswordHash: "",
	}
}
func UserModelMocker(n int) []UserModel {
	var offset int
	test_db.Model(&UserModel{}).Count(&offset)
	var ret []UserModel
	for i := offset + 1; i <= offset+n; i++ {
		userModel := UserModel{
			Username: fmt.Sprintf("user%v", i),
			Email:    fmt.Sprintf("user%v@linkedin.com", i),
		}
		userModel.setPassword("password123")
		test_db.Create(&userModel)
		ret = append(ret, userModel)
	}
	return ret
}

func TestUserModel(t *testing.T) {
	asserts := assert.New(t)

	userModel := newUserModel()
	err := userModel.checkPassword("")
	asserts.Error(err, "empty password should return err")

	userModel = newUserModel()
	err = userModel.setPassword("")
	asserts.Error(err, "empty password can not be set null")

	userModel = newUserModel()
	err = userModel.setPassword("asd123!@#ASD")
	asserts.NoError(err, "password should be set successful")
	asserts.Len(userModel.PasswordHash, 60, "password hash length should be 60")

	err = userModel.checkPassword("sd123!@#ASD")
	asserts.Error(err, "password should be checked and not validated")

	err = userModel.checkPassword("asd123!@#ASD")
	asserts.NoError(err, "password should be checked and validated")
}