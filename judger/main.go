package main

import (
	"fmt"
	"log"
	"net"
	"snail/judger/core"
	"snail/judger/dao"
	"snail/judger/grpcServer/proto"
	"snail/judger/settings"
	"snail/judger/zk"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	settings.Init()
	err := dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
	conn := zk.InitZK(settings.Conf.ZKConfig, settings.Conf.Host, settings.Conf.Port)
	defer zk.Close(conn)
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Printf("fail to lisent: %v\n", err)
		return
	}
	core.RunJudgeTask()
	s := grpc.NewServer()
	proto.RegisterJudgeServerServer(s, &core.JudgeServer{})
	reflection.Register(s)
	err = s.Serve(listen)
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
}
