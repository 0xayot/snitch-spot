package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	dbUrl  = os.Getenv("MONGODB_URI")
)

// Connect initializes the MongoDB connection
func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set up the MongoDB connection options
	clientOptions := options.Client().ApplyURI(dbUrl)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("MongoDB connection error: %v", err)
	}

	// Check if the connection was successful
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("MongoDB connection failed: %v", err)
	}

	log.Println("DB connection successful!")
	return nil
}

// Close closes the MongoDB connection
func CloseDBConnection() {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Printf("Error closing MongoDB connection: %v", err)
		}
	}
}

// GetCollection returns a reference to the MongoDB collection
func GetCollection(dbName string, collection string) *mongo.Collection {
	return client.Database(dbName).Collection(collection)
}
