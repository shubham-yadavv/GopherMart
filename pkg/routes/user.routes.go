package routes

import (
	"net/http"

	"github.com/shubham-yadavv/go-ecommerce/pkg/controllers"
)

func UserRoutes() {
	http.HandleFunc("/users/signup", controllers.Signup())
	http.HandleFunc("/users/login", controllers.Login())
	http.HandleFunc("/admin/addproduct", controllers.ProductViewerAdmin())
	http.HandleFunc("/users/productview", controllers.SearchProduct())
	http.HandleFunc("/users/search", controllers.SearchProductByQuery())

}
