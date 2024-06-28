package usecase

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/usecase/repo"
)

type WalletUseCase struct {
	walletRepo repo.WalletRepo
}

func NewWalletUseCase(walletRepo repo.WalletRepo) *WalletUseCase {
	return &WalletUseCase{walletRepo: walletRepo}
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
