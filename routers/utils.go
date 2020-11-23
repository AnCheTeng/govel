package routers

import "github.com/gin-gonic/gin"

// Ping - Pong
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
