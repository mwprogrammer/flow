package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("failed to load .env file!")
	}

	server := gin.Default()

	server.GET("/demo")
	server.POST("/demo")

	server.Run()
}
