package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"traceroute/handler"
	"traceroute/helper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/traceroute/:host", handler.Trace)

	router.Run(helper.GetEnv("SERVER_HOST", ":8080"))
}
