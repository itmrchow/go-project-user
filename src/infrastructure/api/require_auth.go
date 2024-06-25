package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
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
	token, parseErr := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			key := []byte(viper.GetString("privatekey"))
			return key, nil
		})

	if parseErr != nil {
		log.Println(parseErr)
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			respdto.ApiErrorResp{
				Title:  "Unauthorized",
				Detail: parseErr.Error(),
			},
		)
		return
	}

	checkClaimsErr := CheckClaims(token.Claims)
	if checkClaimsErr != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			respdto.ApiErrorResp{
				Title:  "Unauthorized",
				Detail: checkClaimsErr.Error(),
			},
		)
		return
	}

	c.Set("AuthUser", getAuthUser((token.Claims).(jwt.MapClaims)))

	c.Next()
}

func CheckClaims(claims jwt.Claims) error {
	if claims, ok := claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return errors.New("token has expired")
		}

	} else {
		return errors.New("claims is not mapclaims")
	}

	return nil
}

func getAuthUser(claims jwt.MapClaims) reqdto.AuthUser {

	return reqdto.AuthUser{
		Id:       claims["id"].(string),
		UserName: claims["userName"].(string),
		Account:  claims["account"].(string),
		Email:    claims["email"].(string),
		Phone:    claims["phone"].(string),
	}
}
