package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/muhuchah/traceroute/handler"
	"github.com/muhuchah/traceroute/helper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/traceroute/:host", handler.Trace)

	router.Run(helper.GetEnv("SERVER_HOST", "localhost:8080"))
}
