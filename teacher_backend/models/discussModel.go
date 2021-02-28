package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
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

func GetQuestion(que *Question, request *helper.PageRequest) (queList []*Question, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	if err := dao.DB.Where(&que).Limit(pageSize).Offset((page - 1) * pageSize).Find(&queList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}

func GetSingleQuestion(que *Question) (err error) {
	err = dao.DB.Where(&que).First(&que).Error
	return
}

func GetSearchQuestion(pageRequest *helper.PageRequest, searchName string, courseID string) (questionList []Question, total int, err error) {
	page := pageRequest.Page
	pageSize := pageRequest.PageSize
	if err = dao.DB.Where("title like ? AND course_id = ?", "%"+searchName+"%", courseID).Limit(pageSize).Offset((page - 1) * pageSize).Find(&questionList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}

func AddAnswer(ans *Answer) (err error) {
	err = dao.DB.Create(&ans).Error
	return
}

func GetAnswer(ans *Answer, request *helper.PageRequest) (ansList []*Answer, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	if err := dao.DB.Where(&ans).Limit(pageSize).Offset((page - 1) * pageSize).Find(&ansList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}
