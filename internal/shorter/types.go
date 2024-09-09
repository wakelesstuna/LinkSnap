package shorter

type ShortenUrl struct {
	Id          int    `json:"id"`
	ShortenUrl  string `json:"shorten_url"`
	OriginalUrl string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
}
