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
	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/public/mysort"
)

// Coupon 优惠券相关
type Coupon struct {
}

// GetPromotionCoupon 获取促销优惠券
func (p *Coupon) GetPromotionCoupon(c *api.Context, req *caoguo.GetPromotionCouponReq) (*caoguo.GetPromotionCouponResp, error) {
	orm := core.Dao.GetDBr()
	tmp, err := model.CouponTblMgr(orm.DB).GetFromVaild(true)
	if err != nil {
		if orm.IsNotFound(err) { // 空
			return &caoguo.GetPromotionCouponResp{}, nil
		}
		return nil, err
	}

	// 过滤我的优惠券
	var myCouponList []*model.CouponMyTbl
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) != 0 {
		myCouponList, _ = model.CouponMyTblMgr(orm.DB).GetFromUserID(session.Openid)
	}
	myCouponMp := make(map[int64]bool, len(myCouponList))
	for _, v := range myCouponList {
		myCouponMp[v.CouponID] = true
	}
	// ------end
	var resp []*model.CouponTbl

	for _, v := range tmp {
		if _, ok := myCouponMp[v.ID]; !ok {
			resp = append(resp, v)
		}
	}

	sortkey := &mysort.Fifo{}
	mp := make(map[string][]*model.CouponTbl)
	for _, v := range resp {
		mp[v.GroupName] = append(mp[v.GroupName], v)
		sortkey.Push(v.GroupName)
	}

	result := &caoguo.GetPromotionCouponResp{}
	for _, key := range sortkey.Gets() {
		var tmp caoguo.CouponGroupList
		tmp.Key = key.(string)
		for _, v := range mp[tmp.Key] {
			tmp.Items = append(tmp.Items, &caoguo.CouponInfo{
				Id:         v.ID,                                 // coupon id
				GroupName:  v.GroupName,                          // 分组名
				Title:      v.Title,                              // 优惠券名字
				Validity:   getValidity(v.Validity),              // 有效截止日期
				Gpid:       v.Gpid,                               // 商品优惠券
				Denom:      int64(v.Denom),                       // 面额
				CouponType: int32(v.CouponType),                  // 优惠券类型(1:全场，2:单个商品)
				GreatMoney: GetGreatMoney(float32(v.GreatMoney)), // 满多少可用
				Describle:  v.Describle,                          // 优惠券描述
				Status:     0,                                    // 0：默认，1:未使用(有效),2:已使用,-1:已过期
				Type:       1,
			})
			result.TotalFee += int64(v.Denom)
			result.Ids = append(result.Ids, v.ID)
		}
		result.Items = append(result.Items, &tmp)
	}
	return result, nil
}

// getValidity 获取有效截止日期
func getValidity(offset int) string {
	return fmt.Sprintf("有效期至 %v", tools.FormatTime(time.Now().Add(time.Duration(offset*24)*time.Hour), "2006-01-02"))
}

// GoGetCoupon 一键领取优惠券
func (p *Coupon) GoGetCoupon(c *api.Context, req *caoguo.GoGetCouponReq) (bool, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()
	mgr := model.CouponMyTblMgr(orm.DB)
	resp, err := mgr.GetFromUserID(session.Openid)
	if err != nil {
		if !orm.IsNotFound(err) {
			return false, err
		}
	}
	existsMp := make(map[int64]bool)
	for _, v := range resp {
		existsMp[v.CouponID] = true
	}

	var ids []int64
	for _, v := range req.Ids {
		if !existsMp[v] {
			ids = append(ids, v)
		}
	}

	if len(ids) > 0 { // 开始插入操作
		mp := make(map[int64]*model.CouponTbl)
		couponList, err := model.CouponTblMgr(orm.DB).GetBatchFromID(ids)
		if err != nil {
			return false, nil
		}
		for _, v := range couponList {
			mp[v.ID] = v
		}
		for _, v := range ids {
			mgr.Save(&model.CouponMyTbl{
				UserID:      session.Openid,                                               // 用户id
				CouponID:    v,                                                            // 优惠券id
				ExpiresTime: time.Now().Add(time.Duration(mp[v].Validity*24) * time.Hour), // 过期时间
				Status:      1,                                                            // 当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)
			})
		}
	}

	return true, nil
}

