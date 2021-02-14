package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
	"student_bakcend/dao"
	"student_bakcend/models"
	"student_bakcend/vo"
)

func GetSelect(blog string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	key := fmt.Sprintf(BlogSelectProblem, blog)
	allProblem, err := dao.RedisDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of select all problem fail"
		log.Printf("get redis of select all problem fail : %v\n", err)
		return
	}
	if allProblem == "" {
		GetProblem(blog)
		allProblem, err = dao.RedisDB.Get(context.Background(), key).Result()
		if err != nil && err != redis.Nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "get redis of select all problem fail"
			log.Printf("get redis of select all problem fail : %v\n", err)
			return
		}
	}
	var allProblemJson []models.SelectProblem
	if err := json.Unmarshal([]byte(allProblem), &allProblemJson); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "unmarshal data select value fail"
		log.Printf("unmarshal data select value fail: %v\n", err)
		return
	}
	var allProblemReturn []models.SelectProblem
	for i := range allProblemJson {
		allProblemReturn = append(allProblemReturn, models.SelectProblem{
			ID:          allProblemJson[i].ID,
			Description: allProblemJson[i].Description,
			Choices:     allProblemJson[i].Choices,
			Type:        allProblemJson[i].Type,
			CategoryID:  allProblemJson[i].CategoryID,
			CreateBy:    allProblemJson[i].CreateBy,
		})
	}
	baseResponse.Data = allProblemReturn
	return
}

func GetSelectScore(answers string,blog string,userID string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	num := 0
	if answers != "" {
		key := fmt.Sprintf(BlogSelectProblem,blog)
		scoresValue := strings.Split(answers,",")
		answers,answersError:= dao.RedisDB.Get(context.Background(),key).Result()
		if answersError == redis.Nil{
			GetProblem(blog)
			answers,_ = dao.RedisDB.Get(context.Background(),key).Result()
		}
		if answers != ""{
			var answersValues []models.SelectProblem
			answerValuesError := json.Unmarshal([]byte(answers),&answersValues)
			if answerValuesError != nil{
				baseResponse.Code = vo.Error
				log.Printf("answer values error")
				baseResponse.Msg = "answer values error"
				return
			}
			for i := range answersValues{
				if scoresValue[i] == answersValues[i].Answer{
					num += answersValues[i].Score
				}
			}
		}
	}else{
		baseResponse.Code = vo.Error
		log.Printf("scores is empty")
		baseResponse.Msg = "scores is empty"
	}
	key := fmt.Sprintf(SelectScores,blog,userID)
	if _,err:= dao.RedisDB.Set(context.Background(),key,num,0).Result();err!=nil{
		baseResponse.Code = vo.Error
		log.Printf("set value select score fail:%v\n",err)
		baseResponse.Msg = "set value select score fail"
		return
	}
	return
}
