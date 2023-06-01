package services

import (
	"context"
	"fmt"
	"time"

	"github.com/shubham-yadavv/go-ecommerce/internal/database"
	"github.com/shubham-yadavv/go-ecommerce/internal/models"
	"github.com/shubham-yadavv/go-ecommerce/internal/repositories"
	"github.com/shubham-yadavv/go-ecommerce/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user models.User) error
	Login(username string, password string) (string, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepo
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{repo: repositories.NewUserRepo(database.ConnectPostgres())}
}

func (u *UserServiceImpl) RegisterUser(user models.User) error {
	// Check if the user already exists in the repository
	exists, err := u.repo.UserExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user already exists")
	}

	// Generate a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update the user with the hashed password and registration timestamp
	user.Password = string(hashedPassword)
	user.RegisteredAt = time.Now()

	// Register the user in the repository
	err = u.repo.RegisterUser(user)
	if err != nil {
		return err
	}

	// Store the user ID in Redis
	err = database.ConnectRedis().Set(context.Background(), user.Username, user.UserID, 0).Err()
	if err != nil {
		// Rollback the user registration if storing in Redis fails
		return err
	}

	return nil
}

func (u *UserServiceImpl) Login(username string, password string) (string, error) {

	userData, err := u.repo.GetUserWithUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := utils.GenrateToken(userData.Username, int(userData.UserID))

	return token, nil

}
