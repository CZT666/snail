package models

import (
	"snail/student_bakcend/dao"
	"time"
)

type Question struct {
	ID         int       `json:"id"`
	CourseID   int       `json:"course_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	CreateBy   string    `json:"create_by"`
}

type Answer struct {
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
	AnswerTime time.Time `json:"answer_time"`
	CreateBy   string    `json:"create_by"`
}

func AddQuestion(que *Question) (err error) {
	err = dao.DB.Create(&que).Error
	return
}

func GetQuestion(que *Question) (queList []*Question, err error) {

	if err := dao.DB.Where(&que).Find(&queList).Error; err != nil {
		return nil,  err
	}
	return
}

func GetSingleQuestion(que *Question) (err error) {
	err = dao.DB.Where(&que).First(&que).Error
	return
}

func GetSearchQuestion(searchName string, courseID string) (questionList []Question,  err error) {
	if err = dao.DB.Where("title like ? AND course_id = ?", "%"+searchName+"%", courseID).Find(&questionList).Error; err != nil {
		return nil, err
	}
	return
}

func AddAnswer(ans *Answer) (err error) {
	err = dao.DB.Create(&ans).Error
	return
}

func GetAnswer(ans *Answer) (ansList []*Answer,err error) {
	if err := dao.DB.Where(&ans).Find(&ansList).Error; err != nil {
		return nil, err
	}
	return
}
