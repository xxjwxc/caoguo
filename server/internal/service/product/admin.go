package product

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/oplog"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/xxjwxc/public/myglobal"
	"github.com/xxjwxc/public/mylog"
	"gorm.io/gorm"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/public/message"
)

// Product 商品相关
type Product struct {
}

// GetProductList admin 获取商品列表
func (p *Product) GetProductList(c *api.Context, req *caoguo.GetProductListReq) (*caoguo.GetProductListResp, error) {
	if req.Index < 0 {
		req.Index = 0
	}
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBr()

	db := orm.Offset(int(req.Index * _ProductPageLen)).Limit(_ProductPageLen).Order("created_at desc")
	if len(req.Gpid) > 0 {
		db = orm.Where("gpid = ?", req.Gpid)
	}
	if userinfo.UserName != "admin" { // 非admin
		db = db.Where("platform_id = ?", userinfo.UserName)
	}

	list, err := model.ProductTblMgr(db).Gets()
	if err != nil {
		return nil, err
	}

	var resp = &caoguo.GetProductListResp{}
	for _, v := range list {
		imgs, videos := GetContext(v.Context)
		resp.ProductList = append(resp.ProductList, &caoguo.Product{
			Gpid:          v.Gpid,
			Gtid:          v.Gtid,
			Name:          v.Name,
			Price:         v.Price,
			OriginalPrice: v.OriginalPrice,
			DistAmount:    v.DistAmount,
			PlatformID:    v.PlatformID,
			Qty:           v.Qty,
			ShipmentFee:   v.ShipmentFee,
			ImgDetail:     imgs,
			VideoDetail:   videos,
			IsSelect:      v.IsSelect,
			Vaild:         v.Vaild,
			UpdatedAt:     tools.GetTimeStr(v.UpdatedAt),
			CreatedAt:     tools.GetTimeStr(v.CreatedAt),

			Sales:      v.ProductInfoTbl.Qty,
			Img:        GetImages(v.ProductInfoTbl.Img),
			Icon:       v.ProductInfoTbl.Icon,
			Service:    v.ProductInfoTbl.Service,
			ShareTit:   "分享活动获取更多奖励",
			Skus:       getProductSkus(v.Gpid),
			SkuPrice:   getProductSkuPrice(v.Gpid),
			CouponList: getProductCouponList(v.Gpid),
		})
	}

	return resp, nil
}

// ReplaceProduct 添加一个产品(可能更新)
func (p *Product) ReplaceProduct(c *api.Context, req AddReq) (bool, error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()

	//先查找类型是否存在
	ptinfo, err := model.ProductTypeTblMgr(orm.DB).GetFromGtid(req.GType)
	if err != nil && !orm.IsNotFound(err) {
		return false, err
	}

	if ptinfo.ID == 0 { // 没有则添加
		ptinfo.Gtid = req.GType
		model.ProductTypeTblMgr(orm.DB).Save(&ptinfo)
	}

	// 判断是否有这个商品
	info, err := model.ProductTblMgr(orm.DB).FetchUniqueIndexByNameID(req.Name, userinfo.UserName)
	if err != nil && !orm.IsNotFound(err) {
		return false, err
	}

	// 更新或者创建
	if len(info.Gpid) == 0 {
		info.Gpid = myglobal.GetNode().GetIDStr()
		info.ProductInfoTbl.Gpid = info.Gpid
		info.PlatformID = userinfo.UserName
		info.Vaild = true
		info.CreatedBy = userinfo.UserName
		info.CreatedAt = time.Now()
		info.ProductInfoTbl.CreatedBy = userinfo.UserName
		info.ProductInfoTbl.CreatedAt = time.Now()
	}
	info.Name = req.Name
	info.Gtid = req.GType
	info.Price = req.Price
	info.Qty = req.Qty
	info.Context = tools.JSONDecode(req.Contexts)
	info.UpdatedBy = userinfo.UserName
	info.UpdatedAt = time.Now()

	info.ProductInfoTbl.UpdatedBy = userinfo.UserName
	info.ProductInfoTbl.UpdatedAt = time.Now()
	info.ProductInfoTbl.Img = tools.JSONDecode(req.Image)
	info.ProductInfoTbl.Icon = req.Icon

	model.ProductInfoTblMgr(orm.DB).Save(&info.ProductInfoTbl)
	model.ProductTblMgr(orm.DB).Save(&info)

	var logs model.OpLogTbl
	logs.Operator = userinfo.UserName
	logs.Topic = "产品系统"
	logs.Bundle = "修改"
	logs.Pid = req.Name
	logs.IPAddr = c.GetGinCtx().ClientIP()
	oplog.OpLogBaseAdd(&logs) //保存日志

	return true, nil
}

