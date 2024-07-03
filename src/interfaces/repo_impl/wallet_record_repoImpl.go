package repo_impl

import (
	"log"

	"github.com/gin-gonic/gin"
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

func (w *WalletRecordRepoImpl) Create(ctx *gin.Context, wallet *domain.WalletRecord) error {
	return w.DB.WithContext(ctx).Create(wallet).Error
}

func (w *WalletRecordRepoImpl) Get(ctx *gin.Context, id uint) (*domain.WalletRecord, error) {
	var walletRecord = domain.WalletRecord{}
	walletRecord.ID = id

	result := w.DB.WithContext(ctx).Preload("Wallet").First(&walletRecord)
	return &walletRecord, result.Error
}

func (w *WalletRecordRepoImpl) Update(ctx *gin.Context, record *domain.WalletRecord) (int64, error) {
	result := w.DB.Model(&domain.WalletRecord{}).Where("id = ?", record.ID).Updates(record)
	return result.RowsAffected, result.Error
}

func (w *WalletRecordRepoImpl) WithTrx(trxHandle *gorm.DB) repo.WalletRecordRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return w
	}

	return &WalletRecordRepoImpl{
		DB: trxHandle,
	}
}

func (w *WalletRecordRepoImpl) Migrate() error {
	panic("TODO: Implement")
}
