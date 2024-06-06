package repository

import (
	"itmrchow/go-project/user/model"

)

type Repository interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id string) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(id string) error
	// GetUserBy
}
