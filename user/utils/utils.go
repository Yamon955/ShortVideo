package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5 使用 md5 进行加密
func Md5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func GenerateTokn() string{
	return ""
}