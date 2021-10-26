package product

import (
	"caoguo/internal/api"
	"caoguo/internal/core"
	"caoguo/internal/model"
	"caoguo/rpc/caoguo"
	"caoguo/rpc/common"
)

// GetAdInfo 获取广告主信息列表(可能更新)
func (p *Product) GetAdInfo(c *api.Context, req *common.Empty) (*caoguo.GetAdResp, error) {
	var resp caoguo.GetAdResp
	// 用户判断

	orm := core.Dao.GetDBr()
	db := orm.Order("sort_id desc")
	ads, err := model.ProductAdTblMgr(db).GetFromVaild(true)
	if err != nil {
		return nil, err
	}

	for _, v := range ads {
		switch v.Type {
		case 1: // 1:轮播图广告，
			resp.RotationList = append(resp.RotationList, &caoguo.AdInfo{
				Url:    v.URL,
				Img:    v.Img,
				SortId: int32(v.SortID),
				Attach: v.Attach,
			})
		case 2: // 2:类型广告
			resp.TypeList = append(resp.TypeList, &caoguo.AdInfo{
				Url:    v.URL,
				Img:    v.Img,
				SortId: int32(v.SortID),
				Attach: v.Attach,
			})
		case 3: //3:主销广告
			resp.MasterInfo = &caoguo.AdInfo{
				Url:    v.URL,
				Img:    v.Img,
				SortId: int32(v.SortID),
				Attach: v.Attach,
			}
		}
	}

	return &resp, nil
}

// GetBoutiqueGroup 获取团购列表
func (p *Product) GetBoutiqueGroup(c *api.Context, req *common.Empty) (*caoguo.BoutiqueGroupResp, error) {
	orm := core.Dao.GetDBr()
	adgrop, err := model.ProductAdGroupTblMgr(orm.Order("sort_id desc").Where("vaild = true")).GetFromVaild(true)
	if err != nil {
		return nil, err
	}

	// 获取商品信息
	var gpids []string
	for _, v := range adgrop {
		gpids = append(gpids, v.MainGpid)
		gpids = append(gpids, v.SubGpid)
	}

	mp, err := GetProductByGpids(gpids)
	if err != nil {
		return nil, err
	}

	resp := &caoguo.BoutiqueGroupResp{}
	for _, v := range adgrop { // 团购列表
		resp.Groups = append(resp.Groups, &caoguo.BoutiqueGroupInfo{
			MainProduct: GetSampleProductInfo(mp[v.MainGpid]),
			SubProduct:  GetSampleProductInfo(mp[v.SubGpid]),
		})
	}

	// 获取喜欢列表(随机取数)
	lkinfo, err := model.ProductTblMgr(orm.Order("RAND()").Limit(4)).GetFromVaild(true)
	if err != nil {
		return nil, err
	}
	for _, v := range lkinfo {
		resp.Likes = append(resp.Likes, GetSampleProductInfo(v))
	}
	// -----------end

	return resp, nil
}

// GetSampleProductInfo 转换product info
func GetSampleProductInfo(in *model.ProductTbl) *caoguo.SampleProductInfo {
	if in == nil {
		return nil
	}

	return &caoguo.SampleProductInfo{
		Gpid:          in.Gpid,
		Name:          in.Name,
		Price:         in.Price,
		OriginalPrice: in.OriginalPrice,
		Sales:         in.ProductInfoTbl.Qty,
		Img:           GetImages(in.ProductInfoTbl.Img),
		Icon:          in.ProductInfoTbl.Icon,
		Percent:       50,
	}

}
