package main

import (
	// WAFのGinをインポート
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})
	router.Run(":8081")
}
