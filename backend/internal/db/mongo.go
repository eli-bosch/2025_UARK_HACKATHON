// Connection with database
package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName         string
	Client         *mongo.Client
	userCollection *mongo.Collection
	noteCollection *mongo.Collection
)

func InitMongoDB(uri string, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	Client = client
	dbName = name
	log.Println("Connected to MongoDB!")

	userCollection = client.Database(dbName).Collection("users")
	noteCollection = client.Database(dbName).Collection("notes")

}
