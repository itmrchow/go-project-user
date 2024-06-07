package repo

import "itmrchow/go-project/user/src/domain"

type UserRepo interface {
	Create(user *domain.User) error
}
