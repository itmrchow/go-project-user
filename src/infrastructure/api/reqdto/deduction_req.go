package reqdto

type DeductionReq struct {
	WalletId    uint    `json:"walletId"     example:"12"`
	Amount      float64 `json:"amount"     example:"50"`
	EventName   string  `json:"eventName"     example:"Deduction"`
	Description string  `json:"description"     example:"幫你扣個錢"`
}
