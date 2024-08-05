package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"traceroute/handler"
	"traceroute/helper"
)

func main() {
	// Initialize Zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		// Fallback to a simple logger if unable to create the production logger
		logger = zap.NewExample()
	}
	defer logger.Sync() // Flushes buffer

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("Error loading .env file", zap.Error(err))
	} else {
		logger.Info("Successfully loaded .env file")
	}

	router := gin.Default()
	router.GET("/traceroute/:host", handler.Trace)

	// Start the server
	serverHost := helper.GetEnv("SERVER_HOST", ":8080")
	logger.Info("Starting server", zap.String("host", serverHost))
	if err := router.Run(serverHost); err != nil {
		logger.Fatal("Unable to start server", zap.Error(err))
	}
}
