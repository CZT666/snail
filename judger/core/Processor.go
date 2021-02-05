package core

import (
	"bytes"
	"log"
	"os/exec"
	proto "snail/judger/grpcServer"
	"snail/judger/model"
	"snail/judger/settings"
	"strconv"
)

var (
	taskQueue = make(chan *proto.NewSubmissionReq, settings.Conf.MaxTask)
)

func NewTask(req *proto.NewSubmissionReq) {
	taskQueue <- req
}

func RunJudgeTask() {
	go func() {
		for {
			req := <-taskQueue
			ProcessJudge(req)
		}
	}()
}

func ProcessJudge(req *proto.NewSubmissionReq) {
	submission := new(model.Submission)
	submission.ID = int(req.SubmissionId)
	if err := model.GetOneSubmission(submission); err != nil {
		log.Printf("Get submission failed: %v\n", err)
		OnErrorOccurred(err.Error())
		return
	}
	workPath, err := MakeWorkPlace(submission)
	if err != nil {
		log.Printf("make word place failed: %v\n", err)
		return
	}
	err = GenScript(submission, workPath)
	if err != nil {
		log.Printf("gen script failed: %v\n", err)
		return
	}
	result, msg := compile(workPath)
	switch result {
	case 0:
		log.Printf("compile success.\n")
		OnCompileFinished()
	case -1:
		log.Printf("compile failed.\n")
		OnErrorOccurred(msg)
	}
	if result == 0 {
		msg := runJudge(submission, workPath)
		OnAllCheckPointFinished(msg)
	}
}

func compile(wordPath string) (result int, msg string) {
	command := wordPath + "/run.sh"
	log.Printf("compile commond: %v\n", command)
	cmd := exec.Command("sh", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	err := cmd.Run()
	log.Printf("result of cmd: %v\n", out.String())
	log.Printf("err of cmd: %v\n", errOut.String())
	if err != nil {
		log.Printf("run cmd error: %v\n", err)
		return -1, errOut.String()
	}
	return 0, "编译成功"
}

func runJudge(submission *model.Submission, workPath string) (msg string) {
	queId := submission.ProblemId
	checkPointList, err := model.GetCheckPointByProblemId(queId)
	if err != nil {
		log.Printf("get check points failed: %v\n", err)
		return err.Error()
	}
	languageId := submission.LanguageId
	language := new(model.Language)
	language.ID = languageId
	if err := model.GetOneLanguage(language); err != nil {
		log.Printf("get one language failed: %v\n", err)
		return err.Error()
	}
	for index, checkPoint := range checkPointList {
		scriptName, err := genCheckScript(index, language.RunCommand, language.ExeFileName, checkPoint.Input, workPath)
		if err != nil {
			OnErrorOccurred("执行第" + strconv.Itoa(index) + "个测试用例失败" + err.Error())
			continue
		}
		cmd := exec.Command("sh", scriptName)
		var out bytes.Buffer
		cmd.Stdout = &out
		var errOut bytes.Buffer
		cmd.Stderr = &errOut
		err = cmd.Run()
		log.Printf("result of cmd: %v\n", out.String())
		if err != nil {
			log.Printf("run cmd error: %v\n", err)
			OnErrorOccurred("执行第" + strconv.Itoa(index) + "个测试用例失败" + err.Error())
		}
		OnOneCheckPointFinished("执行第" + strconv.Itoa(index) + "成功")
	}
	return "success"
}
