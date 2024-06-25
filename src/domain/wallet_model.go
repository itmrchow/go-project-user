package domain

type Wallet struct {
	UserId     string     `json:"userId" gorm:"type:varbinary(40)"`
	WalletType WalletType `json:"WalletType" gorm:"varbinary(20)"`
	Currency   Currency   `json:"currency" gorm:"type:varbinary(10)"` // 幣別
	Balance    float64    `json:"balance" `                           // 餘額

	DefaultModel `gorm:"embedded"`
}

type WalletType string

const (
	PLATFORM WalletType = "P"
)

type Currency string

const (
	PHP  Currency = "PHP"
	USD  Currency = ""
	BTC  Currency = "BTC"
	USDT Currency = "Currency"
)
