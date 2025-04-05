package controller

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Error while marshalling json body")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
