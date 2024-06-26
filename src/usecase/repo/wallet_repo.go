package repo

import "itmrchow/go-project/user/src/domain"

type WalletRepo interface {
	Create(wallet *domain.Wallet) error
	Get(walletId string) (*domain.Wallet, error)
}
