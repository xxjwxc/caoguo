package product

import (
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/internal/service/history"
	"caoguo/internal/service/weixin"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
	"fmt"
	"strconv"
	"strings"
	"time"

	"sort"

	"github.com/jinzhu/gorm"
	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/tools"

	"caoguo/internal/api"

	"github.com/xxjwxc/public/mylog"
)

func init() {
	message.GetErrorMsg(message.UserNotExisted)
	message.GetErrorMsg(message.NotFindError)
	message.GetErrorMsg(message.NotifyIsNotMatch)
	message.GetErrorMsg(message.ParameterInvalid)
	message.GetErrorMsg(message.Overdue)
}

// GetProductDetails 获取订单详情
func (p *Product) GetProductDetails(c *api.Context, req *caoguo.GetProductDetailsReq) (*caoguo.GetProductDetailsRsp, error) {
	orm := core.Dao.GetDBr()
	resp, err := model.ProductTblMgr(orm.Where("vaild = true")).GetFromGpid(req.Gpid)
	if err != nil {
		return nil, err
	}

	if resp.ID == 0 { // not fond
		return nil, fmt.Errorf("失效商品或已下架")
	}

	couponlist, _ := model.CouponTblMgr(orm.Where("vaild = true")).GetFromGpid(resp.Gpid)
	var price int
	var conList []string
	for _, v := range couponlist {
		conList = append(conList, fmt.Sprintf("订单满 %v 减 %v", getMoneyStr(int64(v.GreatMoney)), getMoneyStr(int64(v.Denom))))
		price += v.Denom
	}
	if resp.ShipmentFee == 0 {
		conList = append(conList, "单笔购买免邮费")
	}
	conList = append(conList, "参加平台限时限量优惠")

	couponName := "获取更多优惠"
	if price > 0 {
		couponName = fmt.Sprintf("点击获取 (%v元) 优惠券", getMoneyStr(int64(price)))
	}

	go func() { // 添加一个历史数据
		sessionID := c.GetSessionToken()
		mylog.Infof("sessionid:%v", sessionID)
		session := weixin.GetSessionkey(sessionID)
		if len(session.Openid) == 0 {
			return
		}
		history.Add(session.Openid, resp.Gpid, resp.ProductInfoTbl.Icon)
	}()

	if resp.OriginalPrice < resp.Price { // 原价低于折扣价
		resp.OriginalPrice = resp.Price
	}

	imgs, videos := GetContext(resp.Context)
	return &caoguo.GetProductDetailsRsp{
		Info: &caoguo.ProductInfo{
			Gpid:          resp.Gpid,
			Img:           GetImages(resp.ProductInfoTbl.Img),
			Name:          resp.Name,
			Price:         resp.Price,
			OriginalPrice: resp.OriginalPrice,
			CouponPct:     int32((resp.Price / resp.OriginalPrice * 10)), // 折扣
			Sales:         resp.ProductInfoTbl.Qty,                       // 商品销量
			Stock:         getNumberStr(resp.Qty),                        // 库存
			Views:         getNumberStr(0),
			ShareTit:      "分享活动获取更多奖励",
			Sku:           getSkuInfo(req.Gpid),
			SkuPrice:      getSkuPrice(req.Gpid),
			CouponName:    couponName,
			ConList:       conList,
			BzList:        []string{resp.ProductInfoTbl.Service},
			ImgDetail:     imgs,
			VdDetail:      videos,
			RichText:      "",
			// 	 `
			// 	<div style="width:100%">
			// 		<video src="https://xxx.com/upload/sys/file/60/3e37775cdb97659a01f4a9fdfafe59.mp4" controls="controls">视频</video>
			// 		<img style="width:100%;display:block;" src="https://gd3.alicdn.com/imgextra/i4/479184430/O1CN01nCpuLc1iaz4bcSN17_!!479184430.jpg_400x400.jpg" />
			// 		<img style="width:100%;display:block;" src="https://gd2.alicdn.com/imgextra/i2/479184430/O1CN01gwbN931iaz4TzqzmG_!!479184430.jpg_400x400.jpg" />
			// 		<img style="width:100%;display:block;" src="https://gd3.alicdn.com/imgextra/i3/479184430/O1CN018wVjQh1iaz4aupv1A_!!479184430.jpg_400x400.jpg" />
			// 		<img style="width:100%;display:block;" src="https://gd4.alicdn.com/imgextra/i4/479184430/O1CN01tWg4Us1iaz4auqelt_!!479184430.jpg_400x400.jpg" />
			// 		<img style="width:100%;display:block;" src="https://gd1.alicdn.com/imgextra/i1/479184430/O1CN01Tnm1rU1iaz4aVKcwP_!!479184430.jpg_400x400.jpg" />
			// 	</div>
			// `,
			Icon:       resp.ProductInfoTbl.Icon,
			IsFavorite: getIsFavorite(c, resp.Gpid),
		},
	}, nil
}

