package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func courseGroup(engine *gin.Engine) {
	courseID := "id"
	group := engine.Group("/course")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/add", controller.AddCourse)
		group.GET("/list", controller.QueryCourseList)
		group.GET("/detail", middleware.CourseOperationMiddleware(courseID, true), controller.QueryCourseDetail)
		group.POST("/update", middleware.CourseOperationMiddleware(courseID, false), controller.UpdateCourse)
		group.POST("/delete", middleware.CourseOperationMiddleware(courseID, false), controller.DeleteCourse)
	}
}
