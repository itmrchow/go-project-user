package reqdto

type CreateWalletReq struct {
	UserId     string  `json:"userId"     example:"Jeff"`
	WalletType string  `json:"walletType" example:"P"`
	Currency   string  `json:"currency"   example:"PHP"`
	Balance    float64 `json:"balance"    example:"0.0"`
}
