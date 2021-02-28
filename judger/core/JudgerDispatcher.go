package core

import (
	"snail/judger/grpcServer/proto"
)

// 接收到新任务
func OnSubmissionCreate(req *proto.NewSubmissionReq) {
	NewTask(req)
}

// 系统发生错误
func OnErrorOccurred(address string, msg string, errorCode int) {
	SendMessage(address, msg, errorCode)
}

// 编译完成
func OnCompileFinished(address string) {
	SendMessage(address, "compile finished.", 0)
}

// 完成一个测试点
func OnOneCheckPointFinished(address string, msg string) {
	SendMessage(address, "one check point finished.", 1)
}

// 完成所有测试点
func OnAllCheckPointFinished(address string, msg string) {
	SendMessage(address, "all check point finished.", 2)
}
