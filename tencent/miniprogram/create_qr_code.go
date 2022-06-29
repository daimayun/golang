package miniprogram

import (
	"encoding/json"
	"github.com/daimayun/golang/https"
)

// CreateQRCode 获取小程序二维码
func CreateQRCode(accessToken, path string, width int) (b []byte, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=" + accessToken
	jsonByte, _ := json.Marshal(struct {
		Path  string `json:"path"`
		Width int    `json:"width"`
	}{
		Path:  path,
		Width: width,
	})
	b, err = https.PostJson(url, string(jsonByte))
	if err != nil {
		return
	}
	return
}