func getMoneyStr(money int64) string {
	return fmt.Sprintf("￥%.2f", float32(money)*0.01)
}

// 获取无穷
func getNumberStr(qty int64) string {
	if qty <= 0 {
		return "+∞"
	}

	return fmt.Sprintf("%v", qty)
}

func getSkuInfo(gpid string) *caoguo.ProducktSkuInfo {
	result := &caoguo.ProducktSkuInfo{
		Gpid: gpid,
	}

	orm := core.Dao.GetDBr()
	resp, err := model.ProductSkuTblMgr(orm.DB.Order("id asc")).GetFromGpid(gpid)
	if err != nil {
		return result
	}

	// 聚合
	var typeOrder []string
	mp := make(map[string][]*caoguo.SkuInfo, len(resp))
	for _, v := range resp {
		mp[v.TypeName] = append(mp[v.TypeName], &caoguo.SkuInfo{
			Id:       int64(v.ID),
			TypeName: v.TypeName,
			TagName:  v.TagName,
			// Premium:  v.Premium,
		})

		isFind := false
		for _, v1 := range typeOrder {
			if v1 == v.TypeName {
				isFind = true
				break
			}
		}
		if !isFind {
			typeOrder = append(typeOrder, v.TypeName)
		}
	}

	for _, v := range typeOrder {
		result.Groups = append(result.Groups, &caoguo.SkuGroup{
			TypeName: v,
			Items:    mp[v],
		})
	}
	// ------end

	return result
}

func getSkuPrice(gpid string) *caoguo.ProducktSkuPriceInfo {
	result := &caoguo.ProducktSkuPriceInfo{
		Gpid: gpid,
	}

	orm := core.Dao.GetDBr()
	resp, err := model.ProductSkuPriceTblMgr(orm.DB).GetFromGpid(gpid)
	if err != nil {
		return result
	}
	for _, v := range resp {
		result.List = append(result.List, &caoguo.SkuPriceGroup{
			SkuList: v.SkuList,
			Premium: v.Premium,
		})
	}
	return result
}

