package src

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var clients []*websocket.Conn

func handleLobby(c *websocket.Conn) {
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
		players = append(players, responseSkeleton{Data: fmt.Sprint("player ", i+1)})
	}

	playerStruct := struct {
		Players []responseSkeleton `json:"players"`
	}{
		Players: players,
	}

	for _, client := range clients {
		client.WriteJSON(playerStruct)
	}

}
