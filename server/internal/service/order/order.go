package order

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/coupon"
	"caoguo/internal/service/weixin"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/tools"
)

// Order 订单相关
type Order struct {
}

// CartTmp 购物车模板
type CartTmp struct {
	ID       int64
	Gpid     string // 商品id
	UserID   string // 用户id
	UserType int    // 用户类型(1:微信)
	Number   int    // 数量
	SkuList  string // 属性列表
}

// ReckonFee 费用计算
func (p *Order) ReckonFee(c *api.Context, req *caoguo.ReckonFeeReq) (*caoguo.ReckonFeeResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}
	cartTmp, err := getCartTmp(session.Openid, req.BuyType, req.Ids)
	if err != nil {
		return nil, err
	}

	var gpids, skuids []string
	for _, v := range cartTmp {
		gpids = append(gpids, v.Gpid)
		if len(v.SkuList) > 0 {
			skuids = append(skuids, strings.Split(v.SkuList, ",")...)
		}
	}
	orm := core.Dao.GetDBr()
	// 获取商品信息
	result, err := model.ProductTblMgr(orm.DB).GetBatchFromGpid(gpids)
	if err != nil {
		if !orm.IsNotFound(err) {
			return nil, fmt.Errorf(message.NotFindError.String())
		}
	}
	productMp := make(map[string]*model.ProductTbl, len(result))
	for _, v := range result {
		productMp[v.Gpid] = v
	}

	// 获取sku信息
	skuMp := make(map[string]*model.ProductSkuTbl, len(skuids))
	if len(skuids) > 0 {
		result1, err := model.ProductSkuTblMgr(orm.Where("id IN (?)", skuids)).Gets()
		if err != nil {
			if !orm.IsNotFound(err) {
				return nil, fmt.Errorf(message.NotFindError.String())
			}
		}
		for _, v := range result1 {
			skuMp[fmt.Sprintf("%v", v.ID)] = v
		}
	}

	// 开始计算费用
	resp := &caoguo.ReckonFeeResp{}
	var couponFee int64   // 优惠券
	var shipmentFee int64 // 运费
	// var discountFee int64 // 优惠金额
	// var totalFee int64    // 总费用

	for _, v := range cartTmp {
		product := productMp[v.Gpid]
		price := product.Price
		distAmount := product.DistAmount
		if len(v.SkuList) > 0 {
			skuprice, _ := model.ProductSkuPriceTblMgr(orm.DB).GetFromSkuList(v.SkuList)
			price = skuprice.Premium
			distAmount = skuprice.DistAmount // 走skuprice中拿
		}
		skulist := strings.Split(v.SkuList, ",")
		var skuStr []string
		for _, v := range skulist {
			if len(v) > 0 {
				skuStr = append(skuStr, skuMp[v].TagName)
			}
		}

		resp.OrderInfo = append(resp.OrderInfo, &caoguo.OrderProductInfo{
			Gpid:       v.Gpid,
			Name:       product.Name,
			Price:      price,
			Icon:       product.ProductInfoTbl.Icon,
			SkuVal:     strings.Join(skuStr, " "),
			SkuList:    v.SkuList,
			Number:     int64(v.Number),
			DistAmount: distAmount,
		})
		resp.TotalFee += price * int64(v.Number)
		shipmentFee += product.ShipmentFee * int64(v.Number)
		// discountFee += (product.OriginalPrice - price) * int64(v.Number)
	}

	resp.TotalFee += shipmentFee // 总费用添加运费
	// 开始制作费用明细
	// 优惠券
	resp.CouponDetail = "暂无可用优惠券"
	resp.PromotionDetail = "平台限时优惠"
	resp.VendorName = "订单列表"
	resp.VendorImg = "/static/temp/h1.png"
	// resp.FeeDetails = append(resp.FeeDetails, &caoguo.FeeDetail{
	// 	Key:   "原价",
	// 	Value: getMoneyStr(resp.TotalFee + discountFee),
	// })

	resp.FeeDetails = append(resp.FeeDetails, &caoguo.FeeDetail{
		Key:   "商品金额",
		Value: getMoneyStr(resp.TotalFee),
	})

	// resp.FeeDetails = append(resp.FeeDetails, &caoguo.FeeDetail{
	// 	Key:   "已优惠金额",
	// 	Value: getMoneyStr(-discountFee),
	// 	Color: "red",
	// }) // 总金额

	var shipmentFeeStr string
	if shipmentFee == 0 {
		shipmentFeeStr = "免运费"
	} else {
		shipmentFeeStr = getMoneyStr(shipmentFee)
	}
	resp.FeeDetails = append(resp.FeeDetails, &caoguo.FeeDetail{
		Key:   "运费",
		Value: shipmentFeeStr,
	})
	//----- end
	resp.ShipmentFee = shipmentFee

	// 开始获取 coupon
	couponlist := coupon.GetCartVaildCoupon(c, resp)
	for _, v := range couponlist {
		resp.CouponList = append(resp.CouponList,
			&caoguo.CouponInfo{
				Id:         v.ID,                                                                  // coupon id
				GroupName:  v.CouponTbl.GroupName,                                                 // 分组名
				Title:      v.CouponTbl.Title,                                                     // 优惠券名字
				Validity:   fmt.Sprintf("有效期至 %v", tools.FormatTime(v.ExpiresTime, "2006-01-02")), // 有效截止日期
				Gpid:       v.CouponTbl.Gpid,                                                      // 商品优惠券
				Denom:      int64(v.CouponTbl.Denom),                                              // 面额
				CouponType: int32(v.CouponTbl.CouponType),                                         // 优惠券类型(1:全场，2:单个商品)
				GreatMoney: coupon.GetGreatMoney(float32(v.CouponTbl.GreatMoney)),                 // 满多少可用
				Describle:  v.CouponTbl.Describle,                                                 // 优惠券描述
				Status:     1,                                                                     // 0：默认，1:未使用(有效),2:已使用,-1:已过期
				Type:       2,
			})
		resp.CouponDetail = "请选择优惠券"
	}

	if req.CouponId == 0 { // 默认设置一个
		max := 0
		for _, v := range couponlist {
			if v.CouponTbl.Denom > max {
				req.CouponId = v.ID
				max = v.CouponTbl.Denom
			}
		}
	}

	// 计算优惠金额
	couponFee, err = coupon.ModifyCoupon(c, req.CouponId, resp)
	if err != nil {
		return nil, err
	}
	if couponFee > 0 {
		resp.CouponFee = couponFee
		resp.TotalFee -= couponFee
		resp.FeeDetails = append(resp.FeeDetails, &caoguo.FeeDetail{
			Key:   "抵用券",
			Value: getMoneyStr(-couponFee),
			Color: "red",
		}) // 总金额

		for _, v := range couponlist {
			if v.ID == req.CouponId {
				resp.CouponDetail = fmt.Sprintf("%v %v", v.CouponTbl.Title, getMoneyStr(couponFee))
				resp.CouponId = v.ID
				break
			}
		}

	}

	return resp, nil
}

