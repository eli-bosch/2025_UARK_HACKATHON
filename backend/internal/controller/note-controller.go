package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/utils"
	"github.com/gorilla/mux"
)

func GetUserNotes(w http.ResponseWriter, r *http.Request) {
	userReq := &models.User{}
	utils.ParseBody(r, userReq)

	user := db.FindUserByUsername(userReq.Username)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	notes := db.FindNotesbyUser(user.ID)
	if notes == nil {
		http.Error(w, "No notes found for user", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	userReq := &models.User{}
	utils.ParseBody(r, userReq)

	user := db.FindUserByUsername(userReq.Username)
	if user.Password != userReq.Password { //Security
		w.WriteHeader(401)
		w.Write(nil)
		return
	}

	newNote := models.Note{}
	newNote.Header = "Untitled(Header)"
	newNote.Body = "Untitled(Body)"

	note := db.InsertNote(user.ID, newNote)

	res, err := json.Marshal(note)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	updatedNote := &models.Note{}
	utils.ParseBody(r, updatedNote)

	note := db.UpdateNote(updatedNote.ID, *updatedNote)
	if note == nil {
		w.WriteHeader(404)
		w.Write(nil)
		return
	}

	res, err := json.Marshal(note)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]

	delNote := &models.Note{}
	utils.ParseBody(r, delNote)

	user := db.FindUserByUsername(userName)
	if user == nil {
		w.WriteHeader(404)
		w.Write(nil)
		return
	}
	fmt.Println("MADE IT HERE")
	note := db.DeleteNote(delNote.ID, user.ID)
	if note == nil {
		w.WriteHeader(404)
		w.Write(nil)
		return
	}

	res, err := json.Marshal(note)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
