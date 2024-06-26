package repo

import "itmrchow/go-project/user/src/domain"

type UserRepo interface {
	Create(user *domain.User) error
	Get(userId string) (*domain.User, error)
	ExistsByAccountOrEmailOrPhone(account string, email string, phone string) (bool, error)
	GetByAccountOrEmail(account string, email string) (*domain.User, error)
}
