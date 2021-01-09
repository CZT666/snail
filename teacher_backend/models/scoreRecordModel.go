package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
)

type ScoreRecord struct {
	ID       int    `json:"id"`
	StuID    string `json:"stu_id"`
	CourseID int    `json:"course_id"`
	Score    int    `json:"score"`
}

func GetScoreRecord(record *ScoreRecord, request *helper.PageRequest) (recordList []*ScoreRecord, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	var list []*ScoreRecord
	if err := dao.DB.Where(record).Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	getScoreRecordSubList(&list, &recordList, page, pageSize)
	return
}

func getScoreRecordSubList(org *[]*ScoreRecord, dst *[]*ScoreRecord, page int, pageSize int) {
	orgSize := len(*org)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > orgSize {
		return
	}
	if end > orgSize {
		end = orgSize
	}
	a := *org
	*dst = a[start:end]
}
