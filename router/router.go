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
		ctx.Abort()
		return
	}

	token, err := modules.GetTokenInfo(tokenStr)
	if err != nil {
		controllers.ReturnError(ctx, err_code.ErrInvalidRequest, err.Error())
		ctx.Abort()
		return
	}

	if token.Expire < time.Now().Unix() {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		ctx.Abort()
		return
	}

	if token.State == 0 {
		controllers.ReturnError(ctx, err_code.ErrInvalidToken, "token expired")
		ctx.Abort()
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

	category := router.Group("/category")
	{
		category.POST("/create", ValidateToken, controllers.CategoryController{}.CreateCategory)
		category.POST("/update", ValidateToken, controllers.CategoryController{}.UpdateCategory)
		category.GET(":id", ValidateToken, controllers.CategoryController{}.GetCategoryById)
		category.GET("/list", ValidateToken, controllers.CategoryController{}.GetCateGoryList)
	}
	return router
}
