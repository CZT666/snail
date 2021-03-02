package model

import "snail/judger/dao"

type Language struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CompileCommand string `json:"compile_command"`
	RunCommand     string `json:"run_command"`
	ExeFileName    string `json:"exe_file_name"`
}

func GetOneLanguage(language *Language) (err error) {
	err = dao.DB.Where(&language).First(&language).Error
	return
}
