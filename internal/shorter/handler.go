package shorter

import (
	"github.com/gin-gonic/gin"
)

type ShorterHandler struct {
	ss *ShorterService
}

func NewShorterHandler(ss *ShorterService) *ShorterHandler {
	return &ShorterHandler{ss}
}

func (sh *ShorterHandler) generateShortUrlHandler(ctx *gin.Context) {
	url := ctx.PostForm("url")
	shortenUrl := sh.ss.GenerateShortUrl(ctx.Request.Context(), "http", ctx.Request.Host, url)

	ctx.JSON(200, gin.H{
		"url":        url,
		"shortenUrl": shortenUrl,
	})
}

func (sh *ShorterHandler) getUrlHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (sh *ShorterHandler) deleteShortUrlHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
