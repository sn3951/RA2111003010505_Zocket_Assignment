
# Zocket-Product-Management-App

## Project Description

Zocket-Product-Management-App is a backend application built using the Go programming language. It is designed to manage product information and user data for an e-commerce platform. This application provides API endpoints to create users, add products, and retrieve product details. It uses MongoDB as the data storage backend.

---

## Key Features

### User Management:
- Create users with essential details such as name, contact number, and location coordinates.

### Product Management:
- Add products with details like name, description, images, and price.
  
### Data Storage:
- MongoDB is used to store and manage user and product information persistently.

### API Endpoints:
- The backend exposes various RESTful API endpoints for product and user management.

### Testing:
- The project includes comprehensive unit tests to ensure code quality and application stability.

---

## Technologies Used

- **Go (Golang)**: The primary programming language for building the application logic.
- **MongoDB**: A NoSQL database for storing user and product information.
- **Postman**: For testing API endpoints.

---

## Project Structure

```plaintext
RA2111003010505_Zocket_Assignment/
├── api/
│   ├── handlers.go
│   └── routes.go
├── cmd/
│   └── main.go
├── db/
│   ├── database.go
│   └── db_test.go
├── models/
│   ├── models.go
├── .gitignore
├── go.mod
├── go.sum
├── README.md
```

- **cmd/main.go**: The entry point of the application, responsible for starting the server.
- **api/handlers.go**: Contains the logic for handling API requests and responses.
- **api/routes.go**: Defines the routes for API endpoints.
- **db/database.go**: Contains functions to interact with the MongoDB database.
- **models/models.go**: Defines the data models for Users and Products.
- **go.mod**: Defines the Go modules and dependencies.
- **go.sum**: Contains the checksum of Go modules.

---

## Getting Started

Follow the instructions below to set up and run the project on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or above)
- [MongoDB](https://www.mongodb.com/try/download/community) (locally or a cloud instance like MongoDB Atlas)
- [Postman](https://www.postman.com/) (optional, for testing the APIs)

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/your-username/zocket-assignment01.git
cd zocket-assignment01
```

### 2. Install Dependencies

Run the following command to install all Go dependencies:

```bash
go mod tidy
```

### 3. Configure MongoDB

Ensure you have a MongoDB instance running. You can either run MongoDB locally or use a cloud instance like [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

Update the MongoDB connection string in `db/database.go` to match your MongoDB setup:

```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
```

For MongoDB Atlas, you will use the URI provided by Atlas, like:

```go
clientOptions := options.Client().ApplyURI("mongodb+srv://<user>:<password>@cluster0.mongodb.net")
```

### 4. Run the Application

Once everything is set up, run the application:

```bash
go run cmd/main.go
```

The server will start, and you should see something like:

```
Bravo, Product Management Application is running
```

Your API will be available at `http://localhost:8080`.

---

## API Endpoints

### Create User

- **Endpoint**: `POST /users`
- **Payload**:
  ```json
  {
    "name": "John Doe",
    "mobile": "1234567890",
    "latitude": "12.9716",
    "longitude": "77.5946"
  }
  ```
  
### Create Product

- **Endpoint**: `POST /products`
- **Payload**:
  ```json
  {
    "product_name": "Example Product",
    "product_description": "This is an example product.",
    "product_images": ["image_url1", "image_url2"],
    "product_price": 100
  }
  ```

### Get Products

- **Endpoint**: `GET /products`
- **Description**: Retrieves a list of all products.

### Get Product by ID

- **Endpoint**: `GET /products/{productID}`
- **Description**: Retrieves product details by its ID.

---

## Testing

To run the unit tests for the project:

```bash
go test -v ./...
```

This will execute all the tests defined in your project and give you detailed output.

---

## How to Contribute

Feel free to fork the repository and create a pull request. Contributions are always welcome!

### Steps to Contribute:
1. Fork the repository.
2. Clone your fork locally.
3. Create a new branch for your feature or bug fix.
4. Make your changes and test them.
5. Commit and push your changes.
6. Submit a pull request.

---
