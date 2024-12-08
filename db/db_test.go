package db

import (
	"testing"
	"zocket-api/models"

	"github.com/stretchr/testify/assert"
)

func TestInsertProduct(t *testing.T) {
	// Initialize a new MongoDB client for testing
	Connect()

	// Create a product for testing
	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Images:      []string{"image1.jpg", "image2.jpg"},
		Price:       100,
	}

	// Insert the product into the database
	err := InsertProduct(product)
	assert.Nil(t, err)

	// Retrieve the inserted product
	insertedProducts, err := GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, insertedProducts)
	assert.NotEmpty(t, insertedProducts)

	// Ensure the inserted product matches the expected values
	insertedProduct := insertedProducts[len(insertedProducts)-1]
	assert.Equal(t, product.Name, insertedProduct.Name)
	assert.Equal(t, product.Description, insertedProduct.Description)

}

func TestGetProducts(t *testing.T) {
	// Initialize a new MongoDB client for testing
	Connect()

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

	// Insert products into the database
	_ = InsertProduct(product1)
	_ = InsertProduct(product2)

	// Retrieve products from the database
	products, err := GetProducts()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.NotEmpty(t, products)

}
