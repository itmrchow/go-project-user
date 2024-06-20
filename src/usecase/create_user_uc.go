package usecase

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
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
}

type CreateUserUseCase struct {
	userRepo          repo.UserRepo
	encryptionHandler handler.EncryptionHandler
}

func NewCreateUserUseCase(userRepo repo.UserRepo, encryptionHandler handler.EncryptionHandler) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo, encryptionHandler: encryptionHandler}
}

func (c CreateUserUseCase) CreateUser(input CreateUserInput) (*CreateUserOutput, error) {

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
	}

	if err := c.userRepo.Create(userModel); err == nil {
		return &CreateUserOutput{
			Id:       userModel.Id,
			UserName: userModel.UserName,
			Account:  userModel.Account,
			Email:    userModel.Email,
			Phone:    userModel.Phone,
		}, nil
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errors.Join(ErrDbInsertFail, err)
	} else {
		return nil, errors.Join(ErrDbFail, err)
	}

}
