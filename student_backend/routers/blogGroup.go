package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
	"student_bakcend/middleware"
)

func blogGroup(engine *gin.Engine) {
	group := engine.Group("/blog")
	group.Use(middleware.JWTAuthMiddleware())
	{
		//搜索博客、博客列表推荐、博客详情
		group.POST("/list", controller.QueryBlogList)
		group.POST("/detail",controller.QueryBlogDetail)
		//group.GET("/search/:name",controller.SearchBlog)
	}
}
