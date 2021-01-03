package utils

import (
	"encoding/json"
	"errors"
	"strings"
)

func GetJsonValue(jsonData []byte, key string) (value string, err error) {
	valid := json.Valid(jsonData)
	if !valid {
		return "", errors.New("Invalid_Json")
	}
	jsonString := string(jsonData)
	jsonString = strings.TrimPrefix(jsonString, "{")
	jsonString = strings.TrimSuffix(jsonString, "}")
	jsonString = strings.TrimSpace(jsonString)
	for _, pair := range strings.Split(jsonString, ",") {
		pair = strings.TrimSpace(pair)
		part := strings.Split(pair, ":")
		keyPair := strings.TrimSpace(part[0])
		valuePair := strings.TrimSpace(part[1])
		ketTmp := strings.Trim(keyPair, "\"")
		if ketTmp == key {
			return strings.Trim(valuePair, "\""), nil
		}
	}
	return "", errors.New("Json_Key_Not_Exist")
}
