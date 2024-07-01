package reqdto

type TransferFundsReq struct {
	ToWalletInfo   WalletInfoReq `json:"toWalletInfo"   `
	FromWalletInfo WalletInfoReq `json:"fromWalletInfo"   `
	Amount         float64       `json:"amount" example:"100.00"`
	Description    string        `json:"description" example:"Transfer Funds"`
}

type WalletInfoReq struct {
	Id         string `json:"account" example:"1234567890"`
	WalletType string `form:"walletType" json:"walletType" `
	Currency   string `form:"currency"   json:"currency"   `
}
