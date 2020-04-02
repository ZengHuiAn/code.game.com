package server

import (
	"github.com/micro/go-micro/v2"
	"log"
)

type MicroServer struct {
	name        string
	gatewayAddr string

	svr micro.Service
}

func (this *MicroServer) InitMicroServer() {
	if err := this.GetService().Run(); err != nil {
		log.Println(err)
	}
}

func (this *MicroServer) GetService() micro.Service {
	return this.svr
}

func (this *MicroServer) RunMicroServer() {

}

func (this *MicroServer) GetServerName() string {
	return this.name
}

func (this *MicroServer) GetServerAddr() string {
	return "server." + this.GetServerName()
}

func NEWMicroServer(name string) *MicroServer {
	ser := &MicroServer{
		name: name,
	}
	ser.InitMicroServer()
	return ser
}