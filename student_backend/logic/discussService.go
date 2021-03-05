package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"log"
	"snail/student_bakcend/dao"
	"snail/student_bakcend/models"
	"snail/student_bakcend/models/helper"
	"snail/student_bakcend/vo"
	"time"
)

type RedPoint struct {
	Question       models.Question
	DataVersion    int
	CurrentVersion int
	IsRed          bool
}

func AddQuestion(question *models.Question, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	question.CreateBy = user.GetName()
	//question.CreateBy = "canruichen"
	question.CreateTime = time.Now()
	if err := models.AddQuestion(question); err != nil {
		log.Printf("discuss service add question failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
	}

	tempCourse := models.Course{ID: question.CourseID}
	if err := models.GetSingleCourse(&tempCourse); err != nil {
		log.Printf("discuss service get single course failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	RedPointKey := fmt.Sprintf(RedisRedPoint, tempCourse.CreateBy)
	var teacherValues []RedPoint
	teacherRedis, err := dao.RedisDB.Get(context.Background(), RedPointKey).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of teacher values fail"
		log.Printf("get redis of teacher values fail : %v\n", err)
		return
	}
	if teacherRedis != "" {
		if err := json.Unmarshal([]byte(teacherRedis), &teacherValues); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "unmarshal teacher redis fail"
			log.Printf("unmarshal teacher redis fail: %v\n", err)
			return
		}
	}
	teacherValues = append(teacherValues, RedPoint{Question: *question, DataVersion: -1, CurrentVersion: 0, IsRed: true})
	value, err := json.Marshal(teacherValues)
	if err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "marshal value fail"
		log.Printf("marshal value fail:: %v\n", err)
		return
	}
	if _, err := dao.RedisDB.Set(context.Background(), RedPointKey, cast.ToString(value), 0).Result(); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "redis set marshal data fail"
		log.Printf("redis set marshal data fail:: %v\n", err)
		return
	}
	tempAssistance := models.Assistance{
		CourseID: question.CourseID,
	}
	allAssistance, err := models.GetAssistance(&tempAssistance)
	for i := range allAssistance {
		student := models.Student{StudentID: allAssistance[i].StuID}
		if err := models.GetSingleStudent(&student); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "discuss get single student fail"
			log.Printf("discuss get single student fail: %v\n", err)
			return
		}
		fmt.Printf("student msg : %v\n",student)
		RedPointKey = fmt.Sprintf(RedisRedPoint, student.StudentID)
		var assistanceValues []RedPoint
		assistanceRedis, err := dao.RedisDB.Get(context.Background(), RedPointKey).Result()
		if err != nil && err != redis.Nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "get assistance redis of teacher values fail"
			log.Printf("get assistance redis of teacher values fail : %v\n", err)
			return
		}
		if assistanceRedis != "" {
			if err := json.Unmarshal([]byte(assistanceRedis), &assistanceValues); err != nil {
				baseResponse.Code = vo.Error
				baseResponse.Msg = "unmarshal assistance redis fail"
				log.Printf("unmarshal assistance redis fail: %v\n", err)
				return
			}
		}
		assistanceValues = append(assistanceValues, RedPoint{Question: *question, DataVersion: -1, CurrentVersion: 0, IsRed: true})
		fmt.Printf("assistanceValues is :%v\n",assistanceValues)
		value, err := json.Marshal(assistanceValues)
		if err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "assistance marshal value fail"
			log.Printf("assistance marshal value fail:: %v\n", err)
			return
		}
		if _, err := dao.RedisDB.Set(context.Background(), RedPointKey, cast.ToString(value), 0).Result(); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "redis set marshal data fail"
			log.Printf("redis set marshal data fail:: %v\n", err)
			return
		}
	}
	return
}

