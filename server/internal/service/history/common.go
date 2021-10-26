package history

import (
	"fmt"
	"path"
	"strings"
	"sync"

	"github.com/xxjwxc/public/mysort"

	"github.com/xxjwxc/public/myleveldb"
	"github.com/xxjwxc/public/myqueue"
	"github.com/xxjwxc/public/tools"
)

var _que *myqueue.MyQueue
var lv myleveldb.MyLevelDB
var once sync.Once

func myinit() {
	_que = myqueue.New()
	lv = (myleveldb.OnInitDB(path.Join(tools.GetCurrentDirectory(), "file/database/leveldb_history")))

	go onDeal() // 开始消费
}

func getLevelDB() *myleveldb.MyLevelDB {
	once.Do(myinit)
	return &lv
}

func onDeal() {
	leveldb := getLevelDB()
	// defer leveldb.OnDestoryDB()

	for {
		info := _que.Pop().(queInfo)
		var tmp []string
		leveldb.Get(info.UserID, &tmp)
		var sort mysort.Lifo
		for _, v := range tmp {
			sort.Add(v)
		}
		sort.PushGrab(fmt.Sprintf("%v@%v", info.Gpid, info.Icon))

		its := sort.Gets()
		items := make([]string, 0, len(its))
		for _, v := range its {
			items = append(items, v.(string))
		}
		if len(items) > 10 {
			items = items[:10]
		}
		leveldb.Add(info.UserID, items)
	}
}

// gets 获取
func gets(userID string) (resp []queInfo) {
	leveldb := getLevelDB()
	// defer leveldb.OnDestoryDB()
	var items []string
	leveldb.Get(userID, &items)
	for _, v := range items {
		s := strings.Split(v, "@")
		if len(s) > 1 {
			resp = append(resp, queInfo{
				UserID: userID,
				Gpid:   s[0],
				Icon:   v[len(s[0])+1:],
			})
		}
	}

	return resp
}
