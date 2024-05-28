package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/req_dto"
)

func addUserRoutes(rg *gin.RouterGroup) {
	// user API
	rg.GET("/user/:userId", getUser)
	// apiV1.GET("/users", getUser)
	rg.POST("/user", createUser)
	// apiV1.PUT("/user/:userId", getUser)
	// apiV1.PATCH("/user/:userId", getUser)
	// apiV1.DELETE("/user/:userId", getUser)
}

func getUser(c *gin.Context) {
	userId := c.Param("userId")

	// query by userId

	c.JSON(http.StatusOK, gin.H{"hello": userId})
}

func createUser(c *gin.Context) {
	userReq := new(req_dto.UserReq)
	c.BindJSON(&userReq)

	fmt.Println(userReq)
	// create user

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
