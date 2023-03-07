package src

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var clients []*websocket.Conn
var host *websocket.Conn

func _setHost() {
	if len(clients) > 0 {
		host = clients[0]
	}
}

func handleLobby(c *websocket.Conn) {
	clients = append(clients, c)

	fmt.Println("connected to lobby")
	go handlePlayersInLobby(clients, false) // Notify all clients of new connection

	_setHost()
	for {
		_, message, err := c.ReadMessage()
		if string(message) == "start" {
			go handlePlayersInLobby(clients, true) // notify start button

		}
		fmt.Println(string(message))
		if err != nil {
			fmt.Println("lobby err", err)
			removeClient(c)
			_setHost()

			go handlePlayersInLobby(clients, false) // Notify all clients of disconnection

			break
		}
	}
}

func handlePlayersInLobby(clients []*websocket.Conn, startCheck bool) {
	var players []playerLobbyResponse
	for i, client := range clients {
		players = append(players, playerLobbyResponse{
			Player: fmt.Sprint("Player ", i+1),
			Host:   host == client,
		})
	}

	for _, client := range clients {
		client.WriteJSON(struct {
			Data  []playerLobbyResponse `json:"data"`
			Host  bool                  `json:"host"`
			You   int                   `json:"you"`
			Start bool                  `json:"start"` //flag to tell the clients to start timer
		}{
			Data: players,
			Host: host == client,
			You:  indexOf(clients, client),
			Start: func() bool {
				if startCheck {
					return true
				} else {
					return false
				}
			}(),
		})
	}
}
