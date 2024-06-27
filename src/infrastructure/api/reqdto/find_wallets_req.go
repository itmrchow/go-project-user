package reqdto

type FindWalletsReq struct {
	WalletType string `form:"walletType" json:"walletType" binding:"required"`
	Currency   string `form:"currency"   json:"currency"   binding:"required"`
}
