package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/shubham-yadavv/go-ecommerce/pkg/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())

}
