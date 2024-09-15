package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ShowHello(c *gin.Context) {

	msg := os.Getenv("SAMPLE_MESSAGE")

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
