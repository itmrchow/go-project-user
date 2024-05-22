package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/resp_dto"
)

func addPingRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", getPingMsg)
}

func getPingMsg(c *gin.Context) {
	data := new(resp_dto.PingResp)
	data.Msg = "pong!"
	c.JSON(http.StatusOK, data)
}
