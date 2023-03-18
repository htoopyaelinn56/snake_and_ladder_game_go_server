package src

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var _players []*websocket.Conn = make([]*websocket.Conn, 4)
var currentTurn = -1
var totalPlayers = -1
var _canStart = false

func _selfAssign(position int, c *websocket.Conn) {
	//assign clients' turn
	_players[position] = c
	fmt.Println(_players)
	fmt.Println("len ", totalPlayers)

	if _countNotNil() == totalPlayers {
		fmt.Println("can start")
		_canStart = true
	}

	_handleDices(-1, nil, _canStart)
}

func _countNotNil() int {
	count := 0
	for _, p := range _players {
		if p != nil {
			count += 1
		}
	}
	return count
}

func handleGameWs(c *websocket.Conn) {

	for {
		_, message, err := c.ReadMessage()
		var inputData *payload
		fmt.Println(string(message))
		decodeErr := json.Unmarshal(message, &inputData)

		if decodeErr != nil {
			fmt.Println("decode error")
			break
		}
		if err != nil {
			fmt.Println("read error:", err)
			removeClient(&_players, c)
			break
		}
		fmt.Printf("received: %s\n", message)
		if inputData.Dice == -1 {
			go _selfAssign(inputData.AssignTurn, c)
		} else {
			go _handleDices(inputData.Dice, c, _canStart)
		}

	}
}

func _handleDices(diceNum int, sender *websocket.Conn, canStart bool) {

	if canStart {
		fmt.Println("send, current turn ", currentTurn)
		currentTurn += 1
		if currentTurn >= totalPlayers { //reset turn after round
			currentTurn = 0
		}
	}
	for _, client := range _players {
		if client != sender && client != nil {
			err := client.WriteJSON(response{
				Current:  currentTurn,
				Dice:     diceNum,
				CanStart: canStart,
			})
			if err != nil {
				fmt.Println("write error:", err)
				break
			}
		}
	}

}
