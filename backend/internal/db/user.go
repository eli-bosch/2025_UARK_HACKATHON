package db

import (
	"context"
	"log"
	"time"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert into userCollection
func InsertUser(user models.User) *models.User {
	//Connects to collection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Makes sure that there are no duplicaton
	if FindUserByUsername(user.Username) != nil {
		log.Println("ERROR: username aalready present")
		return nil
	}

	//Fills out the user struct
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	user.CurrentNotes = []primitive.ObjectID{}

	//Inserts struct into the collection
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("InsertUser error:", err)
		return nil
	}

	return &user
}

// Deletes user and cascade deletes notes
func DeleteUserAndNotes(userID primitive.ObjectID) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var deletedUser models.User
	err := userCollection.FindOneAndDelete(ctx, bson.M{"_id": userID}).Decode(&deletedUser)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return nil
	}

	_, err = noteCollection.DeleteMany(ctx, bson.M{"user_id": userID})
	if err != nil {
		log.Println("Cascade delete notes error:", err)
	}

	return &deletedUser
}

// Finds user by username
func FindUserByUsername(username string) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Println("FindUserByUsername error:", err)
		return nil
	}

	return &user
}

// Finds user by ID
func FindUserByID(userID primitive.ObjectID) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Println("FindUserByID error:", err)
		return nil
	}

	return &user
}

// Finds the array of notes using the userID
func FindNotesbyUser(userID primitive.ObjectID) *[]models.Note {
	user := FindUserByID(userID)
	length := len(user.CurrentNotes)

	notes := make([]models.Note, length)

	for index, noteID := range user.CurrentNotes {
		note := FindNoteByID(noteID)

		notes[index] = *note
	}

	return &notes
}

// Returns all users
func FindAllUsers() *[]models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
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

// Deletes all users
func DeleteAllUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := userCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("DeleteAllUsers error:", err)
		return
	}

	log.Printf("Deleted %d users\n", result.DeletedCount)
}

// Updates User
func UpdateUser(username string, newUser models.User) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set update fields
	update := bson.M{
		"$set": bson.M{
			"password":      newUser.Password,
			"current_notes": newUser.CurrentNotes,
			"updated_at":    time.Now().UTC(),
		},
	}

	// Find the note and update it in one step
	var updatedUser models.User
	err := userCollection.FindOneAndUpdate(
		ctx,
		bson.M{"username": username},
		update,
	).Decode(&updatedUser)

	if err != nil {
		log.Println("UpdateUser error:", err)
		return nil
	}

	log.Printf("Updated user %s\n", updatedUser.ID.Hex())
	return &updatedUser
}
