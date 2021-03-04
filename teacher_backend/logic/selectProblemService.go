package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/mirrors/go-xls"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
	"log"
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"strconv"
	"strings"
)

type choicesHelper struct {
	Data []string
}

func AddSelectProblem(req *vo.AddSelectProblemReq, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	req.SelectProblem.CreateBy = user.GetIdentity()
	if err := models.CreateSelectProblem(req.SelectProblem); err != nil {
		log.Printf("Select problem serice create selece problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加题目失败"
		return
	}
	if !(req.BlogID > 0) {
		return
	}
	err := models.AppendQueSetSelectProblem(req.BlogID, req.SelectProblem.ID)
	if err != nil {
		log.Printf("Select problem service append queSet select problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "题目添加成功，但未添加至题集，请尝试手动添加"
	}
	key := fmt.Sprintf(BlogSelectProblem,cast.ToString(req.BlogID))
	if err = dao.RedisDB.Del(context.Background(),key).Err();err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除 redis 失败"
	}
	return
}

func AppendSelectProblem(req *vo.AppendSelectProblemReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.AppendQueSetSelectProblem(req.BlogId, req.QueId); err != nil {
		log.Printf("Select problem service append single select problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "增加题目失败"
	}
	key := fmt.Sprintf(BlogSelectProblem,cast.ToString(req.BlogId))
	if err := dao.RedisDB.Del(context.Background(),key).Err();err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除 redis 失败"
	}
	return
}

func AddSelectProblemBatch(req *vo.AddSelectProblemBatchReq, filePath string, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	fileType, err := utils.GetFileType(req.File.Filename)
	if err != nil {
		log.Printf("Select problem service get file type failed: %v\n", err)
		baseResponse.Code = vo.FileError
		return
	}
	var queList []*models.SelectProblem
	switch fileType {
	case "xlsx":
		queList = getQueFromXlsx(filePath, user)
	//case "xls":
	//	queList = getQueFromXLS(filePath, user)
	default:
		baseResponse.Code = vo.FileError
		baseResponse.Msg = "仅支持xlsx文件"
		return
	}
	for index, que := range queList {
		if err := models.CreateSelectProblem(que); err != nil {
			log.Printf("Select problem serice create selece problem failed: %v, index: %v\n", err, index)
			continue
		}
		if req.BlogID > 0 {
			if err := models.AppendQueSetSelectProblem(req.BlogID, que.ID); err != nil {
				log.Printf("Select problem service append single select problem failed: %v\n", err)
			}
		}
	}
	err = utils.DeleteFile(filePath)
	if err != nil {
		log.Printf("Delete file failed: %v\n", err)
	}
	key := fmt.Sprintf(BlogSelectProblem,cast.ToString(req.BlogID))
	if err = dao.RedisDB.Del(context.Background(),key).Err();err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除 redis 失败"
	}
	return
}

func getQueFromXlsx(filePath string, user helper.User) (queList []*models.SelectProblem) {
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
			que := new(models.SelectProblem)
			length := len(row.Cells)
			for row.Cells[length-1].String() == "" {
				length -= 1
			}
			var choices []string
			for index, cell := range row.Cells[:length] {
				switch index {
				case 0:
					que.Description = cell.String()
				case length - 4:
					que.Answer = cell.String()
				case length - 3:
					que.Score, _ = strconv.Atoi(cell.String())
				case length - 2:
					que.Type, _ = strconv.Atoi(cell.String())
				case length - 1:
					que.CategoryID = getCategoryId(cell.String(), user.GetIdentity())
				default:
					choices = append(choices, cell.String())
				}
			}
			tmp := &choicesHelper{
				choices,
			}
			jsonBytes, err := json.Marshal(tmp)
			if err != nil {
				log.Printf("Select problem service gen json failed: %v\n", err)
				return nil
			}
			que.Choices = string(jsonBytes)
			que.CreateBy = user.GetIdentity()
			queList = append(queList, que)
		}
	}
	return
}

func getCategoryId(categoryName string, createBy string) int {
	selectCategory := new(models.SelectCategory)
	selectCategory.CategoryName = categoryName
	selectCategory.CreateBy = createBy
	if err := models.GetSingleCategory(selectCategory); err == nil {
		return selectCategory.ID
	}
	if err := models.CreateSelectCategory(selectCategory); err == nil {
		return selectCategory.ID
	}
	log.Println("Select problem service get category id failed")
	return -1
}

