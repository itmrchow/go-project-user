package repo_impl

import (
	"errors"

	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"

)

type UserRepoImpl struct {
	handler database.DB_Handler
}

func NewUserRepoImpl(handler database.DB_Handler) UserRepoImpl {
	return UserRepoImpl{
		handler: handler,
	}
}

func (u UserRepoImpl) Create(user *domain.User) error {
	return u.handler.DB.Create(*user).Error
}

func (u UserRepoImpl) Get(userId string) (*domain.User, error) {
	var user = domain.User{Id: userId}
	result := u.handler.DB.First(&user)

	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return nil, result.Error
	}
}

func (u UserRepoImpl) ExistsByAccountOrEmailOrPhone(account string, email string, phone string) (bool, error) {
	var count int64
	queryStr := "account = ? OR email = ? OR phone = ?"
	err := u.handler.DB.Model(&domain.User{}).Where(queryStr, account, email, phone).Count(&count).Error

	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}
