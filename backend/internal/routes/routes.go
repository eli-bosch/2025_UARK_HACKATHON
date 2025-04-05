package routes

import (
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/login", controller.UserLogin).Methods("POST")
}
