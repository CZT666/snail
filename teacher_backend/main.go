package main

import (
	"fmt"
	"snail/teacher_backend/dao"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/routers"
	"snail/teacher_backend/settings"
)

func main() {
	settings.Init()
	err := dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
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
