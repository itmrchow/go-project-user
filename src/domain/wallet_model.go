package domain

type Wallet struct {
	UserId     string  `json:"userId" gorm:"uniqueIndex:uq_userid_type; type:varbinary(40)"`
	WalletType string  `json:"WalletType" gorm:"uniqueIndex:uq_userid_type; type:varbinary(20)"`
	Currency   string  `json:"currency" gorm:"type:varbinary(10)"` // 幣別
	Balance    float64 `json:"balance"`                            // 餘額

	DefaultModel `gorm:"embedded"`
}
