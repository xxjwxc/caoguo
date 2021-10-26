package notify

import (
	"caoguo/internal/config"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"fmt"
	"time"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/public/myemail"
)

// NotifyEmail 邮件提醒
func NotifyEmail(billId string) {
	orm := core.Dao.GetDBr()
	orders, _ := model.BillOrderTblMgr(orm.DB).GetFromBillID(billId)

	var gpids []string
	for _, v := range orders {
		gpids = append(gpids, v.Gpid)
	}

	mgr := model.ProductTblMgr(orm.DB)
	mgr.SetIsRelated(false) // 关闭外键查询
	products, _ := mgr.GetBatchFromGpid(gpids)

	mpname := make(map[string]string)
	for _, v := range products {
		mpname[v.Gpid] = v.PlatformID
	}

	mp := make(map[string][]*model.BillOrderTbl)
	for _, v := range orders {
		mp[mpname[v.Gpid]] = append(mp[mpname[v.Gpid]], v)
	}

	// 开始组织发送
	email := config.GetEmail()
	myemail := myemail.New(email.User, email.Password, email.Host, "LUCHAO")
	for k, v := range mp {
		to := getEmailList(k)
		if len(to) > 0 {
			var lplist string
			for _, v1 := range v {
				lplist += fmt.Sprintf("<p>商品名:(%v),商品属性:(%v),数量:(%v)</p>", v1.Name, v1.SkuArrt, v1.Number)
			}
			b, _ := myemail.SendMail(to, fmt.Sprintf(subject, tools.GetUtcTime(time.Now())),
				fmt.Sprintf(body, billId, lplist, tools.GetTimeStr(time.Now())))
			if b {
				mylog.Infof("send email %v success.", to)
			}
		}
	}
	// -------------end
}
