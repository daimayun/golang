package function

import (
	"encoding/base64"
	"net/url"
)

// Base64Encode base64_encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode base64_decode
func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

// UrlEncode url_encode
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// UrlDecode url_decode
func UrlDecode(str string) string {
	data, _ := url.QueryUnescape(str)
	return data
}
