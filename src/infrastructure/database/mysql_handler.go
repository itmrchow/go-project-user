package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"itmrchow/go-project/user/src/domain"
)

const (
	DnsStr = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Password  string
	Name      string
	IsShowLog bool
}

var mysqlHandler MysqlHandler

type MysqlHandler struct {
	DB     *gorm.DB
	config DbConfig
}

func NewMySqlHandler() (*MysqlHandler, error) {
	if mysqlHandler.DB == nil {
		connectErr := mysqlHandler.Connect()

		if connectErr != nil {
			return nil, connectErr
		}
	}

	return &mysqlHandler, nil
}

func (h *MysqlHandler) Connect() error {

	h.config = DbConfig{
		Host:      viper.GetString("mysql.host"),
		Port:      viper.GetInt("mysql.port"),
		User:      viper.GetString("mysql.user"),
		Password:  viper.GetString("mysql.password"),
		Name:      viper.GetString("mysql.name"),
		IsShowLog: viper.GetBool("mysql.isShowLog"),
	}

	dsn := fmt.Sprintf(
		DnsStr,
		h.config.User,
		h.config.Password,
		h.config.Host,
		h.config.Port,
		h.config.Name,
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: getLogger(h.config.IsShowLog),
		},
	)

	if err != nil {
		return err
	} else {
		h.DB = db
		return nil
	}
}

func (m *MysqlHandler) Migrate() error {
	return m.DB.AutoMigrate(
		new(domain.User),
		new(domain.Wallet),
		new(domain.WalletRecord),
	)
}

func getLogger(isShowLog bool) logger.Interface {
	if isShowLog {
		return logger.Default.LogMode(logger.Info)
	} else {
		return logger.Default.LogMode(logger.Silent)
	}
}
