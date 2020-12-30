package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/utils"
	"strconv"
	"strings"
)

type QueSet struct {
	BlogID        int    `json:"blog_id"`
	SelectProblem string `json:"select_problem"`
	CodeProblem   string `json:"code_problem"`
}

func CreateQueSet(set *QueSet) (err error) {
	err = dao.DB.Create(&set).Error
	return
}

func UpdateQueSet(set *QueSet) (err error) {
	err = dao.DB.Model(&QueSet{}).Updates(&set).Error
	return
}

func GetQueSet(set *QueSet, pageRequest *helper.PageRequest) (queSetList []*QueSet, total int, err error) {
	page := pageRequest.Page
	pageSize := pageRequest.PageSize
	if err := dao.DB.Where(&set).Limit(pageSize).Offset((page - 1) * pageSize).Find(&queSetList).Count(&total).Error; err != nil {
		return nil, 0, nil
	}
	return
}

func GetSingleQueSet(set *QueSet) (err error) {
	err = dao.DB.Where(&set).First(&set).Error
	return
}

func DeleteQueSet(set *QueSet) (err error) {
	err = dao.DB.Delete(&set).Error
	return
}

func AppendQueSetSelectProblem(blogID int, pid int) (err error) {
	queSet := new(QueSet)
	queSet.BlogID = blogID
	count := 0
	dao.DB.Where(&queSet).First(&queSet).Count(&count)
	var builder strings.Builder
	addPid := strconv.Itoa(pid)
	if count != 0 {
		add := queSet.SelectProblem
		set := utils.String2Set(add, ",")
		if set.Contains(addPid) {
			return
		}
		builder.WriteString(add)
		builder.WriteString(addPid)
		builder.WriteString(",")
		queSet.SelectProblem = builder.String()
		err = dao.DB.Model(&QueSet{}).Updates(queSet).Error
	} else {
		builder.WriteString(strconv.Itoa(pid))
		builder.WriteString(",")
		queSet.SelectProblem = builder.String()
		err = dao.DB.Create(&queSet).Error
	}
	return
}
