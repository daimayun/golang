package wechat

// GetAuthCode 获取Auth2.0 Code的链接地址
func GetAuthCode(appId, redirectUri, state string, scope ...AuthScopeType) string {
	s := AuthScopeSnsApiBase
	if len(scope) > 0 {
		s = scope[0]
	}
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + appId + "&redirect_uri=" + redirectUri + "&response_type=code&scope=" + string(s) + "&state=" + state + "#wechat_redirect"
}
