package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

func EncodeMD5(strList ...string) string {
	var stringBuffer bytes.Buffer
	for _, str := range strList {
		stringBuffer.WriteString(str)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(stringBuffer.String())))
}
