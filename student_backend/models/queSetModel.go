package models

import (
	"snail/student_bakcend/dao"
)

type QueSet struct {
	BlogID        int    `json:"blog_id"`
	SelectProblem string `json:"select_problem"`
	CodeProblem   string `json:"code_problem"`
}

func GetSingleQueSet(set *QueSet) (err error) {
	err = dao.DB.Where(&set).First(&set).Error
	return
}
