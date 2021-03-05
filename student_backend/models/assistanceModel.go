package models

import (
	"snail/student_bakcend/dao"
	"time"
)

type Assistance struct {
	StuID       string    `json:"stu_id"`
	CourseID    int       `json:"course_id"`
	ExpiredTime time.Time `json:"expired_time"`
}

func GetAssistance(assistance *Assistance) (assistanceList []Assistance, err error) {
	if err = dao.DB.Where(&assistance).Find(&assistanceList).Error; err != nil {
		return nil, err
	}
	return
}

