package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"itmrchow/go-project/user/config"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/interfaces/api/controllers"
)

func addUserRoutes(rg *gin.RouterGroup) {

	userController, _ := config.InitUserController()

	// user API
	rg.GET("/user/:userId", func(c *gin.Context) {
		getUser(c, userController)
	})
	rg.GET("/users", func(c *gin.Context) {
		getUsers(c, userController)
	})
	rg.POST("/user", func(c *gin.Context) {
		createUser(c, userController)
	})
	rg.PUT("/user/:userId", func(c *gin.Context) {
		putUser(c, userController)
	})
	rg.PATCH("/user/:userId", func(c *gin.Context) {
		patchUser(c, userController)
	})
	rg.DELETE("/user/:userId", func(c *gin.Context) {
		deleteUser(c, userController)
	})

	rg.POST("/login", loginUser)
}

// @Summary 刪除用戶 by Id
// @Produce json
// @Tags User
// @Param userId path string true "User Id"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Router /user/{userId} [delete]
func deleteUser(c *gin.Context, userController *controllers.UserController) {
	panic("unimplemented")
}

// @Summary 部分更新用戶 by Id
// @Produce json
// @Tags User
// @Param userId path string true "User Id"
// @Param body body reqdto.PatchUserReq true "Patch user sample"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Router /user/{userId} [patch]
func patchUser(c *gin.Context, userController *controllers.UserController) {
	panic("unimplemented")
}

// @Summary 完整更新用戶 by Id
// @Produce json
// @Tags User
// @Param userId path string true "User Id"
// @Param body body reqdto.PutUserReq true "Put user sample"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Router /user/{userId} [Put]
func putUser(c *gin.Context, userController *controllers.UserController) {
	panic("unimplemented")
}

// @Summary 查詢用戶
// @Produce json
// @Tags User
// @Parameters.QueryParams
// @Param userName query string false "User Name"
// @Param email query string false "User Email"
// @Param phone query string false "User Phone"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Router /users [GET]
func getUsers(c *gin.Context, controller *controllers.UserController) {

	panic("unimplemented")
}

// @Summary 查詢用戶 by Id
// @Produce json
// @Tags User
// @Param userId path string true "User Id"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Router /user/{userId} [get]
func getUser(c *gin.Context, controller *controllers.UserController) {
	userId := c.Param("userId")
	resp, _ := controller.GetUser(userId)

	c.JSON(http.StatusOK, resp)
}

// @Summary 建立用戶
// @Produce json
// @Tags User
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Param body body reqdto.CreateUserReq true "Create user sample"
// @Router /user [post]
func createUser(c *gin.Context, controller *controllers.UserController) {

	// context to dto
	userReq := new(reqdto.CreateUserReq) // bind bto
	c.BindJSON(&userReq)

	// call controller
	response := controller.CreateUser(userReq)

	c.JSON(http.StatusOK, response)
}

// @Summary 登入
// @Produce json
// @Tags User
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Router /login [post]
func loginUser(c *gin.Context) {
	loginReq := new(reqdto.LoginReq)
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
