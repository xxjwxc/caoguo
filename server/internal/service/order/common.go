package order

import (
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/notify"
	"caoguo/internal/service/weixin"
	"math/rand"
	"time"

	"caoguo/rpc/caoguo"

	"github.com/jinzhu/gorm"
	"github.com/xxjwxc/public/mylog"
	wx "github.com/xxjwxc/public/weixin"
)

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

// WxQueryCallBack 支付确认定时任务
func WxQueryCallBack() {
	//获取15分之前的数据
	offTimes := time.Now().Add(-30 * time.Minute) // 获取15分钟之前的数据(超过15分钟自动取消)
	mylog.Info("start WxQueryCallBack")
	orm := core.Dao.GetDBr()
	mgr := model.BillTblMgr(orm.Where("created_at < ?", offTimes))
	mgr.SetIsRelated(false)
	bills, err := mgr.GetFromStatus(1)
	if err != nil {
		return
	}

	for _, v := range bills {
		if len(v.BillID) > 0 {
			checkOrder(v)
		}
	}
}

func checkOrder(bill *model.BillTbl) bool {
	ret, msg := weixin.SelectOrder(bill.UserID, bill.BillID)
	mylog.Info(msg)
	switch ret {
	case wx.PAY_SUCCESS: //支付成功
		updateBillStatus(bill.BillID, int(caoguo.BillStatus_Paid))
		offsetUserPoints(bill.UserID, "system", bill.PayAmount) // 添加用户积分
		return true
	case wx.PAY_REFUND, wx.PAY_CLOSED, wx.PAY_NOTPAY, wx.PAY_ERROR: //转入退款 已关闭 未支付
		if bill.CreatedAt.Before(time.Now().Add(-2 * time.Hour)) { // 超过15分钟自动取消
			updateBillStatus(bill.BillID, int(caoguo.BillStatus_Cancelled))
		}
		return true
	}
	return false
}

// 更新bill状态 支付状态(-1:已退款,0:待支付,1:已支付,2:已取消)
func updateBillStatus(billID string, status int) {
	orm := core.Dao.GetDBw()
	orm.Model(&model.BillTbl{}).Where("bill_id = ?", billID).Updates(map[string]interface{}{"status": status, "updated_at": time.Now()})

	if status == int(caoguo.BillStatus_Cancelled) { // 取消，可能退优惠券
		coupon, err := model.BillCouponTblMgr(orm.DB).GetFromBillID(billID)
		if err != nil {
			return
		}
		if coupon.CouponID > 0 {
			orm.Model(&model.CouponTbl{}).Where("id = ?", coupon.CouponID).Update("status", 1)
		}
	}

	if status == int(caoguo.BillStatus_Paid) { // 已支付
		go notify.NotifyEmail(billID) // 邮件提醒
	}
}

// 设置用户积分
func offsetUserPoints(userid, updateBy string, points int64) {
	orm := core.Dao.GetDBw()
	orm.Model(&model.UserTbl{}).Where("user_id = ?", userid).Updates(map[string]interface{}{"points": gorm.Expr("points + ?", points), "updated_at": time.Now(), "updated_by": updateBy})
}
