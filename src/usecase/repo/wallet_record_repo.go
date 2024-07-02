//go:generate mockery --name WalletRepo
package repo

import (
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
)

type WalletRecordRepo interface {
	Create(wallet *domain.WalletRecord) error
	Get(id string) (*domain.WalletRecord, error)
	WithTrx(*gorm.DB) WalletRecordRepo
	Migrate() error
}
