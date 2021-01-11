package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/vo"
)

func QueryCourseRecord(ctx *gin.Context) {
	req := new(vo.CourseRecordReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Query course record bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QueryCourseRecord(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func QueryPracticeRecord(ctx *gin.Context) {
	req := new(vo.PracticeRecordReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Query blog record bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QueryBlogRecord(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}
