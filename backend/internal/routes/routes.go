package routes

import (
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/controller"
	"github.com/gorilla/mux"
)

var TestRoutes = func(router *mux.Router) {
	//router.HandleFunc("/test", controller.TestConnection).Methods("POST")
	return
}

var NoteRoutes = func(router *mux.Router) {
	//router.HandleFunc("/api/notes/{userID}", controller.GetUserNotesHandler).Methods("GET")
	router.HandleFunc("/api/notes/{user_id}", controller.GetUserNotesHandler).Methods("GET")
}
