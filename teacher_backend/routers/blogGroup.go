package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func blogGroup(engine *gin.Engine) {
	courseID := "course_id"
	group := engine.Group("/blog")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/add", middleware.CourseOperationMiddleware(courseID, false), controller.AddBlog)
		group.POST("/list", middleware.CourseOperationMiddleware(courseID, true), controller.QueryBlogList)
		group.POST("/detail", middleware.CourseOperationMiddleware(courseID, true), controller.QueryBlogDetail)
		group.POST("/update", middleware.CourseOperationMiddleware(courseID, false), controller.UpdateBlog)
		group.POST("/delete", middleware.CourseOperationMiddleware(courseID, false), controller.DeleteBlog)
	}
}
