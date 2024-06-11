package reqdto

type PutUserReq struct {
	UserName string `json:"userName" example:"Jeff"`
	Account  string `json:"account" example:"jeff7777"`
	Password string `json:"password" example:"jeffpwd"`
	Email    string `json:"email" example:"jeff@gmail.com"`
	Phone    string `json:"phone" example:"+886955555555"`
}
