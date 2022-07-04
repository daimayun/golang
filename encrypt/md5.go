package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 MD5加密
func Md5(str string) string {
	o := md5.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Md5S16 返回16位MD5加密字符串
func Md5S16(str string) string {
	o := md5.New()
	o.Write([]byte(str))
	s := hex.EncodeToString(o.Sum(nil))
	return s[8:24]
}
