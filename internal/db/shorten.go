package db

import (
	"time"
)

type ShortenUrl struct {
	Id          int64     `json:"id"`
	ShortenUrl  string    `json:"shorten_url"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}
