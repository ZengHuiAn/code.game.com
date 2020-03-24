package api

import (
	"code.game.com/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

var serverList []proto.ServerInfo = []proto.ServerInfo{
	{ID: 1, Gateway: "gateway11", Name: "测试"},
}

func PullServerList(c *gin.Context) {
	c.JSON(http.StatusOK, serverList)
}
