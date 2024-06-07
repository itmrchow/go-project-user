package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"itmrchow/go-project/user/src/domain"
)

type Db_Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type DB_Handler struct {
	DB *gorm.DB
}

func NewSqlHandler() DB_Handler {
	var config Db_Config
	if err := viper.UnmarshalKey("mysql", &config); err != nil {
		panic("read config error: " + err.Error())
	}

	formatStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(formatStr, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gorm connection error: " + err.Error())
	}

	if err := db.AutoMigrate(new(domain.User)); err != nil {
		panic("gorm migration error: " + err.Error())
	}

	handler := new(DB_Handler)
	handler.DB = db

	return *handler
}
