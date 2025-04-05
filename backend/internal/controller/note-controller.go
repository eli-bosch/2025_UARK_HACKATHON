package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/utils"
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
