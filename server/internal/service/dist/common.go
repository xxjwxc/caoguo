package dist

import (
	"caoguo/internal/core"
	"caoguo/internal/model"
	"fmt"
	"time"

	"github.com/xxjwxc/public/mylog"
)

func getMoneyStr(money int64) string {
	return fmt.Sprintf("￥%.2f", float32(money)*0.01)
}

// ConnectionDistUserID 关联分销
func ConnectionDistUserID(userID, parntID string) {
	orm := core.Dao.GetDBr()
	// 查找是否已经有关联分销关系了
	distUser, err := model.DistUserinfoTblMgr(orm.DB).GetFromUserID(userID)
	if err != nil && !orm.IsNotFound(err) {
		mylog.Error(err)
		return
	}
	if distUser.ID > 0 { // 已经有关联关系了
		return
	}

	// 判断是否有分销权限
	user, err := model.UserTblMgr(orm.DB).GetFromUserID(parntID)
	if err != nil && !orm.IsNotFound(err) {
		mylog.Error(err)
		return
	}
	if !user.DistVip { // 没有分销权限
		return
	}
	// 添加关联关系
	distUserInfo := &model.DistUserinfoTbl{
		UserID:    userID,     // 子节点
		ParntID:   parntID,    // 父节点
		CreatedAt: time.Now(), // 创建时间
	}
	err = model.DistUserinfoTblMgr(orm.DB).Save(distUserInfo).Error
	if err != nil {
		mylog.Error(err)
	}
	return
}

// DistOrderMoney 开始分销
func DistOrderMoney(billID string) {
	// 查询账单信息
	orm := core.Dao.GetDBw()
	// 查询是否有分销过
	var count int64
	model.DistOrderTblMgr(orm.DB).Where("bill_id = ?", billID).Count(&count)
	if count > 0 { // 已做过分销 ，禁止
		return
	}

	billInfo, err := model.BillTblMgr(orm.DB).GetFromBillID(billID)
	if err != nil {
		mylog.Error(err)
		return
	}

	if billInfo.Status != 2 && billInfo.Status != 4 && billInfo.Status != 5 && billInfo.Status != 6 { // 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
		mylog.Error(fmt.Errorf("bill dist status error :%v, %v", billID, billInfo.Status))
		return
	}

	// 获取分销用户层级
	distUserList := getDistUserList(billInfo.UserID)
	if len(distUserList) == 0 { // 无分销
		return
	}
	// ------------end

	// 开始分销
	billOrders, err := model.BillOrderTblMgr(orm.DB).GetFromBillID(billID)
	if err != nil {
		mylog.Error(err)
		return
	}

	status := 1 // 状态(-1:失效,1:待确认,2:可提现，3:已提现)
	if billInfo.Status == 4 || billInfo.Status == 5 || billInfo.Status == 6 {
		status = 2
	}

	var distOrderlist []model.DistOrderTbl
	for _, v := range billOrders {
		if v.DistAmount > 0 {
			totalDistMoney := v.DistAmount * int64(v.Number)
			remainMoney := totalDistMoney
			for i := 0; i < len(distUserList); i++ {
				price := remainMoney / 2
				if price == 0 { // offset
					price = remainMoney
				}

				distOrderlist = append(distOrderlist, model.DistOrderTbl{
					BillID:     v.BillID,        // 账单id
					UserID:     distUserList[i], // 分销给的账号
					Gpid:       v.Gpid,          // 商品id
					Name:       v.Name,          // 商品名字
					Price:      price,           // 分销金额
					Level:      i + 1,           // 分销等级
					TotalPrice: totalDistMoney,  // 总分销金额
					Status:     status,          // 状态(-1:失效,1:待确认,2:可提现，3:已提现)
					Number:     v.Number,        // 商品数量
					CreatedAt:  time.Now(),      // 创建时间
				})

				remainMoney -= price
				if remainMoney <= 0 {
					break
				}
			}
		}
	}

	if len(distOrderlist) > 0 {
		err := model.DistOrderTblMgr(orm.DB).Create(&distOrderlist).Error
		if err != nil {
			mylog.Error(err)
		}
	}
}

// ChangeDistStatus 修改分销状态
// 状态(-1:失效,1:待确认,2:可提现，3:已提现)
func ChangeDistStatus(billID string, status int) error {
	if status == 1 { // 不允许 1:待确认状态修改
		return nil
	}
	orm := core.Dao.GetDBw()

	// 直接暴力更新
	return model.DistOrderTblMgr(orm.DB).Where("bill_id = ?", billID).Updates(map[string]interface{}{"status": status, "updated_at": time.Now()}).Error
}

// getDistUserList 获取分销层级
func getDistUserList(userID string) []string {
	var list []string
	orm := core.Dao.GetDBr()
	for {
		tmp, err := model.DistUserinfoTblMgr(orm.DB).GetFromUserID(userID)
		if err != nil || tmp.ID == 0 {
			break
		}
		list = append(list, tmp.ParntID)
		userID = tmp.ParntID
	}
	return list
}
