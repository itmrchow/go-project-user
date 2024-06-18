//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/google/wire"

	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
	"itmrchow/go-project/user/src/interfaces/repo_impl"
	"itmrchow/go-project/user/src/usecase/repo"
)

var dbSet = wire.NewSet(
	database.NewMySqlHandler,
)

var repoSet = wire.NewSet(
	repo_impl.NewUserRepoImpl,
	wire.Bind(new(repo.UserRepo), new(repo_impl.UserRepoImpl)),
)

var controllerSet = wire.NewSet(
	controllers.NewUserController,
)

var usecaseSet = wire.NewSet()

func InitUserController() (*controllers.UserController, error) {
	wire.Build(dbSet, controllers.NewUserController)

	return &controllers.UserController{}, nil
}
