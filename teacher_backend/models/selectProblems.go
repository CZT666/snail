package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
)

type SelectProblem struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Choices     string `json:"choices"`
	Answer      string `json:"answer"`
	Score       int    `json:"score"`
	Type        int    `json:"type"`
	CategoryID  int    `json:"category_id"`
	CreateBy    string `json:"create_by"`
}

func CreateSelectProblem(problem *SelectProblem) (err error) {
	err = dao.DB.Create(&problem).Error
	return
}

func UpdateSelectProblem(problem *SelectProblem) (err error) {
	err = dao.DB.Model(&SelectProblem{}).Updates(&problem).Error
	return
}

func GetSelectProblem(problem *SelectProblem, pageRequest *helper.PageRequest) (selectProblemList []*SelectProblem, total int, err error) {
	page := pageRequest.Page
	pageSize := pageRequest.PageSize
	var totalList []*SelectProblem
	if err = dao.DB.Where(&problem).Find(&totalList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	getSubList(&totalList, &selectProblemList, page, pageSize)
	return
}

func GetSingleSelectProblem(problem *SelectProblem) (err error) {
	err = dao.DB.Where(&problem).First(&problem).Error
	return
}

func DeleteSelectProblem(problem *SelectProblem) (err error) {
	err = dao.DB.Delete(&problem).Error
	return
}

func FindSelectProblem(keyWord string, categoryID int, createBy string) (problemList []*SelectProblem, err error) {
	word := "%" + keyWord + "%"
	if categoryID > 0 {
		err = dao.DB.Where("category_id = ? and description like ? and create_by = ? or create_by = 'system'", categoryID, word, createBy).Find(&problemList).Error
	} else {
		err = dao.DB.Where("description like ? and create_by = ? or create_by = 'system'", word, createBy).Find(&problemList).Error
	}
	if err != nil {
		return nil, err
	}
	return
}

func getSubList(org *[]*SelectProblem, dst *[]*SelectProblem, page int, pageSize int) {
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
