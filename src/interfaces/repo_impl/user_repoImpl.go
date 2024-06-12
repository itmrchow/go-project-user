package repo_impl

import (
	"errors"
	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	handler database.DB_Handler
}

func NewUserRepoImpl(handler database.DB_Handler) UserRepoImpl {
	return UserRepoImpl{
		handler: handler,
	}
}

func (r UserRepoImpl) Create(user *domain.User) error {
	return r.handler.DB.Create(*user).Error
}

func (r UserRepoImpl) Get(userId string) (*domain.User, error) {
	var user = domain.User{Id: userId}
	result := r.handler.DB.First(&user)

	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return nil, result.Error
	}
}
