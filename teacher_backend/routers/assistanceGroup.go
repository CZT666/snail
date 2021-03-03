package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func AssistanceGroup(engine *gin.Engine) {
	courseID := "course_id"
	group := engine.Group("/assistance")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("find", controller.FindStudent)
		group.POST("/add", middleware.CourseOperationMiddleware(courseID, false), controller.AddAssistance)
		group.POST("/delete", middleware.CourseOperationMiddleware(courseID, false), controller.DeleteAssistance)
		group.POST("/get", controller.GetAllAssistance)
	}
}
