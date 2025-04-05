package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.InitMongoDB("mongodb://localhost:27017") //Need to change once website is hosted
	r := mux.NewRouter()
	routes.TestRoutes(r)

	fmt.Println("server listening...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
