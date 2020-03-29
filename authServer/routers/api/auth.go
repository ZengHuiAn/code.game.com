package api

import (
	"bytes"
	"code.game.com/proto/amf"
	proto "code.game.com/server_proto/proto"
	"context"
	"log"
)

type Auth struct {
	tokens []string
}

func (g *Auth) AuthClient(ctx context.Context, req *proto.AuthRequest, resp *proto.AuthRespone) error {
	var send = [2]interface{}{}
	send[0] = 0
	send[1] = "123456789"
	var buffer bytes.Buffer

	amf.Encode(&buffer,send)
	resp.AuthToken = buffer.Bytes()
	resp.ErrorCode = 200
	return nil
}

func init()  {
	s :=GetInstance().Server()
	proto.RegisterAuthHandler(s, new(Auth))
	log.Println("register auth service")
}