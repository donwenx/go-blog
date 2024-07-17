package router

import (
	"blog/controllers"
	err_code "blog/errcode"
	"blog/modules"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("token")
	if tokenStr == "" {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, "user not login")
		return
	}

	token, err := modules.GetTokenInfo(tokenStr)
	if err != nil {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, err.Error())
		return
	}

	if token.Expire < time.Now().Unix() {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		return
	}

	ctx.Next()
}

func Router() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
		user.POST("/logout", ValidateToken, controllers.UserController{}.LogOut)
	}
	return router
}
