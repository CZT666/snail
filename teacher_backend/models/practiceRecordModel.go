package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
	"time"
)

type PracticeRecord struct {
	ID           int        `json:"id"`
	BlogID       int        `json:"blog_id"`
	StuID        string     `json:"stu_id"`
	QueID        int        `json:"que_id"`
	FinishTime   *time.Time `json:"finish_time"`
	PracticeTime int        `json:"practice_time"`
	Status       int        `json:"status"`
}

func GetPracticeRecord(record *PracticeRecord, request *helper.PageRequest) (recordList []*PracticeRecord, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	var list []*PracticeRecord
	if err := dao.DB.Where(&record).Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	getPracticeRecordSubList(&list, &recordList, page, pageSize)
	return
}

func getPracticeRecordSubList(org *[]*PracticeRecord, dst *[]*PracticeRecord, page int, pageSize int) {
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
