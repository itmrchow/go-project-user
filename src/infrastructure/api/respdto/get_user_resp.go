package respdto

import "time"

type GetUserResp struct {
	Id        string    `json:"id"`
	UserName  string    `json:"userName"`
	Account   string    `json:"account"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedBy string    `json:"CreatedBy"`
	UpdatedBy string    `json:"UpdatedBy"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
