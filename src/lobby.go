package src

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var clients []*websocket.Conn
var host *websocket.Conn

func handleLobby(c *websocket.Conn) {
	clients = append(clients, c)

	fmt.Println("connected to lobby")
	go handlePlayersInLobby(clients) // Notify all clients of new connection
	if len(clients) > 0 {
		host = clients[0]
	}
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			fmt.Println("lobby err", err)
			removeClient(c)
			if len(clients) > 0 {
				host = clients[0]
			}
			go handlePlayersInLobby(clients) // Notify all clients of disconnection
			break
		}
	}
}

func handlePlayersInLobby(clients []*websocket.Conn) {
	var players []playerLobbyResponse
	for i, client := range clients {
		players = append(players, playerLobbyResponse{
			Player: fmt.Sprint("Player ", i+1),
			Host:   host == client,
		})
	}

	for _, client := range clients {
		client.WriteJSON(struct {
			Data []playerLobbyResponse `json:"data"`
			Host bool                  `json:"host"`
			You  int                   `json:"you"`
		}{
			Data: players,
			Host: host == client,
			You:  indexOf(clients, client),
		})
	}
}
