package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Printf("Cannot load env file: %v", err)
		}

		msg := os.Getenv("SAMPLE_MESSAGE")

		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	})
	engine.Run(":8080")
}
