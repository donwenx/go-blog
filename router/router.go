package router

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
		user.POST("/logout", controllers.UserController{}.LogOut)
	}
	return router
}
