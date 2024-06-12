package usecase

import (
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

func NewGetUserUseCase(userRepo repo.UserRepo) GetUserUseCase {
	return GetUserUseCase{userRepo: userRepo}
}

func (c GetUserUseCase) GetUser(userId string) *GetUserOutput {
	user, _ := c.userRepo.Get(userId)
	return &GetUserOutput{
		Id:       user.Id,
		UserName: user.UserName,
		Account:  user.Account,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}
