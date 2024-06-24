package controllers

import (
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/usecase"
)

type UserController struct {
	createUserUC *usecase.CreateUserUseCase
	getUserUC    *usecase.GetUserUseCase
}

func NewUserController(
	createUserUc *usecase.CreateUserUseCase,
	getUserUc *usecase.GetUserUseCase,
) *UserController {
	return &UserController{
		createUserUC: createUserUc,
		getUserUC:    getUserUc,
	}
}

func (controller *UserController) CreateUser(createUserReq *reqdto.CreateUserReq) (*respdto.CreateUserResp, error) {

	input := new(usecase.CreateUserInput)
	input.Account = createUserReq.Account
	input.Email = createUserReq.Email
	input.Phone = createUserReq.Phone
	input.Password = createUserReq.Password
	input.UserName = createUserReq.UserName

	out, err := controller.createUserUC.CreateUser(*input)

	if err != nil {
		return nil, err
	}

	return &respdto.CreateUserResp{
		Id:       out.Id,
		UserName: out.UserName,
		Account:  out.Account,
		Email:    out.Email,
		Phone:    out.Phone,
	}, err
}

func (controller *UserController) GetUser(userId string) (*respdto.GetUserResp, error) {

	out, err := controller.getUserUC.GetUser(userId)

	if out == nil {
		return nil, err
	}

	return &respdto.GetUserResp{
		Id:       out.Id,
		UserName: out.UserName,
		Account:  out.Account,
		Email:    out.Email,
		Phone:    out.Phone,
	}, nil
}

func (controller *UserController) Login(loginReq *reqdto.LoginReq) error {
	return nil
}
