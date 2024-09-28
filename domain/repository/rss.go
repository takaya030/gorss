package repository

import (
	"github.com/takaya030/gorss/domain/model"
)

type RssRepository interface {
	GetFeedItems() ([]*model.News, error)
}
