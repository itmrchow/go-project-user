package initialization

import (
	"github.com/spf13/viper"

	"itmrchow/go-project/user/src/infrastructure/database"
)

func SetConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic("read config error: " + err.Error())
	}
}

func SetDbConfig() {
	mysqlHandler, dbCreateErr := database.NewMySqlHandler()
	if dbCreateErr != nil {
		panic("failed to connect to database: " + dbCreateErr.Error())
	}

	migrateErr := mysqlHandler.Migrate()
	if migrateErr != nil {
		panic("database migrate error: " + migrateErr.Error())
	}
}

// func SetI18n() {
// 	count := 2

// 	bundle := i18n.NewBundle(language.English)
// 	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
// 	bundle.LoadMessageFile("i18n/active.en.toml")
// 	bundle.LoadMessageFile("i18n/active.es.toml")

// 	localizer := i18n.NewLocalizer(bundle, "es")

// 	buying := localizer.MustLocalize(&i18n.LocalizeConfig{
// 		MessageID:   "BuyingCookies",
// 		PluralCount: count,
// 	})

// 	fmt.Print(buying)
// }
