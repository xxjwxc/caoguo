package dist

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/weixin"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"fmt"
	"strings"

	"github.com/xxjwxc/public/message"
)

// Dist 分销
type Dist struct {
}

// GetDistList 获取可分销列表
func (d *Dist) GetDistList(c *api.Context, req *common.Empty) (*caoguo.GetDistListResp, error) {
	resp := &caoguo.GetDistListResp{
		Info: &caoguo.DistDoc{
			Detail: `<p>分销金:%50<p>`,
			Imgs:   []string{},
		},
		Items: []*caoguo.DistProductInfo{},
	}
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return resp, fmt.Errorf(message.UserNotExisted.String())
	}

	// 判断是否有分销权限
	orm := core.Dao.GetDBr()
	user, err := model.UserTblMgr(orm.DB).GetFromUserID(session.Openid)
	if err != nil {
		if !orm.IsNotFound(err) {
			return resp, err
		}
	}
	resp.Info.IsVip = user.DistVip

	// 查找 有设置分销的商品列表
	price, err := model.ProductSkuPriceTblMgr(orm.Where("share_amount > 0")).Gets()
	if err != nil && !orm.IsNotFound(err) {
		return resp, err
	}

	var gpids, skuids []string
	for _, v := range price {
		gpids = append(gpids, v.Gpid)
		if len(v.SkuList) > 0 {
			skuids = append(skuids, strings.Split(v.SkuList, ",")...)
		}
	}

	// 获取sku信息
	skuMp := make(map[string]*model.ProductSkuTbl, len(skuids))
	if len(skuids) > 0 {
		result1, err := model.ProductSkuTblMgr(orm.Where("id IN (?)", skuids)).Gets()
		if err != nil && !orm.IsNotFound(err) {
			return nil, fmt.Errorf(message.NotFindError.String())
		}
		for _, v := range result1 {
			skuMp[fmt.Sprintf("%v", v.ID)] = v
		}
	}

	gpidMp := make(map[string][]string)
	for _, v := range price {
		sps := strings.Split(v.SkuList, ",")
		var tmp string
		for _, v1 := range sps {
			if _, ok := skuMp[v1]; ok {
				tmp += skuMp[v1].TagName + " "
			}
		}
		if len(tmp) > 0 {
			tmp += "分享收益:" + getMoneyStr(v.DistAmount)
			gpidMp[v.Gpid] = append(gpidMp[v.Gpid], tmp)
		}
	}

	// 获取商品信息
	result, err := model.ProductTblMgr(orm.DB).GetBatchFromGpid(gpids)
	if err != nil && !orm.IsNotFound(err) {
		return nil, fmt.Errorf(message.NotFindError.String())
	}

	for _, v := range result {
		if _, ok := gpidMp[v.Gpid]; ok {
			resp.Items = append(resp.Items, &caoguo.DistProductInfo{
				Gpid:          v.Gpid,                // 商品gpid
				Name:          v.Name,                // 商品名
				Price:         v.Price,               // 商品价格
				OriginalPrice: v.OriginalPrice,       // 商品原始价格
				Icon:          v.ProductInfoTbl.Icon, // 商品图标
				ShareAmounts:  gpidMp[v.Gpid],
			})
		}
	}
	// ----------end
	return resp, nil
}

// RequestDist 申请分销
func (d *Dist) RequestDist(c *api.Context, req *common.Empty) (*common.ResultResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return &common.ResultResp{}, fmt.Errorf(message.UserNotExisted.String())
	}

	// 判断是否有分销权限
	orm := core.Dao.GetDBr()
	user, err := model.UserTblMgr(orm.DB).GetFromUserID(session.Openid)
	if err != nil && !orm.IsNotFound(err) {
		return &common.ResultResp{}, err
	}
	if !user.DistVip { // 设置用户分销权限
		err = model.UserTblMgr(orm.DB).Where("user_id = ?", session.Openid).Update("dist_vip", true).Error
		if err != nil {
			return &common.ResultResp{}, err
		}
	}

	return &common.ResultResp{Result: true}, nil
}
