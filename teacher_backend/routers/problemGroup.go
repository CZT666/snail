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
		groupSelect.POST("/detail", controller.QuerySelectProblemDetail)
		groupSelect.GET("/category", controller.QuerySelectProblemCategory)
		groupSelect.GET("/find", controller.FindSelectProblem)
		groupSelect.POST("/template", controller.DownloadTemplate)
	}
	groupCode := engine.Group("/codeProblem")
	groupCode.Use(middleware.JWTAuthMiddleware())
	{
		groupCode.POST("/add", controller.AddCodeProblem)
		groupCode.POST("/addCheckPoint", controller.AddCheckPoint)
		groupCode.POST("/addAllCheckPoint", controller.AddCheckPointBatch)
		groupCode.POST("/updateCheckPoint", controller.UpdateCheckPoint)
		groupCode.POST("deleteCheckPoint", controller.DeleteCheckPoint)
		groupCode.POST("/update", controller.UpdateCodeProblem)
		groupCode.POST("/detail", controller.FindCodeProblem)
		groupCode.POST("/delete", controller.DeleteCodeProblem)
		groupCode.POST("/append", controller.AppendCodeProblem)
		groupCode.POST("/deleteFromSet", controller.DeleteCodeProblemFromSet)
		groupCode.GET("/categories", controller.QueryCodeCategories)
		groupCode.POST("/template", controller.CheckPointTemplate)
	}
	groupSet := engine.Group("/queSet")
	groupSet.Use(middleware.JWTAuthMiddleware())
	{
		groupSet.POST("/query", controller.QueryQueSet)
	}
}
