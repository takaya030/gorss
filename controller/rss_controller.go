package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

func ShowHello(c *gin.Context) {

	msg := os.Getenv("SAMPLE_MESSAGE")

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func ParseFeed(c *gin.Context) {
	rss_url := os.Getenv("RSS_URL")
	feed, err := gofeed.NewParser().ParseURL(rss_url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(feed.Title)
	fmt.Println(feed.FeedType, feed.FeedVersion)
	for _, item := range feed.Items {
		if item == nil {
			break
		}
		fmt.Println(item.Title)
		fmt.Println("\t->", item.Link)
		fmt.Println("\t->", item.PublishedParsed.Format(time.RFC3339))
		fmt.Println("\t->", item.PublishedParsed.Unix())
	}
}