func GetAllQuestion(courseID string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	tempQuestion := models.Question{CourseID: cast.ToInt(courseID)}
	result, err := models.GetQuestion(&tempQuestion)
	if err != nil {
		log.Printf("queans service get all question failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	baseResponse.Data = result
	return
}

func GetSingleQuestion(questionID string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	question := models.Question{ID: cast.ToInt(questionID)}
	if err := models.GetSingleQuestion(&question); err != nil {
		log.Printf("queans service get single question failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	baseResponse.Data = question
	return
}

func SearchQuestion( searchName string, courseID string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	courseList, err := models.GetSearchQuestion(searchName, courseID)
	if err != nil {
		log.Printf("search question list failed: %v\n", err)
		baseResponse.Code = vo.ServerError
		return
	}
	baseResponse.Data = courseList
	return
}

func AddAnswer(answer *models.Answer, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	answer.CreateBy = user.GetName()
	//answer.CreateBy = "canruichen"
	answer.AnswerTime = time.Now()
	if err := models.AddAnswer(answer); err != nil {
		log.Printf("disscuss service add answer failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
	}
	question := models.Question{
		ID: answer.QuestionID,
	}
	if err := models.GetSingleQuestion(&question);err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get single question  fail"
		log.Printf("get single question fail : %v\n", err)
		return
	}
	course := models.Course{
		ID: question.CourseID,
	}
	RedPointKey := fmt.Sprintf(RedisRedPoint,course.CreateBy)
	var teacherValues []RedPoint
	teacherRedis, err := dao.RedisDB.Get(context.Background(), RedPointKey).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of teacher values fail"
		log.Printf("get redis of teacher values fail : %v\n", err)
		return
	}
	if teacherRedis != "" {
		if err := json.Unmarshal([]byte(teacherRedis), &teacherValues); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "unmarshal teacher redis fail"
			log.Printf("unmarshal teacher redis fail: %v\n", err)
			return
		}
	}
	for i := range teacherValues{
		if teacherValues[i].Question.ID == answer.QuestionID{
			teacherValues[i].CurrentVersion += 1
			teacherValues[i].IsRed = true
		}
	}
	value, err := json.Marshal(teacherValues)
	if err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "marshal value fail"
		log.Printf("marshal value fail:: %v\n", err)
		return
	}
	if _, err := dao.RedisDB.Set(context.Background(), RedPointKey, cast.ToString(value), 0).Result(); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "redis set marshal data fail"
		log.Printf("redis set marshal data fail:: %v\n", err)
		return
	}
	tempAssistance := models.Assistance{
		CourseID: question.CourseID,
	}
	allAssistance, err := models.GetAssistance(&tempAssistance)
	for i := range allAssistance {
		student := models.Student{StudentID: allAssistance[i].StuID}
		if err := models.GetSingleStudent(&student); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "discuss get single student fail"
			log.Printf("discuss get single student fail: %v\n", err)
			return
		}
		fmt.Printf("student msg : %v\n",student)
		RedPointKey = fmt.Sprintf(RedisRedPoint, student.StudentID)
		var assistanceValues []RedPoint
		assistanceRedis, err := dao.RedisDB.Get(context.Background(), RedPointKey).Result()
		if err != nil && err != redis.Nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "get assistance redis of teacher values fail"
			log.Printf("get assistance redis of teacher values fail : %v\n", err)
			return
		}
		if assistanceRedis != "" {
			if err := json.Unmarshal([]byte(assistanceRedis), &assistanceValues); err != nil {
				baseResponse.Code = vo.Error
				baseResponse.Msg = "unmarshal assistance redis fail"
				log.Printf("unmarshal assistance redis fail: %v\n", err)
				return
			}
		}
		for i := range assistanceValues{
			if assistanceValues[i].Question.ID == answer.QuestionID{
				assistanceValues[i].CurrentVersion += 1
				assistanceValues[i].IsRed = true
			}
		}
		fmt.Printf("assistanceValues is :%v\n",assistanceValues)
		value, err := json.Marshal(assistanceValues)
		if err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "assistance marshal value fail"
			log.Printf("assistance marshal value fail:: %v\n", err)
			return
		}
		if _, err := dao.RedisDB.Set(context.Background(), RedPointKey, cast.ToString(value), 0).Result(); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "redis set marshal data fail"
			log.Printf("redis set marshal data fail:: %v\n", err)
			return
		}
	}
	return
}

func GetAnswer(questionID string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	tempAnswer := models.Answer{QuestionID: cast.ToInt(questionID)}
	result,err := models.GetAnswer(&tempAnswer)
	if err != nil {
		log.Printf("discuss service get answer failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	baseResponse.Data = result
	return
}
