package models

import "snail/teacher_backend/dao"

type SelectCategory struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CreateBy     string `json:"create_by"`
}

func CreateSelectCategory(category *SelectCategory) (err error) {
	err = dao.DB.Create(&category).Error
	return
}

func GetSelectCategories(createBy string) (categories []*SelectCategory, err error) {
	if err = dao.DB.Where("create_by = ? or create_by = 'system'", createBy).Find(&categories).Error; err != nil {
		return nil, err
	}
	return
}

func GetSingleCategory(category *SelectCategory) (err error) {
	err = dao.DB.Where("create_by = ? or create_by = 'system'", category.CreateBy).First(&category).Error
	return
}

func DeleteSelectCategory(category *SelectCategory) (err error) {
	err = dao.DB.Delete(&category).Error
	return
}
