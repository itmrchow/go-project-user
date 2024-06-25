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
	ViperKey          = "mysql"
	ViperIsShowLogKey = "mysql.isShowLog"
	DnsStr            = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Password  string
	Name      string
	IsShowLog bool
}

type MysqlHandler struct {
	DB     *gorm.DB
	config DbConfig
}

func NewMySqlHandler() (*MysqlHandler, error) {

	handler := MysqlHandler{}
	connectErr := handler.Connect()

	if connectErr != nil {
		return nil, connectErr
	}

	migrateErr := handler.Migrate()
	if migrateErr != nil {
		return nil, migrateErr
	}

	return &handler, nil
}

func (h *MysqlHandler) Connect() error {

	if err := viper.UnmarshalKey(ViperKey, &h.config); err != nil {
		panic("read config error: " + err.Error())
	}

	h.config.IsShowLog = viper.GetBool(ViperIsShowLogKey)

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
	)
}

func getLogger(isShowLog bool) logger.Interface {
	if isShowLog {
		return logger.Default.LogMode(logger.Info)
	} else {
		return logger.Default.LogMode(logger.Silent)
	}
}
