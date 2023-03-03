package src

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

var currentTurn int = 0
var yourTurn int = 0
var currentPlayerCount = 0

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
			removeClient(c)
			break
		}
		fmt.Printf("received: %s\n", message)
		go handleDices(inputData.Dice, c)
	}
}

func handleDices(diceNum int, sender *websocket.Conn) {
	for _, client := range clients {
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
