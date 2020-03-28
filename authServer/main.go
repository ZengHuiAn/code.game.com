package main

import "code.game.com/authServer/routers"

func main() {
	r := routers.InitRouter()
	r.Run(":9090")
}
