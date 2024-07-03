//go:generate mockery --name WalletRepo
package repo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
)

type WalletRecordRepo interface {
	Create(ctx *gin.Context, record *domain.WalletRecord) error
	Get(ctx *gin.Context, id uint) (*domain.WalletRecord, error)
	Update(ctx *gin.Context, record *domain.WalletRecord) (int64, error)
	WithTrx(*gorm.DB) WalletRecordRepo
	Migrate() error
}
