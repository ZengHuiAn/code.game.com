package main

import "code.game.com/micro_server/routers"

func main() {
	//routers.GetGinInstance().
	r := routers.InitRouter()
	_ = r.Run(":10001")
}
