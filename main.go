package main

import (
	"itmrchow/go-project/user/src/infrastructure/api"
	"itmrchow/go-project/user/src/infrastructure/initialization"
)

// @title           User Service API
// @version         1.0
// @description     User Service API

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @tag.name User
// @tag.description User API
// @tag.name Example
// @tag.description Example API
// @tag.name Wallet
// @tag.description Wallet API
// @tag.name Other
// @tag.description Other description

func main() {

	initialization.SetConfig()
	initialization.SetDbConfig()

	api.Run()
}
