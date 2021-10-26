package login

import (
	"github.com/xxjwxc/public/message"
)

// LoginReq 登陆请求
type LoginReq struct {
	UserName string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码(用户输入的md5值,大写)
}

// RefreshReq reflashToken
type RefreshReq struct {
	Token string `json:"refresh_token"` // token
}

type ReqChangepwd struct {
	AccessToken string `json:"access_token"` //access_token
	UserName    string `json:"username"`     //用户名
	Password    string `json:"password"`     //新密码
}

// LoginTokenMsg ...
type LoginTokenMsg struct {
	message.MessageBody
	Data *LoginToken `json:"data,omitempty"`
}

// LoginToken ...
type LoginToken struct {
	AccessToken  string `json:"access_token,omitempty"`  //access_token
	ExpireTime   string `json:"expire_time,omitempty"`   //
	RefreshToken string `json:"refresh_token,omitempty"` //
}

//
// type Req_doaction struct {
// 	mysign.Sing_head
// 	UserName string `json:"username"`           //用户名
// 	Password string `json:"password"`           //密码
// 	Verify   string `json:"verify,omitempty"`   //验证码
// 	Flag     int    `json:"flag,omitempty"`     //1:注册，2:修改密码
// 	Bundleid string `json:"bundleid,omitempty"` //bundleid
// }
