package core

import (
	"log"
	"os"
	"snail/judger/model"
	"strconv"
	"strings"
	"time"
)

func MakeWorkPlace(submission *model.Submission) (workPath string, err error) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Printf("get work dir failed: %v\n", err)
		return "", err
	}
	workDir = strings.ReplaceAll(workDir, "\\", "/")
	stringBuilder := new(strings.Builder)
	currentTime := time.Now().Unix()
	stringBuilder.WriteString(workDir)
	stringBuilder.WriteString("/")
	stringBuilder.WriteString(strconv.Itoa(int(currentTime)) + "_")
	stringBuilder.WriteString(strconv.Itoa(submission.ID))
	err = os.Mkdir(stringBuilder.String(), os.ModePerm)
	if err != nil {
		log.Printf("mk dir faled: %v\n", err)
		return "", err
	}
	return stringBuilder.String(), nil
}
