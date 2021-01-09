package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
)

type CheckPoint struct {
	ID        int    `json:"id"`
	ProblemId int    `json:"problem_id"`
	Score     int    `json:"score"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}

func CreateCheckpoint(point *CheckPoint) (err error) {
	err = dao.DB.Create(&point).Error
	return
}

func GetCheckPoint(point *CheckPoint, request *helper.PageRequest) (checkPointList []*CheckPoint, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	var list []*CheckPoint
	if err = dao.DB.Where(&point).Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	getCheckPointSubList(&list, &checkPointList, page, pageSize)
	return
}

func GetSingleCheckPoint(point *CheckPoint) (err error) {
	err = dao.DB.Where(&point).First(&point).Error
	return
}

func UpdateCheckpoint(point *CheckPoint) (err error) {
	err = dao.DB.Model(&CheckPoint{}).Where("id = ?", point.ID).Updates(&point).Error
	return
}

func DeleteCheckPoint(point *CheckPoint) (err error) {
	err = dao.DB.Delete(&point).Error
	return
}

func getCheckPointSubList(org *[]*CheckPoint, dst *[]*CheckPoint, page int, pageSize int) {
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
