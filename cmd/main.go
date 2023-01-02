package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-ecommerce/pkg/database"
	"github.com/shubham-yadavv/go-ecommerce/pkg/routes"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

	database.DBSet()
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":" + port)

}
