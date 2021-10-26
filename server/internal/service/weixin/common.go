package weixin

import (
	"caoguo/internal/config"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/xxjwxc/public/dev"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/public/myglobal"
	wx "github.com/xxjwxc/public/weixin"

	"github.com/xxjwxc/public/mycache"
)

var _wx wx.WxTools

func init() {
	info := config.GetWxInfo()
	t, err := wx.New(wx.WxInfo{
		AppID:     info.AppID,
		AppSecret: info.AppSecret,
		APIKey:    info.APIKey,
		MchID:     info.MchID,
		NotifyURL: info.NotifyURL,
		ShearURL:  info.ShearURL,
	})
	if err != nil {
		mylog.Fatal(err)
	}
	_wx = t
}

//生成sessionId
func createSessionkey() (sessionKey string) {
	cache := mycache.NewCache(cacheSessionkey)

	sessionKey = myglobal.GetNode().GetIDStr()
	if cache.IsExist(sessionKey) { // 如果有,再找一次
		sessionKey = myglobal.GetNode().GetIDStr()
	}
	return
}

// GetOpenID 微信授权
func GetOpenID(jscode string) (result OauthResp) {
	body := _wx.SmallAppOauth(jscode)
	result.SessionID, result.OpenID, result.OverdueTime = SaveSessionKey([]byte(body))
	return
}

// SaveSessionKey 保存用户sessionkey
func SaveSessionKey(body []byte) (sessionID, openID string, overdueTime int64) {
	v := WxSessionKey{}
	json.Unmarshal(body, &v)
	if len(v.Openid) > 0 {
		orm := core.Dao.GetDBr()

		var tmp model.WxUserinfo
		orm.Where("openid = ?", v.Openid).Find(&tmp)
		if tmp.ID == 0 {
			tmp.Openid = v.Openid
			orm.Create(&tmp)
		}
		sessionID = createSessionkey()
		openID = v.Openid
		//保存缓存
		cache := mycache.NewCache(cacheSessionkey)
		cache.Add(sessionID, v, 2*time.Hour)
		overdueTime = time.Now().Add(2 * time.Hour).Unix()
	}
	return
}

// GetSessionkey 获取session_key信息
func GetSessionkey(sessionID string) (info WxSessionKey) {
	cache := mycache.NewCache(cacheSessionkey)
	if dev.IsDev() && len(sessionID) == 0 {
		cache.Add("1234567890", WxSessionKey{
			Openid:     "oXebs4mR1hZEj6XTAWCISS5gYSIo",
			SessionKey: "1234567890",
		}, 2*time.Hour)
		sessionID = "1234567890"
	}

	cache.Value(sessionID, &info)
	return
}

// updateInfo 更新用户信息
func updateInfo(info model.WxUserinfo) bool {
	dbmgr := model.WxUserinfoMgr(core.Dao.GetDBw().DB)
	_tmp, err := dbmgr.GetFromOpenid(info.Openid)
	if err != nil {
		mylog.Error(err)
		return false
	}

	if _tmp.ID > 0 { // 更新
		info.ID = _tmp.ID
		num := dbmgr.Save(&info).RowsAffected
		if num > 0 { //清除缓存
			mycache.NewCache(cacheWxUser).Delete(info.Openid)
		}
		return num > 0
	}

	return false
}

// createUnique 创建billid
func createUnique(channal string) string {
	tmp := time.Now().Format("20060102150405") + channal
	tmp = tmp + GetRandomString(32-len(tmp))
	return tmp
}

// GetRandomString 生成随机数字字符串
func GetRandomString(n int) string {
	var _bytes = []byte("0123456789")
	var r *rand.Rand

	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	for i := 0; i < n; i++ {
		result = append(result, _bytes[r.Intn(len(_bytes))])
	}
	return string(result)
}
