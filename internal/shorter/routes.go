package shorter

import (
	"github.com/gin-gonic/gin"
	"github.com/wakelesstuna/link-snap/internal/db"
)

func InitRoutes(gin *gin.Engine) {
	db := db.NewDb()
	ss := NewShorterService(db)
	sh := NewShorterHandler(ss)

	gin.GET("/shorten/:hex", sh.redirectUrlHandler)

	api := gin.Group("/api")
	{
		api.GET("/shorten/:hex", sh.getUrlHandler)
		api.POST("/shorten", sh.generateShortUrlHandler)
		api.DELETE("/shorten/:shortenUrl", sh.deleteShortUrlHandler)
	}
}
