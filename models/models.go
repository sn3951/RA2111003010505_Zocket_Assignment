package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the user model.
type User struct {
	ID        primitive.ObjectID ` json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Mobile    string             `json:"mobile" bson:"mobile"`
	Latitude  string             `json:"latitude" bson:"latitude"`
	Longitude string             `json:"longitude" bson:"longitude"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// Product represents the product model.
type Product struct {
	ID          primitive.ObjectID `json:"product_id" bson:"_id,omitempty"`
	Name        string             `json:"product_name" bson:"name"`
	Description string             `json:"product_description" bson:"description"`
	Images      []string           `json:"product_images" bson:"images"`
	Price       int                `json:"product_price" bson:"price"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
