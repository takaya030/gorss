package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", ShowHello)
	r.GET("/rss", ParseFeed)
	r.GET("/feed", FetchFeed)
	return r
}
