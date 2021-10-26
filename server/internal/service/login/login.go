package login

import (
	"caoguo/internal/api"
	"caoguo/internal/config"
	"caoguo/internal/model"
	"caoguo/internal/service/oplog"
	"log"
	"strconv"
	"time"

	"encoding/json"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/tools"
)

// Login 登陆相关
type Login struct {
}

// Login 登录
// @Router /login/login [POST]
func (l *Login) Login(c *api.Context, req LoginReq) {
	username := req.UserName
	password := req.Password
	times := strconv.FormatInt(tools.GetUtcTime(time.Now()), 10)

	if len(username) == 0 || len(password) == 0 || len(times) == 0 {
		c.WriteJSON(message.GetErrorMsg(message.ParameterInvalid))
		return
	}
	//-------------------------end

	tmp := GetAccessTokenByUser(username, password, times)
	if tmp.State {
		msg := message.GetSuccessMsg(message.NormalMessageID)
		msg.Data = tmp.Data
		//设置客户端cookie
		tm, _ := strconv.Atoi(tmp.Data.ExpireTime)
		c.GetGinCtx().SetCookie(api.UserToken, tmp.Data.AccessToken, tm, "/", c.GetGinCtx().Request.Host, false, true)
		c.WriteJSON(msg)
		var logs model.OpLogTbl
		logs.Operator = username
		logs.Topic = "帐号系统"
		logs.Bundle = "登录"
		logs.Pid = "pc登录"
		logs.IPAddr = c.GetGinCtx().ClientIP()
		oplog.OpLogBaseAdd(&logs) //保存日志
	} else {
		c.WriteJSON(tmp)
	}
	return
}

// CheckToken 验证token
func CheckToken(c *api.Context, req map[string]string) {
	token := req["token"]
	userInfo, b := api.GetUserFromToken(token)
	log.Println(userInfo.UserName, b)
	if !b {
		c.WriteJSON(message.GetErrorMsg(message.TokenFailure, "please login"))
	} else {
		msg := message.GetSuccessMsg(message.NormalMessageID)
		parm := make(map[string]string)
		parm["username"] = userInfo.UserName
		msg.Data = parm
		c.WriteJSON(msg)
	}
}

// RefreshToken 刷新token管理员
func (l *Login) RefreshToken(c *api.Context, req *RefreshReq) {
	token := req.Token
	if len(token) == 0 { //参数检测
		c.WriteJSON(message.GetErrorMsg(message.ParameterInvalid))
		return
	}

	parm := make(map[string]string)
	parm["refresh_token"] = token

	b, _ := json.Marshal(parm)
	rBody, _ := myhttp.OnPostJSON(config.GetRefreshTokenURL(), string(b))
	if len(rBody) == 0 {
		c.WriteJSON(message.GetErrorMsg(message.UnknownError))
		//return
	} else {
		var msg api.MapMessageBody
		json.Unmarshal([]byte(rBody), &msg)
		if msg.State {
			tm, _ := strconv.Atoi(msg.Data["expire_time"])
			c.GetGinCtx().SetCookie(api.UserToken, msg.Data["access_token"], tm, "/", c.GetGinCtx().Request.Host, false, true)
			//保存缓存
			api.SaveCacheBody(msg.Data["access_token"], msg.Data["username"], msg.Data["expire_time"])
			//返回结果
			c.WriteJSON(msg)
		} else {
			c.WriteJSON(message.GetErrorMsg(msg.Code))
		}
	}
}

// ChangePwd 修改用户密码
func ChangePwd(c *api.Context, req ReqChangepwd) {
	//参数检测
	if !tools.CheckParam(req.UserName, req.Password) {
		c.WriteJSON(message.GetErrorMsg(message.ParameterInvalid, req))
		return
	}

	//验证token，并获取用户名
	// _, bfind := user.GetUserDetail(req.UserName)
	// if !bfind {
	// 	w.WriteJson(message.GetErrorMsg(message.NotFindError))
	// 	return
	// }

	//oauth2修改密码
	if OnChangePwd(req.UserName, req.Password) {
		c.WriteJSON(message.GetSuccessMsg())
		var logs model.OpLogTbl
		logs.Operator = req.UserName
		logs.Topic = "帐号系统"
		logs.Bundle = "修改密码"
		logs.Pid = "pc修改密码"
		logs.IPAddr = c.GetGinCtx().ClientIP()
		oplog.OpLogBaseAdd(&logs) //保存日志
	} else {
		c.WriteJSON(message.GetErrorMsg(message.UserNameDoNotExist))
	}
	return
}

// CheckUser 选择用户
func CheckUser(c *api.Context) {
	_, b := c.GetUserInfo()
	if !b {
		c.WriteJSON(message.GetErrorMsg(message.TokenFailure, "please login"))
		return
	}

	c.WriteJSON(message.GetSuccessMsg())
}
