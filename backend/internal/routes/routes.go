package routes

import (
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/login", controller.UserLogin).Methods("POST")
}

var NoteRoutes = func(router *mux.Router) {
	//router.HandleFunc("/api/notes/{userID}", controller.GetUserNotesHandler).Methods("GET")
	router.HandleFunc("/api/notes/user", controller.GetUserNotes).Methods("GET")
}
