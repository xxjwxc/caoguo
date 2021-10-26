package order

import (
	"caoguo/internal/api"
	"caoguo/internal/config"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"fmt"
	"strconv"
	"time"

	"github.com/xxjwxc/public/mykdniao"
	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/public/tools"
)

// GetOrdertrackInfo 获取物流轨迹信息
func (p *Order) GetOrdertrackInfo(c *api.Context, req *caoguo.GetOrdertrackInfoReq) (*caoguo.GetOrdertrackInfoResp, error) {
	checkOrderTrack(req.BillId) // 先检查一次
	orm := core.Dao.GetDBr()
	shipments, err := model.ShipmentTblMgr(orm.DB).GetFromBillID(req.BillId)
	if err != nil {
		return nil, err
	}

	if len(shipments) == 0 {
		return nil, fmt.Errorf("工作人员正在录入中,请耐心等待")
	}

	var shipment *model.ShipmentTbl // 在途中 快递单
	for _, v := range shipments {
		if v.Status != 3 {
			shipment = v
			break
		}
	}
	if shipment == nil { // 无在途中订单 直接取第一个
		shipment = shipments[0]
	}

	addr, err := model.BillAddressTblMgr(orm.DB).GetFromBillID(req.BillId)
	if err != nil {
		return nil, err
	}

	addr.Name = tools.UnicodeEmojiDecode(addr.Name)
	addr.AddressName = tools.UnicodeEmojiDecode(addr.AddressName)
	addr.Address = tools.UnicodeEmojiDecode(addr.Address)
	addr.Area = tools.UnicodeEmojiDecode(addr.Area)

	resp := &caoguo.GetOrdertrackInfoResp{
		BillId:         req.BillId,
		DeliveryStatus: int32(shipment.Status),                                             // 快递状态  2-在途中,3-签收,4-问题件
		PostName:       shipment.PostTbl.Name,                                              // 快递名称
		Logo:           shipment.PostTbl.Icon,                                              // 快递logo
		ExpPhone:       shipment.PostTbl.ExpPhone,                                          // 快递电话
		PostNo:         shipment.ShipmentID,                                                // 快递单号
		Addr:           fmt.Sprintf("%v,%v,%v", addr.Address, addr.AddressName, addr.Area), // 收货地址
	}

	logistics, err := model.LogisticsTblMgr(orm.Order("times desc")).GetFromShipmentID(shipment.ShipmentID)
	if err != nil {
		return nil, err
	}

	for _, v := range logistics {
		resp.List = append(resp.List, &caoguo.Logistics{
			Time:    tools.GetTimeStr(v.Times),                                                 // 时间
			TimeArr: []string{tools.GetDayStr(v.Times), tools.FormatTime(v.Times, "15:04:05")}, // 时间分割(['2020-04-12', '13:00:57'])
			Context: v.Context,                                                                 // 描述信息
		})
	}

	return resp, nil
}

// GetShipmentPost 获取快递列表
func (p *Order) GetShipmentPost(c *api.Context, req *common.Empty) (*caoguo.GetShipmentPostResp, error) {
	orm := core.Dao.GetDBr()
	posts, err := model.PostTblMgr(orm.DB).Gets()
	if err != nil {
		return nil, err
	}
	resp := &caoguo.GetShipmentPostResp{}
	for _, v := range posts {
		resp.Items = append(resp.Items, &caoguo.PostInfo{
			Name:     v.Name, // 快递名称
			Id:       v.Code, // 快递id
			Icon:     v.Icon, // 快递logo地址
			ExpPhone: v.ExpPhone,
		})
	}
	return resp, nil
}

// CheckOrderTrack 物流抓取定时任务
func CheckOrderTrack() {
	checkOrderTrack("")
}

// checkOrderTrack 物流抓取定时任务
func checkOrderTrack(billID string) {
	orm := core.Dao.GetDBr()
	db := orm.DB
	if len(billID) > 0 {
		db = db.Where("bill_id = ?", billID)
	}
	shipments, err := model.ShipmentTblMgr(db).GetBatchFromStatus([]int{0, 1, 2, 4})
	if err != nil {
		mylog.Error(err)
		return
	}

	for _, v := range shipments {
		if err := UpdateShipmentInfo(v); err != nil {
			mylog.Error(err)
		}
	}
}

// UpdateShipmentInfo 更新运单及账单信息
func UpdateShipmentInfo(shipment *model.ShipmentTbl) error {
	if len(shipment.ShipmentID) == 0 { // 没有快递单，返回
		return nil
	}
	orm := core.Dao.GetDBw()
	logistics, err := model.LogisticsTblMgr(orm.Order("times desc")).GetFromShipmentID(shipment.ShipmentID)
	if err != nil {
		mylog.Error(err)
		return err
	}
	billAddr, _ := model.BillAddressTblMgr(orm.DB).GetFromBillID(shipment.BillID)
	var phone string

	if shipment.PostKey == "SF" { // 顺丰添加快递
		phone = billAddr.Mobile
		if len(phone) > 4 {
			phone = phone[len(phone)-4:]
		}
		phone = ""
	}

	var newTime time.Time
	for _, v := range logistics {
		if v.Context != "添加快递单" {
			newTime = v.Times
			break
		}
	}

	kdn := mykdniao.New(config.GetKdniao().EBusinessID, config.GetKdniao().AppKey)

	result := kdn.GetLogisticsTrack(shipment.ShipmentID, shipment.PostKey, phone)
	if result.Success {
		tx := orm.Begin()
		// 更新运单信息
		for _, v1 := range result.Traces {
			tm := tools.StrToTime(v1.AcceptTime, "2006-01-02 15:04:05", nil)
			if tm.After(newTime) {
				model.LogisticsTblMgr(tx).Save(&model.LogisticsTbl{
					ShipmentID: result.LogisticCode, // 快递单号
					Times:      tm,                  // 时间
					Context:    v1.AcceptStation,    // 描述信息
					CreatedAt:  time.Now(),
				})
			}
		}
		status, _ := strconv.Atoi(result.State)
		if status > 0 && shipment.Status != status { // 快递状态  2在途中 3已签收 4 问题件
			shipment.Status = status
			model.ShipmentTblMgr(tx).Where("shipment_id = ?", shipment.ShipmentID).Updates(map[string]interface{}{"status": status, "updated_at": time.Now()})
			if status == 3 { // 已签收 (修改bill状态 已完成)
				isUpdateBill := true
				sps, err := model.ShipmentTblMgr(tx.Where("status != ?", 3)).GetFromBillID(shipment.BillID)
				if err != nil {
					if !orm.IsNotFound(err) {
						tx.Rollback()
						return err
					}
				}
				for _, v := range sps {
					if v.Status == 2 {
						isUpdateBill = false
					}
				}
				if isUpdateBill {
					err = model.BillTblMgr(tx).Where("bill_id = ?", shipment.BillID).Updates(map[string]interface{}{"status": 5, "updated_at": time.Now()}).Error
					if err != nil {
						mylog.Error(err)
						tx.Rollback()
						return err // 跳过
					}
				}
			}
		}
		tx.Commit()
	} else {
		mylog.Errorf("抓取物流信息失败:%v,%v", shipment.ShipmentID, result.Reason)
	}

	return nil
}
