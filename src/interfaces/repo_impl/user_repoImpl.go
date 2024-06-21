package repo_impl

import (
	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"
)

type UserRepoImpl struct {
	handler *database.MysqlHandler
}

func NewUserRepoImpl(handler *database.MysqlHandler) *UserRepoImpl {
	return &UserRepoImpl{
		handler: handler,
	}
}

func (u *UserRepoImpl) Create(user *domain.User) error {
	return u.handler.DB.Create(*user).Error
}

func (u *UserRepoImpl) Get(userId string) (*domain.User, error) {
	var user = domain.User{Id: userId}
	result := u.handler.DB.First(&user)

	return &user, result.Error
}

func (u *UserRepoImpl) ExistsByAccountOrEmailOrPhone(account string, email string, phone string) (bool, error) {
	var count int64
	queryStr := "account = ? OR email = ? OR phone = ?"
	err := u.handler.DB.Model(&domain.User{}).Where(queryStr, account, email, phone).Count(&count).Error

	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}
