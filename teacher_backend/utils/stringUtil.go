package utils

import (
	"snail/teacher_backend/common"
	"strings"
)

func String2Set(str string, reg string) (set *common.Set) {
	set = common.NewSet()
	itemList := strings.Split(str, reg)
	for _, item := range itemList {
		set.Add(item)
	}
	return
}

func List2String(strList []string, reg string) string {
	var stringBuilder strings.Builder
	for _, item := range strList {
		if item == "" {
			continue
		}
		stringBuilder.WriteString(item)
		stringBuilder.WriteString(reg)
	}
	return stringBuilder.String()
}
