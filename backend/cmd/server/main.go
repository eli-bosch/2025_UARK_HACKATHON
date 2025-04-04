package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.TestRoutes(r)
	
	fmt.Println("server listening...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
