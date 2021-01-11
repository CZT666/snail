package models

import (
	"fmt"
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
	err = dao.DB.Model(&QueSet{}).Where("blog_id = ?", set.BlogID).Updates(&set).Error
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
	fmt.Printf("queSet: %v\n", queSet)
	dao.DB.Where(&queSet).First(&queSet).Count(&count)
	fmt.Printf("queSet: %v\n", queSet)
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
		err = dao.DB.Model(&QueSet{}).Where("blog_id = ?", queSet.BlogID).Updates(&queSet).Error
	} else {
		builder.WriteString("0,")
		builder.WriteString(addPid)
		builder.WriteString(",")
		queSet.SelectProblem = builder.String()
		err = dao.DB.Create(&queSet).Error
	}
	return
}

func AppendQueSetCodeProblem(blogID int, pid int) (err error) {
	queSet := new(QueSet)
	queSet.BlogID = blogID
	count := 0
	dao.DB.Where(&queSet).First(&queSet).Count(&count)
	var builder strings.Builder
	addPid := strconv.Itoa(pid)
	if count != 0 {
		add := queSet.CodeProblem
		set := utils.String2Set(add, ",")
		if set.Contains(addPid) {
			return
		}
		builder.WriteString(add)
		builder.WriteString(addPid)
		builder.WriteString(",")
		queSet.CodeProblem = builder.String()
		err = dao.DB.Model(&QueSet{}).Where("blog_id = ?", queSet.BlogID).Updates(&queSet).Error
	} else {
		builder.WriteString("0,")
		builder.WriteString(addPid)
		builder.WriteString(",")
		queSet.CodeProblem = builder.String()
		err = dao.DB.Create(&queSet).Error
	}
	return
}
