package repo_impl

import (
	"context"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"
)

type WalletRepoImpl struct {
	handler *database.MysqlHandler
}

func NewWalletRepoImpl(handler *database.MysqlHandler) *WalletRepoImpl {
	return &WalletRepoImpl{
		handler: handler,
	}
}

func (w *WalletRepoImpl) Create(wallet *domain.Wallet) error {

	return w.handler.DB.Create(wallet).Error
}

func (w *WalletRepoImpl) Get(ctx context.Context, walletId string) (*domain.Wallet, error) {
	// tx :=w.handler.DB.WithContext(ctx)

	panic("TODO: Implement")
}

func (w *WalletRepoImpl) Find(query interface{}, args ...interface{}) ([]domain.Wallet, error) {
	wallets := []domain.Wallet{}

	result := w.handler.DB.Where(query, args...).Find(&wallets)

	if result.Error != nil {
		return nil, result.Error
	}
	return wallets, nil
}

func (w *WalletRepoImpl) GetByUserIdAndWalletType(ctx context.Context, userId, walletType string) (*domain.Wallet, error) {
	var wallet = domain.Wallet{}
	tx := w.handler.DB.WithContext(ctx)
	result := tx.First(&wallet, "user_id = ? AND wallet_type = ?", userId, walletType)

	return &wallet, result.Error
}
