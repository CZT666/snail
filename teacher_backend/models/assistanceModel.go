package models

import (
	"snail/teacher_backend/dao"
	"time"
)

type Assistance struct {
	StuID       string    `json:"stu_id"`
	CourseID    int       `json:"id"`
	ExpiredTime time.Time `json:"expired_time"`
}

func CreateAssistance(assistance *Assistance) (err error) {
	err = dao.DB.Create(&assistance).Error
	return
}

func GetAssistance(assistance *Assistance) (assistanceList []Assistance, err error) {
	if err = dao.DB.Where(&assistance).Find(&assistanceList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteAssistance(assistance *Assistance) (err error) {
	err = dao.DB.Delete(&assistance).Error
	return
}