func getMoneyStr(money int64) string {
	return fmt.Sprintf("￥%.2f", float32(money)*0.01)
}

func getCartTmp(openid string, buyType int32, ids []int64) ([]*CartTmp, error) {
	var tmp []*CartTmp
	orm := core.Dao.GetDBr()
	switch buyType {
	case 1: // 1,来自直接购买
		resp, err := model.CartTmpTblMgr(orm.Where("user_id = ?", openid)).GetBatchFromID(ids)
		if err != nil {
			return nil, err
		}
		for _, v := range resp {
			tmp = append(tmp, &CartTmp{
				ID:       v.ID,
				Gpid:     v.Gpid,     // 商品id
				UserID:   v.UserID,   // 用户id
				UserType: v.UserType, // 用户类型(1:微信)
				Number:   v.Number,   // 数量
				SkuList:  v.SkuList,  // 属性列表
			})
		}
	case 2: // 2:来自购物车
		resp, err := model.CartTblMgr(orm.Where("user_id = ?", openid)).GetBatchFromID(ids)
		if err != nil {
			return nil, err
		}
		for _, v := range resp {
			tmp = append(tmp, &CartTmp{
				ID:       v.ID,
				Gpid:     v.Gpid,     // 商品id
				UserID:   v.UserID,   // 用户id
				UserType: v.UserType, // 用户类型(1:微信)
				Number:   v.Number,   // 数量
				SkuList:  v.SkuList,  // 属性列表
			})
		}
	}

	return tmp, nil
}

