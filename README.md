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

# Clean Architecture




# 參考
https://ithelp.ithome.com.tw/users/20120647/ironman/3110
https://dongstudio.medium.com/clean-architecture-%E4%BA%8C-%E6%95%B4%E6%BD%94%E5%BC%8F%E6%9E%B6%E6%A7%8B-be4010ee62d4