package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
)

func registerGroup(engine *gin.Engine) {
	engine.POST("/teacher/register", controller.RegisterTeacher)
}
