package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"

	"itmrchow/go-project/user/delivery/api"
	"itmrchow/go-project/user/setting"
)

func main() {
	setI18n()

	setConfig()
	setting.MySqlORMSetting()
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
