package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shubham-yadavv/go-ecommerce/internal/models"
	"github.com/shubham-yadavv/go-ecommerce/internal/services"
)

// UserHandler interface
type UserHandler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
}

// UserHandlerImpl struct
type UserHandlerImpl struct {
	service services.UserService
}

// NewUserHandler function
func NewUserHandler() *UserHandlerImpl {
	return &UserHandlerImpl{service: services.NewUserService()}
}

// CreateUser function
func (u *UserHandlerImpl) RegisterUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = u.service.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (u *UserHandlerImpl) Login(c *gin.Context) {

	type LoginCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	credentials := LoginCredentials{}
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := u.service.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
