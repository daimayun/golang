package encrypt

import (
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
)

// Ripemd160 RIPEMD-160加密[40位16进制数字]
func Ripemd160(str string) string {
	o := ripemd160.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}