// Favorite 收藏 or 取消收藏
func (p *Product) Favorite(c *api.Context, req *caoguo.FavoriteReq) (bool, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	mgr := model.FavoriteTblMgr(core.Dao.GetDBw().DB)
	resp, err := mgr.FetchUniqueIndexByUserGpid(req.Gpid, session.Openid)
	if err != nil {
		if err != gorm.ErrRecordNotFound { // 没找到
			return false, err
		}
	}

	if resp.ID == 0 {
		resp.Gpid = req.Gpid
		resp.UserID = session.Openid
		resp.UserType = 1
		resp.CreatedAt = time.Now()
		if req.IsFavorite { // 添加
			if err := mgr.Save(&resp).Error; err != nil {
				return false, err
			}
		}
	} else {
		if !req.IsFavorite { // 删除
			if err := mgr.Delete(resp).Error; err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// GetFavorite 获取收藏列表
func (p *Product) GetFavorite(c *api.Context, req *common.Empty) (*caoguo.GetFavoriteResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBr()
	favorites, err := model.FavoriteTblMgr(orm.DB).GetFromUserID(session.Openid)
	if err != nil {
		if err != gorm.ErrRecordNotFound { // 没找到
			return nil, err
		}
	}

	var gpids []string
	for _, v := range favorites {
		gpids = append(gpids, v.Gpid)
	}

	products, err := model.ProductTblMgr(orm.DB).GetBatchFromGpid(gpids)
	if err != nil {
		if err != gorm.ErrRecordNotFound { // 没找到
			return nil, err
		}
	}

	resp := &caoguo.GetFavoriteResp{}
	for _, v := range products {
		resp.Items = append(resp.Items, GetSampleProductInfo(v))
	}

	return resp, nil
}

// AddToCart 加入购物车
func (p *Product) AddToCart(c *api.Context, req *caoguo.AddCartReq) (bool, error) {
	sessionID := c.GetSessionToken()
	mylog.Infof("sessionid:%v", sessionID)
	session := weixin.GetSessionkey(sessionID)
	if len(session.Openid) == 0 {
		return false, fmt.Errorf(message.UserNotExisted.String())
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

	mgr := model.CartTblMgr(core.Dao.GetDBw().DB)
	resp, _ := mgr.FetchUniqueIndexByUserGpid(req.Gpid, session.Openid, skuList)
	if req.Number <= 0 { // 删除
		if resp.ID > 0 {
			if err := mgr.Delete(&resp).Error; err != nil {
				return false, err
			}
		}
	} else { // 设置
		resp.Gpid = req.Gpid
		resp.UserID = session.Openid
		resp.UserType = 1
		resp.Number = int(req.Number)
		resp.SkuList = skuList
		if resp.CreatedAt.Unix() > 0 {
			resp.UpdatedAt = time.Now()
		} else {
			resp.CreatedAt = time.Now()
		}
		if err := mgr.Save(&resp).Error; err != nil {
			return false, err
		}
	}

	return true, nil
}

// GetCartList 获取购物车内容
func (p *Product) GetCartList(c *api.Context, req *caoguo.GetCartListReq) (*caoguo.GetCartListResp, error) {
	sessionID := c.GetSessionToken()
	mylog.Infof("sessionid:%v", sessionID)
	session := weixin.GetSessionkey(sessionID)
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	orm := core.Dao.GetDBw()
	mgr := model.CartTblMgr(orm.DB)
	list, err := mgr.GetFromUserID(session.Openid)
	if err != nil {
		if !orm.IsNotFound(err) {
			return nil, fmt.Errorf(message.NotFindError.String())
		}
	}

	cartMp := make(map[int64]*model.CartTbl, len(list))
	var gpids, skuids []string
	resp := &caoguo.GetCartListResp{}
	for _, v := range list {
		resp.Items = append(resp.Items, &caoguo.CartInfo{
			Id:     int64(v.ID),
			Number: int64(v.Number),
		})
		cartMp[int64(v.ID)] = v
		gpids = append(gpids, v.Gpid)
		if len(v.SkuList) > 0 {
			skuids = append(skuids, strings.Split(v.SkuList, ",")...)
		}
	}

	if len(resp.Items) == 0 { // 空
		return resp, nil
	}

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

	// 开始封装
	for _, v := range resp.Items {
		gpid := cartMp[v.Id].Gpid
		info := productMp[gpid]

		v.Icon = info.ProductInfoTbl.Icon

		// sku
		var offsetPrice int64
		if len(cartMp[v.Id].SkuList) > 0 {
			skuprice, _ := model.ProductSkuPriceTblMgr(orm.DB).GetFromSkuList(cartMp[v.Id].SkuList)
			offsetPrice += skuprice.Premium
		}

		skulist := strings.Split(cartMp[v.Id].SkuList, ",")
		var skuStr []string
		for _, v := range skulist {
			if len(v) > 0 {
				if _, ok := skuMp[v]; ok {
					skuStr = append(skuStr, skuMp[v].TagName)
				}
			}
		}
		v.SkuVal = strings.Join(skuStr, " ")
		// ------ end

		v.Stock = info.Qty // 库存
		v.Name = info.Name
		v.Price = info.Price + offsetPrice
		v.OriginalPrice = info.OriginalPrice + offsetPrice
	}

	return resp, nil
}

// ChangeCart 改变
func (p *Product) ChangeCart(c *api.Context, req *caoguo.ChangeCardReq) (bool, error) {
	sessionID := c.GetSessionToken()
	mylog.Infof("sessionid:%v", sessionID)
	session := weixin.GetSessionkey(sessionID)
	if len(session.Openid) == 0 {
		return false, fmt.Errorf(message.UserNotExisted.String())
	}

	mgr := model.CartTblMgr(core.Dao.GetDBw().DB)

	if req.Id == 0 { // 清空购物车
		mgr.Where("user_id = ?", session.Openid).Delete(&model.CartTbl{})
		return true, nil
	}

	resp, err := mgr.GetFromID(req.Id)
	if err != nil {
		return false, err
	}

	if req.Number <= 0 { // 删除
		if err := mgr.Delete(&resp).Error; err != nil {
			return false, err
		}
	} else { // 修改
		resp.Number = int(req.Number)
		resp.UpdatedAt = time.Now()
		if err := mgr.Save(&resp).Error; err != nil {
			return false, err
		}
	}

	return true, nil
}

// AddToBuyTmpCart 添加到直接购买虚拟购物车里面
func (p *Product) AddToBuyTmpCart(c *api.Context, req *caoguo.AddToBuyTmpCartReq) (*caoguo.AddToBuyTmpCartResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	mgr := model.CartTmpTblMgr(core.Dao.GetDBw().DB)

	resp, err := mgr.GetFromUserID(session.Openid)
	if err != nil {
		if err != gorm.ErrRecordNotFound { // 没找到
			return nil, err
		}
	}

	// 添加或者更新
	resp.Gpid = req.Gpid
	resp.UserID = session.Openid
	resp.UserType = 1
	resp.Number = int(req.Number)

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

	resp.SkuList = skuList
	resp.CreatedAt = time.Now()
	err = mgr.Save(&resp).Error
	if err != nil {
		return nil, err
	}

	mylog.Info(resp)

	return &caoguo.AddToBuyTmpCartResp{
		Id: int64(resp.ID),
	}, err
}

// AddAddress 添加或者修改地址
func (p *Product) AddAddress(c *api.Context, req *caoguo.AddAddressReq) (_i int64, _err error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return 0, fmt.Errorf(message.UserNotExisted.String())
	}
	// if req.Addr.AddressName == "在地图选择" {
	// 	return 0, fmt.Errorf("请填写小区信息")
	// }

	req.Addr.Name = tools.UnicodeEmojiCode(req.Addr.Name)
	req.Addr.AddressName = tools.UnicodeEmojiCode(req.Addr.AddressName)
	req.Addr.Address = tools.UnicodeEmojiCode(req.Addr.Address)
	req.Addr.Area = tools.UnicodeEmojiCode(req.Addr.Area)

	orm := core.Dao.GetDBw()
	tx := orm.Begin()
	defer func() {
		if _err != nil {
			tx.AddError(_err)
		}
		orm.Commit(tx)
	}()

	mgr := model.AddressTblMgr(tx)
	var addr model.AddressTbl
	if req.Addr.Id > 0 { // find
		tmp, err := mgr.GetFromID(req.Addr.Id)
		if err != nil {
			return 0, err
		}
		if tmp.UserID != session.Openid { // 失败
			return 0, fmt.Errorf(message.NotifyIsNotMatch.String())
		}
		addr = tmp
	}

	callResetDefalt := func() error { // 重置默认值
		if req.Addr.Default == true {
			return mgr.Model(model.AddressTbl{}).Where("user_id = ? and `default` = true", session.Openid).Update("default", false).Error
		}
		return nil
	}

	switch req.Op {
	case 1: // 添加
		if err := callResetDefalt(); err != nil {
			return 0, err
		}

		addr = model.AddressTbl{
			UserID:      session.Openid,
			Name:        req.Addr.Name,
			Mobile:      req.Addr.Mobile,
			AddressName: req.Addr.AddressName,
			Address:     req.Addr.Address,
			Area:        req.Addr.Area,
			Default:     req.Addr.Default,
			CreatedAt:   time.Now(),
		}
		err := mgr.Save(&addr).Error
		if err != nil {
			return 0, err
		}
	case 2: // 修改
		if req.Addr.Id <= 0 {
			return 0, fmt.Errorf(message.ParameterInvalid.String())
		}
		if err := callResetDefalt(); err != nil {
			return 0, err
		}
		addr = model.AddressTbl{
			ID:          req.Addr.Id,
			UserID:      session.Openid,
			Name:        req.Addr.Name,
			Mobile:      req.Addr.Mobile,
			AddressName: req.Addr.AddressName,
			Address:     req.Addr.Address,
			Area:        req.Addr.Area,
			Default:     req.Addr.Default,
			CreatedAt:   time.Now(),
		}
		err := mgr.Save(&addr).Error
		if err != nil {
			return 0, err
		}
	case 3: // 删除
		if req.Addr.Id <= 0 {
			return 0, fmt.Errorf(message.ParameterInvalid.String())
		}
		err := mgr.Model(model.AddressTbl{}).Where("id = ? and user_id = ?", req.Addr.Id, session.Openid).Delete(&model.AddressTbl{}).Error
		if err != nil {
			return 0, err
		}
	}

	// 到此全部成功
	// tx.Commit()

	return addr.ID, nil
}

// GetAddress 获取地址
func (p *Product) GetAddress(c *api.Context, req *caoguo.GetAddressReq) (*caoguo.GetAddressResp, error) {
	session := weixin.GetSessionkey(c.GetSessionToken())
	if len(session.Openid) == 0 {
		return nil, fmt.Errorf(message.UserNotExisted.String())
	}

	db := core.Dao.GetDBr().DB
	db = db.Where("user_id = ?", session.Openid)
	// if req.IsDefault {
	// 	db = db.Where("default = true")
	// }
	mgr := model.AddressTblMgr(db)

	resp, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	result := &caoguo.GetAddressResp{}
	for _, v := range resp {
		if v.Default || !req.IsDefault {
			v.Name = tools.UnicodeEmojiDecode(v.Name)
			v.AddressName = tools.UnicodeEmojiDecode(v.AddressName)
			v.Address = tools.UnicodeEmojiDecode(v.Address)
			v.Area = tools.UnicodeEmojiDecode(v.Area)
			result.Addrs = append(result.Addrs, &caoguo.AddressInfo{
				Id:          v.ID,
				Name:        v.Name,
				Mobile:      v.Mobile,
				AddressName: v.AddressName,
				Address:     v.Address,
				Area:        v.Area,
				Default:     v.Default,
			})
		}
	}

	if len(result.Addrs) == 0 && len(resp) > 0 { // 随机获取一个
		v := resp[0]
		v.Name = tools.UnicodeEmojiDecode(v.Name)
		v.AddressName = tools.UnicodeEmojiDecode(v.AddressName)
		v.Address = tools.UnicodeEmojiDecode(v.Address)
		v.Area = tools.UnicodeEmojiDecode(v.Area)
		result.Addrs = append(result.Addrs, &caoguo.AddressInfo{
			Id:          v.ID,
			Name:        v.Name,
			Mobile:      v.Mobile,
			AddressName: v.AddressName,
			Address:     v.Address,
			Area:        v.Area,
			Default:     v.Default,
		})
	}

	return result, nil
}

// GetProductByType 通过商品类型获取商品列表
func (p *Product) GetProductByType(c *api.Context, req *caoguo.GetProductByTypeReq) (*caoguo.GetProductByTypeResp, error) {
	resp := &caoguo.GetProductByTypeResp{
		TypeId:     req.TypeId,
		PageNumber: req.PageNumber, // 当前页数
	}

	id, _ := strconv.Atoi(req.TypeId)
	// 先通过类型获取
	orm := core.Dao.GetDBr()
	typeInfo, _ := model.ProductTypeTblMgr(orm.DB).GetFromID(id)
	db := orm.Offset(int(req.PageNumber * _ProductPageLen)).Limit(_ProductPageLen).Order("created_at desc").Where("vaild = true")
	if len(typeInfo.Gtid) > 0 {
		db = db.Where("gtid = ?", typeInfo.Gtid)
	}
	products, err := model.ProductTblMgr(db).Gets()
	if err != nil {
		if orm.IsNotFound(err) {
			return resp, nil
		}
		return nil, err
	}

	if len(products) == 0 { // not found
		return resp, nil
	}

	for _, v := range products {
		resp.Items = append(resp.Items, GetSampleProductInfo(v))
	}

	// 找到部分
	if len(products) < _ProductPageLen {
		_len := _ProductPageLen - len(products) // 其它来填充
		db := orm.Order("RAND()").Limit(_len).Order("created_at desc").Where("vaild = true")
		if len(typeInfo.Gtid) > 0 {
			db = db.Where("gtid != ?", typeInfo.Gtid)
		}

		// 获取喜欢列表(随机取数)
		lkinfo, _ := model.ProductTblMgr(db).Gets()
		for _, v := range lkinfo {
			resp.Items = append(resp.Items, GetSampleProductInfo(v))
		}
	}

	return resp, nil
}
