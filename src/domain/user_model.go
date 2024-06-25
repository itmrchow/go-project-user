package domain

type User struct {
	Id           string `json:"id" gorm:"column:id;<-:create;type:varbinary(40)"`
	UserName     string `json:"userName" gorm:"varbinary(20)"`
	Account      string `json:"account" gorm:"varbinary(20)"`
	Password     string `json:"password" gorm:"varbinary(20)"`
	Email        string `json:"email" gorm:"varbinary(40)"`
	Phone        string `json:"phone" gorm:"varbinary(12)"`
	DefaultModel `gorm:"embedded"`

	Wallets []Wallet `gorm:"foreignKey:UserId"`
}

func (user *User) CheckFieId() {
	// TODO: 補其他邏輯
	if user.Account == "" {
		panic("account is empty")
	}
}
