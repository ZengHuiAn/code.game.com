package main

import (
	proto "code.game.com/server_proto/proto"
	"context"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	cli := micro.NewService(micro.Name("greeter.client"))
	cli.Init()

	g := proto.NewGreeterService("greeter", cli.Client())

	resp, err := g.Hello(context.TODO(), &proto.HelloRequest{Name: "测试"})

	if err != nil {
		log.Println(err)
	}

	log.Println(resp)
}
