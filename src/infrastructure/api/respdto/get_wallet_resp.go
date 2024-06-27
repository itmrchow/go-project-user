package respdto

import "time"

type GetWalletResp struct {
	UserId     string    `json:"userId"`
	WalletType string    `json:"WalletType"`
	Currency   string    `json:"currency"` // 幣別
	Balance    float64   `json:"balance"`  // 餘額
	CreatedBy  string    `json:"createdBy"`
	UpdatedBy  string    `json:"updatedBy"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
