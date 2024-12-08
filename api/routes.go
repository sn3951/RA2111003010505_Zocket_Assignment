package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes API routes.
func SetupRoutes() {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", CreateUser).Methods("POST")

	// Product routes
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{productID}", GetProductByID).Methods("GET")

	// Start the server
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
