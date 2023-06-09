package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// for localhost mongoDB
// const connectionString = "mongodb://localhost:27017"
const connectionString = "mongodb://localhost:27017"

var credential = options.Credential{
	Username: "root",
	Password: "admin",
}

// Database Name
const dbName = "todolist"

// Collection name
const collName = "tasks"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func Init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(credential)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func GetCollection() *mongo.Collection {
	return collection
}
