package order

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/rpc/caoguo"
	"fmt"
	"strings"
	"time"

	"caoguo/internal/model"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/tools"
)

// AddBillShipment 订单添加运单号
func (p *Order) AddBillShipment(c *api.Context, req *caoguo.AddBillShipmentReq) (bool, error) {
	user, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()
	billInfo, err := model.BillTblMgr(orm.DB).GetFromBillID(req.BillId)
	if err != nil {
		return false, err
	}

	if billInfo.Status != int(caoguo.BillStatus_Paid) && billInfo.Status != int(caoguo.BillStatus_Received) {
		return false, fmt.Errorf("状态不匹配")
	}

	tx := orm.Begin()

	if len(req.ShipmentId) > 0 {
		// 先查询
		_info, _ := model.ShipmentTblMgr(tx).FetchUniqueIndexByBillShipment(req.BillId, req.ShipmentId)
		if _info.ID > 0 {
			return false, fmt.Errorf("请不要重复添加。")
		}
		// --------end
		shipment := model.ShipmentTbl{
			BillID:     req.BillId,
			ShipmentID: req.ShipmentId,
			PostKey:    req.PostKey,
			Status:     2, // 快递状态  2在途中 3已签收 4 问题件
			CreatedAt:  time.Now(),
		}

		err = model.ShipmentTblMgr(tx).Save(&shipment).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}

		resp, _ := model.LogisticsTblMgr(orm.DB).GetFromShipmentID(req.ShipmentId)
		if len(resp) == 0 { // 没有就添加一个
			logistics := model.LogisticsTbl{
				ShipmentID: req.ShipmentId, // 快递单号
				Times:      time.Now(),     // 时间
				Context:    "添加快递单",        // 描述信息
				CreatedBy:  user.UserName,  // 创建者
				CreatedAt:  time.Now(),     // 创建时间
			}
			err = model.LogisticsTblMgr(orm.DB).Save(&logistics).Error
			if err != nil {
				tx.Rollback()
				return false, err
			}
		}
	}

	if billInfo.Status == int(caoguo.BillStatus_Paid) {
		err = model.BillTblMgr(tx).Where("bill_id = ?", req.BillId).Updates(map[string]interface{}{"status": int(caoguo.BillStatus_Received), "updated_at": time.Now(), "updated_by": user.UserName}).Error
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	// 到此成功
	tx.Commit()

	if len(req.ShipmentId) == 0 {
		return false, fmt.Errorf("快递单为空,订单已锁定，已发货。")
	}
	return true, nil
}

