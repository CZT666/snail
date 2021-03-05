package models

import (
	"snail/student_bakcend/dao"
	"time"
)

type PracticeRecord struct {
	ID           int        `json:"id"`
	BlogID       int        `json:"blog_id"`
	StuID        string     `json:"stu_id"`
	QueID        int        `json:"que_id"`
	FinishTime   time.Time `json:"finish_time"`
	PracticeTime int        `json:"practice_time"`
	Status       int        `json:"status"`
}

func GetSinglePracticeRecord(pra *PracticeRecord)(err error){
	err = dao.DB.Where(&pra).First(&pra).Error
	return
}

func AddPracticeRecord(pra *PracticeRecord) (err error) {
	err = dao.DB.Create(&pra).Error
	return
}

func UpdatePracticeRecord(pra *PracticeRecord)(err error){
	err = dao.DB.Exec("update practice_records set finish_time=?,practice_time=?,status=? where id=?",pra.FinishTime,pra.PracticeTime,pra.Status,pra.ID).Error
	return
}


