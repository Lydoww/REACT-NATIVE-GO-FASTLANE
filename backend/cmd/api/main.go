package main

import (
	"fmt"

	"github.com/Lydoww/react-native-go-fastlane/config"
	"github.com/Lydoww/react-native-go-fastlane/db"
	"github.com/Lydoww/react-native-go-fastlane/handlers"
	"github.com/Lydoww/react-native-go-fastlane/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Fastlane",
		ServerHeader: "Fiber",
	})

	// Repository

	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)

	// Router

	server := app.Group("/api")

	// Handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	fmt.Println("🚀 Fastlane server started on port 3000")
	app.Listen(":" + envConfig.ServerPort)

}