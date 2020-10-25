package main

import (
	"fmt"
	"snail/teacher_backend/dao"
	"snail/teacher_backend/routers"
	"snail/teacher_backend/settings"
)

func main() {
	// test
	settings.Init()
	err := dao.InitMySQL(settings.Conf)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
	r := routers.SetupRouter()
	r.Run(":8080")
}
