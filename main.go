package main

import "github.com/AnCheTeng/govel/routers"

func main() {
	r := routers.InitRouter()
	r.Run(":7788")
}
