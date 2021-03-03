package model

import "snail/judger/dao"

type CheckPoint struct {
	Id        int    `json:"id"`
	ProblemId int    `json:"problem_id"`
	Score     int    `json:"score"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}

func GetCheckPointByProblemId(id int) (checkPointList []*CheckPoint, err error) {
	if err = dao.DB.Where("problem_id = ?", id).Find(&checkPointList).Error; err != nil {
		return nil, err
	}
	return
}
