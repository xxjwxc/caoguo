package oplog

import (
	"caoguo/internal/model"
	"time"

	"caoguo/internal/core"

	"github.com/xxjwxc/public/tools"
)

// AddOpLog	添加日志数据
func AddOpLog(operator, receive, topic, bundle, pid, ipAddr string, num0 int, Attach0 string) bool {
	var info model.OpLogTbl
	info.Operator = operator
	info.Receive = receive
	info.CreateTime = time.Now()
	info.Topic = topic
	info.Bundle = bundle
	info.Pid = pid
	info.IPAddr = ipAddr
	info.Num0 = num0
	info.Attach0 = Attach0

	return OpLogBaseAdd(&info)
}

// OpLogBaseAdd 添加日志数据
func OpLogBaseAdd(info *model.OpLogTbl) bool {
	info.CreateTime = time.Now() //默认时间
	core.Dao.GetDBw().Create(info)
	return true
}

// GetWxLog 1：收藏 2：点赞 3：订单 0:全部
func GetWxLog(openID string, action int) (result []*Wx_user_log_view) {
	orm := core.Dao.GetDBw()
	if action == 0 { //全部
		orm.Where("open_id = ?", openID).Order("create_time desc").Find(&result)
	} else {
		orm.Where("open_id = ? and type = ?", openID, action).Order("create_time desc").Find(&result)
	}
	if len(result) > 0 {
		for _, v := range result {
			tools.JSONEncode(v.Pic_url_str, &v.Pic_url)
		}
	}
	return
}
