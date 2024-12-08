package db

import (
	"context"
	"log"

	"zocket-api/models"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect initializes a connection to the MongoDB database.
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
}

// InsertUser inserts a new user into the database.
func InsertUser(user models.User) error {
	collection := client.Database("zocket").Collection("users")

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	log.Println("User inserted successfully")
	return nil
}

// InsertProduct inserts a new product into the database.
func InsertProduct(product models.Product) error {
	collection := client.Database("zocket").Collection("products")

	_, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		log.Println("Error inserting product:", err)
		return err
	}

	log.Println("Product inserted successfully")
	return nil
}

// GetProducts retrieves all products.
func GetProducts() ([]models.Product, error) {
	collection := client.Database("zocket").Collection("products")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching all products:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	if err := cursor.All(context.Background(), &products); err != nil {
		log.Println("Error decoding all products:", err)
		return nil, err
	}

	log.Println("All products retrieved successfully")
	return products, nil
}

// GetProductByID retrieves a single product by its ID.
func GetProductByID(productID string) (models.Product, error) {
	collection := client.Database("zocket").Collection("products")

	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Println("Invalid product ID:", err)
		return models.Product{}, err
	}

	filter := bson.M{"_id": objID}

	var product models.Product
	if err := collection.FindOne(context.Background(), filter).Decode(&product); err != nil {
		log.Println("Error fetching product by ID:", err)
		return models.Product{}, err
	}

	log.Println("Product retrieved by ID successfully")
	return product, nil
}
