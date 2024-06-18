package controllers

import (
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/interfaces/handlerimpl"
	"itmrchow/go-project/user/src/interfaces/repo_impl"
	"itmrchow/go-project/user/src/usecase"
)

type UserController struct {
	createUserUC *usecase.CreateUserUseCase
	getUserUC    *usecase.GetUserUseCase
}

func NewUserController(handler *database.MysqlHandler) *UserController {

	userRepo := repo_impl.NewUserRepoImpl(handler)
	encryptionHandler := new(handlerimpl.BcryptHandler)
	createUserUC := usecase.NewCreateUserUseCase(userRepo, encryptionHandler)
	getUserUC := usecase.NewGetUserUseCase(userRepo)

	return &UserController{
		createUserUC: createUserUC,
		getUserUC:    getUserUC,
	}
}

func (controller *UserController) CreateUser(createUserReq *reqdto.CreateUserReq) *respdto.CreateUserResp {

	input := new(usecase.CreateUserInput)
	input.Account = createUserReq.Account
	input.Email = createUserReq.Email
	input.Phone = createUserReq.Phone
	input.Password = createUserReq.Password
	input.UserName = createUserReq.UserName

	out, _ := controller.createUserUC.CreateUser(*input)

	return &respdto.CreateUserResp{
		Id:       out.Id,
		UserName: out.UserName,
		Account:  out.Account,
		Email:    out.Email,
		Phone:    out.Phone,
	}
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
