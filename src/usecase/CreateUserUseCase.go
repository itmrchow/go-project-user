package usecase

import (
	"itmrchow/go-project/user/src/domain"
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
	userRepo repo.UserRepo
}

func NewCreateUserUseCase(userRepo repo.UserRepo) CreateUserUseCase {
	return CreateUserUseCase{userRepo: userRepo}
}

func (c CreateUserUseCase) CreateUser(input CreateUserInput) (*CreateUserOutput, error) {
	// Input 寫入 Model
	userModel := &domain.User{
		Id:       "UUID",
		UserName: input.UserName,
		Account:  input.Account,
		Password: input.Password,
		Email:    input.Email,
		Phone:    input.Phone,
	}

	// 欄位檢查
	userModel.CheckFieId()

	// 資料庫檢查
	// 資料庫新增
	if err := c.userRepo.Create(userModel); err == nil {
		return &CreateUserOutput{
			Id:       userModel.Id,
			UserName: userModel.UserName,
			Account:  userModel.Account,
			Email:    userModel.Email,
			Phone:    userModel.Phone,
		}, nil
	} else {
		return nil, err
	}

}
