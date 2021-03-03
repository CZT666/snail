package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"log"
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/vo"
	"time"
)

type RedPoint struct {
	Question       models.Question
	DataVersion    int
	CurrentVersion int
	IsRed          bool
}

type Wrapper struct {
	redPoint []RedPoint
	by  func(p, q *RedPoint) bool
}

func GetRedPoint(user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	key := fmt.Sprintf(RedisRedPoint,user.GetIdentity())
	//key := fmt.Sprintf(RedisRedPoint,"1074596965@qq.com")
	result, err := dao.RedisDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of red point message values fail"
		log.Printf("get redis of red point message values fail : %v\n", err)
		return
	}
	var resultValue []RedPoint
	if result != "" {
		if err := json.Unmarshal([]byte(result), &resultValue); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "unmarshal result redis fail"
			log.Printf("unmarshal result redis fail: %v\n", err)
			return
		}
	}
	for i := range resultValue{
		if resultValue[i].IsRed{
			baseResponse.Data = 1
			return
		}
	}
	baseResponse.Data = 0
	return
}

func GetAllQuestion(user helper.User, pageRequest *helper.PageRequest) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	key := fmt.Sprintf(RedisRedPoint,user.GetIdentity())
	//key := fmt.Sprintf(RedisRedPoint,"1074596965@qq.com")
	result, err := dao.RedisDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of red point message values fail"
		log.Printf("get redis of red point message values fail : %v\n", err)
		return
	}
	var resultValue,res []RedPoint
	total := 0
	if result != "" {
		if err := json.Unmarshal([]byte(result), &resultValue); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "unmarshal result redis fail"
			log.Printf("unmarshal result redis fail: %v\n", err)
			return
		}
		var resultTrue,resultFalse []RedPoint
		for i := range resultValue{
			if resultValue[i].CurrentVersion != resultValue[i].DataVersion{
				resultValue[i].IsRed = true
			}else{
				resultValue[i].IsRed = false
			}
			if resultValue[i].IsRed{
				resultTrue = append(resultTrue,resultValue[i])
			}else{
				resultFalse =append(resultFalse,resultValue[i])
			}
			total +=1
		}
		res = append(res, resultTrue...)
		res = append(res, resultFalse...)
	}
	baseResponse.Data = helper.NewPageResponse(total, res)
	return
}

func GetSingleQuestion(questionID string,user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	question := models.Question{ID: cast.ToInt(questionID)}
	if err := models.GetSingleQuestion(&question); err != nil {
		log.Printf("queans service get single question failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	key := fmt.Sprintf(RedisRedPoint,user.GetIdentity())
	//key := fmt.Sprintf(RedisRedPoint,"1074596965@qq.com")
	result, err := dao.RedisDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of red point message values fail"
		log.Printf("get redis of red point message values fail : %v\n", err)
		return
	}
	var resultValue []RedPoint
	if result != "" {
		if err := json.Unmarshal([]byte(result), &resultValue); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "unmarshal result redis fail"
			log.Printf("unmarshal result redis fail: %v\n", err)
			return
		}
		for i := range resultValue{
			if resultValue[i].Question.ID == cast.ToInt(questionID){
				resultValue[i].IsRed = false
				resultValue[i].DataVersion = resultValue[i].CurrentVersion
			}
		}
		res,err:= json.Marshal(resultValue)
		if err!= nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "marshal result redis fail"
			log.Printf("marshal result redis fail: %v\n", err)
			return
		}
		if _, err := dao.RedisDB.Set(context.Background(), key, cast.ToString(res), 0).Result(); err != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "redis set marshal res fail"
			log.Printf("redis set marshal res fail:: %v\n", err)
			return
		}
	}
	baseResponse.Data = question
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
	fmt.Printf("question value:%v\n",question)
	if err := models.GetSingleQuestion(&question);err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get single question  fail"
		log.Printf("get single question fail : %v\n", err)
		return
	}
	course := models.Course{
		ID: question.CourseID,
	}
	if err := models.GetSingleCourse(&course);err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get single course  fail"
		log.Printf("get single course fail : %v\n", err)
		return
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

func GetAnswer(questionID string, pageRequest *helper.PageRequest) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	tempAnswer := models.Answer{QuestionID: cast.ToInt(questionID)}
	result, total, err := models.GetAnswer(&tempAnswer, pageRequest)
	if err != nil {
		log.Printf("discuss service get answer failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	pageResponse := helper.NewPageResponse(total, result)
	baseResponse.Data = pageResponse
	return
}