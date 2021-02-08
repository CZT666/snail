package model

import "snail/judger/dao"

type CodeProblem struct {
	Id           int    `json:"id"`
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

func GetProblemById(id int) (problem *CodeProblem, err error) {
	if err = dao.DB.Where("id = ?", id).First(&problem).Error; err != nil {
		return nil, err
	}
	return
}
