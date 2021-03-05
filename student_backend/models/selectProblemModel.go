package models

import (
	"snail/student_bakcend/dao"
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

func GetSingleSelectProblem(problem *SelectProblem) (err error) {
	err = dao.DB.Where(&problem).First(&problem).Error
	return
}