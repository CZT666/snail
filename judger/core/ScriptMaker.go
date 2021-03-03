package core

import (
	"log"
	"os"
	"snail/judger/model"
	"strconv"
	"strings"
)

func GenScript(submission *model.Submission, workPath string) error {
	language := new(model.Language)
	language.ID = submission.LanguageId
	if err := model.GetOneLanguage(language); err != nil {
		log.Printf("get language failed: %v\n", err)
		return err
	}
	fileName, err := genCode(submission.Code, language, workPath)
	if err != nil {
		log.Printf("gen code file failed: %v\n", err)
		return err
	}
	if submission.LanguageId == 3 {
		return nil
	}
	err = genCompileSh(fileName, language, workPath)
	if err != nil {
		log.Printf("gen run sh failed: %v\n", err)
		return err
	}
	return nil
}

func genCode(code string, language *model.Language, workPath string) (string, error) {

	var fileName string
	switch language.Name {
	case "Java":
		fileName = "Main.java"
	case "Go":
		fileName = "main.go"
	case "Python":
		fileName = "main.py"
	}
	filePath := workPath + "/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("create file failed: %v\n", err)
		return "", err
	}
	_, err = file.WriteString(code)
	defer file.Close()
	if err != nil {
		log.Printf("write code file failed: %v\n", err)
		return "", err
	}
	return fileName, nil
}

func genCompileSh(fileName string, language *model.Language, workPath string) error {
	stringBuilder := new(strings.Builder)
	stringBuilder.WriteString("#!/bin/bash\n")
	stringBuilder.WriteString("cd " + workPath + "\n")
	stringBuilder.WriteString(language.CompileCommand)
	stringBuilder.WriteString(" " + workPath + "/" + fileName + "\n")
	log.Printf(stringBuilder.String())
	filePath := workPath + "/compile.sh"
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Printf("create run.sh failed: %v\n", err)
		return err
	}
	_, err = file.WriteString(stringBuilder.String())
	if err != nil {
		log.Printf("write run.sh failed: %v\n", err)
	}
	return nil
}

func genCheckScript(index int, runCommand string, exeFileName string, input string, problem *model.CodeProblem, workPath string) (fileName string, err error) {
	fileName = workPath + "/check_" + strconv.Itoa(index) + ".sh"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Printf("gen check script main file failed: %v\n", err)
		return "", err
	}
	watchFileName := workPath + "/watch_" + strconv.Itoa(index) + ".sh"
	watchFile, err := os.Create(watchFileName)
	defer watchFile.Close()
	if err != nil {
		log.Printf("gen check script watch file failed: %v\n", err)
		return "", err
	}
	// 生成主程序脚本
	var stringBuilder strings.Builder
	stringBuilder.WriteString("#!/bin/bash\nmaxTime=" + strconv.Itoa(problem.TimeLimit) + "\nmaxMemory=" +
		strconv.Itoa(problem.MemoryLimit) + "\ncIndex=" + strconv.Itoa(index) + "\n")
	stringBuilder.WriteString("touch " + workPath +"/result_${cIndex}.txt\n")
	stringBuilder.WriteString("touch " + workPath +"/memory_log_${cIndex}.txt\n")
	stringBuilder.WriteString("touch " + workPath +"/time_log_${cIndex}.txt\n")
	stringBuilder.WriteString("source " + watchFileName + " $$ $maxTime $maxMemory $cIndex &\n")
	stringBuilder.WriteString("chipid=$!\n")
	stringBuilder.WriteString("cd " + workPath + "\n")
	if runCommand != "" {
		stringBuilder.WriteString(runCommand + " " + exeFileName)
	} else {
		stringBuilder.WriteString("./" + exeFileName)
	}
	for _, param := range strings.Split(input, ",") {
		stringBuilder.WriteString(" " + param)
	}
	stringBuilder.WriteString(" >> " + workPath +"/result_${cIndex}.txt\nkill $chipid")
	_, err = file.WriteString(stringBuilder.String())
	if err != nil {
		log.Printf("gen check write file failed: %v\n", err)
		return "", err
	}

	// 生成监控脚本
	var sb strings.Builder
	sb.WriteString("#!/bin/bash\nfpid=$1\nmaxTime=$2\nmaxMemory=$3\ncIndex=$4\ntimeGap=0.01\nstartTime=$[$(date +%s%N)/1000000]\n" +
		"while true\ndo\necho \"watch pid $fpid\"\n" +
		"info=$(cat  /proc/$fpid/status|grep -e VmRSS)\necho \"current memory $info\"\n" +
		"currentMemory=`echo $info | tr -cd \"[0-9]\"`\necho \"memory: $currentMemory\"\n" +
		"echo \"max: $maxMemory\"\nif [ $currentMemory -gt $maxMemory ];then\necho \"memory out\"\nkill $fpid\n" +
		"echo \"-1\" >> " + workPath + "/result_${cIndex}.txt\nbreak\nfi\n" +
		"echo $currentMemory >> " + workPath + "/memory_log_${cIndex}.txt\ncurrentTime=$[$(date +%s%N)/1000000]\n" +
		"declare -i tmp=$currentTime-$startTime\necho \"time gap $tmp\"\n" +
		"echo $tmp >> " + workPath + "/time_log_${cIndex}.txt\nif [ $tmp -gt $maxTime ];then\necho \"time out\"\nkill $fpid\n" +
		"echo \"-2\" >> " + workPath + "/result_${cIndex}.txt\nfi\nsleep $timeGap\ndone")
	_, err = watchFile.WriteString(sb.String())
	if err != nil {
		log.Printf("gen watch write file failed: %v\n", err)
		return "", err
	}
	return
}
