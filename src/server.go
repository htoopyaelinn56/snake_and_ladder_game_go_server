package src

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Current int `json:"current"`
	DiceNum int `json:"dice_num"`
}

func dice(c *gin.Context) {

	conn, err := GetWebsocket(c)
	if err != nil {
		return
	}

	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()

		var formattedMsg = string(message)
		var incomingJson *Response

		decodeErr := json.Unmarshal(message, &incomingJson)
		if decodeErr != nil {
			log.Print("json decode failed")
			return
		}
		fmt.Printf("formatted : pos %v dice %v\n", incomingJson.Current, incomingJson.DiceNum)

		fmt.Print(formattedMsg, "\n")
		responseJson := &Response{
			Current: incomingJson.Current,
			DiceNum: incomingJson.DiceNum,
		}

		writeErr := conn.WriteJSON(*responseJson)

		if writeErr != nil {
			log.Print("write failed: ", err)
			return
		}
	}
}

func RunApp() {
	r := gin.Default()
	r.GET("/dice", dice)
	r.Run(":8080")
}
