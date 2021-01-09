package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/vo"
)

func QueryQueSet(ctx *gin.Context) {
	req := new(vo.QueSetQueryReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Query select set failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QueryQueSet(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}
