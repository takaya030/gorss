package repository

import "github.com/mmcdole/gofeed"

type RssRepository interface {
	ParseURL(url string) error
	GetTitle() (string, error)
	GetRssType() (string, error)
	GetFeedItems() ([]*gofeed.Item, error)
}
