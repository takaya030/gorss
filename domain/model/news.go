package model

type News struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	TimeStamp int64  `json:"timestamp"`
	Url       string `json:"url"`
}
