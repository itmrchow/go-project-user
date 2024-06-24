package reqdto

type LoginReq struct {
	Account  string `json:"account" example:"jeff7777" binding:"required_without=Email,omitempty"`
	Email    string `json:"email" example:"jeff@gmail.com" binding:"required_without=Account,omitempty,email"`
	Password string `json:"password" example:"password" binding:"required,min=8,max=20"`
}
