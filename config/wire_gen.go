// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitUserController() (*controllers.UserController, error) {
	mysqlHandler, err := database.NewMySqlHandler()
	if err != nil {
		return nil, err
	}
	userRepoImpl := repo_impl.NewUserRepoImpl(mysqlHandler)
	bcryptHandler := handlerimpl.NewBcryptHandler()
	createUserUseCase := usecase.NewCreateUserUseCase(userRepoImpl, bcryptHandler)
	getUserUseCase := usecase.NewGetUserUseCase(userRepoImpl, bcryptHandler)
	userController := controllers.NewUserController(createUserUseCase, getUserUseCase)
	return userController, nil
}

func InitPingController() (*controllers.PingController, error) {
	mysqlHandler, err := database.NewMySqlHandler()
	if err != nil {
		return nil, err
	}
	userRepoImpl := repo_impl.NewUserRepoImpl(mysqlHandler)
	pingServiceImpl := usecase.NewPingServiceImpl(userRepoImpl)
	pingController := controllers.NewPingController(pingServiceImpl)
	return pingController, nil
}

func InitWalletController() (*controllers.WalletController, error) {
	mysqlHandler, err := database.NewMySqlHandler()
	if err != nil {
		return nil, err
	}
	walletRepoImpl := repo_impl.NewWalletRepoImpl(mysqlHandler)
	walletRecordRepoImpl := repo_impl.NewWalletRecordRepoImpl(mysqlHandler)
	walletUseCase := usecase.NewWalletUseCase(walletRepoImpl, walletRecordRepoImpl)
	walletController := controllers.NewWalletController(walletUseCase)
	return walletController, nil
}

// wire.go:

var dbSet = wire.NewSet(database.NewMySqlHandler)

var repoSet = wire.NewSet(repo_impl.NewUserRepoImpl, wire.Bind(new(repo.UserRepo), new(*repo_impl.UserRepoImpl)), repo_impl.NewWalletRepoImpl, wire.Bind(new(repo.WalletRepo), new(*repo_impl.WalletRepoImpl)), repo_impl.NewWalletRecordRepoImpl, wire.Bind(new(repo.WalletRecordRepo), new(*repo_impl.WalletRecordRepoImpl)))

var handlerSet = wire.NewSet(handlerimpl.NewBcryptHandler, wire.Bind(new(handler.EncryptionHandler), new(*handlerimpl.BcryptHandler)))

var usecaseSet = wire.NewSet(usecase.NewCreateUserUseCase, usecase.NewGetUserUseCase, usecase.NewWalletUseCase, usecase.NewPingServiceImpl, wire.Bind(new(usecase.PingService), new(*usecase.PingServiceImpl)))
