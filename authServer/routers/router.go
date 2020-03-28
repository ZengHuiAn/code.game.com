package routers

import (
	"code.game.com/authServer/routers/api"
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

	instance.GET("/api/auth", api.AuthToServer)
	return instance
}
