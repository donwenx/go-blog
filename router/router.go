package router

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
		user.POST("/logout", middleware.ValidateToken, controllers.UserController{}.LogOut)
		user.POST("/update", middleware.ValidateToken, controllers.UserController{}.UpdateUser)
		user.GET(":id", middleware.ValidateToken, controllers.UserController{}.GetUserById)
		user.GET("/list", middleware.ValidateToken, controllers.UserController{}.GetUserList)
		user.DELETE(":id", middleware.ValidateToken, controllers.UserController{}.DeleteUser)
	}

	category := router.Group("/category")
	{
		category.POST("/create", middleware.ValidateToken, controllers.CategoryController{}.CreateCategory)
		category.POST("/update", middleware.ValidateToken, controllers.CategoryController{}.UpdateCategory)
		category.GET(":id", middleware.ValidateToken, controllers.CategoryController{}.GetCategoryById)
		category.GET("/list", middleware.ValidateToken, controllers.CategoryController{}.GetCateGoryList)
		category.DELETE(":id", middleware.ValidateToken, controllers.CategoryController{}.DeleteCategory)
	}
	return router
}
