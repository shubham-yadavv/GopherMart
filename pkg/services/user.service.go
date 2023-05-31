package services

import (
	"github.com/shubham-yadavv/go-ecommerce/pkg/config"
	"github.com/shubham-yadavv/go-ecommerce/pkg/models"
)

func Signup(user *models.User) {
	config.DB.Create(user)
}

func Login(user *models.User) {
	config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user)
}
