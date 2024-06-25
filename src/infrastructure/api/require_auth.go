package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	"itmrchow/go-project/user/src/infrastructure/api/respdto"
)

func RequireAuth(c *gin.Context) {

	// get token
	tokenStr := strings.Split(c.GetHeader("Authorization"), " ")[1]

	if tokenStr == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			respdto.ApiErrorResp{
				Title:  "Unauthorized",
				Detail: "No token",
			},
		)
		return
	}

	// parse token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		key := []byte(viper.GetString("privatekey"))
		return key, nil
	})
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			respdto.ApiErrorResp{
				Title:  "Unauthorized",
				Detail: err.Error(),
			},
		)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	println("token:" + tokenStr)

	c.Next()
}
