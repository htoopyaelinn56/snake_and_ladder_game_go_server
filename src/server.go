package src

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients []*websocket.Conn

func handWs(c *websocket.Conn) {
	clients = append(clients, c)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		fmt.Printf("received: %s\n", message)

		go handleMessages(message)
	}
}

func handleMessages(data []byte) {
	for _, client := range clients {
		err := client.WriteMessage(1, data)
		if err != nil {
			fmt.Println("write error:", err)
			break
		}
	}
}

func RunApp() {
	app := fiber.New()

	// WebSocket route
	app.Get("/ws", websocket.New(handWs))

	app.Listen(":3000")
}
