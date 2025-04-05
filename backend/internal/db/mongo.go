package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	dbName string
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

	noteCollection = client.Database(dbName).Collection("notes")
}

func GetUserNotes(userID string) ([]models.Note, error) {
	var notes []models.Note
	//filter := bson.M{"user_id": userID}

	// Log the filter to ensure it contains the correct userID
	//fmt.Printf("Fetching notes for userID: %s with filter: %v\n", userID, filter)

	cursor, err := noteCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		// Log the error when the query fails
		fmt.Printf("Error fetching notes for userID: %s - Error: %v\n", userID, err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Check if there are any results
	count := 0
	for cursor.Next(context.TODO()) {
		var note models.Note
		if err := cursor.Decode(&note); err != nil {
			// Log any decoding errors
			fmt.Printf("Error decoding note: %v\n", err)
			return nil, err
		}
		notes = append(notes, note)
		count++
	}

	if err := cursor.Err(); err != nil {
		// Log any cursor errors
		fmt.Printf("Cursor error for userID %s: %v\n", userID, err)
		return nil, err
	}

	// Log the number of notes found
	fmt.Printf("Found %d notes for userID: %s\n", count, userID)

	return notes, nil
}
