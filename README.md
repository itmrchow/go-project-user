# 初始化
go 初始化
```
go mod init itmrchow/go-project/user
```

# 安裝
```
gin
go get -u github.com/gin-gonic/gin

MySQL 安裝
go get -u github.com/go-sql-driver/mysql

ORM 安裝
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

viper
go get github.com/spf13/viper

```

# goi18n 指令
```
goi18n extract -sourceLanguage en
goi18n merge active.en.toml active.es.toml //產生translate file
更新translate file後 , 修改檔名active.es.toml

```