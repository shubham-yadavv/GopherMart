package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-ecommerce/config"
	"github.com/shubham-yadavv/go-ecommerce/internal/database"
	"github.com/shubham-yadavv/go-ecommerce/internal/routes"
)

func main() {

	port := config.GetEnv("PORT")

	database.ConnectPostgres()

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run(":" + port)

}
