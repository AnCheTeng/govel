package main

import (
	"os"

	"github.com/AnCheTeng/govel/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":" + os.Getenv("PORT"))
}
