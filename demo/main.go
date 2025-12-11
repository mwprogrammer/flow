package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("failed to load .env file!")
	}

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Sample Flow app")
	})

	server.GET("/demo", func(c *gin.Context) {

		token := c.Query("hub.verify_token")
		challenge := c.Query("hub.challenge")

		if token == "sample_flow_app" {
			c.String(http.StatusOK, challenge)
		}

	})

	server.POST("/demo", func(c *gin.Context) {

		data, err := io.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read body"})
			return
		}

		io.NopCloser(bytes.NewBuffer(data))

		data_string := string(data)

		fmt.Println("Raw Body Received:", data_string)

		c.JSON(http.StatusOK, gin.H{
			"message":         "Raw body processed",
			"raw_data_string": data_string,
		})

		fmt.Println()

	})

	server.Run(":9000")
}
