package application

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"snail/judger/grpcServer/proto"
)

func SendMessage(address string, msg string, code int) {
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Printf("init connect failed: %v\n", err)
		return
	}
	defer conn.Close()
	client := proto.NewJudgeClientClient(conn)
	req := new(proto.SendMessageReq)
	req.Msg = msg
	req.ResultCode = int32(code)
	rsp, err := client.SendMessage(context.Background(), req)
	if err != nil {
		log.Printf("send message back error: %v\n", err)
		return
	}
	log.Printf("rsp of send message: %v\n", rsp)
	return
}
