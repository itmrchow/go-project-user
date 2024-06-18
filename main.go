package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"

	"itmrchow/go-project/user/src/infrastructure/api"
)

// @title           User Service API
// @version         1.0
// @description     User Service API

// @securityDefinitions.basic  BasicAuth

// @tag.name User
// @tag.description User API
// @tag.name Example
// @tag.description Example API
// @tag.name Other
// @tag.description Other description

func main() {
	setI18n()
	setConfig()
	// mysqlHandler := setMysqlDB()

	// dbHander := setting.NewSqlHandler()
	api.Run()
}

func setConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic("read config error: " + err.Error())
	}
}

func setI18n() {
	count := 2

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("i18n/active.en.toml")
	bundle.LoadMessageFile("i18n/active.es.toml")

	localizer := i18n.NewLocalizer(bundle, "es")

	buying := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "BuyingCookies",
		PluralCount: count,
	})

	fmt.Print(buying)
}
