package usecase

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/usecase/repo"
)

type WalletUseCase struct {
	walletRepo       repo.WalletRepo
	walletRecordRepo repo.WalletRecordRepo
}

func NewWalletUseCase(walletRepo repo.WalletRepo, walletRecord repo.WalletRecordRepo) *WalletUseCase {
	return &WalletUseCase{walletRepo: walletRepo, walletRecordRepo: walletRecord}
}

type CreateWalletInput struct {
	UserId     string
	WalletType string
	Currency   string
	Balance    float64
}

type CreateWalletOutput struct {
	UserId     string
	WalletType string
	Currency   string
	Balance    float64
	CreatedBy  string    `json:"CreatedBy"`
	UpdatedBy  string    `json:"UpdatedBy"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}

func (u *WalletUseCase) WithTrx(trxHandle *gorm.DB) *WalletUseCase {
	u.walletRepo = u.walletRepo.WithTrx(trxHandle)
	u.walletRecordRepo = u.walletRecordRepo.WithTrx(trxHandle)
	return u
}

func (u *WalletUseCase) CreateWallet(input *CreateWalletInput, authUser reqdto.AuthUser) (*CreateWalletOutput, error) {

	wallet := domain.Wallet{
		DefaultModel: domain.DefaultModel{
			CreatedBy: authUser.Id,
			UpdatedBy: authUser.Id,
		},
	}

	if err := copier.Copy(&wallet, &input); err != nil {
		return nil, err
	}

	if err := u.walletRepo.Create(&wallet); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errors.Join(ErrDataExists, err)
		}
		return nil, errors.Join(ErrDbInsertFail, err)
	}

	out := CreateWalletOutput{}
	if err := copier.Copy(&out, &wallet); err != nil {
		return nil, err
	}

	return &out, nil
}

type FindWalletInput struct {
	UserId     string
	WalletType string
	Currency   string
}
type FindWalletOutput struct {
	UserId     string
	WalletType string
	Currency   string
	Balance    float64
	CreatedBy  string
	UpdatedBy  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *WalletUseCase) FindWallet(input *FindWalletInput) (*[]FindWalletOutput, error) {

	query := domain.Wallet{}
	if err := copier.Copy(&query, &input); err != nil {
		return nil, err
	}

	wallets, err := u.walletRepo.Find(query)

	if err != nil {
		return nil, errors.Join(ErrDbFail, err)
	}

	outSlice := []FindWalletOutput{}

	if err := copier.Copy(&outSlice, &wallets); err != nil {
		return nil, errors.Join(ErrDbFail, err)
	}

	return &outSlice, nil
}

type GetWalletInput struct {
	Id         uint
	UserId     string
	WalletType string
}

type GetWalletOut struct {
	Id         uint
	UserId     string
	WalletType string
	Currency   string
	Balance    float64
	CreatedBy  string
	UpdatedBy  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *WalletUseCase) GetWallet(ctx *gin.Context, input *GetWalletInput) (*GetWalletOut, error) {

	var wallet *domain.Wallet
	var err error

	if input.Id == 0 {
		// get by UserId & Wallet Type
		wallet, err = u.walletRepo.GetByUserIdAndWalletType(ctx, input.UserId, input.WalletType)
	} else {
		// get by id
		wallet, err = u.walletRepo.Get(ctx, input.Id)
	}

	if err != nil {
		return nil, errors.Join(ErrDbFail, err)
	}

	out := &GetWalletOut{}
	if err := copier.Copy(out, wallet); err != nil {
		return nil, err
	}

	return out, nil
}

type TransferFundsInput struct {
	ToWalletInfo   TransferFundsWalletInfoInput
	FromWalletInfo TransferFundsWalletInfoInput
	Amount         float64
	Description    string
}

type TransferFundsWalletInfoInput struct {
	Id         uint   `json:"account" example:"1234567890"`
	UserId     string `json:"userId" example:"1234567890"`
	WalletType string `form:"walletType" json:"walletType" `
}

// 轉帳
func (u *WalletUseCase) TransferFunds(ctx *gin.Context, input *TransferFundsInput, authUser reqdto.AuthUser) error {
	toInfo := input.ToWalletInfo
	fromInfo := input.FromWalletInfo
	eventName := "轉帳扣款"

	// 扣款
	if err := u.DecrementMoney(ctx, fromInfo.Id, input.Amount, eventName, input.Description, authUser.Id); err != nil {
		return err
	}

	// 付款
	if err := u.IncrementMoney(ctx, toInfo.Id, input.Amount, eventName, input.Description, authUser.Id); err != nil {
		return err
	}

	return nil
}

// DecrementMoney
func (u *WalletUseCase) DecrementMoney(ctx *gin.Context, walletId uint, amount float64, eventName string, depiction string, updateUserId string) error {
	amount = amount * -1
	record, err := u.CreateWalletRecord(ctx, walletId, amount, eventName, depiction, updateUserId)
	if err != nil {
		return nil
	}

	db := ctx.MustGet("DB").(*gorm.DB)
	return db.Transaction(func(tx *gorm.DB) error {

		if err := u.UpdateWalletByRecord(ctx, tx, record); err != nil {
			return err
		}

		return u.UpdateWalletRecord(ctx, tx, record)
	})
}

// IncrementMoney
func (u *WalletUseCase) IncrementMoney(ctx *gin.Context, walletId uint, amount float64, eventName string, depiction string, updateUserId string) error {
	record, err := u.CreateWalletRecord(ctx, walletId, amount, eventName, depiction, updateUserId)
	if err != nil {
		return nil
	}

	db := ctx.MustGet("DB").(*gorm.DB)
	return db.Transaction(func(tx *gorm.DB) error {

		if err := u.UpdateWalletByRecord(ctx, tx, record); err != nil {
			return err
		}

		return u.UpdateWalletRecord(ctx, tx, record)
	})
}

func (u *WalletUseCase) CreateWalletRecord(ctx *gin.Context, walletId uint, amount float64, eventName string, depiction string, updateUserId string) (*domain.WalletRecord, error) {
	// get wallet
	wallet, err := u.walletRepo.Get(ctx, walletId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Join(ErrDataNotExists, err)
		} else {
			return nil, errors.Join(ErrDbFail, err)
		}
	}

	// create wallet record
	walletRecord := &domain.WalletRecord{
		DefaultModel: domain.DefaultModel{
			CreatedBy: updateUserId,
			UpdatedBy: updateUserId,
		},
		WalletId:    wallet.ID,
		RecordName:  eventName,
		Currency:    wallet.Currency,
		Amount:      amount,
		Status:      domain.WALLET_RECORD_STATUS_PENDING,
		Description: depiction,
	}

	if err := u.walletRecordRepo.Create(ctx, walletRecord); err != nil {
		return nil, errors.Join(ErrDbFail, err)
	}

	return walletRecord, nil
}

func (u *WalletUseCase) UpdateWalletByRecord(ctx *gin.Context, tx *gorm.DB, record *domain.WalletRecord) error {

	txWalletRepo := u.walletRepo.WithTrx(tx)

	// get record
	wallet, err := txWalletRepo.GetWalletWithLock(ctx, record.WalletId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// update record , status = failed
		record.Status = domain.WALLET_RECORD_STATUS_FAILED
		return nil
	}
	if err != nil {
		// go retry
		return errors.Join(ErrDbFail, err)
	}

	// check balance
	if !wallet.CheckDecrementAmount(record.Amount) {
		record.Status = domain.WALLET_RECORD_STATUS_FAILED
		return nil
	}

	// update wallet
	wallet.UpdatedBy = record.CreatedBy

	if updatedCount, err := txWalletRepo.Update(wallet, record.Amount); err != nil || updatedCount == 0 {
		if err != nil {
			return errors.Join(ErrDbFail, err)
		} else {
			record.Status = domain.WALLET_RECORD_STATUS_FAILED
			return nil
		}
	}

	// update record
	record.Status = domain.WALLET_RECORD_STATUS_SUCCESS
	record.WalletBalance = wallet.Balance

	return nil

}

func (u *WalletUseCase) UpdateWalletRecord(ctx *gin.Context, tx *gorm.DB, record *domain.WalletRecord) error {
	repo := u.walletRecordRepo

	if tx != nil {
		repo = repo.WithTrx(tx)
	}

	if updatedCount, err := repo.Update(ctx, record); err != nil || updatedCount == 0 {
		if err != nil {
			return errors.Join(ErrDbFail, err)
		} else {
			msg := fmt.Sprintf("wallet record not exist:%v", record)
			log.Println(msg)
			// no
			return errors.Join(ErrDataNotExists, errors.New(msg))
		}
	}
	return nil
}
