package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"snail/judgerClient/grpcServer/proto"
	"snail/judgerClient/judger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (receiver *server) SendMessage(ctx context.Context, req *proto.SendMessageReq) (*proto.SendMessageRsp, error) {
	// TODO 具体实现
	fmt.Printf("req: %v\n", req)
	rsp := new(proto.SendMessageRsp)
	rsp.Result = 0
	return rsp, nil
}

func initServer() {
	fmt.Print("hello\n")
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Printf("fail to lisent: %v\n", err)
		return
	}
	fmt.Print("listen\n")
	s := grpc.NewServer()
	proto.RegisterJudgeClientServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(listen)
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
}

func main() {
	go initServer()
	fmt.Print("start...\n")
	judger.NewSubmission(1, "127.0.0.1:9090")
	fmt.Print("end...\n")
	var tmp string
	fmt.Scanln(&tmp)
}
