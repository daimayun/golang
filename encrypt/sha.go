package encrypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Sha1 sha1加密
func Sha1(str string) string {
	o := sha1.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha256 sha256加密[64位16进制数字]
func Sha256(str string) string {
	o := sha256.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha512 sha512加密[128位16进制数字]
func Sha512(str string) string {
	o := sha512.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}
