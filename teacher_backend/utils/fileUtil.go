package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	pathLock sync.Mutex
)

func GetFileType(fileName string) (string, error) {
	index := strings.Index(fileName, ".")
	if index == -1 {
		return "", errors.New("文件名错误")
	}
	return fileName[index+1:], nil
}

func GenFilePath(fileName string) (string, error) {
	pathLock.Lock()
	defer pathLock.Unlock()
	currentTime := time.Now().Unix()
	var stringBuilder strings.Builder
	workPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	workPath = strings.ReplaceAll(workPath, "\\", "/")
	stringBuilder.WriteString(workPath)
	stringBuilder.WriteString("/cache/tmp_")
	stringBuilder.WriteString(strconv.FormatInt(currentTime, 10))
	stringBuilder.WriteString(".")
	fileType, err := GetFileType(fileName)
	stringBuilder.WriteString(fileType)
	return stringBuilder.String(), err
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	return err
}
