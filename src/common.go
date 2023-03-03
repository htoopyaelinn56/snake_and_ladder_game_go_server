package src

import (
	"github.com/gofiber/websocket/v2"
)

func removeClient(ws *websocket.Conn) {
	// Remove the client from the clients slice
	for i, client := range clients {
		if client == ws {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}
