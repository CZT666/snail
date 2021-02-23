package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
)

func problemGroup(engine *gin.Engine) {
	groupProblem := engine.Group("/problem")
	//groupProblem.Use(middleware.JWTAuthMiddleware())
	{
		groupProblem.GET("/getProblem/:blogID", controller.GetProblem)
	}

	groupSelect := engine.Group("/selectProblem")
	//groupSelect.Use(middleware.JWTAuthMiddleware())
	{
		groupSelect.GET("/getSelect/:blogID", controller.GetSelect)
		groupSelect.POST("/getSelectScore",controller.GetSelectScore)
	}
	groupCode := engine.Group("/codeProblem")
	//groupCode.Use(middleware.JWTAuthMiddleware())
	{
		groupCode.GET("/getCode/:blogID", controller.GetCode)
	}
}
