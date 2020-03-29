package api
import (
	"github.com/micro/go-micro/v2"
)
var authService micro.Service

func GetInstance() micro.Service {
	if authService != nil {
		return authService
	}
	authService := micro.NewService(
		micro.Name("server.auth"),
		micro.Version("laste"),
	)
	authService.Init()

	return authService
}