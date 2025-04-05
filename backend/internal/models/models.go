package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Username     string               `bson:"username"`
	Password     string               `bson:"password"`
	CurrentNotes []primitive.ObjectID `bson:"current_notes"`
	CreatedAt    time.Time            `bson:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at"`
}

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Header    string             `bson:"header"`
	Body      string             `bson:"body"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Message struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Body     string             `bson:"body"`
	Sender   time.Time          `bson:"sender"`
	Receiver time.Time          `bson:"receiver"`
	DateTime primitive.DateTime `bson:"datetime"`
}
