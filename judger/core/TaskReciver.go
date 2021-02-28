package core

import (
	"context"
	"snail/judger/grpcServer/proto"
)

type JudgeServer struct {
}

func (server *JudgeServer) NewSubmission(ctx context.Context, req *proto.NewSubmissionReq) (*proto.NewSubmissionRsp, error) {
	OnSubmissionCreate(req)
	rsp := new(proto.NewSubmissionRsp)
	rsp.Result = 0
	return rsp, nil
}
