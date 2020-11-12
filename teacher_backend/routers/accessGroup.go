package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
)

func accessGroup(engine *gin.Engine) {
	engine.POST("/register", controller.TeacherRegister)
	engine.GET("/login", controller.TeacherLogin)
	engine.GET("/resetPwd/:mail", controller.ResetPwdReq)
	engine.POST("/resetPwd", controller.UpdatePwd)
}
