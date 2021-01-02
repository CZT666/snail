package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func problemGroup(engine *gin.Engine) {
	courseID := "course_id"
	queID := "que_id"
	groupSelect := engine.Group("/selectProblem")
	groupSelect.Use(middleware.JWTAuthMiddleware())
	{
		groupSelect.POST("/add", controller.AddSelectProblem)
		groupSelect.POST("/append", middleware.CourseOperationMiddleware(courseID, false), controller.AppendSelectProblem)
		groupSelect.POST("/addAll", controller.AddSelectProblemBatch)
		groupSelect.POST("/delete", middleware.SelectQuestionMiddleware(queID), controller.DeleteSelectProblem)
		groupSelect.POST("/deleteFromSet", middleware.CourseOperationMiddleware(courseID, false), controller.DeleteSelectProblemFromSet)
		groupSelect.GET("/list", controller.QuerySelectProblemList)
		groupSelect.GET("/detail", controller.QuerySelectProblemDetail)
		groupSelect.GET("/category", controller.QuerySelectProblemCategory)
		groupSelect.GET("/find", controller.FindSelectProblem)
	}
}
