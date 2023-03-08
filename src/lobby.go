package src

import (
	"fmt"
	"time"

	"github.com/gofiber/websocket/v2"
)

var clients []*websocket.Conn
var _host *websocket.Conn
var _originalTime = 10
var _tick = _originalTime
var _startCheck = false

func _setHost() {
	if len(clients) > 0 {
		_host = clients[0]
	}
}

func handleLobby(c *websocket.Conn) {
	clients = append(clients, c)

	fmt.Println("connected to lobby")
	go _handlePlayersInLobby(clients) // Notify all clients of new connection
	_startCheck = false

	_setHost()
	for {
		_, message, err := c.ReadMessage()
		if string(message) == "start" {
			go _handlePlayersInLobby(clients) // notify start button
			_startCheck = true
		}
		fmt.Println(string(message))
		if err != nil {
			fmt.Println("lobby err", err)
			removeClient(c)
			_setHost()
			go _handlePlayersInLobby(clients) // Notify all clients of disconnection
			_startCheck = false
			break
		}
	}
}

func _handlePlayersInLobby(clients []*websocket.Conn) {
	var players []playerLobbyResponse
	for i, client := range clients {
		players = append(players, playerLobbyResponse{
			Player: fmt.Sprint("Player ", i+1),
			Host:   _host == client,
		})
	}

	fmt.Println("just check", _startCheck)
	if _startCheck {
		fmt.Println("it is true")
		for i := 0; i <= 10; i++ {
			if !_startCheck {
				fmt.Println("timer should not tick")
				break
			}
			for _, client := range clients {
				client.WriteJSON(_getResponse(players, client, _startCheck, _tick))
			}
			time.Sleep(time.Second)
			_tick -= 1
			fmt.Println("this state ticking", _startCheck)
		}
	} else {
		for _, client := range clients {
			if _tick != 0 {
				client.WriteJSON(_getResponse(players, client, _startCheck, -1))
				_tick = _originalTime
			}
		}

	}
}

func _getResponse(players []playerLobbyResponse, client *websocket.Conn, _startCheck bool, tick int) lobbyResponse {
	return lobbyResponse{
		Data: players,
		Host: _host == client,
		You:  indexOf(clients, client),
		Start: func() bool {
			if _startCheck {
				return true
			} else {
				return false
			}
		}(),
		Timer: tick,
	}
}
