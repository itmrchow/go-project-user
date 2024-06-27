package respdto

import "time"

type CreateWalletResp struct {
	UserId     string    `json:"userId" example:"Jeff"`
	WalletType string    `json:"WalletType" example:"P"`
	Currency   string    `json:"currency" example:"PHP"` // 幣別
	Balance    float64   `json:"balance" example:"0.0"`  // 餘額
	CreatedBy  string    `json:"createdBy" example:"Jeff"`
	UpdatedBy  string    `json:"updatedBy" example:"Jeff"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