// ReplaceSku 添加一个产品Sku(可能更新)
func (p *Product) ReplaceSku(c *api.Context, req AddSkuReq) (_b bool, _err error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()
	tx := orm.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()

	// defer tx.Commit()
	product, err := model.ProductTblMgr(tx).FetchUniqueByGpid(req.Gpid)
	if err != nil {
		return false, err
	}

	if product.PlatformID != userinfo.UserName {
		return false, errors.New(message.NotVaildError.String())
	}

	// 找到老的
	skuDbMgr := model.ProductSkuTblMgr(tx)
	skus, err := skuDbMgr.GetFromGpid(req.Gpid)
	if err != nil && !orm.IsNotFound(err) {
		return false, err
	}

	mp := make(map[string]*model.ProductSkuTbl)
	for _, v := range skus {
		mp[fmt.Sprintf("%v_%v", v.TypeName, v.TagName)] = v
	}

	for _, v := range req.Skus {
		key := fmt.Sprintf("%v_%v", v.TypeName, v.TagName)
		if v1, ok := mp[key]; ok { // edit
			v1.TypeName = v.TypeName
			v1.UpdatedBy = userinfo.UserName
			v1.UpdatedAt = time.Now()
		} else { // add
			mp[key] = &model.ProductSkuTbl{
				Gpid:     req.Gpid,
				TypeName: v.TypeName,
				TagName:  v.TagName,
				// Premium:   v.Premium,
				CreatedBy: userinfo.UserName,
				Model: gorm.Model{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				UpdatedBy: userinfo.UserName,
			}
		}
	}

	// ok save
	for _, v := range mp {
		skuDbMgr.Save(v)
	}

	// tx.Commit() // 提交

	return true, nil
}

// UpsetSkuPrice 更新/设置产品sku价格
func (p *Product) UpsetSkuPrice(c *api.Context, req *caoguo.UpsetSkuPriceReq) (_b bool, _err error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	if req.Id == 0 && len(req.Gpid) == 0 { // 入口参数检查
		return false, fmt.Errorf(message.ParameterInvalid.String())
	}

	orm := core.Dao.GetDBw()

	// 试图更新
	if req.Id == 0 { // 无sku 时只在product上更新操作
		err := model.ProductTblMgr(orm.Where("gpid = ?", req.Gpid)).
			Updates(map[string]interface{}{"price": req.Premium, "dist_amount": req.DistAmount, "updated_by": userinfo.UserName, "updated_at": time.Now()}).Error
		return err == nil, err
	}

	var tmp []int
	for _, v := range req.SkuList {
		tmp = append(tmp, int(v))
	}
	sort.Ints(tmp) // 排序
	var tmpSrc []string
	for _, v := range tmp {
		tmpSrc = append(tmpSrc, fmt.Sprintf("%v", v))
	}
	skuList := strings.Join(tmpSrc, ",")

	switch req.Tag {
	case -1: // 删除
		_err = orm.Delete(&model.ProductSkuPriceTbl{}, req.Id).Error
	case 0: // 添加
		tmp := &model.ProductSkuPriceTbl{
			Gpid:       req.Gpid,          // 商品id
			SkuList:    skuList,           // 属性列表
			Premium:    req.Premium,       // 商品单价
			DistAmount: req.DistAmount,    // 分享收益
			CreatedBy:  userinfo.UserName, // 创建者
			// CreatedAt:  time.Now(),
		}
		tmp.CreatedAt = time.Now()
		_err = orm.Create(tmp).Error
	case 1: // 更新
		_err = model.ProductSkuPriceTblMgr(orm.Where("id = ?", req.Id)).
			Updates(map[string]interface{}{"premium": req.Premium, "dist_amount": req.DistAmount, "updated_by": userinfo.UserName, "updated_at": time.Now()}).Error

	}

	return _err == nil, _err
}

// UpdateSelect 修改商品推荐信息
func (p *Product) UpdateSelect(c *api.Context, req UpdateSelectReq) (bool, error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()

	// 试图更新
	err := model.ProductTblMgr(orm.Where("gpid = ?", req.Gpid)).
		Updates(map[string]interface{}{"is_select": req.IsSelect, "updated_by": userinfo.UserName, "updated_at": time.Now()}).Error

	return err == nil, err
}

func getProductSkus(gpid string) (result []*caoguo.SkuInfo) {
	orm := core.Dao.GetDBr()
	resp, err := model.ProductSkuTblMgr(orm.DB.Order("id asc")).GetFromGpid(gpid)
	if err != nil {
		return
	}

	for _, v := range resp {
		result = append(result, &caoguo.SkuInfo{
			Id:       int64(v.ID),
			TypeName: v.TypeName, // 类型名字
			TagName:  v.TagName,
		})
	}

	return
}

func getProductSkuPrice(gpid string) (result []*caoguo.SkuPriceGroup) {
	orm := core.Dao.GetDBr()
	resp, err := model.ProductSkuPriceTblMgr(orm.DB).GetFromGpid(gpid)
	if err != nil {
		return
	}
	for _, v := range resp {
		result = append(result, &caoguo.SkuPriceGroup{
			Id:         int64(v.ID),
			SkuList:    v.SkuList,
			Premium:    v.Premium,
			DistAmount: v.DistAmount,
		})
	}
	return
}

func getProductCouponList(gpid string) (result []*caoguo.Coupon) {
	orm := core.Dao.GetDBr()
	resp, err := model.CouponTblMgr(orm.DB).GetFromGpid(gpid)
	if err != nil {
		return
	}
	for _, v := range resp {
		result = append(result, &caoguo.Coupon{
			Id:         v.ID,                // coupon id
			GroupName:  v.GroupName,         // 分组名
			Title:      v.Title,             // 优惠券名字
			Validity:   int32(v.Validity),   // 有效截止日期
			Gpid:       v.Gpid,              // 商品优惠券
			Denom:      int64(v.Denom),      // 面额
			CouponType: int32(v.CouponType), // 优惠券类型(1:全场，2:单个商品)
			GreatMoney: int32(v.GreatMoney), // 满多少可用
			Describle:  v.Describle,         // 优惠券描述
			Vaild:      v.Vaild,             // 0：默认，1:未使用(有效),2:已使用,-1:已过期
			Type:       int32(v.CouponType), // 优惠券类型，1:促销优惠券，2：用户已领取的优惠券
		})
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// 广告相关  ////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GetAdminAdInfo admin 获取首页信息(轮播图广告，类型广告，主销广告,精选列表)
func (p *Product) GetAdminAdInfo(c *api.Context, req *common.Empty) (*caoguo.GetAdminAdInfo, error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	if !strings.EqualFold(userinfo.UserName, "admin") {
		return nil, fmt.Errorf("非admin用户不允许查询")
	}

	var resp caoguo.GetAdminAdInfo

	orm := core.Dao.GetDBr()

	ads, err := model.ProductAdTblMgr(orm.DB).Gets() // 首页广告
	if err != nil && !orm.IsNotFound(err) {
		return nil, err
	}

	for _, v := range ads {
		resp.AdList = append(resp.AdList, &caoguo.AdminAdInfo{
			Id:     int64(v.ID),
			Url:    v.URL,
			Img:    v.Img,
			SortId: int32(v.SortID),
			Type:   int32(v.Type),
			Attach: v.Attach,
			Vaild:  v.Vaild,
		})
	}

	group, err := model.ProductAdGroupTblMgr(orm.DB).Gets()
	if err != nil && !orm.IsNotFound(err) {
		return nil, err
	}
	for _, v := range group {
		resp.GroupList = append(resp.GroupList, &caoguo.AdminAdGroupInfo{
			Id:       int64(v.ID),
			MainGpid: v.MainGpid,
			SubGpid:  v.SubGpid,
			SortId:   int32(v.SortID),
			Vaild:    v.Vaild,
		})
	}

	return &resp, nil
}

// UpsetAdminAdInfo admin 更新首页信息(轮播图广告，类型广告，主销广告,精选列表)(id为0,添加，id>0 删除)
func (p *Product) UpsetAdminAdInfo(c *api.Context, req *caoguo.GetAdminAdInfo) (b bool, _err error) {
	// 用户判断
	userinfo, b := c.GetUserInfo()
	if !b {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	if !strings.EqualFold(userinfo.UserName, "admin") {
		return false, fmt.Errorf("非admin用户不允许修改")
	}

	orm := core.Dao.GetDBw()
	tx := orm.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()

	for _, v := range req.AdList { // 更新或者修改
		var tmp model.ProductAdTbl
		tmp.ID = uint(v.Id)
		//		tmp.CreatedAt = time.Now()
		tmp.UpdatedAt = time.Now()
		tmp.URL = v.Url
		tmp.Img = v.Img
		tmp.SortID = int(v.SortId)
		tmp.Type = int(v.Type)
		tmp.Attach = v.Attach
		tmp.Vaild = v.Vaild
		tmp.UpdatedBy = userinfo.UserName
		if v.Id > 0 {
			tmp.CreatedAt = time.Now()
			tmp.CreatedBy = userinfo.UserName
		}
		_err = model.ProductAdTblMgr(tx).Save(&tmp).Error
		if _err != nil {
			mylog.Error(_err)
			return false, _err
		}
	}

	for _, v := range req.GroupList { // 更新或者修改
		var tmp model.ProductAdGroupTbl
		tmp.ID = int(v.Id)
		tmp.UpdatedAt = time.Now()
		tmp.MainGpid = v.MainGpid
		tmp.SubGpid = v.SubGpid
		tmp.SortID = int(v.SortId)
		tmp.Vaild = v.Vaild
		tmp.UpdatedBy = userinfo.UserName
		if v.Id > 0 {
			tmp.CreatedAt = time.Now()
			tmp.CreatedBy = userinfo.UserName
		}
		_err = model.ProductAdGroupTblMgr(tx).Save(&tmp).Error
		if _err != nil {
			mylog.Error(_err)
			return false, _err
		}
	}

	return true, nil
}
