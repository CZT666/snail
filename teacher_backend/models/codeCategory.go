package models

import "snail/teacher_backend/dao"

type CodeCategory struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CreateBy     string `json:"create_by"`
}

func CreateCodeCategory(category *CodeCategory) (err error) {
	err = dao.DB.Create(&category).Error
	return
}

func GetCodeCategory(createBy string) (categoryList []*CodeCategory, err error) {
	err = dao.DB.Where("create_by = ? or create_by = 'system'", createBy).Find(&categoryList).Error
	return
}

func DeleteCodeCategory(category *CodeCategory) (err error) {
	err = dao.DB.Delete(&category).Error
	return
}
