package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/config"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
)

func addWalletRoutes(rg *gin.RouterGroup) {

	walletController, err := config.InitWalletController()

	if err != nil {
		panic(err)
	}

	rg.Use(RequireAuth)

	rg.POST("/wallet", func(c *gin.Context) {
		createWallet(c, walletController)
	})

	rg.GET("/wallets", func(c *gin.Context) {
		findWallets(c, walletController)
	})

}

// @Summary 取得錢包
// @Produce json
// @Tags Wallet
// @Param userId     path string true "User Id"
// @Param walletType path string true "Wallet Type"
// @Success 200 {object} respdto.GetWalletResp "返回錢包訊息"
// @response default {object} respdto.ApiErrorResp "error response"
// @Router /wallet/{userId}/{walletType} [GET]
func getWallet(c *gin.Context, controller *controllers.WalletController) {
	panic("unimplemented")
}

// @Summary 查詢錢包
// @Description "查找User所屬的錢包"
// @Produce json
// @Tags Wallet
// @Parameters.QueryParams
// @Param walletType query string false  "錢包類型" Enums(P)
// @Param currency   query string false  "幣別"    Enums(PHP,USD,BTC,USDT)
// @Success 200 {array} respdto.FindWalletResp "ok" "返回錢包查詢訊息"
// @response default {object} respdto.ApiErrorResp "error response"
// @Router /wallets [GET]
func findWallets(c *gin.Context, controller *controllers.WalletController) {

	// context to dto
	req := new(reqdto.FindWalletsReq) // bind bto
	if err := c.BindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	authUser := GetAuthUser(c)

	resp, err := controller.FindWallets(req, authUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary 建立錢包
// @Produce json
// @Tags Wallet
// @Param body body reqdto.CreateWalletReq true "Create wallet sample"
// @Success 200 {object} respdto.CreateWalletResp "返回建立錢包訊息"
// @response default {object} respdto.ApiErrorResp "error response"
// @Router /wallet [post]
func createWallet(c *gin.Context, controller *controllers.WalletController) {

	// context to dto
	walletReq := new(reqdto.CreateWalletReq) // bind bto
	if err := c.BindJSON(&walletReq); err != nil {
		c.Error(err)
		return
	}

	authUser := GetAuthUser(c)

	// call controller
	response, err := controller.CreateWallet(walletReq, authUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
