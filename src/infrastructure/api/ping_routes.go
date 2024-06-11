package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func addExampleRoutes(rg *gin.RouterGroup) {
	controller := controllers.NewPingController()

	rg.GET("/ping", func(c *gin.Context) {
		respMsg := controller.Ping()
		data := new(respdto.PingResp)
		data.Msg = respMsg
		c.JSON(http.StatusOK, data)
	})

	rg.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, "helloworld")
	})
}
