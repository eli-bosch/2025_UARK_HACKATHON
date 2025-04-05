package routes

import (
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/login", controller.UserLogin).Methods("POST")     //Takes in username and password
	router.HandleFunc("/user/signup", controller.UserSignUp).Methods("POST")   //Takes in username
	router.HandleFunc("/user/all", controller.GetAllUsers).Methods("GET")      //takes in nothing
	router.HandleFunc("/user/delete", controller.DeleteUser).Methods("DELETE") //takes in username and password
}

var RegisterNoteRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/notes/user", controller.GetUserNotes).Methods("GET") //Takes in username
}
