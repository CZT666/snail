package vo

import (
	"mime/multipart"
	"snail/teacher_backend/models"
)

type AddSelectProblemReq struct {
	BlogID        int                   `json:"blog_id"`
	CourseID      int                   `json:"course_id"`
	SelectProblem *models.SelectProblem `json:"select_problem"`
}

type AppendSelectProblemReq struct {
	BlogId   int `json:"blog_id"`
	QueId    int `json:"que_id"`
	CourseID int `json:"course_id"`
}

type AddSelectProblemBatchReq struct {
	CourseID int                   `form:"course_id"`
	BlogID   int                   `form:"blog_id"`
	File     *multipart.FileHeader `form:"file"`
}

type DeleteSelectProblemReq struct {
	QueID    int `json:"que_id"`
	CourseID int `json:"course_id"`
}

type DeleteSelectProblemFromSetReq struct {
	BlogID   int `json:"blog_id"`
	QueID    int `json:"que_id"`
	CourseId int `json:"course_id"`
}

type ProblemDetailReq struct {
	QueID int `json:"que_id"`
}

type FindSelectReq struct {
	KeyWord    string `json:"key_word"`
	CategoryID int    `json:"category_id"`
}

type AddCodeProblemReq struct {
	BlogID      int                 `json:"blog_id"`
	CourseID    int                 `json:"course_id"`
	CodeProblem *models.CodeProblem `json:"code_problem"`
}

type AddCheckPointBatchReq struct {
	QueID int                   `form:"que_id"`
	File  *multipart.FileHeader `form:"file"`
}

type AppendCodeProblemReq struct {
	BlogID   int `json:"blog_id"`
	QueId    int `json:"que_id"`
	CourseID int `json:"course_id"`
}

type DeleteCodeProblemFromSetReq struct {
	BlogID   int `json:"blog_id"`
	QueID    int `json:"que_id"`
	CourseId int `json:"course_id"`
}
