package services

import (
	"errors"

	"go-api/models"
	"go-api/repositories"
	"go-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input map[string]string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input["password"]), 10)

	user := models.User{
		Name:     input["name"],
		Email:    input["email"],
		Password: string(hash),
		Role:     "user", // default
	}

	return repositories.CreateUser(&user)
}

func LoginUser(input map[string]string) (string, error) {
	user, err := repositories.GetUserByEmail(input["email"])
	if err != nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input["password"]))
	if err != nil {
		return "", errors.New("invalid password")
	}

	return utils.GenerateToken(user.ID, user.Role)
}
