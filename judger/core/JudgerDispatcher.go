package core

import (
	proto "snail/judger/grpcServer"
)

// 接收到新任务
func OnSubmissionCreate(req *proto.NewSubmissionReq) {
	NewTask(req)
}

// 系统发生错误
func OnErrorOccurred(msg string) {

}

// 编译完成
func OnCompileFinished() {

}

// 完成一个测试点
func OnOneCheckPointFinished(msg string) {

}

// 完成所有测试点
func OnAllCheckPointFinished(msg string) {

}
