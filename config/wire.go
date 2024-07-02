//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/google/wire"

	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
	"itmrchow/go-project/user/src/interfaces/handlerimpl"
	"itmrchow/go-project/user/src/interfaces/repo_impl"
	"itmrchow/go-project/user/src/usecase"
	"itmrchow/go-project/user/src/usecase/handler"
	"itmrchow/go-project/user/src/usecase/repo"
)

var dbSet = wire.NewSet(
	database.NewMySqlHandler,
)

var repoSet = wire.NewSet(
	repo_impl.NewUserRepoImpl, wire.Bind(new(repo.UserRepo), new(*repo_impl.UserRepoImpl)),
	repo_impl.NewWalletRepoImpl, wire.Bind(new(repo.WalletRepo), new(*repo_impl.WalletRepoImpl)),
	repo_impl.NewWalletRecordRepoImpl, wire.Bind(new(repo.WalletRecordRepo), new(*repo_impl.WalletRecordRepoImpl)),
)

var handlerSet = wire.NewSet(
	handlerimpl.NewBcryptHandler, wire.Bind(new(handler.EncryptionHandler), new(*handlerimpl.BcryptHandler)),
)

var usecaseSet = wire.NewSet(
	usecase.NewCreateUserUseCase,
	usecase.NewGetUserUseCase,
	usecase.NewWalletUseCase,
	usecase.NewPingServiceImpl, wire.Bind(new(usecase.PingService), new(*usecase.PingServiceImpl)),
)

func InitUserController() (*controllers.UserController, error) {
	wire.Build(dbSet, repoSet, handlerSet, usecaseSet, controllers.NewUserController)
	return &controllers.UserController{}, nil
}

func InitPingController() (*controllers.PingController, error) {

	wire.Build(dbSet, repoSet, usecaseSet, controllers.NewPingController)

	return &controllers.PingController{}, nil
}

func InitWalletController() (*controllers.WalletController, error) {

	wire.Build(dbSet, repoSet, usecaseSet, controllers.NewWalletController)

	return &controllers.WalletController{}, nil
}
