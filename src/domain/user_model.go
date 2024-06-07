package domain

type User struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (user *User) CheckFieId() {
	// TODO: 補其他邏輯
	if user.Account == "" {
		panic("account is empty")
	}
}
