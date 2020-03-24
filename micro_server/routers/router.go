package routers

import (
	"code.game.com/micro_server/routers/api"
	"github.com/gin-gonic/gin"
)

var ginDefalt *gin.Engine;

func GetGinInstance() *gin.Engine {
	if ginDefalt == nil {
		ginDefalt = gin.Default()
	}
	return ginDefalt
}

func InitRouter() *gin.Engine {
	instance := GetGinInstance()

	instance.GET("/tools/api/pullServerList", api.PullServerList)
	return instance
}
