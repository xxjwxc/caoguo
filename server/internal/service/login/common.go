package login

import (
	"caoguo/internal/api"
	"caoguo/internal/config"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/tools"
)

// GetAccessTokenByUser 登录获取token
func GetAccessTokenByUser(username, password, times string) (mess LoginTokenMsg) {
	if len(username) == 0 || len(password) == 0 || len(times) == 0 {
		mess.MessageBody = message.GetErrorMsg(message.ParameterInvalid)
		return
	}
	tokenType := config.GetTokenType()
	appid, appsecret := config.GetAppInfo()
	original := appid + username + password + times + appsecret
	token := tools.Md5Encoder(original)

	parm := make(map[string]string)
	parm["appid"] = appid
	parm["token"] = token
	parm["token_type"] = tokenType
	parm["password"] = password
	parm["username"] = username
	parm["time"] = times

	myhttp.SendPost(parm, &mess, config.GetLoginURL())
	if mess.State {
		api.SaveCacheBody(mess.Data.AccessToken, username, mess.Data.ExpireTime) //保存缓存
	}
	return
}

// GetAccessTokenByUserNpwd 登录获取token 非用户密码登录
func GetAccessTokenByUserNpwd(username, times string) (mess LoginTokenMsg) {
	if len(username) == 0 || len(times) == 0 {
		mess.MessageBody = message.GetErrorMsg(message.ParameterInvalid)
		return
	}
	tokenType := config.GetTokenType()
	appid, appsecret := config.GetAppInfo()
	original := appid + username + times + appsecret
	token := tools.Md5Encoder(original)

	parm := make(map[string]string)
	parm["appid"] = appid
	parm["token"] = token
	parm["token_type"] = tokenType
	parm["username"] = username
	parm["time"] = times

	myhttp.SendPost(parm, &mess, config.GetLoginNoPwdURL())
	if mess.State {
		api.SaveCacheBody(mess.Data.AccessToken, username, mess.Data.ExpireTime) //保存缓存
	}

	return

}
