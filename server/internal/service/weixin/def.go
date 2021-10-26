package weixin

const (
	cacheSessionkey = "session_key"
	cacheWxUser     = "wx_user_cache"
)

// OauthReq 微信授权请求信息
type OauthReq struct {
	Jscode string `json:"jscode" binding:"required"` // 微信端jscode
}

// OauthResp 微信授权返回信息
type OauthResp struct {
	SessionID   string `json:"session_id"`   // 用户sessionid 用于当前交互使用
	OpenID      string `json:"open_id"`      // openid
	OverdueTime int64  `json:"overdue_time"` // 逾期时间点(时间戳)
}

// WxSessionKey ...
type WxSessionKey struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

// BaseSessionID 基础类型（sessionid）
type BaseSessionID struct {
	SessionID string `json:"session_id" binding:"required"`
}

// WxUserinfo 用户信息
type WxUserinfo struct {
	NickName  string `json:"nickName"`  //用户昵称
	AvatarURL string `json:"avatarUrl"` //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Gender    int    `json:"gender"`    //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City      string `json:"city"`      //用户所在城市
	Province  string `json:"province"`  //用户所在省份
	Country   string `json:"country"`   //用户所在国家
	Language  string `json:"language"`  //用户的语言，简体中文为zh_CN
}
