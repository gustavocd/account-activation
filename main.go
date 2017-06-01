package main

import (
	"fmt"
	"net/http"

	"github.com/gustavocd/confirmation-account/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()

	mux.GET("/", controllers.Index)
	mux.POST("/signup", controllers.Create)

	fmt.Println("Server running at localhost:8080")
	http.ListenAndServe(":8080", mux)
}