// GetBillInfoByAdmin 获取订单信息
func (p *Order) GetBillInfoByAdmin(c *api.Context, req *caoguo.GetMyOrdersReq) (*caoguo.GetMyOrdersResp, error) {
	user, b := c.GetUserInfo()
	if !b {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}
	resp := &caoguo.GetMyOrdersResp{
		TotalPages: 1,
		PageNumber: 0,
	}

	orm := core.Dao.GetDBr()
	db := orm.Offset(int(req.PageNumber * 10)).Limit(10).Order("bill_order_tbl.created_at desc")

	var condition []string
	condition = append(condition, "bill_tbl.status != 1")
	if req.Status != 0 {
		condition = append(condition, fmt.Sprintf("bill_tbl.status = %v", req.Status))
	}
	if user.UserName != "admin" {
		condition = append(condition, fmt.Sprintf("product_tbl.platform_id = %v", user.UserName))
	}
	if user.UserName == "13795896867" {
		condition = append(condition, "bill_tbl.created_at > '2020-9-13 00:00:00'")
	}
	if len(req.Search) > 0 {
		condition = append(condition, fmt.Sprintf("(bill_address_tbl.name like('%v') or bill_address_tbl.mobile = '%v')", "%"+req.Search+"%", req.Search))
	}

	db = db.Raw(fmt.Sprintf(`SELECT bill_order_tbl.*,bill_tbl.status,bill_coupon_tbl.coupon_id,bill_coupon_tbl.title FROM bill_order_tbl 
		inner join product_tbl on product_tbl.gpid = bill_order_tbl.gpid
		left join bill_coupon_tbl on bill_coupon_tbl.bill_id = bill_order_tbl.bill_id 
		INNER JOIN bill_address_tbl on bill_address_tbl.bill_id = bill_order_tbl.bill_id
		inner join bill_tbl on  bill_tbl.bill_id = bill_order_tbl.bill_id and bill_tbl.deleted_at IS NULL and %v`, strings.Join(condition, " and "))) // 过滤掉代付款订单

	var orders []*BillOrderTbl
	err := db.Scan(&orders).Error
	if err != nil {
		if orm.IsNotFound(err) {
			return resp, nil
		}
		return nil, err
	}

	for _, v := range orders {
		tm := v.UpdatedAt.Unix()
		if tm == 0 {
			tm = v.CreatedAt.Unix()
		}
		billInfo := &caoguo.BillOrderInfo{
			Status: int32(v.Status), // 状态 1:成功，-1:失败
			BillId: v.BillID,        // 账单id
			// TotalFee:    v.Price,         //  总费用
			Time:        tools.GetTimeStr(tools.UnixToTime(tm)),
			CheckStatus: getCheckState(int64(v.Status)),
			CouponTitle: v.Title,
		}
		billInfo.StateTip, billInfo.StateTipColor = getStateTip(int64(v.Status), v.CreatedAt)

		skuArrt := v.SkuArrt
		if len(skuArrt) == 0 {
			skuArrt = " "
		}
		item := &caoguo.OrderProductInfo{
			Gpid:    v.Gpid,          // 商品gpid
			Name:    v.Name,          // 商品名
			Price:   v.Price,         // 商品价格
			Icon:    v.Icon,          // 商品图标
			SkuVal:  skuArrt,         // sku描述
			SkuList: v.SkuList,       // sku信息
			Number:  int64(v.Number), // 购买数量
		}
		isFind := false
		for _, v1 := range resp.Items {
			if v.BillID == v1.BillId {
				v1.Items = append(v1.Items, item)
				v1.Number += item.Number
				isFind = true
				break
			}
		}

		if !isFind { // 没找到，添加新的
			billInfo.Items = append(billInfo.Items, item)
			billInfo.Number += item.Number
			// 获取shipment 信息
			sps, _ := model.ShipmentTblMgr(orm.DB).GetFromBillID(v.BillID)
			for _, v := range sps {
				billInfo.ShipmentInfo = append(billInfo.ShipmentInfo, &caoguo.ShipmentInfo{
					ShipmentId:   v.ShipmentID,
					ShipmentName: v.PostTbl.Name,
				})
			}
			// -------------end
			// 获取地址信息
			addinfo, _ := model.BillAddressTblMgr(orm.DB).GetFromBillID(v.BillID)

			if addinfo.ID > 0 {
				addinfo.Name = tools.UnicodeEmojiDecode(addinfo.Name)
				addinfo.AddressName = tools.UnicodeEmojiDecode(addinfo.AddressName)
				addinfo.Address = tools.UnicodeEmojiDecode(addinfo.Address)
				addinfo.Area = tools.UnicodeEmojiDecode(addinfo.Area)
				billInfo.Addr = &caoguo.AddressInfo{
					Name:    addinfo.Name,                                                                // 名字
					Mobile:  addinfo.Mobile,                                                              // 手机号
					Address: fmt.Sprintf("%v,%v,%v", addinfo.Address, addinfo.AddressName, addinfo.Area), // 详细地址
				}
			}
			// -------------end
			resp.Items = append(resp.Items, billInfo)
		}
	}

	return resp, nil
}

// GetAfterMsg 获取回复信息
func (p *Order) GetAfterMsg(c *api.Context, req *caoguo.AfterMsgReq) (*caoguo.AfterMsgResp, error) {
	orm := core.Dao.GetDBr()
	list, err := model.BillDealLogTblMgr(orm.Where("type = 1").Order("created_at desc")).GetFromBillID(req.BillId)
	if err != nil {
		if orm.IsNotFound(err) {
			return nil, fmt.Errorf("暂无售后信息")
		}
	}

	resp := &caoguo.AfterMsgResp{
		BillId: req.BillId,
	}
	for _, v := range list {
		var img []string
		tools.JSONEncode(v.Imgs, &img)
		resp.Items = append(resp.Items, &caoguo.AfterMsg{
			Contact: v.Contact,
			Remak:   v.Remak,
			Time:    tools.GetTimeStr(v.CreatedAt),
			Imgs:    img,
		})
	}
	return resp, nil
}
