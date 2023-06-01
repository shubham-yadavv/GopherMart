package repositories

import (
	"github.com/shubham-yadavv/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	RegisterUser(user models.User) error
	GetUserWithUsername(username string) (*models.User, error)
	UserExists(username string) (bool, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db}
}

func (u *UserRepoImpl) RegisterUser(user models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepoImpl) GetUserWithUsername(username string) (*models.User, error) {

	user := models.User{}
	result := u.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepoImpl) UserExists(username string) (bool, error) {
	var count int64
	result := u.db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
