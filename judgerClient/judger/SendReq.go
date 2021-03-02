package judger

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"snail/judgerClient/grpcServer/proto"
	"snail/judgerClient/zk"
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
	// address := "127.0.0.1:8081"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("connnect server failed: %v\n", err)
		return err
	}
	defer conn.Close()
	log.Print("connect")
	client := proto.NewJudgeServerClient(conn)
	req := new(proto.NewSubmissionReq)
	req.SubmissionId = int32(submissionId)
	req.OriginIp = originIp
	ret, err := client.NewSubmission(context.Background(), req)
	if err != nil {
		log.Printf("new submission failed: %v\n", err)
		return err
	}
	log.Printf("result of new submission: %v\n", ret.Result)
	return nil
}
