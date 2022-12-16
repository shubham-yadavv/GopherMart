package main

import (
	"fmt"
	"net/http"

	"github.com/shubham-yadavv/go-ecommerce/pkg/config"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Server is running on port " + config.Config("PORT"))
	http.ListenAndServe(":"+config.Config("PORT"), nil)

}
