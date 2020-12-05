package main

import (
	"os"

	"github.com/AnCheTeng/govel/routers"
	"github.com/gin-gonic/contrib/static"
)

func main() {
	r := routers.InitRouter()
	r.Use(static.Serve("/", static.LocalFile("./build", true)))

	r.Run(":" + os.Getenv("PORT"))
}
