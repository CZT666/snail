package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
	"student_bakcend/middleware"
)

func blogGroup(engine *gin.Engine) {
	courseID := "course_id"
	group := engine.Group("/blog")
	group.Use(middleware.JWTAuthMiddleware())
	{
		//搜索博客、博客列表推荐、博客详情
		group.GET("/list", middleware.CourseOperationMiddleware(courseID, true), controller.QueryBlogList)
		group.GET("/detail",  middleware.CourseOperationMiddleware(courseID, true),controller.QueryBlogDetail)
		//group.GET("/search/:name",controller.SearchBlog)
	}
}
