package wechat

// ErrCodeIsSuccess 微信接口请求返回成功
var ErrCodeIsSuccess int64 = 0

var (
	// AuthScopeSnsApiBase 不弹出授权页面，直接跳转，只能获取用户openid
	AuthScopeSnsApiBase AuthScopeType = "snsapi_base"
	// AuthScopeSnsApiUserInfo 弹出授权页面，可通过 openid 拿到昵称、性别、所在地。并且， 即使在未关注的情况下，只要用户授权，也能获取其信息
	AuthScopeSnsApiUserInfo AuthScopeType = "snsapi_userinfo"
)
