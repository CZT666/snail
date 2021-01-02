package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
)

func blogGroup(engine *gin.Engine) {
	group := engine.Group("/blog")
	//group.Use(middleware.JWTAuthMiddleware())
	{
		//搜索博客、博客列表推荐、博客详情
		group.GET("/list", controller.QueryBlogList)
		group.GET("/detail", controller.QueryBlogDetail)
		group.GET("/search/:name",controller.SearchBlog)
	}
}
