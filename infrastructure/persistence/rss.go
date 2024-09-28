package persistence

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/takaya030/gorss/domain/define"
	"github.com/takaya030/gorss/domain/model"
	"github.com/takaya030/gorss/domain/repository"
)

type rssClient struct {
	feed *gofeed.Feed
}

func NewRssClient(rss_url string) (repository.RssRepository, error) {
	fd, err := initFeed(rss_url)
	if err != nil {
		return nil, err
	}
	return &rssClient{feed: fd}, nil
}

func initFeed(rss_url string) (*gofeed.Feed, error) {
	feed, err := gofeed.NewParser().ParseURL(rss_url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	return feed, nil
}

func (r *rssClient) GetFeedItems() ([]*model.News, error) {
	if r.feed == nil {
		fmt.Fprintln(os.Stderr, define.ErrNilFeed)
		return nil, define.ErrNilFeed
	}

	var newsItems []*model.News
	for _, item := range r.feed.Items {
		if item == nil {
			break
		}
		n := new(model.News)
		n.Title = item.Title
		n.TimeStamp = item.PublishedParsed.Unix()
		n.Url = item.Link
		newsItems = append(newsItems, n)
	}
	return newsItems, nil
}
