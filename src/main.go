package main

import (
	"api/productapi"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/product/findAll", productapi.FindAll).Methods("GET")
	r.HandleFunc("/api/product/findAllUsers", productapi.FindAllUsers).Methods("GET")
	r.HandleFunc("/api/product/find/{id}", productapi.FindProduct).Methods("GET")
	r.HandleFunc("/api/product/createProduct", productapi.CreateProduct).Methods("POST")

	// corsObj := handlers.AllowedOrigins([]string{"*"})
	err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Listening 3000...")
	}
}
