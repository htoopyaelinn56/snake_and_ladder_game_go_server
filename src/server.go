package src

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type response struct {
	Current int `json:"current"`
	Dice    int `json:"dice_num"`
}

var clients []*websocket.Conn

func handleWebSocket(c *gin.Context) {
	// Upgrade the connection to a websocket connection
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Error upgrading connection to websocket: ", err)
	}

	// Send a welcome message to the client
	err = ws.WriteMessage(websocket.TextMessage, []byte("Welcome!"))
	if err != nil {
		log.Println("Error sending welcome message to client: ", err)
	}

	// Add the new client to the clients slice
	clients = append(clients, ws)

	// Start listening for incoming messages from the client
	go dice(ws)
}

func dice(ws *websocket.Conn) {
	for {
		// Read a message from the client
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client: ", err)
			removeClient(ws)
			break
		}
		var incomingJson *response
		decodeErr := json.Unmarshal(message, &incomingJson)
		// Send the message to all connected clients
		if decodeErr != nil {
			log.Println("json decode error")
		}
		sendToAllClients(incomingJson, ws)
	}
}

func sendToAllClients(message *response, ws *websocket.Conn) {
	for _, client := range clients {
		if ws != client {
			err := client.WriteJSON(message)

			if err != nil {
				log.Println("Error sending message to client: ", err)
				removeClient(client)
			}
		}

	}
}

func removeClient(ws *websocket.Conn) {
	// Remove the client from the clients slice
	for i, client := range clients {
		if client == ws {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func RunApp() {
	// Configure routes
	r := gin.Default()
	r.GET("/dice", handleWebSocket)

	// Start the server
	log.Println("Starting server...")
	err := r.Run()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
