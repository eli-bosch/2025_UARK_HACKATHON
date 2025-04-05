package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/utils"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	userLogin := &models.User{}
	utils.ParseBody(r, userLogin)

	user := db.FindUserByUsername(userLogin.Username)
	if user == nil { //User is not in the database
		w.WriteHeader(404)
		w.Write(nil)
		return
	}

	if user.Password != userLogin.Password {
		w.WriteHeader(401) //Username is in the db, but the passwords don't match
		w.Write(nil)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)

	user := db.InsertUser(*newUser)
	if user == nil {
		w.WriteHeader(401)
		w.Write(nil)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := db.FindAllUsers()

	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	delUser := &models.User{}
	utils.ParseBody(r, delUser)

	user := db.FindUserByUsername(delUser.Username)
	if delUser.Password != user.Password {
		w.WriteHeader(401)
		w.Write(nil)
		return
	}

	deletedUser := db.DeleteUserAndNotes(user.ID)
	if deletedUser == nil {
		w.WriteHeader(404)
		w.Write(nil)
		return
	}

	res, err := json.Marshal(deletedUser)
	if err != nil {
		http.Error(w, "Failed to marshal notes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
