package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/config"
	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
)

func addExampleRoutes(rg *gin.RouterGroup) {
	controller, _ := config.InitPingController()

	rg.GET("/ping", func(c *gin.Context) {
		getPingHandler(c, controller)
	})

	rg.GET("/helloworld", getHelloWorldHandler)
}

// @Summary Ping Server
// @Schemes
// @Description 確認服務正常
// @Tags Example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @response default {object} respdto.ApiErrorResp "error response"
// @Router /ping [get]
func getPingHandler(c *gin.Context, controller *controllers.PingController) {
	respMsg, err := controller.Ping()

	if err != nil {
		c.Error(err)
		return
	}

	data := new(respdto.PingResp)
	data.Msg = respMsg
	c.JSON(http.StatusOK, data)
}

// @Summary HelloWorld example
// @Schemes
// @Description do ping
// @Tags Example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func getHelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
}
