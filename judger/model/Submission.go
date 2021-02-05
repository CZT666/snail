package model

import (
	"snail/judger/dao"
	"time"
)

type Submission struct {
	ID          int        `json:"id"`
	ProblemId   int        `json:"problem_id"`
	UserId      string     `json:"user_id"`
	LanguageId  int        `json:"language_id"`
	SubmitTime  *time.Time `json:"submit_time"`
	ExecuteTime *time.Time `json:"execute_time"`
	UsedTime    *time.Time `json:"used_time"`
	UsedMemory  int        `json:"used_memory"`
	JudgeResult int        `json:"judge_result"`
	JudgeScore  int        `json:"judge_score"`
	JudgeLog    string     `json:"judge_log"`
	Code        string     `json:"code"`
}

func GetOneSubmission(submission *Submission) (err error) {
	err = dao.DB.Where(&submission).First(&submission).Error
	return
}

func UpdateSubmission(submission *Submission) (err error) {
	err = dao.DB.Model(&Submission{}).Where("id = ?", submission.ID).Updates(&submission).Error
	return
}
