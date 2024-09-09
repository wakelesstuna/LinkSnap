package shorter

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ShorterService struct {
}

func NewShorterService() *ShorterService {
	return &ShorterService{}
}

func (ss *ShorterService) GenerateShortUrl(ctx context.Context, scheme, host, url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("%s://%s/shorten/%s", scheme, host, hash[:8])
}

func (ss *ShorterService) GetUrl(shortUrl string) string {
	return "https://wakelesstuna.github.io/link-snap/shorten/" + shortUrl
}

func (ss *ShorterService) DeleteShortUrl(shortUrl string) string {
	return "https://wakelesstuna.github.io/link-snap/shorten/" + shortUrl
}
