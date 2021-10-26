package coupon

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"fmt"
	"strings"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
)

// GetAdminCouponInfo admin 获取优惠券列表
func (p *Coupon) GetAdminCouponInfo(c *api.Context, req *common.Empty) (*caoguo.GetAdminCouponInfo, error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	if !strings.EqualFold(userinfo.UserName, "admin") {
		return nil, fmt.Errorf("非admin用户不允许查询")
	}

	var resp caoguo.GetAdminCouponInfo

	orm := core.Dao.GetDBr()
	ads, err := model.CouponTblMgr(orm.DB).Gets() // 首页广告
	if err != nil && !orm.IsNotFound(err) {
		return nil, err
	}

	for _, v := range ads {
		resp.CouponList = append(resp.CouponList, &caoguo.AdminCouponInfo{
			Id:         int64(v.ID),
			GroupName:  v.GroupName,
			Title:      v.Title,
			Validity:   int32(v.Validity),
			Gpid:       v.Gpid,
			Denom:      int64(v.Denom),
			CouponType: int32(v.CouponType),
			GreatMoney: int32(v.GreatMoney),
			Describle:  v.Describle,
			Vaild:      v.Vaild,
		})
	}

	return &resp, nil
}

// UpsetAdminCouponInfo admin 更新/添加 优惠券列表
func (p *Coupon) UpsetAdminCouponInfo(c *api.Context, req *caoguo.GetAdminCouponInfo) (b bool, _err error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	if !strings.EqualFold(userinfo.UserName, "admin") {
		return false, fmt.Errorf("非admin用户不允许操作")
	}

	orm := core.Dao.GetDBw()
	tx := orm.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()

	for _, v := range req.CouponList {
		tmp := &model.CouponTbl{
			ID:         v.Id,
			GroupName:  v.GroupName,
			Title:      v.Title,
			Validity:   int(v.Validity),
			Gpid:       v.Gpid,
			Denom:      int(v.Denom),
			CouponType: int(v.CouponType),
			GreatMoney: int(v.GreatMoney),
			Describle:  v.Describle,
			Vaild:      v.Vaild,
			UpdatedBy:  userinfo.UserName,
			UpdatedAt:  time.Now(),
		}
		if v.Id == 0 {
			tmp.CreatedAt = time.Now()
			tmp.CreatedBy = userinfo.UserName
		}

		_err = model.ProductAdTblMgr(tx).Save(&tmp).Error
		if _err != nil {
			mylog.Error(_err)
			return false, _err
		}
	}

	return true, nil
}
