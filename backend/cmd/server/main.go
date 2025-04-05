package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	db.InitMongoDB(uri, dbName) //Need to change once website is hosted

	r := mux.NewRouter()
	//routes.TestRoutes(r)
	fmt.Println("server listening...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
