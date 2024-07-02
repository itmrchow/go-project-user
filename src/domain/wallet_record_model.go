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
}

type WalletRecordStatus uint

const (
	WALLETRECORDSTATUS_PENDING = iota
	WALLETRECORDSTATUS_SUCCESS
	WALLETRECORDSTATUS_FAILED
)
