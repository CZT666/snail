package main

import (
	"fmt"
	"log"
	"snail/student_bakcend/dao"
	"snail/student_bakcend/routers"
	"snail/student_bakcend/settings"
	"snail/student_bakcend/logic"
)

func main() {
	// test
	settings.Init()
	err := dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
	err = dao.InitRedis(settings.Conf.RedisConfig)
	if err != nil {
		log.Printf("Init redis failed, err: %v", err)
		return
	}
	defer dao.CloseRedis()
	err = dao.InitResetPwdNSQ(settings.Conf.NSQConfig)
	if err != nil {
		fmt.Printf("init reset pwd nsq failed, err:%v\n", err)
		return
	}
	err = logic.InitResetPwdConsumer(settings.Conf.ResetPwdConsumerConfig)
	if err != nil {
		fmt.Printf("init reset pwd consumer failed, err:%v\n", err)
		return
	}
	r := routers.SetupRouter()
	r.Run(":8080")
}
