package router

import (
	"PersonBlog/controller"
	"PersonBlog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		//auth.POST("/query", controller.QuertyUserInfo)
		auth.POST("/create_post", controller.CreatPost)
		auth.GET("/get_post_list", controller.GetPostList)
		auth.GET("/get_post_info", controller.GetPostInfo)
		auth.PUT("/update_post", controller.UpdatePost)
		auth.DELETE("/delete_post", controller.DeletePost)

		auth.POST("/make_comment", controller.CreateComment)
		auth.GET("/get_comment_list", controller.GetComment)
	}
	return r
}
