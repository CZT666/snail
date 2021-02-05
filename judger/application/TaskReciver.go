package application

import (
	"context"
	"snail/judger/core"
	proto "snail/judger/grpcServer"
)

type JudgeServer struct {
}

func (server *JudgeServer) NewSubmission(ctx context.Context, req *proto.NewSubmissionReq) (*proto.NewSubmissionRsp, error) {
	core.OnSubmissionCreate(req)
	rsp := new(proto.NewSubmissionRsp)
	rsp.Result = 0
	return rsp, nil
}
