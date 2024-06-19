package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/src/infrastructure/api/respdto"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 沒有錯就直接回傳
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0]
		var status int

		var errResp = respdto.ApiErrorResp{}
		// 轉換 API error & Domain error

		switch err.Type {
		case gin.ErrorTypePublic:
			status = http.StatusBadRequest
		case gin.ErrorTypeBind:
			status = http.StatusBadRequest
		case gin.ErrorTypeRender:
			status = http.StatusInternalServerError
		default:
			status = http.StatusInternalServerError
			errResp.Title = "Other Error"
			errResp.Detail = err.Error()
		}

		c.JSON(status, errResp)
	}
}
