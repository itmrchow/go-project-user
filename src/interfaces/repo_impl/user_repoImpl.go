package repo_impl

import (
	"itmrchow/go-project/user/src/domain"
	"itmrchow/go-project/user/src/infrastructure/database"

)

type UserRepoImpl struct {
	handler database.DB_Handler
}

// Create implements repo.UserRepo.
func (repo UserRepoImpl) Create(user *domain.User) error {
	return repo.handler.DB.Create(*user).Error
}

func NewUserRepoImpl(handler database.DB_Handler) UserRepoImpl {
	return UserRepoImpl{
		handler: handler,
	}
}

// func (repo *UserRepoImpl) Create(user *domain.User) error {
// 	return repo.handler.DB.Create(*user).Error
// }
