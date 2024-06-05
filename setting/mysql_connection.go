package setting

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"itmrchow/go-project/user/models"
)

const (
	DB_USER = "root"
	DB_PASS = "abc123!"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_NAME = "Project"
)

func MySqlORMSetting() {
	formatStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(formatStr, DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gorm connection error: " + err.Error())
	}

	if err := db.AutoMigrate(new(models.User)); err != nil {
		panic("gorm migration error: " + err.Error())
	}

}
