package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// md5加密
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	chipherStr := h.Sum(nil)
	return hex.EncodeToString(chipherStr)
}

// md5加密大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 验证密码,plainpwd明文密码,slat随机数,passwd密文密码
func ValidatePasswd(plainpwd, slat, passwd string) bool {
	return Md5Encode(plainpwd+slat) == passwd
}

// 创建密码,plainpwd明文密码,slat随机数
func MakePasswd(plainpwd, slat string) string {
	return Md5Encode(plainpwd + slat)
}
