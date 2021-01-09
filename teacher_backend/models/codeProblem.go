package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
)

type CodeProblem struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	TimeLimit    int    `json:"time_limit"`
	MemoryLimit  int    `json:"memory_limit"`
	Description  string `json:"description"`
	CategoryId   int    `json:"category_id"`
	InputFormat  string `json:"input_format"`
	OutputFormat string `json:"output_format"`
	SampleInput  string `json:"sample_input"`
	SampleOutput string `json:"sample_output"`
	CreateBy     string `json:"create_by"`
}

func CreateCodeProblem(problem *CodeProblem) (err error) {
	err = dao.DB.Create(&problem).Error
	return
}

func GetCodeProblem(problem *CodeProblem, request *helper.PageRequest) (problemList []*CodeProblem, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	var list []*CodeProblem
	if err = dao.DB.Where(&problem).Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	getCodeSubList(&list, &problemList, page, pageSize)
	return
}

func GetSingleCodeProblem(problem *CodeProblem) (err error) {
	err = dao.DB.Where(&problem).First(&problem).Error
	return
}

func UpdateCodeProblem(problem *CodeProblem) (err error) {
	err = dao.DB.Model(&CodeProblem{}).Where("id = ?", problem.ID).Updates(&problem).Error
	return
}

func DeleteCodeProblem(problem *CodeProblem) (err error) {
	err = dao.DB.Delete(&problem).Error
	return
}

func getCodeSubList(org *[]*CodeProblem, dst *[]*CodeProblem, page int, pageSize int) {
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
