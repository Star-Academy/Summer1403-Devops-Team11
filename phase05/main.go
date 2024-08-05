package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"traceroute/handler"
	"traceroute/helper"
)

func main() {
	// Initialize logger
	logger := helper.InitLogger()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("Warning: failed to load .env file")
	} else {
		logger.Info("Successfully loaded .env file")
	}

	router := gin.Default()
	router.GET("/traceroute/:host", handler.Trace)

	// Start the server
	serverHost := helper.GetEnv("SERVER_HOST", ":8080")
	logger.Info("Starting server")
	if err := router.Run(serverHost); err != nil {
		logger.Error("Unable to start server")
	}
}
