package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/utils"
)

func GetUserNotes(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	fmt.Println("objID: ", user.ID)

	notes := db.FindNotesbyUser(user.ID)
	if notes == nil {
		w.WriteHeader(404)
		w.Write(nil)
		return
	}

	res, err := json.Marshal(notes)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
