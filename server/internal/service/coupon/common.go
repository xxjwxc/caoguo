package coupon

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/weixin"
	"caoguo/rpc/caoguo"
	"fmt"
	"time"

	"github.com/xxjwxc/public/message"
)

// ModifyCoupon 计算优惠券费用明细
func ModifyCoupon(c *api.Context, couponID int64, resp *caoguo.ReckonFeeResp) (int64, error) {
	if couponID == 0 {
		return 0, nil
	}

	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return 0, fmt.Errorf(message.UserNotExisted.String())
	}

	now := time.Now()
	orm := core.Dao.GetDBr()
	info, err := model.CouponMyTblMgr(orm.Where("status = 1 and user_id = ?", session.Openid)).GetFromID(couponID)
	if err != nil {
		return 0, err
	}
	if now.After(info.ExpiresTime) || !info.CouponTbl.Vaild { // 过期
		return 0, fmt.Errorf(message.Overdue.String())
	}

	greatMoney := int64(info.CouponTbl.GreatMoney)
	switch info.CouponTbl.CouponType {
	case 1: // 全场
		{
			if resp.TotalFee >= greatMoney { // 可用
				return int64(info.CouponTbl.Denom), nil
			}
		}
	case 2: // 单个商品
		{
			var price int64
			for _, order := range resp.OrderInfo {
				if order.Gpid == info.CouponTbl.Gpid { // 同一个商品
					price += order.Price * order.Number
				}
			}
			if price >= greatMoney { // 找到 & 可用
				return int64(info.CouponTbl.Denom), nil
			}
		}
	case 3: // 免邮卷
		if resp.ShipmentFee >= greatMoney { // 可用
			return int64(info.CouponTbl.Denom), nil
		}
	}

	return 0, nil // 不可用 or 未找到
}

// GetCartVaildCoupon 根据购物车获取可用优惠券
func GetCartVaildCoupon(c *api.Context, resp *caoguo.ReckonFeeResp) (result []*model.CouponMyTbl) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return
	}

	mycoupons := getVaildCoupon(session.Openid) // 获取可用优惠券
	for _, v := range mycoupons {
		switch v.CouponTbl.CouponType {
		case 1: // 全场
			{
				if resp.TotalFee >= int64(v.CouponTbl.GreatMoney) { // 可用
					result = append(result, v)
				}
			}

		case 2: // 单个商品
			{
				var price int64
				for _, order := range resp.OrderInfo {
					if order.Gpid == v.CouponTbl.Gpid { // 同一个商品
						price += order.Price * order.Number
					}
				}
				if price >= int64(v.CouponTbl.GreatMoney) { // 找到 & 可用
					result = append(result, v)
				}
			}
		case 3: // 免邮卷
			if resp.ShipmentFee >= int64(v.CouponTbl.GreatMoney) { // 可用
				result = append(result, v)
			}
		}
	}

	return result
}

// 获取可用优惠券
func getVaildCoupon(userID string) (result []*model.CouponMyTbl) {
	now := time.Now()
	orm := core.Dao.GetDBr()
	resp, _ := model.CouponMyTblMgr(orm.Where("status = 1")).GetFromUserID(userID)
	for _, v := range resp {
		if now.Before(v.ExpiresTime) && v.CouponTbl.Vaild { // 未过期
			result = append(result, v)
		}
	}
	return
}

// GetGreatMoney 满多少可用
func GetGreatMoney(offset float32) string {
	if offset <= 0 {
		return ""
	}
	return fmt.Sprintf("满 %.2f 可用", offset*0.01)
}
