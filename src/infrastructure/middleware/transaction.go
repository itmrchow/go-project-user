package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()

		log.Println("beginning database transaction")

		defer func() {
			if r := recover(); r != nil { // recover 捕獲執行過程的異常
				txHandle.Rollback() // 有異常Rollback
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next() // 執行具體邏輯

		if lo.Contains([]int{http.StatusOK, http.StatusCreated}, c.Writer.Status()) { //200 or 201 commit
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			log.Print("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}

	}
}
