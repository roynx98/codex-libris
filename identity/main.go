package main

import (
	"identity/lib/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, config.Default.ServiceName+" is running")
	})

	r.Run(":" + config.Default.Port)
}