// PlaceOrder 下单
func (p *Order) PlaceOrder(c *api.Context, req *caoguo.PlaceOrderReq) (_resp *caoguo.PayResp, _err error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	if req.AddrId <= 0 {
		return nil, fmt.Errorf("请选择地址")
	}

	// 费用计算
	resp, err := p.ReckonFee(c, &caoguo.ReckonFeeReq{
		BuyType:  req.BuyType, // 购买类型:1,来自直接购买，2:来自购物车
		Ids:      req.Ids,     // 购买列表
		AddrId:   req.AddrId,  // 地址id
		CouponId: req.CouponId,
	})

	if err != nil {
		return nil, err
	}

	if req.CouponId != resp.CouponId { // 自动选择优惠券
		req.CouponId = resp.CouponId
	}

	if resp.TotalFee <= 0 {
		return nil, fmt.Errorf(message.UnknownError.String())
	}

	orm := core.Dao.GetDBw()
	tx := orm.DB.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()
	billMgr := model.BillTblMgr(tx)

	var isTrue = false
	var bill model.BillTbl
	var times [5][0]int
	for range times {
		bill.BillID = createUnique("")
		_, err := billMgr.GetFromBillID(bill.BillID)
		if orm.IsNotFound(err) {
			isTrue = true
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if !isTrue {
		return nil, fmt.Errorf(message.NotFindError.String())
	}
	bill.UserID = session.Openid
	bill.PayPlatform = 1           // 支付类型(1:微信支付)
	bill.PayAmount = resp.TotalFee // 支付金额
	bill.Status = 1                // 支付状态(-1:已退款,1:待支付,2:已支付,3:已取消)
	bill.CreatedAt = time.Now()
	bill.Desc = req.Desc

	// 删除购物车
	switch req.BuyType {
	case 1: // 1,来自直接购买
		err = model.CartTmpTblMgr(tx.Where("id in (?)", req.Ids)).Delete(&model.CartTmpTbl{}).Error
		if err != nil {
			return nil, err
		}
	case 2: // 2:来自购物车
		err := model.CartTblMgr(tx.Where("id in (?)", req.Ids)).Delete(&model.CartTbl{}).Error
		if err != nil {
			return nil, err
		}
	}
	// ---------- end

	// 派送信息
	addrMgr := model.AddressTblMgr(tx)
	addr, err := addrMgr.GetFromID(req.AddrId)
	if err != nil {
		return nil, err
	}
	var billAddr model.BillAddressTbl
	billAddr.AddrID = addr.ID               // 地址id
	billAddr.UserID = session.Openid        // 用户id
	billAddr.BillID = bill.BillID           // 账单id
	billAddr.Name = addr.Name               // 用户名
	billAddr.Mobile = addr.Mobile           // 手机号
	billAddr.AddressName = addr.AddressName // 地址名称
	billAddr.Address = addr.Address         // 详细地址
	billAddr.Area = addr.Area               // 单元门牌号
	billAddr.CreatedAt = time.Now()         // 创建时间
	// ------end

	// 账单订单列表
	for _, v := range resp.OrderInfo {
		billOrder := model.BillOrderTbl{
			BillID:     bill.BillID,    // 账单id
			Gpid:       v.Gpid,         // 商品id
			Name:       v.Name,         // 商品名字
			Price:      v.Price,        // 商品价格(分)
			Icon:       v.Icon,         // 图标信息
			UserID:     session.Openid, // 用户id
			Number:     int(v.Number),  // 数量
			DistAmount: v.DistAmount,   // 分销金额
			SkuList:    v.SkuList,      // 属性列表
			SkuArrt:    v.SkuVal,       // 属性描述
			CreatedAt:  time.Now(),
		}
		err = tx.Save(&billOrder).Error
		if err != nil {
			return nil, err
		}

		// 更新已购数量
		err = tx.Model(&model.ProductInfoTbl{}).Where("gpid = ?", v.Gpid).UpdateColumn("qty", gorm.Expr("qty + ?", v.Number)).Error
		if err != nil {
			return nil, err
		}
	}

	err = tx.Model(&model.BillAddressTbl{}).Save(&billAddr).Error
	if err != nil {
		return nil, err
	}

	err = billMgr.Save(&bill).Error
	if err != nil {
		return nil, err
	}

	// 优惠券列表池
	if req.CouponId > 0 {
		// 修改优惠券成已使用
		coupon, err := model.CouponMyTblMgr(tx.Where("user_id = ?", session.Openid)).GetFromID(req.CouponId)
		if err != nil {
			return nil, err
		}
		coupon.UpdatedAt = time.Now()
		coupon.Status = 2
		err = tx.Save(&coupon).Error
		if err != nil {
			return nil, err
		}

		// 添加优惠券记录
		billCoupon := model.BillCouponTbl{
			BillID:     bill.BillID,
			CouponID:   coupon.ID,                   // 我领取的优惠券id
			Title:      coupon.CouponTbl.Title,      // 优惠券名字
			Denom:      coupon.CouponTbl.Denom,      // 面额
			CouponType: coupon.CouponTbl.CouponType, // 优惠券类型(1:全场，2:单个商品)
			CreatedAt:  time.Now(),                  // 创建时间
		}
		err = tx.Save(&billCoupon).Error
		if err != nil {
			return nil, err
		}
	}

	// 开始开始调起支付
	mp, err := weixin.PayPlaceOrder(session.Openid, bill.PayAmount, "用户下单支付", bill.BillID, c.GetClientIP())
	if err != nil {
		tx.AddError(err)
		return nil, err
	}

	// 到此完成
	return &caoguo.PayResp{
		Status: true,        // 状态 1:成功，-1:失败
		BillId: bill.BillID, // 账单id
		Info:   mp,
	}, nil
}

// Pay 支付
func (p *Order) Pay(c *api.Context, req *caoguo.PayReq) (*caoguo.PayResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBr()
	bill, err := model.BillTblMgr(orm.DB).GetFromBillID(req.BillId)
	if err != nil {
		return nil, err
	}

	switch bill.Status { // 支付状态(-1:已退款,1:待支付,2:已支付,3:已取消)
	case -1:
		return nil, fmt.Errorf("该订单已退款,不能发起支付")
	case 2, 4, 5, 6:
		return nil, fmt.Errorf("该笔订单已支付,请不要重复支付")
	case 3:
		return nil, fmt.Errorf("该订单已取消,请重新下单")
	}

	// 开始开始调起预支付
	mp, err := weixin.PayPlaceOrder(session.Openid, bill.PayAmount, "用户下单支付", bill.BillID, c.GetClientIP())
	if err != nil {
		return nil, err
	}

	// 到此完成
	return &caoguo.PayResp{
		Status: true,        // 状态 1:成功，-1:失败
		BillId: bill.BillID, // 账单id
		Info:   mp,
	}, nil
}

// CheckPay 支付成功check一次
func (p *Order) CheckPay(c *api.Context, req *caoguo.PayReq) (*caoguo.PayResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}
	orm := core.Dao.GetDBr()
	mgr := model.BillTblMgr(orm.DB)
	mgr.SetIsRelated(false)
	bill, err := mgr.GetFromBillID(req.BillId)
	if err != nil {
		return nil, err
	}

	b := checkOrder(&bill)
	return &caoguo.PayResp{
		Status: b,
		BillId: req.BillId,
	}, nil
}

// GetMyOrders 获取所有订单
func (p *Order) GetMyOrders(c *api.Context, req *caoguo.GetMyOrdersReq) (*caoguo.GetMyOrdersResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}
	orm := core.Dao.GetDBr()

	db := orm.Offset(int(req.PageNumber * 10)).Limit(10).Order("created_at desc")
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	resp := &caoguo.GetMyOrdersResp{
		TotalPages: 1,
		PageNumber: 0,
	}

	bills, err := model.BillTblMgr(db).GetFromUserID(session.Openid)
	if err != nil {
		if orm.IsNotFound(err) {
			return resp, nil
		}
		return nil, err
	}

	for _, v := range bills {
		tm := v.UpdatedAt.Unix()
		if tm == 0 {
			tm = v.CreatedAt.Unix()
		}
		billInfo := &caoguo.BillOrderInfo{
			Status:      int32(v.Status), // 状态 1:成功，-1:失败
			BillId:      v.BillID,        // 账单id
			TotalFee:    v.PayAmount,     //  总费用
			Time:        tools.GetTimeStr(tools.UnixToTime(tm)),
			CheckStatus: getCheckState(int64(v.Status)),
		}
		billInfo.StateTip, billInfo.StateTipColor = getStateTip(int64(v.Status), v.CreatedAt)

		for _, v1 := range v.BillOrderTblList {
			skuArrt := v1.SkuArrt
			if len(skuArrt) == 0 {
				skuArrt = " "
			}
			billInfo.Items = append(billInfo.Items, &caoguo.OrderProductInfo{
				Gpid:    v1.Gpid,          // 商品gpid
				Name:    v1.Name,          // 商品名
				Price:   v1.Price,         // 商品价格
				Icon:    v1.Icon,          // 商品图标
				SkuVal:  skuArrt,          // sku描述
				SkuList: v1.SkuList,       // sku信息
				Number:  int64(v1.Number), // 购买数量
			})
			billInfo.Number += int64(v1.Number)
			// 获取地址信息
			addinfo, _ := model.BillAddressTblMgr(orm.DB).GetFromBillID(v.BillID)
			if addinfo.ID > 0 {
				// addinfo.Name = tools.UnicodeEmojiDecode(addinfo.Name)
				// addinfo.AddressName = tools.UnicodeEmojiDecode(addinfo.AddressName)
				// addinfo.Address = tools.UnicodeEmojiDecode(addinfo.Address)
				// addinfo.Area = tools.UnicodeEmojiDecode(addinfo.Area)
				billInfo.Addr = &caoguo.AddressInfo{
					Name:    addinfo.Name,                                                                                            // 名字
					Mobile:  addinfo.Mobile,                                                                                          // 手机号
					Address: fmt.Sprintf("%v,%v", tools.UnicodeEmojiDecode(addinfo.Address), tools.UnicodeEmojiDecode(addinfo.Area)), // 详细地址
				}
			}
			// -------------end
		}
		resp.Items = append(resp.Items, billInfo)
	}

	return resp, nil
}

