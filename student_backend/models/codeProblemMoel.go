package models

import (
	"student_bakcend/dao"
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

func GetSingleCodeProblem(problem *CodeProblem) (err error) {
	err = dao.DB.Where(&problem).First(&problem).Error
	return
}
