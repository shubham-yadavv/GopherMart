package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-ecommerce/internal/handlers"
)

func UserRoutes(router *gin.Engine) {

	router.POST("/register", handlers.NewUserHandler().RegisterUser)
	router.POST("/login", handlers.NewUserHandler().Login)
}
