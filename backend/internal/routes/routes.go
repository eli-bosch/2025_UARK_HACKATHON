package routes

import (
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/controller"
	"github.com/gorilla/mux"
)

var TestRoutes = func(router *mux.Router) {
	router.HandleFunc("/test/{content}", controller.TestConnection).Methods("PUT")
}
