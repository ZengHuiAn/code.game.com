package api

import (
	proto "code.game.com/server_proto/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

type APIClient struct {
	svr micro.Service
	client client.Client
	AuthService proto.AuthService
}

func NewAPIClient(name string) *APIClient {
	cli := &APIClient{}
	cli.svr = micro.NewService(micro.Name("server." + name))
	cli.svr.Init()
	cli.client = cli.svr.Client()
	return cli
}

func (this *APIClient) Init() {
	this.AuthService = proto.NewAuthService("server.auth", this.svr.Client())
}
