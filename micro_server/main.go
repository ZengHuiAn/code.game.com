package main

import "code.game.com/micro_server/routers"

func main() {
	//routers.GetGinInstance().
	r := routers.InitRouter()
	r.Run(":8081")
}
