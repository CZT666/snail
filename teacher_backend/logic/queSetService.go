package logic

import (
	"log"
	"snail/teacher_backend/models"
	"snail/teacher_backend/vo"
	"strconv"
	"strings"
)

func QueryQueSet(req *vo.QueSetQueryReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	queSet := new(models.QueSet)
	queSet.BlogID = req.BlogID
	if err := models.GetSingleQueSet(queSet); err != nil {
		log.Printf("QueSet service get single queSet failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	var queList interface{}
	if req.Type == 0 {
		queList = getSelectProblemFromSet(queSet.SelectProblem)
	} else {
		queList = getCodeProblemFromSet(queSet.CodeProblem)
	}
	baseResponse.Data = queList
	return
}

func getSelectProblemFromSet(queListStr string) (res []*models.SelectProblem) {
	queList := strings.Split(queListStr, ",")
	for _, queId := range queList {
		que := new(models.SelectProblem)
		id, err := strconv.Atoi(queId)
		if err != nil {
			log.Printf("convert string to int failed: %v\n", err)
			continue
		}
		que.ID = id
		if err := models.GetSingleSelectProblem(que); err != nil {
			log.Printf("Que set service get single select problem failed: %v\n", err)
			continue
		}
		res = append(res, que)
	}
	return
}

func getCodeProblemFromSet(queListStr string) (res []*models.CodeProblem) {
	queList := strings.Split(queListStr, ",")
	for _, queId := range queList {
		que := new(models.CodeProblem)
		id, err := strconv.Atoi(queId)
		if err != nil {
			log.Printf("convert string to int failed: %v\n", err)
			continue
		}
		que.ID = id
		if err := models.GetSingleCodeProblem(que); err != nil {
			log.Printf("Que set service get single select problem failed: %v\n", err)
			continue
		}
		res = append(res, que)
	}
	return
}
