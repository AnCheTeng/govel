package routers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	APIv1 := r.Group("/api/v1")
	{
		APIv1.GET("/ping", Ping)
		APIv1.GET("/lastUpdate", GetLastUpdate)
	}

	r.Use(static.Serve("/", static.LocalFile("./build", true)))

	return r
}
