package shorter

import "github.com/gin-gonic/gin"

func InitRoutes(gin *gin.Engine) {
	ss := NewShorterService()
	sh := NewShorterHandler(ss)

	gin.GET("/:shortenUrl", sh.getUrlHandler)
	gin.POST("/shorten", sh.generateShortUrlHandler)
	gin.DELETE("/:shortenUrl", sh.deleteShortUrlHandler)
}
