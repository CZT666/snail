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
	stringBuilder.WriteString(language.CompileCommand)
	stringBuilder.WriteString(" " + fileName + "\n")
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

func genCheckScript(index int, runCommand string, exeFileName string, input string, workPath string) (fileName string, err error) {
	fileName = workPath + "/" + strconv.Itoa(index) + ".sh"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Printf("gen check script make file failed: %v\n", err)
		return "", err
	}
	var stringBuilder strings.Builder
	stringBuilder.WriteString("#!/bin/bash\n")
	stringBuilder.WriteString(runCommand + " " + exeFileName)
	for _, param := range strings.Split(input, ",") {
		stringBuilder.WriteString(" " + param)
	}
	_, err = file.WriteString(stringBuilder.String())
	if err != nil {
		log.Printf("gen check write file failed: %v\n", err)
		return "", err
	}
	return
}
