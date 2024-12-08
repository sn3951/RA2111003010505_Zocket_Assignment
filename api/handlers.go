package api

import (
	"encoding/json"
	"net/http"
	"time"
	"zocket-api/db"
	"zocket-api/models"

	"github.com/gorilla/mux"
	// "github.com/google/uuid"
)

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Parsing JSON from the request body
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Insert into the database
	err := db.InsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// CreateProduct handles the creation of a new product.
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse JSON from the request body
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input (you can add more validation logic here)

	// Set timestamps
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	// Insert into the database
	err := db.InsertProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProducts retrieves all products.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProductByID retrieves a single product by its ID.
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productID"]

	product, err := db.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
