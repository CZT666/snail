package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func recordGroup(engine *gin.Engine) {
	courseID := "course_id"
	groupRecord := engine.Group("/record")
	groupRecord.Use(middleware.JWTAuthMiddleware())
	{
		groupRecord.GET("/course", middleware.CourseOperationMiddleware(courseID, false), controller.QueryCourseRecord)
		groupRecord.GET("/practice", middleware.CourseOperationMiddleware(courseID, false), controller.QueryPracticeRecord)
	}
}
