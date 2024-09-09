package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/wakelesstuna/link-snap/internal/shorter"
)

func InitWeb() {
	gin := gin.Default()

	shorter.InitRoutes(gin)

	err := gin.Run(":9000")
	if err != nil {
		panic(err)
	}
}
