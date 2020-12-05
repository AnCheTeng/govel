package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	APIv1 := r.Group("/api/v1")
	{
		APIv1.GET("/ping", Ping)
		APIv1.GET("/lastUpdate", GetLastUpdate)
	}
	return r
}
