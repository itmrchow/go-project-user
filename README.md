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

# Go Style Decisions
https://google.github.io/styleguide/go/decisions.html

# Clean Architecture

1. Entities - 核心邏輯
   1. domain
2. Use Cases - 商業邏輯
   1. usecase
   2. repo - 資料操作介面
3. Interface Adapters - 對外界面來呼叫Use case
   1. api
      1. context (API response format)
      2. controller (呼叫Usecase)
   2. db
      1. repo_impl
4. Frameworks and Drivers - 框架，資料庫等等的把程式串起來的東西
   1. router
   2. db_handler (資料庫連線)

# Todo
- [ ] swagger
- [ ] wire
- [ ] sturct to sturct
- [ ] UUID
- [ ] response format
- [ ] error handle


# 參考
https://ithelp.ithome.com.tw/users/20120647/ironman/3110
https://dongstudio.medium.com/clean-architecture-%E4%BA%8C-%E6%95%B4%E6%BD%94%E5%BC%8F%E6%9E%B6%E6%A7%8B-be4010ee62d4
https://github.com/bxcodec/go-clean-arch?tab=readme-ov-file
https://dev.to/michinoins/building-a-crud-app-with-mysql-gorm-echo-and-clean-architecture-in-go-h6d
