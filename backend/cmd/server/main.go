package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/routes"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/ws"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	address     = "localhost:8080"
	corsHandler http.Handler
)

func main() {
	initDatabase()
	registerRoutes()

	fmt.Println("Server listening on port,", address)
	log.Fatal(http.ListenAndServe(address, corsHandler))
}

// Database Initialization
func initDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	db.InitMongoDB(uri, dbName) //Need to change once website is hosted
}

// Route Registration
func registerRoutes() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)

	ws.RegisterWSRoutes(r)
	routes.RegisterNoteRoutes(r)

	corsHandler = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		}),
		handlers.AllowedHeaders([]string{
			"Content-Type", "Authorization", "Origin", "X-Requested-With",
		}),
	)(r)
}

// Testing Methods
func insertTestingData() {
	db.InsertUser(models.User{Username: "ebosch", Password: "secretPassword101"})
	user := db.FindUserByUsername("ebosch")

	for index := range 3 {
		db.InsertNote(user.ID, models.Note{Header: "Header" + string(index), Body: "Body" + string(index)})
	}

	fmt.Printf("SUCCESS!")
}

func test() {
	users := db.FindAllUsers()
	for index, user := range *users {
		fmt.Printf("%v. %v", index, user)
	}
}
