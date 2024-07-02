package repo_impl

import (
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/usecase/repo"
)

type WalletRecordRepoImpl struct {
	DB *gorm.DB
}

func NewWalletRecordRepoImpl(handler *database.MysqlHandler) *WalletRecordRepoImpl {
	return &WalletRecordRepoImpl{
		DB: handler.DB,
	}
}

func (w *WalletRecordRepoImpl) Create(wallet *domain.WalletRecord) error {
	return w.DB.Create(wallet).Error
}

func (w *WalletRecordRepoImpl) Get(id uint) (*domain.WalletRecord, error) {
	var walletRecord = domain.WalletRecord{}
	walletRecord.ID = id

	result := w.DB.Preload("Wallet").First(&walletRecord)
	return &walletRecord, result.Error
}

func (w *WalletRecordRepoImpl) WithTrx(p0 *gorm.DB) repo.WalletRecordRepo {
	panic("TODO: Implement")
}

func (w *WalletRecordRepoImpl) Migrate() error {
	panic("TODO: Implement")
}
