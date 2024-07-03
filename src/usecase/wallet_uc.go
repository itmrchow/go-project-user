package usecase

import (
	"errors"
	"fmt"
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

func (u *WalletUseCase) GetWallet(ctx *gin.Context, input *GetWalletInput) (*domain.Wallet, error) {

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

	return wallet, nil
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
	// 扣款
	if err := u.DecrementMoney(ctx, fromInfo.Id, fromInfo.UserId, fromInfo.WalletType, input.Amount, "轉帳扣款", input.Description, authUser); err != nil {
		return err
	}

	// 付款
	if err := u.IncrementMoney(ctx, toInfo.Id, toInfo.UserId, toInfo.WalletType, input.Amount, authUser); err != nil {
		return err
	}

	return nil
}

// DecrementMoney
func (u *WalletUseCase) DecrementMoney(ctx *gin.Context, walletId uint, userId string, walletType string, amount float64, eventName string, depiction string, authUser reqdto.AuthUser) error {
	// get wallet
	var wallet *domain.Wallet
	var err error

	if walletId == 0 {
		// get by UserId & Wallet Type
		wallet, err = u.walletRepo.GetByUserIdAndWalletType(ctx, userId, walletType)
	} else {
		// get by id
		wallet, err = u.walletRepo.Get(ctx, walletId)
	}
	if err != nil {
		return errors.Join(ErrDbFail, err)
	}

	// check amount
	if !wallet.CheckDecrementAmount(amount) {
		return errors.Join(ErrPaymentRequired, errors.New("invalid amount"))
	}

	// create wallet record
	walletRecord := domain.WalletRecord{
		DefaultModel: domain.DefaultModel{
			CreatedBy: authUser.Id,
			UpdatedBy: authUser.Id,
		},
		WalletId:    wallet.ID,
		RecordName:  eventName,
		Currency:    wallet.Currency,
		Amount:      amount * -1,
		Status:      domain.WALLETRECORDSTATUS_PENDING,
		Description: depiction,
	}

	if err := u.walletRecordRepo.Create(ctx, &walletRecord); err != nil {
		return errors.Join(ErrDbFail, err)
	}

	// update balance
	wallet.SetDeductBalance(amount)
	wallet.DefaultModel.UpdatedBy = authUser.Id

	if updatedCount, err := u.walletRepo.Update(wallet); err != nil || updatedCount == 0 {
		if err != nil {
			return errors.Join(ErrDbFail, err)
		} else {
			return errors.Join(ErrDbFail, errors.New("update fail"))
		}
	}

	return nil
}

// IncrementMoney
func (u *WalletUseCase) IncrementMoney(ctx *gin.Context, walletId uint, userId string, walletType string, amount float64, authUser reqdto.AuthUser) error {
	// get wallet
	var wallet *domain.Wallet
	var err error

	if walletId == 0 {
		// get by UserId & Wallet Type
		wallet, err = u.walletRepo.GetByUserIdAndWalletType(ctx, userId, walletType)
	} else {
		// get by id
		wallet, err = u.walletRepo.Get(ctx, walletId)
	}
	if err != nil {
		return errors.Join(ErrDbFail, err)
	}

	// create wallet record
	walletRecord := domain.WalletRecord{
		DefaultModel: domain.DefaultModel{
			CreatedBy: authUser.Id,
			UpdatedBy: authUser.Id,
		},
		WalletId:    wallet.ID,
		RecordName:  "Increment Money",
		Currency:    wallet.Currency,
		Amount:      amount,
		Status:      domain.WALLETRECORDSTATUS_PENDING,
		Description: fmt.Sprintf("Increment Money:%f", amount),
	}

	if err := u.walletRecordRepo.Create(ctx, &walletRecord); err != nil {
		return errors.Join(ErrDbFail, err)
	}

	// update balance
	wallet.SetDeductBalance(amount)
	wallet.DefaultModel.UpdatedBy = authUser.Id

	if updatedCount, err := u.walletRepo.Update(wallet); err != nil || updatedCount == 0 {
		if err != nil {
			return errors.Join(ErrDbFail, err)
		} else {
			return errors.Join(ErrDbFail, errors.New("update fail"))
		}
	}

	return nil
}
