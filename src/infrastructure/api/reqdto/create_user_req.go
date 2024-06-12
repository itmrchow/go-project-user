package reqdto

type CreateUserReq struct {
	UserName string `json:"userName" example:"Jeff" binding:"required,min=4,max=20"`
	Account  string `json:"account" example:"jeff7777" binding:"required,min=8,max=20"`
	Password string `json:"password" example:"jeffpwd" binding:"required,min=8,max=20"`
	Email    string `json:"email" example:"jeff@gmail.com" binding:"required,email"`
	Phone    string `json:"phone" example:"+886955555555" binding:"required,e164"`
}
