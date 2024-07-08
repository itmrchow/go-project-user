package initialization

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"itmrchow/go-project/user/src/infrastructure/database"
)

func SetConfig() {

	path, err := os.Getwd()

	if err != nil {
		panic(err.Error())
	}

	index := strings.Index(path, "go-project-user")

	if index > 0 {
		prefix := path[:(strings.Index(path, "go-project-user"))]
		path = prefix + "go-project-user/config"
	} else {
		path += "/config"
	}

	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

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
