package wechat

// ResponseErrorData 请求返回错误时的数据
type ResponseErrorData struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ResponseAccessTokenData 返回access_token数据
type ResponseAccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// ResponseUserInfoData 获取用户基本信息(UnionID机制)
type ResponseUserInfoData struct {
	Subscribe      int     `json:"subscribe"`
	OpenId         string  `json:"openid"`
	Language       string  `json:"language"`
	SubscribeTime  int64   `json:"subscribe_time"`
	UnionId        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupId        int64   `json:"groupid"`
	TagIdList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int64   `json:"qr_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`
}
