package db

import (
	"context"
	"log"
	"time"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertNote(userID primitive.ObjectID, note models.Note) *models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	note.ID = primitive.NewObjectID()
	note.CreatedAt = now
	note.UpdatedAt = now

	_, err := noteCollection.InsertOne(ctx, note)
	if err != nil {
		log.Println("InsertNote error:", err)
		return nil
	}

	update := bson.M{
		"$push": bson.M{
			"current_notes": note.ID,
		},
	}
	_, err = userCollection.UpdateByID(ctx, userID, update)
	if err != nil {
		log.Println("Failed to update user's CurrentNotes:", err)

		noteCollection.DeleteOne(ctx, bson.M{"_id": note.ID})
		return nil
	}

	log.Printf("Inserted note for user %s: %s", userID.Hex(), note.ID.Hex())
	return &note
}

func FindNoteByID(noteID primitive.ObjectID) *models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(dbName).Collection("notes")

	var note models.Note
	err := collection.FindOne(ctx, bson.M{"_id": noteID}).Decode(&note)
	if err != nil {
		log.Println("FindUserByID error:", err)
		return nil
	}

	return &note
}

func UpdateNote(noteID primitive.ObjectID, newNote models.Note) *models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set update fields
	update := bson.M{
		"$set": bson.M{
			"header":     newNote.Header,
			"body":       newNote.Body,
			"updated_at": time.Now().UTC(),
		},
	}

	// Find the note and update it in one step
	var updatedNote models.Note
	err := noteCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": noteID},
		update,
	).Decode(&updatedNote)

	if err != nil {
		log.Println("UpdateNote error:", err)
		return nil
	}

	log.Printf("Updated note %s\n", noteID.Hex())
	return &updatedNote
}

func DeleteNote(userID primitive.ObjectID, noteID primitive.ObjectID) *models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := Client.Database(dbName)
	notesColl := db.Collection("notes")
	usersColl := db.Collection("users")

	filter := bson.M{
		"_id":     noteID,
		"user_id": userID,
	}

	var deletedNote models.Note
	err := notesColl.FindOneAndDelete(ctx, filter).Decode(&deletedNote)
	if err != nil {
		log.Println("DeleteNote error (not found or doesn't belong to user):", err)
		return nil
	}

	update := bson.M{
		"$pull": bson.M{
			"current_notes": noteID,
		},
	}
	_, err = usersColl.UpdateOne(ctx, bson.M{"_id": userID}, update)
	if err != nil {
		log.Println("Failed to update user's CurrentNotes array:", err)
	}

	log.Printf("Deleted note %s and updated user %s\n", noteID.Hex(), userID.Hex())
	return &deletedNote
}

func DeleteAllNotes() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(dbName).Collection("notes")

	result, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("DeleteAllNotes error:", err)
		return
	}

	log.Printf("Deleted %d notes\n", result.DeletedCount)
}
