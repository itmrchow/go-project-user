package controllers

import (
	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/src/infrastructure/api/req_dto"
	"itmrchow/go-project/user/src/infrastructure/api/resp_dto"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/interfaces/repo_impl"
	"itmrchow/go-project/user/src/usecase"

)

type UserController struct {
	createUserUC usecase.CreateUserUseCase
}

// 建構
func NewUserController(handler database.DB_Handler) *UserController {

	userRepo := repo_impl.NewUserRepoImpl(handler)
	createUserUC := usecase.NewCreateUserUseCase(userRepo)

	return &UserController{createUserUC: createUserUC}

}

func (controller *UserController) CreateUser(createUserReq *req_dto.UserReq) *resp_dto.UserResp {
	input := new(usecase.CreateUserInput)
	input.Account = createUserReq.Account
	input.Email = createUserReq.Email
	input.Phone = createUserReq.Phone
	input.Password = createUserReq.Password
	input.UserName = createUserReq.UserName

	out, _ := controller.createUserUC.CreateUser(*input)

	// call use case

	// userController.CreateUser()

	// fmt.Println(userReq)
	// // create user

	// c.JSON(http.StatusOK, gin.H{"msg": "success"})

	return &resp_dto.UserResp{
		Id:       out.Id,
		UserName: out.UserName,
		Account:  out.Account,
		Email:    out.Email,
		Phone:    out.Phone,
	}
}

func (controller *UserController) GetUser(c *gin.Context) {
	// Context 轉DTO
	//
}
