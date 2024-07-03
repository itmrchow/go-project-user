package repo_impl

import (
	"context"
	"log"

	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/usecase/repo"
)

type WalletRepoImpl struct {
	DB *gorm.DB
}

func NewWalletRepoImpl(handler *database.MysqlHandler) *WalletRepoImpl {
	return &WalletRepoImpl{
		DB: handler.DB,
	}
}

func (w *WalletRepoImpl) Create(wallet *domain.Wallet) error {

	return w.DB.Create(wallet).Error
}

func (w *WalletRepoImpl) Get(ctx context.Context, walletId uint) (*domain.Wallet, error) {
	var wallet = domain.Wallet{}
	wallet.ID = walletId
	result := w.DB.First(&wallet)

	return &wallet, result.Error
}

func (w *WalletRepoImpl) Find(query interface{}, args ...interface{}) ([]domain.Wallet, error) {
	wallets := []domain.Wallet{}

	result := w.DB.Where(query, args...).Find(&wallets)

	if result.Error != nil {
		return nil, result.Error
	}
	return wallets, nil
}

func (w *WalletRepoImpl) GetByUserIdAndWalletType(ctx context.Context, userId, walletType string) (*domain.Wallet, error) {
	var wallet = domain.Wallet{}
	tx := w.DB.WithContext(ctx)
	result := tx.First(&wallet, "user_id = ? AND wallet_type = ?", userId, walletType)

	return &wallet, result.Error
}

func (w *WalletRepoImpl) Update(wallet *domain.Wallet) (int64, error) {
	result := w.DB.Model(&domain.Wallet{}).Where("id = ?", wallet.ID).Updates(wallet)
	return result.RowsAffected, result.Error
}

func (w *WalletRepoImpl) WithTrx(trxHandle *gorm.DB) repo.WalletRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return w
	}

	return &WalletRepoImpl{
		DB: trxHandle,
	}
}

func (w *WalletRepoImpl) Migrate() error {
	panic("TODO: Implement")
}
