package router

import (
	"blog/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "test")
		})
		user.POST("/register", controllers.UserController{}.Register)
	}
	return router
}
