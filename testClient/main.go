package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	//cli := micro.NewService(micro.Name("greeter.client"))
	//cli.Init()
	//
	//g := proto.NewGreeterService("greeter", cli.Client())
	//
	//resp, err := g.Hello(context.TODO(), &proto.HelloRequest{Name: "测试"})
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(resp)
	conn, err  := net.Dial("tcp", "127.0.0.1:8091")

	fmt.Println("err = ", err)
	if err != nil {
		return
	}
	var bif = bytes.Buffer{}
	bif.WriteString("测试")
	conn.Write(bif.Bytes())
	defer conn.Close()
}
