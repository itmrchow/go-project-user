package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/interfaces/api/controllers"

)

func addPingRoutes(rg *gin.RouterGroup) {
	controller := controllers.NewPingController()

	rg.GET("/ping", func(c *gin.Context) {
		respMsg := controller.Ping()
		data := new(respdto.PingResp)
		data.Msg = respMsg
		c.JSON(http.StatusOK, data)
	})
}
