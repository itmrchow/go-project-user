package main

import (
	"github.com/spf13/viper"

	"itmrchow/go-project/user/routes"
	"itmrchow/go-project/user/setting"
)

func main() {
	setConfig()

	setting.MySqlORMSetting()
	routes.Run()
}

func setConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic("read config error: " + err.Error())
	}

}
