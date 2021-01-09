package vo

import "snail/teacher_backend/models/helper"

type CourseRecordReq struct {
	CourseID    int                 `json:"course_id"`
	PageRequest *helper.PageRequest `json:"page_request"`
}

type PracticeRecordReq struct {
	CourseID    int                 `json:"course_id"`
	BlogID      int                 `json:"blog_id"`
	QueID       int                 `json:"que_id"`
	PageRequest *helper.PageRequest `json:"page_request"`
}