// GetMyOrdersConfig 获取订单配置
func (p *Order) GetMyOrdersConfig(c *api.Context, req *caoguo.GetMyOrdersConfigReq) (*caoguo.GetMyOrdersConfigResp, error) {
	return &caoguo.GetMyOrdersConfigResp{
		NavList: getNavList(req.IsAdmin),
	}, nil
}

// 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func getNavList(isAdmin bool) (resp []*caoguo.NavList) {
	resp = append(resp, &caoguo.NavList{
		State:       0,
		Text:        "全部",
		LoadingType: "more",
	})

	if !isAdmin {
		resp = append(resp, &caoguo.NavList{
			State:       1,
			Text:        "待付款",
			LoadingType: "more",
		}, &caoguo.NavList{
			State:       4,
			Text:        "待收货",
			LoadingType: "more",
		}, &caoguo.NavList{
			State:       5,
			Text:        "已完成",
			LoadingType: "more",
		}, &caoguo.NavList{
			State:       6,
			Text:        "售后",
			LoadingType: "more",
		})
	} else {
		resp = append(resp, &caoguo.NavList{
			State:       2,
			Text:        "待发货",
			LoadingType: "more",
		}, &caoguo.NavList{
			State:       4,
			Text:        "待收货",
			LoadingType: "more",
		}, &caoguo.NavList{
			State:       6,
			Text:        "售后",
			LoadingType: "more",
		})
	}
	// resp = append(resp, &caoguo.NavList{
	// 	State:       3,
	// 	Text:        "已取消",
	// 	LoadingType: "more",
	// })

	return
}

