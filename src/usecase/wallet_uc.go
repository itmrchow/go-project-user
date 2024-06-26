package usecase

import (
	"errors"
	"time"

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
		UserId:     input.UserId,
		WalletType: input.WalletType,
		Currency:   input.Currency,
		Balance:    input.Balance,
		DefaultModel: domain.DefaultModel{
			CreatedBy: authUser.Id,
			UpdatedBy: authUser.Id,
		},
	}

	err := u.walletRepo.Create(&wallet)

	if err == nil {
		return &CreateWalletOutput{
			UserId:     wallet.UserId,
			WalletType: wallet.WalletType,
			Currency:   wallet.Currency,
			Balance:    wallet.Balance,
			CreatedBy:  wallet.DefaultModel.CreatedBy,
			UpdatedBy:  wallet.DefaultModel.UpdatedBy,
			CreatedAt:  wallet.DefaultModel.CreatedAt,
			UpdatedAt:  wallet.DefaultModel.UpdatedAt,
		}, nil
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errors.Join(ErrDataExists, err)
	} else {
		return nil, errors.Join(ErrDbInsertFail, err)
	}

}
