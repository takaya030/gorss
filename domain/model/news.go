package model

type News struct {
	Id        string `json:"id"`
	TimeStamp int64  `json:"timestamp"`
	Url       string `json:"url"`
}
