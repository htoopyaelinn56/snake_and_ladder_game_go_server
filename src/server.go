package src

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/lobby", websocket.New(handleLobby))
	app.Get("/dice", websocket.New(handleGameWs))
}

func RunApp() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":3000")
}
