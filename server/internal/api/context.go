// Package api The next version of the underlying category will support automatic parsing of a single struct.
package api

import (
	"github.com/gin-gonic/gin"
	ginapi "github.com/gmsec/goplugins/api"
	"github.com/xxjwxc/public/mylog"
)

const (
	//UserToken 用户token
	UserToken = "user_token" // 聚合了用户端token user_token
	// SessionToken session token
	SessionToken = "session_token"
)

// Context Wrapping gin context to custom context
type Context struct { // 包装gin的上下文到自定义context
	ginapi.Context
}

//GetUserInfo 获取用户信息
func (c *Context) GetUserInfo() (u *UserInfo, b bool) {
	accessToken, err := c.GetGinCtx().Cookie(UserToken)
	if err == nil {
		return GetUserFromToken(accessToken)
	}

	accessToken, err = c.GetGinCtx().Cookie(SessionToken)
	if err == nil {
		return GetUserFromToken(accessToken)
	}

	mylog.ErrorString(err.Error()) // output error string
	return nil, false
}

// GetSessionToken find SessionToken by cookie
func (c *Context) GetSessionToken() string {
	sessionToken, _ := c.GetGinCtx().Cookie(SessionToken)

	return sessionToken
}

// GetClientIP get request ip
func (c *Context) GetClientIP() string {
	return c.GetGinCtx().ClientIP()
}

// NewAPIFunc default of custom handlefunc
func NewAPIFunc(c *gin.Context) interface{} {
	return &Context{*ginapi.NewCtx(c)}
}
