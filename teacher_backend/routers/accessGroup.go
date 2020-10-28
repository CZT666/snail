package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
)

func accessGroup(engine *gin.Engine) {
	engine.POST("/teacher/register", controller.TeacherRegister)
	engine.GET("/teacher/login", controller.TeacherLogin)
	engine.GET("/teacher/resetPwd/:mail", controller.ResetPwdReq)
	engine.POST("/teacher/resetPwd", controller.UpdatePwd)
}
