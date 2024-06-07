package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/src/infrastructure/api/req_dto"
	"itmrchow/go-project/user/src/infrastructure/database"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
)

func addUserRoutes(rg *gin.RouterGroup) {
	// user API
	rg.GET("/user/:userId", getUser)
	// apiV1.GET("/users", getUser)
	rg.POST("/user", createUser)
	// apiV1.PUT("/user/:userId", getUser)
	// apiV1.PATCH("/user/:userId", getUser)
	// apiV1.DELETE("/user/:userId", getUser)

	rg.POST("/login", loginUser)
}

func getUser(c *gin.Context) {
	userId := c.Param("userId")

	// query by userId

	c.JSON(http.StatusOK, gin.H{"hello": userId})
}

func createUser(c *gin.Context) {
	userController := controllers.NewUserController(database.NewSqlHandler())

	// context to dto
	userReq := new(req_dto.CreateUserReq) // bind bto
	c.BindJSON(&userReq)

	// call controller
	response := userController.CreateUser(userReq)

	c.JSON(http.StatusOK, response)
}

func loginUser(c *gin.Context) {
	loginReq := new(req_dto.LoginReq)
	// check account
	isAuth := AuthUser(loginReq.Account, loginReq.Password)

	if isAuth {
		c.JSON(http.StatusOK, gin.H{"token": "token"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "account or password is wrong"})
	}
}

func AuthUser(account string, password string) bool {
	return true
}
