package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"log"
	"strings"
	"student_bakcend/dao"
	"student_bakcend/models"
	"student_bakcend/vo"
	"time"
)


func GetProblem(blog string)(baseResponse *vo.BaseResponse){
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	key := fmt.Sprintf(BlogProblem,blog)
	newAllProblem := models.QueSet{}
	allProblem,err:= dao.RedisDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of all problem fail"
		log.Printf("get redis of all problem fail : %v\n",err)
		return
	}

	if allProblem == ""{
		newAllProblem.BlogID = cast.ToInt(blog)
		if err:= models.GetSingleQueSet(&newAllProblem); err!= nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "get single que set fail"
			log.Printf("get single que set fail: %v\n", err)
			return
		}
		marshalData,marshalDataError := json.Marshal(newAllProblem)
		if marshalDataError != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "marshal all problem error"
			log.Printf("marshal all problem error: %v\n",marshalDataError)
			return
		}
		if _,err := dao.RedisDB.Set(context.Background(),key,cast.ToString(marshalData),24*time.Hour).Result();err!=nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "redis set marshal data fail"
			log.Printf("redis set marshal data fail:: %v\n", err)
			return
		}
	}else {
		unmarshalError := json.Unmarshal([]byte(allProblem),&newAllProblem)
		if unmarshalError != nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "unmarshal all problem error "
			log.Printf("unmarshal all problem error  : %v\n",unmarshalError)
			return
		}
	}
	var selectLen,codeLen []string
	var result []int
	var selectValue []models.SelectProblem
	var codeValue []models.CodeProblem
	if newAllProblem.SelectProblem != ""{
		selectLen = strings.Split(newAllProblem.SelectProblem,",")
	}

	for i := range selectLen{
		tempSelect := models.SelectProblem{
			ID: cast.ToInt(selectLen[i]),
		}
		if err:= models.GetSingleSelectProblem(&tempSelect);err!=nil{
			baseResponse.Code = vo.Error
			baseResponse.Msg = "select value get mysql fail"
			log.Printf("select value get mysql fail : %v\n",err)
			return
		}
		selectValue = append(selectValue,tempSelect)
	}
	if len(selectValue) != 0{
		selectKey := fmt.Sprintf(BlogSelectProblem,blog)
		selectValueMarshal,selectValueError := json.Marshal(selectValue)
		if selectValueError != nil{
			baseResponse.Code = vo.Error
			baseResponse.Msg = "marshal select value error"
			log.Printf("marshal select value error : %v\n",err)
			return
		}
		if _,err := dao.RedisDB.Set(context.Background(),selectKey,cast.ToString(selectValueMarshal),24*time.Hour).Result();err!=nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "redis set marshal value fail"
			log.Printf("redis set marshal value fail:: %v\n", err)
			return
		}
	}
	if newAllProblem.CodeProblem != "" {
		codeLen = strings.Split(newAllProblem.CodeProblem, ",")
	}
	for i := range codeLen{
		tempCode := models.CodeProblem{
			ID: cast.ToInt(codeLen[i]),
		}
		if err:= models.GetSingleCodeProblem(&tempCode);err!=nil{
			baseResponse.Code = vo.Error
			baseResponse.Msg = "code value get mysql fail"
			log.Printf("code value get mysql fail : %v\n",err)
			return
		}
		codeValue = append(codeValue,tempCode)
	}
	if len(codeValue) != 0{
		codeKey := fmt.Sprintf(BlogCodeProblem,blog)
		codeValueMarshal,codeValueError := json.Marshal(codeValue)
		if codeValueError != nil{
			baseResponse.Code = vo.Error
			baseResponse.Msg = "marshal code value error"
			log.Printf("marshal code value error : %v\n",err)
			return
		}
		if _,err := dao.RedisDB.Set(context.Background(),codeKey,cast.ToString(codeValueMarshal),24*time.Hour).Result();err!=nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "redis set code marshal value fail"
			log.Printf("redis set code marshal value fail:: %v\n", err)
			return
		}
	}
	result = append(result,len(selectLen))
	result = append(result,len(codeLen))
	baseResponse.Data = result
	return
}

func GetProblemScore(blogID string,studentID string)(baseResponse *vo.BaseResponse) {
	ctx := context.Background()
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	selectKey := fmt.Sprintf(SelectScores,blogID,studentID)
	fmt.Printf("select key value:%v",selectKey)
	codeKey := fmt.Sprintf(CodeScores,blogID,studentID)
	fmt.Printf("code key value:%v",codeKey)
	selectScore,selectError := dao.RedisDB.Get(ctx,selectKey).Result()
	if selectError != nil && selectError != redis.Nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of select score error"
		log.Printf("get redis of select scoree error : %v\n",selectError)
		return
	}
	codeScore,codeError := dao.RedisDB.Get(ctx,codeKey).Result()
	if codeError != nil && codeError != redis.Nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of code score error"
		log.Printf("get redis of code scoree error : %v\n",selectError)
		return
	}
	baseResponse.Data = cast.ToInt(codeScore) + cast.ToInt(selectScore)
	var blog models.Blog
	blog.ID = cast.ToInt(blogID)
	if err := models.GetSingleBlog(&blog);err!=nil{
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get single blog error"
		return
	}
	fmt.Printf("select score:%v, code score:%v",selectScore,codeScore)
	baseResponse.Data = cast.ToInt(selectScore) + cast.ToInt(codeScore)
	scoreRecord := models.ScoreRecord{
		StuID: studentID,
		CourseID: blog.CourseID,
		BlogID: blog.ID,
	}
	tmpErr := models.GetSingleScoreRecord(&scoreRecord)
	if tmpErr !=nil{
		scoreRecord.Score = cast.ToInt(selectScore) + cast.ToInt(codeScore)
		if err:= models.AddScoreRecord(&scoreRecord);err!=nil{
			baseResponse.Code = vo.Error
			baseResponse.Msg = "add score record error"
			return
		}
	}else{
		if err := models.UpdateScoreRecord(&scoreRecord);err!=nil{
			baseResponse.Code = vo.Error
			baseResponse.Msg = "update score record error"
			return
		}
	}
	return
}
