package notify

import (
	"caoguo/internal/core"
	"caoguo/internal/model"
	"sync"
	"time"

	"github.com/xxjwxc/public/mycache"
)

var cache *mycache.MyCache
var once sync.Once
var mtx sync.Mutex

func getEmailList(username string) (resp []string) {
	once.Do(func() {
		cache = mycache.NewCache("_email_list")
	})

	mp := make(map[string][]string)
	mtx.Lock()
	defer mtx.Unlock()
	err := cache.Value("list", &mp)
	if err == nil { // 找到
		resp = append(resp, mp[""]...)
		if _, ok := mp[username]; ok {
			resp = append(resp, mp[username]...)
		}
		return
	}

	mp = make(map[string][]string)
	orm := core.Dao.GetDBr()
	list, _ := model.EmailNotifyTblMgr(orm.DB).Gets()
	for _, v := range list {
		mp[v.UserID] = append(mp[v.UserID], v.Email)
	}

	//添加
	cache.Add("list", mp, 1*time.Hour) // 1小时更新一次
	resp = append(resp, mp[""]...)
	if _, ok := mp[username]; ok {
		resp = append(resp, mp[username]...)
	}
	return resp
}
