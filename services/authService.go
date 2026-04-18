package services

import (
	"errors"

	"go-api/config"
	"go-api/models"
	"go-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input map[string]string) error {

	hash, _ := bcrypt.GenerateFromPassword([]byte(input["password"]), 10)

	user := models.User{
		Name:     input["name"],
		Email:    input["email"],
		Password: string(hash),
	}

	return config.DB.Create(&user).Error
}

func LoginUser(input map[string]string) (string, error) {

	var user models.User

	if err := config.DB.Where("email = ?", input["email"]).First(&user).Error; err != nil {
		return "", errors.New("invalid email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input["password"]))
	if err != nil {
		return "", errors.New("invalid password")
	}

	return utils.GenerateToken(user.ID)
}