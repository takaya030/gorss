package repository

import (
	"context"

	"github.com/takaya030/gorss/domain/model"
)

type FireStoreRepository interface {
	GetNews(ctx context.Context) ([]*model.News, error)
	CreateNews(ctx context.Context, news *model.News) error
	DeleteNews(ctx context.Context, id string) error
}
