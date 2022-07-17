package wechat

import (
	"encoding/json"
	"github.com/daimayun/golang/https"
)

// GetAccessToken 获取AccessToken
func GetAccessToken(appId, appSecret string) (data ResponseAccessTokenData, errData ResponseErrorData, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret
	var b []byte
	b, err = https.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}
	if data.AccessToken != "" && data.ExpiresIn > 0 {
		return
	}
	err = json.Unmarshal(b, &errData)
	return
}

func GetAccessTokenByAuthCode(appId, appSecret, code string) (data, err error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appId + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"
	var b []byte
	b, err = https.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}
	return
}
