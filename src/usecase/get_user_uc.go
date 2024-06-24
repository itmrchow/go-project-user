package usecase

import (
	"errors"

	"gorm.io/gorm"

	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
)

type LoginInput struct {
	Account  string
	Email    string
	Password string
}

// 定義output
type GetUserOutput struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type GetUserUseCase struct {
	userRepo          repo.UserRepo
	encryptionHandler handler.EncryptionHandler
}

func NewGetUserUseCase(userRepo repo.UserRepo, encryptionHandler handler.EncryptionHandler) *GetUserUseCase {
	return &GetUserUseCase{
		userRepo:          userRepo,
		encryptionHandler: encryptionHandler,
	}
}

func (c GetUserUseCase) GetUser(userId string) (*GetUserOutput, error) {
	user, err := c.userRepo.Get(userId)

	if err == nil {
		return &GetUserOutput{
			Id:       user.Id,
			UserName: user.UserName,
			Account:  user.Account,
			Email:    user.Email,
			Phone:    user.Phone,
		}, nil

	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return nil, errors.Join(ErrDbFail, err)
	}
}

func (c GetUserUseCase) Login(loginInput LoginInput) (string, error) {
	// query user
	user, err := c.userRepo.GetByAccountOrEmail(loginInput.Account, loginInput.Email)
	if err != nil {
		return "", errors.Join(ErrUserNotExists, err)
	}

	// check password
	psw := user.Password
	isCorrectPsw := c.encryptionHandler.CheckPasswordHash(loginInput.Password, psw)

	if !isCorrectPsw {
		return "", ErrUnauthorized
	}

	// create token & return

	return "", nil
}
