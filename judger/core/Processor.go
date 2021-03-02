package core

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"snail/judger/grpcServer/proto"
	"snail/judger/model"
	"snail/judger/settings"
	"strconv"
	"strings"
)

var (
	taskQueue = make(chan *proto.NewSubmissionReq, settings.Conf.MaxTask)
)

func NewTask(req *proto.NewSubmissionReq) {
	taskQueue <- req
}

func RunJudgeTask() {
	log.Printf("core task: %v\n", settings.Conf.CoreTask)
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
		OnErrorOccurred(req.OriginIp, err.Error(), -500)
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
	if submission.LanguageId == 3 {
		ret := runJudge(req, submission, workPath)
		OnAllCheckPointFinished(req.OriginIp, ret)
		clearWorkPlace(workPath)
		return
	}
	result, msg := compile(workPath)
	switch result {
	case 0:
		log.Printf("compile success.\n")
		OnCompileFinished(req.OriginIp)
	case -1:
		log.Printf("compile failed.\n")
		saveSubmission(submission, -1, -1, "-1", -3)
		OnErrorOccurred(req.OriginIp, msg, -3)
		return
	}
	ret := runJudge(req, submission, workPath)
	OnAllCheckPointFinished(req.OriginIp, ret)
	clearWorkPlace(workPath)
}

func compile(wordPath string) (result int, msg string) {
	command := wordPath + "/compile.sh"
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

func runJudge(req *proto.NewSubmissionReq, submission *model.Submission, workPath string) (msg string) {
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
	question, err := model.GetProblemById(queId)
	if err != nil {
		log.Printf("get question failed: %v\n", err)
		return err.Error()
	}
	maxMemory := 0
	totalTime := 0
	count := 0
	pass := 0
	for index, checkPoint := range checkPointList {
		count += 1
		scriptName, err := genCheckScript(index, language.RunCommand, language.ExeFileName, checkPoint.Input, question, workPath)
		if err != nil {
			OnErrorOccurred(req.OriginIp, "初始化第"+strconv.Itoa(index)+"个测试用例失败"+err.Error(), -5001)
			continue
		}
		result := workPath + "/result_" + strconv.Itoa(index) + ".txt"
		timeLog := workPath + "/time_log_" + strconv.Itoa(index) + ".txt"
		memoryLog := workPath + "/memory_log_" + strconv.Itoa(index) + ".txt"
		cmd := exec.Command("sh", scriptName)
		var out bytes.Buffer
		cmd.Stdout = &out
		var errOut bytes.Buffer
		cmd.Stderr = &errOut
		err = cmd.Run()
		log.Printf("result of cmd: %v\n", out.String())
		if err != nil {
			log.Printf("run cmd error: %v\n", err)
			OnErrorOccurred(req.OriginIp, "执行第"+strconv.Itoa(index)+"个测试用例失败"+err.Error(), -5002)
		}
		ret, code, err := checkAnswer(result, checkPoint.Output)
		log.Printf("code of run judge: %v\n", code)
		if err != nil {
			log.Printf("check answer failed: %v\n", err)
			saveSubmission(submission, -1, -1, "-1", code)
			return err.Error()
		}
		if code == 0 {
			pass += 1
		}
		OnOneCheckPointFinished(req.OriginIp, "执行第"+strconv.Itoa(index)+"个样例完成,结果:"+ret)
		addCostTime(timeLog, &totalTime)
		fixMemory(memoryLog, &maxMemory)
	}
	averageTime := totalTime / count
	passNum := strconv.Itoa(pass) + "/" + strconv.Itoa(count)
	result := 0
	if pass != count {
		result = 1
	}
	saveSubmission(submission, averageTime, maxMemory, passNum, result)
	return "样例通过情况:" + passNum +
		"; 平均耗时:" + strconv.Itoa(averageTime) +
		"; 最大占用内存" + strconv.Itoa(maxMemory)
}

func checkAnswer(resultPath string, output string) (msg string, code int, err error) {
	result, err := os.Open(resultPath)
	if err != nil {
		log.Printf("Open result file failed: %v\n", err)
		return "no result", -1, err
	}
	defer result.Close()
	index := 0
	var stringBuilder strings.Builder
	scanner := bufio.NewScanner(result)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := scanner.Text()
		if index == 0 {
			switch str {
			case "-1":
				return "超出时间限制", -1, errors.New("-1")
			case "-2":
				return "超出内存限制", -2, errors.New("-2")
			}
		}
		stringBuilder.WriteString(str + "\n")
		index++
	}
	log.Printf("result: %v\n", stringBuilder.String())
	if strings.Trim(stringBuilder.String(), "\n") == strings.Trim(output, "\n") {
		return "答案正确", 0, nil
	} else {
		return "答案错误", -1, nil
	}
}

func addCostTime(timeLogPath string, totalTime *int) {
	result, err := os.Open(timeLogPath)
	if err != nil {
		log.Printf("read time file failed: %v\n", err)
	}
	defer result.Close()
	scanner := bufio.NewScanner(result)
	scanner.Split(bufio.ScanLines)
	last := ""
	for scanner.Scan() {
		str := scanner.Text()
		last = strings.Trim(str, "\n")
	}
	if (last == "") {
		return
	}
	ret, err := strconv.Atoi(last)
	if err != nil {
		log.Printf("convert string failed: %v\n", err)
	}
	*totalTime += ret
}

func fixMemory(memoryLogPath string, maxMemory *int) {
	result, err := os.Open(memoryLogPath)
	if err != nil {
		log.Printf("read memory file failed: %v\n", err)
		return
	}
	defer result.Close()
	scanner := bufio.NewScanner(result)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := scanner.Text()
		tmp, err := strconv.Atoi(strings.Trim(str, "\n"))
		if err != nil {
			log.Printf("convert string failed: %v\n", err)
			return
		}
		if tmp > *maxMemory {
			*maxMemory = tmp
		}
	}
}

func saveSubmission(submission *model.Submission, costTime int, memory int, pass string, result int) {
	submission.JudgeResult = result
	submission.UsedTime = costTime
	submission.UsedMemory = memory
	submission.PassNum = pass
	err := model.UpdateSubmission(submission)
	if err != nil {
		log.Printf("save submission failed: %v\n", err)
	}
}

func clearWorkPlace(workPath string) {
	log.Printf("clear work place %v\n", workPath)
	err := os.RemoveAll(workPath)
	if err != nil {
		log.Printf("clear work place failed: %v\n", err)
	}
}
