package logic

import (
	"log"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/vo"
)

func QueryCourseRecord(req *vo.CourseRecordReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	record := new(models.ScoreRecord)
	record.CourseID = req.CourseID
	recordList, total, err := models.GetScoreRecord(record, req.PageRequest)
	if err != nil {
		log.Printf("Record service get score record failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	pageResponse := new(helper.PageResponse)
	pageResponse.Data = recordList
	pageResponse.Total = total
	baseResponse.Data = pageResponse
	return
}

func QueryBlogRecord(req *vo.PracticeRecordReq) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	record := new(models.PracticeRecord)
	record.BlogID = req.BlogID
	record.QueID = req.QueID
	recordList, total, err := models.GetPracticeRecord(record, req.PageRequest)
	if err != nil {
		log.Printf("Record service get practice record failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	pageResponse := new(helper.PageResponse)
	pageResponse.Data = recordList
	pageResponse.Total = total
	baseResponse.Data = pageResponse
	return
}
