package shorter

import (
	"fmt"

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

	ctx.JSON(200, shortenUrl)
}

func (sh *ShorterHandler) getUrlHandler(ctx *gin.Context) {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	url := fmt.Sprintf("%s://%s/%s", scheme, ctx.Request.Host, ctx.Request.URL.Path)

	fmt.Println(url)
	fmt.Println(ctx.FullPath())
	ShortenUrl, err := sh.ss.GetUrl(ctx.FullPath())
	if err != nil {
		ctx.AbortWithError(404, err)
		return
	}
	ctx.JSON(200, ShortenUrl)
}

func (sh *ShorterHandler) deleteShortUrlHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (sh *ShorterHandler) redirectUrlHandler(ctx *gin.Context) {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	url := fmt.Sprintf("%s://%s%s", scheme, ctx.Request.Host, ctx.Request.URL.Path)
	fmt.Println("URL: " + url)
	ShortenUrl, err := sh.ss.GetUrl(url)
	if err != nil {
		ctx.AbortWithError(404, err)
		return
	}
	ctx.Redirect(302, ShortenUrl.OriginalUrl)
}
