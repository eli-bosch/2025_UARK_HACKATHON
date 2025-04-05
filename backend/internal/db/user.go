package db

import (
	"context"
	"log"
	"time"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.User) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	user.CurrentNotes = []primitive.ObjectID{} // ✅ right here

	collection := Client.Database(dbName).Collection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("InsertUser error:", err)
		return nil
	}

	return &user
}

func DeleteUserAndNotes(userID primitive.ObjectID) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := Client.Database(dbName)
	usersColl := db.Collection("users")
	notesColl := db.Collection("notes")

	var deletedUser models.User
	err := usersColl.FindOneAndDelete(ctx, bson.M{"_id": userID}).Decode(&deletedUser)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return nil
	}

	_, err = notesColl.DeleteMany(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Println("Cascade delete notes error:", err)
	}

	return &deletedUser
}

func FindUserByUsername(username string) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(dbName).Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Println("FindUserByUsername error:", err)
		return nil
	}

	return &user
}

func FindNotesbyUser(userID primitive.ObjectID) *[]models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(dbName).Collection("notes")
	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Println("FindNotesByUser error:", err)
		return nil
	}
	defer cursor.Close(ctx)

	var notes []models.Note
	if err := cursor.All(ctx, &notes); err != nil {
		log.Println("Cursor decode error:", err)
		return nil
	}

	return &notes
}

func FindAllUsers() *[]models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(dbName).Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("FindAllUser erro:", err)
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		log.Println("Error decoding users:", err)
	}

	return &users

}

func DeleteAllUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(dbName).Collection("users")

	result, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("DeleteAllUsers error:", err)
		return
	}

	log.Printf("Deleted %d users\n", result.DeletedCount)
}
