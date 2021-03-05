package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"snail/student_bakcend/dao"
	"snail/student_bakcend/models"
	"snail/student_bakcend/vo"
)

func GetCode(blog string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	key := fmt.Sprintf(BlogCodeProblem, blog)
	allProblem, err := dao.RedisDB.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "get redis of code all problem fail"
		log.Printf("get redis of code all problem fail : %v\n", err)
		return
	}
	if allProblem == "" {
		GetProblem(blog)
		allProblem, err = dao.RedisDB.Get(context.Background(), key).Result()
		if err != nil && err != redis.Nil {
			baseResponse.Code = vo.Error
			baseResponse.Msg = "get redis of code all problem fail"
			log.Printf("get redis of code all problem fail : %v\n", err)
			return
		}
	}
	var allProblemJson []models.CodeProblem
	if err := json.Unmarshal([]byte(allProblem), &allProblemJson); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "unmarshal data code value fail"
		log.Printf("unmarshal data code value fail: %v\n", err)
		return
	}
	baseResponse.Data = allProblemJson
	return
}
