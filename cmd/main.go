package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shubham-yadavv/go-ecommerce/pkg/database"
	"github.com/shubham-yadavv/go-ecommerce/pkg/middleware"
	"github.com/shubham-yadavv/go-ecommerce/pkg/routes"
)

func main() {

	database.ConnectDB()

	routes.UserRoutes()

	middleware.Middleware()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
