package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
)

func AddCodeProblem(ctx *gin.Context) {
	req := new(vo.AddCodeProblemReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Add code problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("CodeProblem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.AddCodeProblem(req, user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func AddCheckPoint(ctx *gin.Context) {
	checkPoint := new(models.CheckPoint)
	if err := ctx.ShouldBindBodyWith(&checkPoint, binding.JSON); err != nil {
		log.Printf("Add check point bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.AddCheckPoint(checkPoint)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func AddCheckPointBatch(ctx *gin.Context) {
	req := new(vo.AddCheckPointBatchReq)
	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Add check point bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	saveFilePath, err := utils.GenFilePath(req.File.Filename)
	if err != nil {
		log.Printf("Add check point gen file path failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	err = ctx.SaveUploadedFile(req.File, saveFilePath)
	if err != nil {
		log.Printf("Add check point save path failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.AddCheckPointBatch(req, saveFilePath)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func UpdateCheckPoint(ctx *gin.Context) {
	checkPoint := new(models.CheckPoint)
	if err := ctx.ShouldBindBodyWith(&checkPoint, binding.JSON); err != nil {
		log.Printf("Update check point bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.UpdateCheckPoint(checkPoint)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteCheckPoint(ctx *gin.Context) {
	checkPoint := new(models.CheckPoint)
	if err := ctx.ShouldBindBodyWith(&checkPoint, binding.JSON); err != nil {
		log.Printf("Delete check point bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteCheckPoint(checkPoint)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func UpdateCodeProblem(ctx *gin.Context) {
	codeProblem := new(models.CodeProblem)
	if err := ctx.ShouldBindBodyWith(&codeProblem, binding.JSON); err != nil {
		log.Printf("Update code problem failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.UpdateCodeProblem(codeProblem)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteCodeProblem(ctx *gin.Context) {
	codeProblem := new(models.CodeProblem)
	if err := ctx.ShouldBindBodyWith(&codeProblem, binding.JSON); err != nil {
		log.Printf("Delete code problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteCodeProblem(codeProblem)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func AppendCodeProblem(ctx *gin.Context) {
	req := new(vo.AppendCodeProblemReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Append Code Problem bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.AppendCodeProblem(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteCodeProblemFromSet(ctx *gin.Context) {
	req := new(vo.DeleteCodeProblemFromSetReq)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Delete code problem from set bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteCodeProblemFromSet(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func QueryCodeCategories(ctx *gin.Context) {
	org, _ := ctx.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("CodeProblem controller get token failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.QueryCodeCategories(user)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}
