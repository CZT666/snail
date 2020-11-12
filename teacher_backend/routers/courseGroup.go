package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func courseGroup(engine *gin.Engine) {
	group := engine.Group("/course")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/add", controller.AddCourse)
		group.GET("/list", controller.QueryCourseList)
		group.GET("/detail", middleware.CourseOperationMiddleware(), controller.QueryCourseDetail)
		group.POST("/update", middleware.CourseOperationMiddleware(), controller.UpdateCourse)
		group.POST("/delete", middleware.CourseOperationMiddleware(), controller.DeleteCourse)
	}
}
