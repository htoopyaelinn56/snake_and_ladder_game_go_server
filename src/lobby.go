package src

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var clients []*websocket.Conn

func handleLobby(c *websocket.Conn) {
	c.WriteJSON(responseSkeleton{Data: "Welcome to lobby"})
	clients = append(clients, c)
	fmt.Println("connected to lobby")
	go handlePlayersInLobby(clients) // Notify all clients of new connection

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			fmt.Println("lobby err", err)
			removeClient(c)
			go handlePlayersInLobby(clients) // Notify all clients of disconnection
			break
		}
		go handlePlayersInLobby(clients)
	}
}

func handlePlayersInLobby(clients []*websocket.Conn) {
	var players []responseSkeleton
	for i := range clients {
		players = append(players, responseSkeleton{Data: fmt.Sprint("player ", i)})
	}

	for _, client := range clients {
		client.WriteJSON(players)
	}

}
