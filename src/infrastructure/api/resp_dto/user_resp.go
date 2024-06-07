package resp_dto

type UserResp struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}