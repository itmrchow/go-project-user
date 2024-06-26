package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/domain/enum"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
)

// 定義input
type CreateUserInput struct {
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// 定義output
type CreateUserOutput struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`

	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserUseCase struct {
	userRepo          repo.UserRepo
	encryptionHandler handler.EncryptionHandler
}

func NewCreateUserUseCase(userRepo repo.UserRepo, encryptionHandler handler.EncryptionHandler) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo, encryptionHandler: encryptionHandler}
}

func (c CreateUserUseCase) CreateUser(input CreateUserInput, authUser reqdto.AuthUser) (*CreateUserOutput, error) {

	// 1. 欄位資料庫檢查 - 是否存在
	isExists, err := c.userRepo.ExistsByAccountOrEmailOrPhone(input.Account, input.Email, input.Phone)
	if err != nil {

		return nil, errors.Join(ErrDbFail, err)
	} else if isExists {
		return nil, ErrUserAlreadyExists
	}

	// 2. 產生UUID
	uuidStr := uuid.New().String()

	// 3. hash password
	hashStr, err := c.encryptionHandler.HashPassword(input.Password)
	if err != nil {
		return nil, errors.Join(ErrPasswordHash, err)
	}

	// 4. 寫入資料庫
	userModel := &domain.User{
		Id:       uuidStr,
		UserName: input.UserName,
		Account:  input.Account,
		Password: hashStr,
		Email:    input.Email,
		Phone:    input.Phone,
		DefaultModel: domain.DefaultModel{
			CreatedBy: authUser.Id,
			UpdatedBy: authUser.Id,
		},
	}

	walletModel := domain.Wallet{
		WalletType: enum.WalletType.PLATFORM,
		Currency:   enum.Currency.PHP,
		DefaultModel: domain.DefaultModel{
			CreatedBy: authUser.Id,
			UpdatedBy: authUser.Id,
		},
	}

	userModel.Wallets = append(userModel.Wallets, walletModel)

	if err := c.userRepo.Create(userModel); err == nil {
		return &CreateUserOutput{
			Id:        userModel.Id,
			UserName:  userModel.UserName,
			Account:   userModel.Account,
			Email:     userModel.Email,
			Phone:     userModel.Phone,
			CreatedBy: userModel.CreatedBy,
			UpdatedBy: userModel.UpdatedBy,
			CreatedAt: userModel.CreatedAt,
			UpdatedAt: userModel.UpdatedAt,
		}, nil
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errors.Join(ErrDbInsertFail, err)
	} else {
		return nil, errors.Join(ErrDbFail, err)
	}

}
