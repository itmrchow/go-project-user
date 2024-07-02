package repo_impl

import (
	"log"

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

func (w *WalletRecordRepoImpl) WithTrx(trxHandle *gorm.DB) repo.WalletRecordRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return w
	}
	w.DB = trxHandle
	return w
}

func (w *WalletRecordRepoImpl) Migrate() error {
	panic("TODO: Implement")
}
