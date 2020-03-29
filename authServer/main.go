package main

import (
	"code.game.com/authServer/routers/api"
	"log"
)


//func (g *Greeter) Hello(context context.Context, req *greeter.HelloRequest, rsp *greeter.HelloRespone) error {
//	rsp.Pid = 1000
//	return nil
//}

func main() {

	service:= api.GetInstance()
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
