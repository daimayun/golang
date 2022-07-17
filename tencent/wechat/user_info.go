package wechat

import (
	"encoding/json"
	"github.com/daimayun/golang/https"
)

// GetUserInfo 获取用户基本信息(UnionID机制)
func GetUserInfo(accessToken, openId string, lang ...string) (data ResponseUserInfoData, errData ResponseErrorData, err error) {
	language := "zh_CN"
	if len(lang) > 0 {
		language = lang[0]
	}
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openId + "&lang=" + language
	var b []byte
	b, err = https.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}
	if data.OpenId != "" {
		return
	}
	err = json.Unmarshal(b, &errData)
	return
}