// GetMyCoupon 获取我的优惠券
func (p *Coupon) GetMyCoupon(c *api.Context, req *caoguo.GetMyCouponReq) (*caoguo.GetMyCouponResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()
	mgr := model.CouponMyTblMgr(orm.DB)
	resp, err := mgr.GetFromUserID(session.Openid)
	if err != nil {
		if !orm.IsNotFound(err) {
			return nil, err
		}
	}

	now := time.Now()
	couponMp := make(map[int64][]*model.CouponMyTbl)
	for _, v := range resp {
		if v.Status == 2 { // 已使用
			couponMp[2] = append(couponMp[2], v)
		} else if v.ExpiresTime.Before(now) { // 已过期
			couponMp[-1] = append(couponMp[-1], v)
		} else { // 未使用
			couponMp[1] = append(couponMp[1], v)
		}
	}

	// 聚合
	resutl := &caoguo.GetMyCouponResp{}

	var tmp1 caoguo.CouponGroupList
	tmp1.Key = "未使用"
	for _, v := range couponMp[1] {
		tmp1.Items = append(tmp1.Items, &caoguo.CouponInfo{
			Id:         v.ID,                                                                  // coupon id
			GroupName:  v.CouponTbl.GroupName,                                                 // 分组名
			Title:      v.CouponTbl.Title,                                                     // 优惠券名字
			Validity:   fmt.Sprintf("有效期至 %v", tools.FormatTime(v.ExpiresTime, "2006-01-02")), // 有效截止日期
			Gpid:       v.CouponTbl.Gpid,                                                      // 商品优惠券
			Denom:      int64(v.CouponTbl.Denom),                                              // 面额
			CouponType: int32(v.CouponTbl.CouponType),                                         // 优惠券类型(1:全场，2:单个商品)
			GreatMoney: GetGreatMoney(float32(v.CouponTbl.GreatMoney)),                        // 满多少可用
			Describle:  v.CouponTbl.Describle,                                                 // 优惠券描述
			Status:     1,                                                                     // 0：默认，1:未使用(有效),2:已使用,-1:已过期
			Type:       2,
		})
	}
	resutl.Items = append(resutl.Items, &tmp1)

	var tmp2 caoguo.CouponGroupList
	tmp2.Key = "已使用"
	for _, v := range couponMp[2] {
		tmp2.Items = append(tmp2.Items, &caoguo.CouponInfo{
			Id:         v.ID,                                                                  // coupon id
			GroupName:  v.CouponTbl.GroupName,                                                 // 分组名
			Title:      v.CouponTbl.Title,                                                     // 优惠券名字
			Validity:   fmt.Sprintf("有效期至 %v", tools.FormatTime(v.ExpiresTime, "2006-01-02")), // 有效截止日期
			Gpid:       v.CouponTbl.Gpid,                                                      // 商品优惠券
			Denom:      int64(v.CouponTbl.Denom),                                              // 面额
			CouponType: int32(v.CouponTbl.CouponType),                                         // 优惠券类型(1:全场，2:单个商品)
			GreatMoney: GetGreatMoney(float32(v.CouponTbl.GreatMoney)),                        // 满多少可用
			Describle:  v.CouponTbl.Describle,                                                 // 优惠券描述
			Status:     2,                                                                     // 0：默认，1:未使用(有效),2:已使用,-1:已过期
			Type:       2,
		})
	}
	resutl.Items = append(resutl.Items, &tmp2)

	var tmp0 caoguo.CouponGroupList
	tmp0.Key = "已过期"
	for _, v := range couponMp[-1] {
		tmp0.Items = append(tmp0.Items, &caoguo.CouponInfo{
			Id:         v.ID,                                                                  // coupon id
			GroupName:  v.CouponTbl.GroupName,                                                 // 分组名
			Title:      v.CouponTbl.Title,                                                     // 优惠券名字
			Validity:   fmt.Sprintf("有效期至 %v", tools.FormatTime(v.ExpiresTime, "2006-01-02")), // 有效截止日期
			Gpid:       v.CouponTbl.Gpid,                                                      // 商品优惠券
			Denom:      int64(v.CouponTbl.Denom),                                              // 面额
			CouponType: int32(v.CouponTbl.CouponType),                                         // 优惠券类型(1:全场，2:单个商品)
			GreatMoney: GetGreatMoney(float32(v.CouponTbl.GreatMoney)),                        // 满多少可用
			Describle:  v.CouponTbl.Describle,                                                 // 优惠券描述
			Status:     -1,                                                                    // 0：默认，1:未使用(有效),2:已使用,-1:已过期
			Type:       2,
		})
	}
	resutl.Items = append(resutl.Items, &tmp0)

	// 完成
	return resutl, nil
}
