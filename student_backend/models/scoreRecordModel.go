package models

import "snail/student_bakcend/dao"

type ScoreRecord struct {
	ID       int    `json:"id"`
	StuID    string `json:"stu_id"`
	CourseID int    `json:"course_id"`
	BlogID	 int    `json:"blog_id"`
	Score    int    `json:"score"`
}

func GetSingleScoreRecord(sco *ScoreRecord)(err error){
	err = dao.DB.Where(&sco).First(&sco).Error
	return
}

func AddScoreRecord(sco *ScoreRecord) (err error) {
	err = dao.DB.Create(&sco).Error
	return
}

func UpdateScoreRecord(sco *ScoreRecord)(err error){
	err = dao.DB.Exec("update score_records set score=? where id=?",sco.Score,sco.ID).Error
	return
}