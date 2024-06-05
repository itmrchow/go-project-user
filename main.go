package main

import (
	"itmrchow/go-project/user/routes"
	"itmrchow/go-project/user/setting"
)

func main() {
	setting.MySqlORMSetting()
	routes.Run()
}