func getStateTip(status int64, tm time.Time) (string, string) {
	switch status {
	case -1:
		return "已退款", "#fa436a"
	case 1:
		return "待支付", "#fa436a"
	case 2:
		if time.Now().After(tm.Add(30 * time.Minute)) {
			return "备货中", "#fa436a"
		}
		return "待发货", "#fa436a"
	case 3:
		return "已取消", "#909399"
	case 4:
		return "待收货", "#909399"
	case 5:
		return "已完成", "#fa436a"
	case 6:
		return "售后", "#fa436a"
	}
	return "未知", "#fa436a"
}

func getCheckState(status int64) int64 {
	switch status {
	case -1: // 已退款
		return -1
	case 1: // 待支付
		return 1 // 取消订单 立即支付
	case 2: // 已支付
		return 2 // 取消订单
	case 3:
		return -1 //
	case 4, 5:
		return 4 // 查看物流 申请售后
	// case 5:
	// 	return 5 // 查看物流 申请售后
	case 6:
		return 6 // 售后回复
	}

	return -1 // 不展示
}

// DealOrder 账单/订单处理
func (p *Order) DealOrder(c *api.Context, req *caoguo.DealOrderReq) (resp *common.Empty, _err error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	if len(req.BillId) == 0 {
		return nil, fmt.Errorf(message.ParameterInvalid.String())
	}

	orm := core.Dao.GetDBw()
	// 找到对应账单信息
	billInfo, err := model.BillTblMgr(orm.DB).GetFromBillID(req.BillId)
	if err != nil {
		if orm.IsNotFound(err) {
			return nil, fmt.Errorf("订单已处理。请刷新")
		}
		return nil, err
	}
	if !strings.EqualFold(billInfo.UserID, session.Openid) {
		return nil, fmt.Errorf(message.InValidOp.String())
	}

	isFix := false
	switch req.Type { // 判断账单是否允许此操作
	case -1: // 取消
		if billInfo.Status == 1 || billInfo.Status == 2 {
			isFix = true
		}
	case 1: // 申请售后
		if billInfo.Status == 4 || billInfo.Status == 5 {
			isFix = true
		} else {
			if billInfo.Status == 6 {
				return nil, fmt.Errorf("亲，我们正在加紧处理中... ")
			}

		}
	case 2: // 删除订单
		if billInfo.Status == -1 || billInfo.Status == 3 || billInfo.Status == 5 || billInfo.Status == 6 {
			isFix = true
		}
	}

	if !isFix {
		return nil, fmt.Errorf("订单状态错误或已更新")
	}

	tx := orm.DB.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()
	billMgr := model.BillTblMgr(tx)

	switch req.Type {
	case -1: // 取消
		if billInfo.Status == 2 { // 已支付,退款
			wx := new(weixin.Weixin)
			_, err := wx.RefundPay(c, &caoguo.RefundPayReq{
				BillId:    billInfo.BillID,
				RefundFee: billInfo.PayAmount,
			})
			if err != nil {
				return nil, err
			}
		}

		// 修改状态已取消
		err = billMgr.Where("bill_id = ?", billInfo.BillID).Updates(map[string]interface{}{"status": 3, "updated_at": time.Now()}).Error
		if err != nil {
			return nil, err
		}
		// 取消，可能退优惠券
		coupon, _ := model.BillCouponTblMgr(tx).GetFromBillID(billInfo.BillID)
		if coupon.CouponID > 0 {
			tx.Model(&model.CouponTbl{}).Where("id = ?", coupon.CouponID).Update("status", 1)
		}
		offsetUserPoints(billInfo.UserID, session.Openid, -billInfo.PayAmount) // 删除用户积分
	case 1: // 申请售后
		if billInfo.Status == 4 || billInfo.Status == 5 {
			// 修改状态 申请售后
			err = billMgr.Where("bill_id = ?", billInfo.BillID).Updates(map[string]interface{}{"status": 6, "updated_at": time.Now()}).Error
			if err != nil {
				return nil, err
			}
		}
	case 2: // 删除订单
		if billInfo.Status == -1 || billInfo.Status == 3 || billInfo.Status == 5 || billInfo.Status == 6 {
			model.BillTblMgr(tx).Delete(&billInfo)
		}
	}

	// 添加日志
	billDealLog := model.BillDealLogTbl{
		UserID:    session.Openid,             // 用户id
		BillID:    billInfo.BillID,            // 账单id
		Type:      int(req.Type),              // 操作：-1：取消，1：申请售后，
		Contact:   req.Contact,                // 联系方式
		Remak:     req.Remak,                  // 备注
		Imgs:      tools.JSONDecode(req.Imgs), // 图片内容
		CreatedAt: time.Now(),                 // 创建时间
	}
	model.BillDealLogTblMgr(tx).Save(&billDealLog)

	return &common.Empty{}, nil
}

// DealSystem 意见与反馈
func (p *Order) DealSystem(c *api.Context, req *caoguo.DealOrderReq) (resp *common.Empty, _err error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	model.ProposalTblMgr(core.Dao.GetDBw().DB).Save(&model.ProposalTbl{
		UserID:    session.Openid,             // 用户id
		Type:      3,                          // 操作：-1：取消，1：申请售后，
		Contact:   req.Contact,                // 联系方式
		Remak:     req.Remak,                  // 备注
		Imgs:      tools.JSONDecode(req.Imgs), // 图片内容
		CreatedAt: time.Now(),                 // 创建时间
	})
	return &common.Empty{}, nil
}
