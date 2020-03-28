package main

import (
	greeter "code.game.com/server_proto/proto"
	"github.com/micro/go-micro/v2"
	"golang.org/x/net/context"
	"log"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, resp *greeter.HelloRespone) error {
	resp.Pid = 200
	return nil
}

//func (g *Greeter) Hello(context context.Context, req *greeter.HelloRequest, rsp *greeter.HelloRespone) error {
//	rsp.Pid = 1000
//	return nil
//}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("laste"),
	)
	service.Init()
	s := service.Server()
	greeter.RegisterGreeterHandler(s, new(Greeter))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
