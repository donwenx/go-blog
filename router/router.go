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
		ctr := new(controllers.UserController)
		user.POST("/register", ctr.Register)
		user.POST("/login", ctr.Login)
		user.POST("/logout", middleware.ValidateToken, ctr.LogOut)
		user.POST("/update", middleware.ValidateToken, ctr.UpdateUser)
		user.GET(":id", middleware.ValidateToken, ctr.GetUserById)
		user.GET("/list", middleware.ValidateToken, ctr.GetUserList)
		user.DELETE(":id", middleware.ValidateToken, ctr.DeleteUser)
	}

	category := router.Group("/category")
	{
		ctr := new(controllers.CategoryController)
		category.POST("/create", middleware.ValidateToken, ctr.CreateCategory)
		category.POST("/update", middleware.ValidateToken, ctr.UpdateCategory)
		category.GET(":id", middleware.ValidateToken, ctr.GetCategoryById)
		category.GET("/list", middleware.ValidateToken, ctr.GetCateGoryList)
		category.DELETE(":id", middleware.ValidateToken, ctr.DeleteCategory)
	}

	article := router.Group("/article")
	{
		ctr := new(controllers.ArticleController)
		article.POST("/create", middleware.ValidateToken, ctr.CreateArticle)
		article.GET(":id", middleware.ValidateToken, ctr.GetArticleById)
		article.GET("/search", middleware.ValidateToken, ctr.GetArticleByKeyword)
		article.GET("/list", middleware.ValidateToken, ctr.GetArticleList)
		article.POST("/update", middleware.ValidateToken, ctr.UpdateArticle)
		article.DELETE(":id", middleware.ValidateToken, ctr.DeleteArticle)
	}

	comment := router.Group("/comment")
	{
		ctr := new(controllers.CommentController)
		comment.POST("/create", middleware.ValidateToken, ctr.CreateComment)
		comment.DELETE(":id", middleware.ValidateToken, ctr.DeleteComment)
	}

	tag := router.Group("/tag")
	{
		ctr := new(controllers.TagController)
		tag.POST("/create", middleware.ValidateToken, ctr.CreateTag)
		tag.GET("/list", middleware.ValidateToken, ctr.GetTagList)
		tag.POST("/update", middleware.ValidateToken, ctr.UpdateTag)
		tag.DELETE(":id", middleware.ValidateToken, ctr.DeleteTag)
	}

	upload := router.Group("/upload")
	{
		ctr := new(controllers.UploadController)
		upload.POST("/create", middleware.ValidateToken, ctr.CreateUpload)
		upload.GET("/:id", middleware.ValidateToken, ctr.GetUploadById)
		upload.GET("/download/:id", middleware.ValidateToken, ctr.Download)
		upload.DELETE("/:id", middleware.ValidateToken, ctr.DeleteUpload)
	}

	like := router.Group("/like")
	{
		ctr := new(controllers.LikeController)
		like.POST("/:type/:id", middleware.ValidateToken, ctr.CreateLike)
		like.DELETE("/:type/:id", middleware.ValidateToken, ctr.CancelLike)
	}

	return router
}
