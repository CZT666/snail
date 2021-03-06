package routers

import (
	"github.com/gin-gonic/gin"
	"snail/student_bakcend/controller"
	"snail/student_bakcend/middleware"
)

func problemGroup(engine *gin.Engine) {
	groupProblem := engine.Group("/problem")
	groupProblem.Use(middleware.JWTAuthMiddleware())
	{
		groupProblem.GET("/getProblem/:blog_id", controller.GetProblem)
		groupProblem.GET("/getProblemScore/:blog_id",controller.GetProblemScore)
	}
	groupSelect := engine.Group("/selectProblem")
	groupSelect.Use(middleware.JWTAuthMiddleware())
	{
		groupSelect.GET("/getSelect/:blog_id", controller.GetSelect)
		groupSelect.POST("/SelectScore",controller.SelectScore)
	}
	groupCode := engine.Group("/codeProblem")
	groupCode.Use(middleware.JWTAuthMiddleware())
	{
		groupCode.GET("/getCode/:blog_id", controller.GetCode)
	}
}
