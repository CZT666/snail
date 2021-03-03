package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"snail/judger/client/judger"
	"snail/judger/grpcServer/proto"
)

type server struct {
}

func (receiver *server) SendMessage(ctx context.Context, req *proto.SendMessageReq) (*proto.SendMessageRsp, error) {
	// TODO 具体实现
	fmt.Printf("req: %v\n", req)
	return nil, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Printf("fail to lisent: %v\n", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterJudgeClientServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(listen)
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
	judger.NewSubmission(1, "127.0.0.1:9090")
}
