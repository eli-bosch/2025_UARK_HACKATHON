package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/routes"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/ws"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var address = "localhost:9010"

	r := mux.NewRouter()
	routes.TestRoutes(r)

	ws.RegisterWSRoutes(r)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		}),
		handlers.AllowedHeaders([]string{
			"Content-Type", "Authorization", "Origin", "X-Requested-With",
		}),
	)(r)
	
	fmt.Println("Server listening on port,", address)
	log.Fatal(http.ListenAndServe(address, corsHandler))
}
