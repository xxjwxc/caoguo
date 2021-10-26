package weixin

import (
	"caoguo/internal/config"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/xxjwxc/public/mylog"

	"caoguo/internal/api"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/caoguo/weixin"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/tools"
)

// Weixin 微信相关操作
type Weixin struct {
}

// Oauth 微信授权获取登录信息
func (w *Weixin) Oauth(c *api.Context, req OauthReq) (OauthResp, error) {
	result := GetOpenID(req.Jscode)
	if len(result.SessionID) == 0 { // fail 失败
		return result, errors.New(message.InValidOp.String())
	}

	//设置客户端cookie
	// tm := int(result.OverdueTime)
	// c.GetGinCtx().SetCookie(api.UserToken, result.SessionID, tm, "/", c.GetGinCtx().Request.Host, false, true)

	return result, nil
}

// UpdateUserInfo 	更新用户信息
func (w *Weixin) UpdateUserInfo(c *api.Context, req *WxUserinfo) (bool, error) {
	sessionID := c.GetSessionToken()
	mylog.Infof("sessionid:%v", sessionID)
	session := GetSessionkey(sessionID)
	if len(session.Openid) == 0 {
		return false, nil
	}

	var info model.WxUserinfo
	info.City = req.City
	info.Country = req.Country
	info.AvatarURL = req.AvatarURL
	info.NickName = tools.UnicodeEmojiCode(req.NickName)
	info.Province = req.Province
	info.Gender = req.Gender
	info.Language = []byte(req.Language)
	info.Openid = session.Openid

	if len(info.AvatarURL) > 0 || len(info.NickName) > 0 {
		updateInfo(info)
	}
	return true, nil
}

// PayPlaceOrder 支付
func PayPlaceOrder(openID string, price int64, priceBody, orderID, clientIP string) (map[string]string, error) {
	msg := _wx.SmallAppUnifiedorder(openID, price, priceBody, orderID, clientIP)
	mylog.Info(msg)

	if msg.State {
		return msg.Data.(map[string]string), nil
	}

	return nil, fmt.Errorf(msg.Error)
}

// SelectOrder 支付查询接口
func SelectOrder(openID, orderID string) (int, message.MessageBody) {
	return _wx.SelectOrder(openID, orderID)
}

// RefundPay 退款
func (w *Weixin) RefundPay(c *api.Context, req *caoguo.RefundPayReq) (resp *caoguo.RefundPayResp, _err error) { // 退款
	orm := core.Dao.GetDBw()
	billInfo, err := model.BillTblMgr(orm.DB).GetFromBillID(req.BillId)
	if err != nil {
		if orm.IsNotFound(err) {
			return nil, fmt.Errorf("订单已处理。请刷新")
		}
		return nil, err
	}
	var refundBill model.BillRefundTbl
	refundBill.BillID = billInfo.BillID
	refundBill.RefundFee = billInfo.PayAmount
	refundBill.RefundID = createUnique("")
	refundBill.CreatedAt = time.Now()

	b, msg := _wx.RefundPay(billInfo.UserID, billInfo.BillID, refundBill.RefundID, refundBill.RefundFee, billInfo.PayAmount)
	mylog.Info(msg) // 退款信息
	if b {
		err := model.BillRefundTblMgr(orm.DB).Save(&refundBill).Error
		if err != nil {
			return &caoguo.RefundPayResp{Status: false}, err
		}
		// success
		return &caoguo.RefundPayResp{Status: true}, nil
	}

	return &caoguo.RefundPayResp{Status: false}, fmt.Errorf(msg.Error)
}

// GetQrcode 获取微信二维码
func (w *Weixin) GetQrcode(c *api.Context, req *weixin.GetQrcodeReq) (resp *weixin.GetQrcodeResp, _err error) {
	page := fmt.Sprintf("/file/qrcode/%v_%v.jpg", strings.Replace(req.Page, "/", "_", -1), req.Scene)
	path := path.Join(tools.GetCurrentDirectory(), page)
	if !tools.CheckFileIsExist(path) {
		ret := _wx.GetShareQrcode(path, req.Scene, req.Page)
		mylog.Info(ret)
		if ret.Errcode != 0 {
			return nil, fmt.Errorf(ret.Errmsg)
		}
	}

	return &weixin.GetQrcodeResp{
		Url: config.GetFileHost() + page,
	}, nil
}
