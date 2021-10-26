package api

import (
	"caoguo/internal/config"
	"encoding/json"
	"strconv"
	"time"

	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mycache"

	"github.com/xxjwxc/public/myhttp"
)

// UserInfo 用户信息
type UserInfo struct {
	AccessToken string // token
	UserName    string // 用户名
	ExpireTime  int    // 超时时间
}

// MapMessageBody req message body
type MapMessageBody struct {
	message.MessageBody
	Data map[string]string `json:"data,omitempty"`
}

// GetUserFromToken 通过token获取用户信息
func GetUserFromToken(token string) (userInfo *UserInfo, b bool) {
	if len(token) == 0 {
		if dev.IsDev() {
			return &UserInfo{
				UserName: "admin",
			}, true
		}
		return
	}

	// 先从缓存中获取
	if tmp, ok := GetCacheBody(token); ok {
		return &UserInfo{
			AccessToken: token,
			UserName:    tmp.UserName,
			ExpireTime:  tmp.ExpireTime}, true
	}

	parm := make(map[string]string)
	parm["token"] = token
	bod, _ := json.Marshal(parm)
	rBody, _ := myhttp.OnPostJSON(config.GetCheckTokenURL(), string(bod))
	if len(rBody) > 0 {
		var msg MapMessageBody
		json.Unmarshal([]byte(rBody), &msg)
		if msg.State {
			SaveCacheBody(token, msg.Data["username"], msg.Data["expire_time"]) // 保存缓存
			return &UserInfo{
				AccessToken: token,
				UserName:    msg.Data["username"],
				ExpireTime:  strToInt(msg.Data["expire_time"])}, true // 返回结果
		}
	}

	return
}

// SaveCacheBody 保存缓存
func SaveCacheBody(accessToken, username, expireTime string) {
	tmp := UserInfo{
		AccessToken: accessToken,
		UserName:    username,
		ExpireTime:  strToInt(expireTime),
	}
	// 保存缓存
	cache := mycache.NewCache("oauth2")
	cache.Add(accessToken, &tmp, time.Duration(tmp.ExpireTime)*time.Second)
	//------------------end
}

// GetCacheBody 获取缓存
func GetCacheBody(token string) (value *UserInfo, b bool) {
	value = &UserInfo{}
	cache := mycache.NewCache("oauth2")
	err := cache.Value(token, value)
	b = (err == nil)

	return
}

func strToInt(src string) int {
	i, _ := strconv.Atoi(src)
	return i
}
