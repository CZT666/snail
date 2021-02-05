package main

import (
	"fmt"
	"snail/judger/core"
	"snail/judger/dao"
	proto "snail/judger/grpcServer"
	"snail/judger/settings"
)

func main() {
	settings.Init()
	err := dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
	req := new(proto.NewSubmissionReq)
	req.SubmissionId = 1
	core.ProcessJudge(req)
}
