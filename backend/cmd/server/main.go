package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/db"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/models"
	"github.com/eli-bosch/2025_UARK_HACKATHON/internal/routes"
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

	//testInsertDB()
	testDB()

	r := mux.NewRouter()
	routes.TestRoutes(r)
	fmt.Println("server listening...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

func testInsertDB() {
	db.InsertUser(models.User{Username: "ebosch", Password: "password"})
	db.InsertUser(models.User{Username: "caspark", Password: "password"})
	db.InsertUser(models.User{Username: "ryland", Password: "password"})
	db.InsertUser(models.User{Username: "juju", Password: "password"})

	id := db.FindUserByUsername("caspark")

	db.DeleteUserAndNotes(id.ID)
}

func testDB() {
	users := db.FindAllUsers()

	for index, user := range *users {
		fmt.Printf("User %d: %+v\n", index+1, user)
	}

}
