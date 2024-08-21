package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Gin
	server := gin.Default()

	// [TODO]: Setup routes

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello from Sensor API.",
		})
	})

	// Setup PORT and start server
	if err := server.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start sensor-api server %s", err)
	}
}
