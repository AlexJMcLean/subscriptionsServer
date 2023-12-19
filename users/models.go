package users

import (
	"errors"

	"github.com/AlexJMcLean/subscriptions/common"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID uint `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Email string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&UserModel{})
}

// Hashes the password to safetly store it in the db
// err := userModel.setPassword(pass)
func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password must not be empty")
	}
	bytePassword := []byte(password)

	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil 
}

// Save UserModel to Db returns error information
// if err := SaveOne(&userModel); err != nil {...}
func SaveOne(data interface{}) error {
	return nil
	// TODO: add db connection
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}