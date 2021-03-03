package logic

import (
	"github.com/tealeg/xlsx"
	"log"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"strconv"
	"strings"
)

func AddCodeProblem(req *vo.AddCodeProblemReq, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	req.CodeProblem.CreateBy = user.GetIdentity()
	if err := models.CreateCodeProblem(req.CodeProblem); err != nil {
		log.Printf("Code problem service create problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
		return
	}
	if !(req.BlogID > 0) {
		return
	}
	if err := models.AppendQueSetCodeProblem(req.BlogID, req.CodeProblem.ID); err != nil {
		log.Printf("Code problem service append queSet code problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "题目添加成功，但未添加至题集，请尝试手动添加"
	}
	baseResponse.Data = req.CodeProblem
	return
}

func AddCheckPoint(point *models.CheckPoint) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.CreateCheckpoint(point); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
	}
	return
}

func AddCheckPointBatch(req *vo.AddCheckPointBatchReq, savePath string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	fileType, err := utils.GetFileType(req.File.Filename)
	if err != nil {
		log.Printf("Check point service get file type failed: %v\n", err)
		baseResponse.Code = vo.FileError
		return
	}
	var checkPointList []*models.CheckPoint
	switch fileType {
	case "xlsx":
		checkPointList = getCheckPointFromXlsx(savePath)
	default:
		baseResponse.Code = vo.FileError
		baseResponse.Msg = "仅支持xlsx文件"
		return
	}
	for index, checkPoint := range checkPointList {
		checkPoint.ProblemId = req.QueID
		if err := models.CreateCheckpoint(checkPoint); err != nil {
			log.Printf("Add checkPoint failed: %v | index: %v\n", err, index)
		}
	}
	err = utils.DeleteFile(savePath)
	if err != nil {
		log.Printf("Delete file failed: %v\n", err)
	}
	return
}

func getCheckPointFromXlsx(filePath string) (checkPointList []*models.CheckPoint) {
	xlsxFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Printf("Open xlsx file failed: %v\n", err)
		return nil
	}
	for _, sheet := range xlsxFile.Sheets {
		for index, row := range sheet.Rows {
			if index == 0 {
				continue
			}
			checkPoint := new(models.CheckPoint)
			for index, cell := range row.Cells {
				switch index {
				case 0:
					checkPoint.Input = cell.String()
				case 1:
					checkPoint.Output = cell.String()
				case 2:
					checkPoint.Score, _ = strconv.Atoi(cell.String())
				}
			}
			checkPointList = append(checkPointList, checkPoint)
		}
	}
	return
}

func UpdateCheckPoint(point *models.CheckPoint) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.UpdateCheckpoint(point); err != nil {
		log.Printf("Code problem service update check point failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "更新失败"
	}
	return
}

func DeleteCheckPoint(point *models.CheckPoint) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.DeleteCheckPoint(point); err != nil {
		log.Printf("Code problem service delete check point failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	return
}

func UpdateCodeProblem(problem *models.CodeProblem) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.UpdateCodeProblem(problem); err != nil {
		log.Printf("Code problem service update code problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "更新失败"
	}
	return
}

func DeleteCodeProblem(problem *models.CodeProblem) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.DeleteCodeProblem(problem); err != nil {
		log.Printf("Code problem service delete code problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	return
}

func AppendCodeProblem(req *vo.AppendCodeProblemReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.AppendQueSetCodeProblem(req.BlogID, req.QueId); err != nil {
		log.Printf("Select problem service append single select problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "增加题目失败"
	}
	return
}

func DeleteCodeProblemFromSet(req *vo.DeleteCodeProblemFromSetReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	queSet := new(models.QueSet)
	queSet.BlogID = req.BlogID
	if err := models.GetSingleQueSet(queSet); err != nil {
		log.Printf("Code problem service get single queSet failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "题目集不存在"
		return
	}
	queList := strings.Split(queSet.CodeProblem, ",")
	var newList []string
	deleteID := strconv.Itoa(req.QueID)
	for _, queID := range queList {
		if deleteID != queID {
			newList = append(newList, queID)
		}
	}
	queSet.CodeProblem = utils.List2String(newList, ",")
	if err := models.UpdateQueSet(queSet); err != nil {
		log.Printf("Code problem service update queSet failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	return
}

func QueryCodeCategories(user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	categoryList, err := models.GetCodeCategory(user.GetIdentity())
	if err != nil {
		log.Printf("Code problem service query code category failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	baseResponse.Data = categoryList
	return
}

func QueryCodeProblemDetail(req *vo.ProblemDetailReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	que := new(models.CodeProblem)
	que.ID = req.QueID
	if err := models.GetSingleCodeProblem(que); err != nil {
		log.Printf("Select problem service get single problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	baseResponse.Data = que
	return
}
