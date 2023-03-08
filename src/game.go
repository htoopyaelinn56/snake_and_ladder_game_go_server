package src

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var _players []*websocket.Conn = make([]*websocket.Conn, 4, 4)
var currentTurn int = 0
var yourTurn int = 0
var currentPlayerCount = 0

func _selfAssign(position int, c *websocket.Conn) {
	//assign clients' turn
	_players[position] = c

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

		if inputData.Dice == -1 {
			_selfAssign(inputData.AssignTurn, c)
		} else {
			fmt.Printf("received: %s\n", message)
			go _handleDices(inputData.Dice, c)
		}

	}
}

func _handleDices(diceNum int, sender *websocket.Conn) {

	for _, client := range _players {
		if client != sender {
			err := client.WriteJSON(response{
				Current: currentTurn,
				Dice:    diceNum,
			})
			if err != nil {
				fmt.Println("write error:", err)
				break
			}
		}
	}

}
