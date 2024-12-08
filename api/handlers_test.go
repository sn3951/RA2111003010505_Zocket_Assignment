// handlers_test.go

package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zocket-api/db"
	"zocket-api/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Initialize a new MongoDB client for testing
	db.Connect()

	// Create a request body for testing
	user := models.User{
		Name:      "Test User",
		Mobile:    "1234567890",
		Latitude:  "12.34",
		Longitude: "56.78",
	}
	userJSON, _ := json.Marshal(user)
	body := bytes.NewReader(userJSON)

	// Create a request and recorder
	req, err := http.NewRequest("POST", "/users", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Verify the response body
	var responseUser models.User
	err = json.Unmarshal(rr.Body.Bytes(), &responseUser)
	assert.Nil(t, err)

}

func TestCreateProduct(t *testing.T) {
	// Initialize a new MongoDB client for testing
	db.Connect()

	// Create a request body for testing
	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Images:      []string{"image1.jpg", "image2.jpg"},
		Price:       100,
	}
	productJSON, _ := json.Marshal(product)
	body := bytes.NewReader(productJSON)

	// Create a request and recorder
	req, err := http.NewRequest("POST", "/products", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(CreateProduct)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Verify the response body
	var responseProduct models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &responseProduct)
	assert.Nil(t, err)

}

func TestGetProducts(t *testing.T) {
	// Initialize a new MongoDB client for testing
	db.Connect()

	// Insert test data into the database
	product1 := models.Product{
		Name:        "Product 1",
		Description: "Description 1",
		Images:      []string{"image1.jpg"},
		Price:       50,
	}
	product2 := models.Product{
		Name:        "Product 2",
		Description: "Description 2",
		Images:      []string{"image2.jpg"},
		Price:       75,
	}

	db.InsertProduct(product1)
	db.InsertProduct(product2)

	// Create a request and recorder
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify the response body
	var responseProducts []models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &responseProducts)
	assert.Nil(t, err)

}
