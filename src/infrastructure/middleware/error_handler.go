package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/src/infrastructure/api/respdto"
	"itmrchow/go-project/user/src/usecase"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 沒有錯就直接回傳
		if len(c.Errors) == 0 {
			return
		}

		ginErr := c.Errors[0]

		var status int
		var errResp = respdto.ApiErrorResp{}

		// bind handle
		if ginErr.Type == gin.ErrorTypeBind {
			status = http.StatusBadRequest

			errResp.Title = "Bind Error"
			errResp.Detail = ginErr.Error()

		} else {
			status, errResp = setErrResp(ginErr.Err)
		}

		c.JSON(status, errResp)
	}
}

func setErrResp(err error) (status int, errResp respdto.ApiErrorResp) {
	switch {
	case errors.Is(err, usecase.ErrUserAlreadyExists), errors.Is(err, usecase.ErrDataExists):
		status = http.StatusBadRequest
		errResp.Title = "Bind Error"
		errResp.Detail = err.Error()

	case errors.Is(err, usecase.ErrDbFail):
		status = http.StatusInternalServerError
		errResp.Title = "Internal Error"
		errResp.Detail = usecase.ErrDbFail.Error()

	case errors.Is(err, usecase.ErrDbInsertFail):
		status = http.StatusConflict
		errResp.Title = "Conflict Error"
		errResp.Detail = err.Error()

	case errors.Is(err, usecase.ErrUserNotExists):
		status = http.StatusNotFound
		errResp.Title = "Not Found"
		errResp.Detail = err.Error()

	case errors.Is(err, usecase.ErrUnauthorized):
		status = http.StatusUnauthorized
		errResp.Title = "Unauthorized"
		errResp.Detail = err.Error()

	case errors.Is(err, usecase.ErrPaymentRequired):
		status = http.StatusPaymentRequired
		errResp.Title = "Payment Required"
		errResp.Detail = err.Error()

	case errors.Is(err, usecase.ErrTimeOut):
		status = http.StatusGatewayTimeout
		errResp.Title = "Time Out"
		errResp.Detail = err.Error()

	default:
		status = http.StatusInternalServerError
		errResp.Title = "Other Error"
		errResp.Detail = err.Error()

		// TODO: Log
	}

	return status, errResp
}