// 他妈的狗东西写得什么鸡巴玩意儿，这种垃圾库也好意思开源
func getQueFromXLS(filePath string, user helper.User) (queList []*models.SelectProblem) {
	xlsFile, err := xls.Open(filePath, "utf-8")
	if err != nil {
		log.Printf("Open xls file failed: %v\n", err)
		return
	}
	for i := 0; i < xlsFile.NumSheets(); i++ {
		workBook := xlsFile.GetSheet(i)
		for j := 0; j < int(workBook.MaxRow); j++ {
			row := workBook.Row(j)
			que := new(models.SelectProblem)
			var choices []string
			for k := 0; k < row.LastCol(); k++ {
				switch k {
				case 0:
					que.Description = row.Col(k)
				case row.LastCol() - 3:
					que.Answer = row.Col(k)
				case row.LastCol() - 2:
					que.Score, _ = strconv.Atoi(row.Col(k))
				case row.LastCol() - 1:
					que.Type, _ = strconv.Atoi(row.Col(k))
				default:
					choices = append(choices, row.Col(k))
				}
			}
			tmp := &choicesHelper{
				choices,
			}
			jsonBytes, err := json.Marshal(tmp)
			if err != nil {
				log.Printf("Select problem service gen json failed: %v\n", err)
				return nil
			}
			que.Choices = string(jsonBytes)
			que.CreateBy = user.GetIdentity()
			queList = append(queList, que)
		}
	}
	return
}

func DeleteSelectProblem(req *vo.DeleteSelectProblemReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	que := new(models.SelectProblem)
	que.ID = req.QueID
	if err := models.DeleteSelectProblem(que); err != nil {
		log.Printf("Select problem service delete question failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	return
}

func DeleteSelectProblemFromSet(req *vo.DeleteSelectProblemFromSetReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	queSet := new(models.QueSet)
	queSet.BlogID = req.BlogID
	if err := models.GetSingleQueSet(queSet); err != nil {
		log.Printf("Select problem service get single queSet failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "题目集不存在"
		return
	}
	queList := strings.Split(queSet.SelectProblem, ",")
	var newList []string
	deleteID := strconv.Itoa(req.QueID)
	for _, queID := range queList {
		if deleteID != queID {
			newList = append(newList, queID)
		}
	}
	queSet.SelectProblem = utils.List2String(newList, ",")
	if err := models.UpdateQueSet(queSet); err != nil {
		log.Printf("Select problem service update queSet failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	key := fmt.Sprintf(BlogSelectProblem,cast.ToString(req.BlogID))
	if err := dao.RedisDB.Del(context.Background(),key).Err();err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除 redis 失败"
	}
	return
}

func UpdateSelectProblem(que *models.SelectProblem) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.UpdateSelectProblem(que); err != nil {
		log.Printf("Select problem service update select problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "更新失败"
	}
	return
}

func QuerySelectProblemList(request *helper.PageRequest, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	que := new(models.SelectProblem)
	que.CreateBy = user.GetIdentity()
	queList, total, err := models.GetSelectProblem(que, request)
	if err != nil {
		log.Printf("Select problem service get problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	pageResponse := new(helper.PageResponse)
	pageResponse.Data = queList
	pageResponse.Total = total
	baseResponse.Data = pageResponse
	return
}

func QuerySelectProblemDetail(req *vo.ProblemDetailReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	que := new(models.SelectProblem)
	que.ID = req.QueID
	if err := models.GetSingleSelectProblem(que); err != nil {
		log.Printf("Select problem service get single problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	baseResponse.Data = que
	return
}

func QuerySelectProblemCategory(user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	categoryList, err := models.GetSelectCategories(user.GetIdentity())
	if err != nil {
		log.Printf("Select problem service get select cateories failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	baseResponse.Data = categoryList
	return
}

func FindSelectProblem(req *vo.FindSelectReq, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	problemList, err := models.FindSelectProblem(req.KeyWord, req.CategoryID, user.GetIdentity())
	if err != nil {
		log.Printf("Select problem service find select problem failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	baseResponse.Data = problemList
	return
}
