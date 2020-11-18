package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func AssistanceGroup(engine *gin.Engine) {
	group := engine.Group("/assistance")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/add", middleware.CourseOperationMiddleware(), controller.AddAssistance)
	}
}
