package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/webdevcaptain/scwm-sensor-api/queues"
	"github.com/webdevcaptain/scwm-sensor-api/routes"
)

func main() {
	// Load the .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Intialize rabbitmq
	cleanupQueue := queues.InitRabbitMQ()
	defer cleanupQueue()

	// Initialize Gin app
	server := gin.Default()

	// Setup routes
	routes.Register(server)

	// Setup PORT and start server
	if err := server.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start sensor-api server %s", err)
	}
}
