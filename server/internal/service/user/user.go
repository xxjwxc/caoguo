package user

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"fmt"
	"time"

	"caoguo/internal/service/history"
	"caoguo/internal/service/weixin"

	"github.com/xxjwxc/public/message"
)

// User 用户相关
type User struct {
}

// GetUserInfo 获取用户相关信息
func (u *User) GetUserInfo(c *api.Context, req *caoguo.GetUserInfoReq) (*caoguo.GetUserInfoResp, error) {
	resp := &caoguo.GetUserInfoResp{}
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return &caoguo.GetUserInfoResp{}, nil
	}

	orm := core.Dao.GetDBr()
	userInfo, err := model.UserTblMgr(orm.DB).GetFromUserID(session.Openid)
	if err != nil {
		if !orm.IsNotFound(err) {
			return nil, err
		}
	} else {
		resp.IsVip = userInfo.Vip       // 是否vip
		resp.Balance = userInfo.Balance // 用户余额
		resp.Points = userInfo.Points   // 用户积分
	}

	// 获取可用优惠券数量
	model.CouponMyTblMgr(orm.DB).Where("user_id = ? and status = 1", session.Openid).Count(&resp.CouponNum)

	// 获取用户浏览历史
	historyInfo := history.Get(session.Openid)
	for _, v := range historyInfo {
		resp.Historys = append(resp.Historys, &caoguo.History{
			Gpid: v.Gpid,
			Icon: v.Icon,
		})
	}

	return resp, nil
}

// RequestVip 申请vip
func (u *User) RequestVip(c *api.Context, req *common.Empty) (*common.Empty, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}
	orm := core.Dao.GetDBw()
	userInfo, err := model.UserTblMgr(orm.DB).GetFromUserID(session.Openid)
	if err != nil {
		if !orm.IsNotFound(err) {
			return nil, err
		}
		userInfo.UserID = session.Openid
	}

	userInfo.Vip = true
	userInfo.UpdatedAt = time.Now()

	model.UserTblMgr(orm.DB).Save(&userInfo)
	return &common.Empty{}, nil
}

// LinkUser 关联用户
func (u *User) LinkUser(c *api.Context, req *caoguo.LinkUserReq) (*common.Empty, error) {
	if req.OpenId == "undefined" || req.OpenId == "" {
		return nil, nil
	}

	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()
	orm.Save(&model.UserLinkLogTbl{
		OpenID:     req.OpenId,
		LinkOpenID: session.Openid,
		CreatedAt:  time.Now(),
	})

	link, err := model.UserLinkTblMgr(orm.DB).GetFromUserID(session.Openid) //找到是否有关联
	if err != nil {
		if !orm.IsNotFound(err) {
			return nil, err
		}
	}
	if link.ID == 0 { // 未找到，开始关联
		orm.Save(&model.UserLinkTbl{
			UserID:       session.Openid, // 被邀请人
			ParentUserID: req.OpenId,     // 邀请人
			CreatedAt:    time.Now(),     // 创建时间
		})
	}

	return &common.Empty{}, nil
}
