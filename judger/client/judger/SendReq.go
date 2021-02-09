package judger

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"snail/judger/client/zk"
	"snail/judger/grpcServer/proto"
)

func getServer() string {
	conn, err := zk.ConnectZK()
	if err != nil {
		return ""
	}
	return zk.GetOneNode(conn)
}

func NewSubmission(submissionId int, originIp string) error {
	address := getServer()
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Printf("connnect server failed: %v\n", err)
		return err
	}
	defer conn.Close()
	client := proto.NewJudgeServerClient(conn)
	req := new(proto.NewSubmissionReq)
	req.SubmissionId = int32(submissionId)
	req.OriginIp = originIp
	ret, err := client.NewSubmission(context.Background(), req)
	if err != nil {
		log.Printf("new submission failed: %v\n", err)
		return err
	}
	log.Printf("result of new submission: %v\n", ret)
	return nil
}
