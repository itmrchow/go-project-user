package reqdto

type FindWalletsReq struct {
	WalletType string `form:"walletType" json:"walletType" `
	Currency   string `form:"currency"   json:"currency"   `
}
