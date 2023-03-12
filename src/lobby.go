package src

import (
	"fmt"
	"time"

	"github.com/gofiber/websocket/v2"
)

var _clients []*websocket.Conn
var _host *websocket.Conn
var _originalTime = 5
var _tick = _originalTime
var _startCheck = false

func _setHost() {
	if len(_clients) > 0 {
		_host = _clients[0]
	}
}

func handleLobby(c *websocket.Conn) {
	totalPlayers = -1
	_clients = append(_clients, c)
	fmt.Println("len of client ", len(_clients))
	fmt.Println("connected to lobby")
	go _handlePlayersInLobby(_clients) // Notify all _clients of new connection
	_startCheck = false

	_setHost()
	for {
		_, message, err := c.ReadMessage()
		if string(message) == "start" {
			go _handlePlayersInLobby(_clients) // notify start button
			totalPlayers = len(_clients)
			_startCheck = true
		}
		fmt.Println(string(message))
		if err != nil {
			fmt.Println("lobby err", err)
			removeClient(&_clients, c)
			_setHost()
			go _handlePlayersInLobby(_clients) // Notify all _clients of disconnection
			_startCheck = false
			break
		}
	}
}

func _handlePlayersInLobby(_clients []*websocket.Conn) {
	var players []playerLobbyResponse
	for i, client := range _clients {
		players = append(players, playerLobbyResponse{
			Player: fmt.Sprint("Player ", i+1),
			Host:   _host == client,
		})
	}

	if _startCheck {
		for i := 0; i <= 10; i++ {
			if !_startCheck {
				break
			}
			for _, client := range _clients {
				client.WriteJSON(_getResponse(players, client, _startCheck, _tick))
			}
			time.Sleep(time.Second)
			_tick -= 1
		}
	} else {
		for _, client := range _clients {
			if _tick != 0 { //dont notify to frontend about client leave when the game starts
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
		You:  indexOf(_clients, client),
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
