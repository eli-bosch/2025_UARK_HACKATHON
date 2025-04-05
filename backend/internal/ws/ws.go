package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// sample chat message object
type ChatMessage struct {
	Message string			`json:"message"`
}

type AuthPayload struct {
	UserID string	`json:"userID"`
}

type Client struct {
	conn 	*websocket.Conn
	userID 	string
	roomID	string
}

// connections = keys
// client pointers = values
// Stores clients by their websocket connections
var clients = make(map[*websocket.Conn]*Client)
var rooms = make(map[string][]*Client) // roomID -> [userID1, userID2]


var upgrader = websocket.Upgrader {
	// allows all connections
	CheckOrigin: func(r *http.Request) bool { return true },
}

func RegisterWSRoutes(r *mux.Router) {
	r.HandleFunc("/ws", WsHandler).Methods("GET") // websockets MUST USE GET
	fmt.Println("WebSocket server is running...")
}

// set up websocket conn and handle clientside message & data
func WsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade HTTP request to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Websocket Connection successfully established.")

	// authenticate user via user ID perhaps
	// TODO: use database to search this i think
	var auth AuthPayload
	err = conn.ReadJSON(&auth)
	if err != nil {
		log.Println("Error reading user ID:", err)
		conn.Close()
		return
	}

	userID := auth.UserID

	// hardcoded for now, have user join a new room
	roomID := "room1" 
	client := &Client{conn: conn, userID: userID, roomID: roomID} // insert new ws conn into clients map

	// add client to clients map
	clients[conn] = client
	
	// join room and add client to the room
	joinRoom(client)

	defer func() {
		conn.Close()
		delete(clients, conn)
		leaveRoom(client)
	}()


	// handle incoming messages
	for {
		var msg ChatMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		// broadcast message to the room
		broadcastMessage(client.roomID, msg)
	}
}

func joinRoom(client *Client) {
	// if room doesn't exist, create it
	if _, exists := rooms[client.roomID]; !exists {
		rooms[client.roomID] = []*Client{}
	}

	// add client to the room
	rooms[client.roomID] = append(rooms[client.roomID], client)
	log.Printf("%s joined room %s", client.userID, client.roomID)
}

// exit client from room

func leaveRoom(client *Client) {
	room := rooms[client.roomID]
	for i, c := range room {
		if c.conn == client.conn {
			// remove client from room
			rooms[client.roomID] = append(room[:i], room[i+1:]...)
			log.Printf("%s left room %s", client.userID, client.roomID)
			break
		}
	}
}

// broadcast a message to all users in the same room (excluding sender)
func broadcastMessage(roomID string, message ChatMessage) {
	room, exists := rooms[roomID]
	if !exists {
		log.Printf("Room %s does not exist", roomID)
		return
	}

	log.Printf("Broadcasting to room %s: %s", roomID, message.Message)

	for _, client := range room {
		// send msg to each client in room
		err := client.conn.WriteJSON(message)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}