//go:generate mockery --name WalletRepo
package repo

import (
	"context"

	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
)

type WalletRepo interface {
	Create(wallet *domain.Wallet) error
	Get(ctx context.Context, walletId uint) (*domain.Wallet, error)
	Find(query interface{}, args ...interface{}) ([]domain.Wallet, error)
	GetByUserIdAndWalletType(ctx context.Context, userId, walletType string) (*domain.Wallet, error)
	Update(wallet *domain.Wallet, amount float64) (int64, error)
	WithTrx(*gorm.DB) WalletRepo
	Migrate() error
	GetWalletWithLock(ctx context.Context, walletId uint) (*domain.Wallet, error)

	// GetWalletId(ctx context.Context, userId, walletType string) (uint, error)
}
