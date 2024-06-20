package usecase

import (
	"errors"

	"gorm.io/gorm"

	"itmrchow/go-project/user/src/usecase/repo"
)

// 定義output
type GetUserOutput struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type GetUserUseCase struct {
	userRepo repo.UserRepo
}

func NewGetUserUseCase(userRepo repo.UserRepo) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
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
