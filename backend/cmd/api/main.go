package main

import (
	"github.com/Lydoww/react-native-go-fastlane/handlers"
	"github.com/Lydoww/react-native-go-fastlane/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Fastlane",
		ServerHeader: "Fiber",
	})

	// Repository

	eventRepository := repositories.NewEventRepository(nil)

	// Router

	server := app.Group("/api")

	// Handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}