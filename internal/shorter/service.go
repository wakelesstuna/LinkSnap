package shorter

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/wakelesstuna/link-snap/internal/db"
)

type ShorterService struct {
	db *db.Database
}

func NewShorterService(db *db.Database) *ShorterService {
	return &ShorterService{
		db: db,
	}
}

func (ss *ShorterService) GenerateShortUrl(ctx context.Context, scheme, host, url string) ShortenUrl {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hash := hex.EncodeToString(hasher.Sum(nil))
	shortUrl := fmt.Sprintf("%s://%s/shorten/%s", scheme, host, hash[:8])
	resp, err := ss.db.AddShortenUrl(shortUrl, url)
	if err != nil {
		panic(err)
	}
	return ShortenUrl{
		Id:          int(resp.Id),
		ShortenUrl:  resp.ShortenUrl,
		OriginalUrl: resp.OriginalUrl,
		CreatedAt:   resp.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (ss *ShorterService) GetUrl(shortUrl string) (*ShortenUrl, error) {
	su, err := ss.db.GetShortenUrl(shortUrl)
	if err != nil {
		return nil, err
	}
	return &ShortenUrl{
		Id:          int(su.Id),
		ShortenUrl:  su.ShortenUrl,
		OriginalUrl: su.OriginalUrl,
		CreatedAt:   su.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (ss *ShorterService) DeleteShortUrl(shortUrl string) string {
	return "https://wakelesstuna.github.io/link-snap/shorten/" + shortUrl
}
