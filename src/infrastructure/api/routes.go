package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"itmrchow/go-project/user/docs"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/infrastructure/middleware"
	"itmrchow/go-project/user/src/usecase"
)

var router = gin.Default()

func Run() {

	dbHandler, _ := database.NewMySqlHandler()
	router.Use(middleware.DBTransactionMiddleware(dbHandler.DB))
	router.Use(middleware.ErrorHandle())

	getRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

func getRoutes() {
	apiV1 := router.Group("api/v1/")

	docs.SwaggerInfo.BasePath = apiV1.BasePath()

	//
	// Example
	addExampleRoutes(apiV1)

	// user
	addUserRoutes(apiV1)

	// wallet
	addWalletRoutes(apiV1)

}

func GetAuthUser(c *gin.Context) *reqdto.AuthUser {
	authUserInfo, isExists := c.Get("AuthUser")
	if !isExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": usecase.ErrUnauthorized})
	}

	authUser, ok := authUserInfo.(reqdto.AuthUser)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": usecase.ErrUnauthorized})
	}

	return &authUser
}
