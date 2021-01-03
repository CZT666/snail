package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"strconv"
)

func SelectQuestionMiddleware(idKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		queID, err := getQueID(ctx, idKey)
		if err != nil {
			log.Printf("Question operation middle ware get course id failed: %v\n", err)
			ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
			ctx.Abort()
			return
		}
		org, _ := ctx.Get("user")
		user, err := models.GetToken(org)
		if err != nil {
			log.Printf("Get token failed: %v\n", err)
			ctx.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
			ctx.Abort()
			return
		}
		que := new(models.SelectProblem)
		que.ID = queID
		que.CreateBy = user.GetIdentity()
		if err := models.GetSingleSelectProblem(que); err != nil {
			log.Printf("No access right")
			ctx.JSON(http.StatusOK, vo.BadResponse(vo.Error))
			ctx.Abort()
			return
		}
		ctx.Next()
		return
	}
}

func getQueID(ctx *gin.Context, idKey string) (queID int, err error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("Question operation middle ware get body failed: %v\n", err)
		return
	}
	id, err := utils.GetJsonValue(body, idKey)
	if err != nil {
		log.Printf("Question operation middle ware parse json failed: %v\n", err)
		return
	}
	queID, err = strconv.Atoi(id)
	if err != nil {
		log.Printf("Question operation middle ware parse json failed: %v\n", err)
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
