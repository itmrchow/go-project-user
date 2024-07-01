//go:generate mockery --name WalletRepo
package repo

import (
	"context"

	"itmrchow/go-project/user/src/domain"
)

type WalletRepo interface {
	Create(wallet *domain.Wallet) error
	Get(ctx context.Context, walletId string) (*domain.Wallet, error)
	Find(query interface{}, args ...interface{}) ([]domain.Wallet, error)
	GetByUserIdAndWalletType(ctx context.Context, userId, walletType string) (*domain.Wallet, error)
}
