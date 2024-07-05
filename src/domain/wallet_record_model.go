package domain

type WalletRecord struct {
	DefaultModel `gorm:"embedded"`

	WalletId      uint               `gorm:"type:int;not null"`
	Wallet        Wallet             `gorm:"foreignKey:WalletId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	RecordName    string             `gorm:"type:varbinary(20);not null"`
	Currency      string             `gorm:"type:varbinary(10);not null"`
	Amount        float64            `gorm:"type:decimal;not null"`
	WalletBalance float64            `gorm:"type:decimal;"`
	Status        WalletRecordStatus `gorm:"type:int;not null"`
	Description   string             `gorm:"type:longtext"`
	// TODO: retry count
	// TODO: msg
}

type WalletRecordStatus uint

const (
	WALLET_RECORD_STATUS_PENDING WalletRecordStatus = iota
	WALLET_RECORD_STATUS_SUCCESS
	WALLET_RECORD_STATUS_FAILED
)
