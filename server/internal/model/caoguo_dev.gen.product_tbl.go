package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductTblMgr struct {
	*_BaseMgr
}

// ProductTblMgr open func
func ProductTblMgr(db *gorm.DB) *_ProductTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductTblMgr) GetTableName() string {
	return "product_tbl"
}

// Get 获取
func (obj *_ProductTblMgr) Get() (result ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ProductTblMgr) Gets() (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增
func (obj *_ProductTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGpid gpid获取 商品id
func (obj *_ProductTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithGtid gtid获取 商品分类id
func (obj *_ProductTblMgr) WithGtid(gtid string) Option {
	return optionFunc(func(o *options) { o.query["gtid"] = gtid })
}

// WithName name获取 商品名字
func (obj *_ProductTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 商品价格(分)
func (obj *_ProductTblMgr) WithPrice(price int64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithOriginalPrice original_price获取 商品原始价格(分)
func (obj *_ProductTblMgr) WithOriginalPrice(originalPrice int64) Option {
	return optionFunc(func(o *options) { o.query["original_price"] = originalPrice })
}

// WithDistAmount dist_amount获取 分享收益
func (obj *_ProductTblMgr) WithDistAmount(distAmount int64) Option {
	return optionFunc(func(o *options) { o.query["dist_amount"] = distAmount })
}

// WithPlatformID platform_id获取 商铺客户id
func (obj *_ProductTblMgr) WithPlatformID(platformID string) Option {
	return optionFunc(func(o *options) { o.query["platform_id"] = platformID })
}

// WithQty qty获取 数量(库存)
func (obj *_ProductTblMgr) WithQty(qty int64) Option {
	return optionFunc(func(o *options) { o.query["qty"] = qty })
}

// WithShipmentTypeID shipment_type_id获取 运输方式类型
func (obj *_ProductTblMgr) WithShipmentTypeID(shipmentTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["shipment_type_id"] = shipmentTypeID })
}

// WithShipmentFee shipment_fee获取 运费
func (obj *_ProductTblMgr) WithShipmentFee(shipmentFee int64) Option {
	return optionFunc(func(o *options) { o.query["shipment_fee"] = shipmentFee })
}

// WithContext context获取 商品详情
func (obj *_ProductTblMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["context"] = context })
}

// WithVolumeWeight volume_weight获取 体积重
func (obj *_ProductTblMgr) WithVolumeWeight(volumeWeight int64) Option {
	return optionFunc(func(o *options) { o.query["volume_weight"] = volumeWeight })
}

// WithActualWeight actual_weight获取 实际重
func (obj *_ProductTblMgr) WithActualWeight(actualWeight int64) Option {
	return optionFunc(func(o *options) { o.query["actual_weight"] = actualWeight })
}

// WithSourceRegion source_region获取 货源地名
func (obj *_ProductTblMgr) WithSourceRegion(sourceRegion string) Option {
	return optionFunc(func(o *options) { o.query["source_region"] = sourceRegion })
}

// WithMaxBuyQty max_buy_qty获取 最大购买数量
func (obj *_ProductTblMgr) WithMaxBuyQty(maxBuyQty int) Option {
	return optionFunc(func(o *options) { o.query["max_buy_qty"] = maxBuyQty })
}

// WithIsSelect is_select获取 是否推荐
func (obj *_ProductTblMgr) WithIsSelect(isSelect bool) Option {
	return optionFunc(func(o *options) { o.query["is_select"] = isSelect })
}

// WithVaild vaild获取 是否有效
func (obj *_ProductTblMgr) WithVaild(vaild bool) Option {
	return optionFunc(func(o *options) { o.query["vaild"] = vaild })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProductTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ProductTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProductTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ProductTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ProductTblMgr) GetByOption(opts ...Option) (result ProductTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductTblMgr) GetByOptions(opts ...Option) (results []*ProductTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 自增
func (obj *_ProductTblMgr) GetFromID(id int64) (result ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找 自增
func (obj *_ProductTblMgr) GetBatchFromID(ids []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_ProductTblMgr) GetFromGpid(gpid string) (result ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_ProductTblMgr) GetBatchFromGpid(gpids []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGtid 通过gtid获取内容 商品分类id
func (obj *_ProductTblMgr) GetFromGtid(gtid string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gtid = ?", gtid).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGtid 批量唯一主键查找 商品分类id
func (obj *_ProductTblMgr) GetBatchFromGtid(gtids []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gtid IN (?)", gtids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容 商品名字
func (obj *_ProductTblMgr) GetFromName(name string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量唯一主键查找 商品名字
func (obj *_ProductTblMgr) GetBatchFromName(names []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPrice 通过price获取内容 商品价格(分)
func (obj *_ProductTblMgr) GetFromPrice(price int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPrice 批量唯一主键查找 商品价格(分)
func (obj *_ProductTblMgr) GetBatchFromPrice(prices []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOriginalPrice 通过original_price获取内容 商品原始价格(分)
func (obj *_ProductTblMgr) GetFromOriginalPrice(originalPrice int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_price = ?", originalPrice).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOriginalPrice 批量唯一主键查找 商品原始价格(分)
func (obj *_ProductTblMgr) GetBatchFromOriginalPrice(originalPrices []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_price IN (?)", originalPrices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDistAmount 通过dist_amount获取内容 分享收益
func (obj *_ProductTblMgr) GetFromDistAmount(distAmount int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_amount = ?", distAmount).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDistAmount 批量唯一主键查找 分享收益
func (obj *_ProductTblMgr) GetBatchFromDistAmount(distAmounts []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_amount IN (?)", distAmounts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPlatformID 通过platform_id获取内容 商铺客户id
func (obj *_ProductTblMgr) GetFromPlatformID(platformID string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("platform_id = ?", platformID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPlatformID 批量唯一主键查找 商铺客户id
func (obj *_ProductTblMgr) GetBatchFromPlatformID(platformIDs []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("platform_id IN (?)", platformIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromQty 通过qty获取内容 数量(库存)
func (obj *_ProductTblMgr) GetFromQty(qty int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qty = ?", qty).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromQty 批量唯一主键查找 数量(库存)
func (obj *_ProductTblMgr) GetBatchFromQty(qtys []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qty IN (?)", qtys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromShipmentTypeID 通过shipment_type_id获取内容 运输方式类型
func (obj *_ProductTblMgr) GetFromShipmentTypeID(shipmentTypeID int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_type_id = ?", shipmentTypeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromShipmentTypeID 批量唯一主键查找 运输方式类型
func (obj *_ProductTblMgr) GetBatchFromShipmentTypeID(shipmentTypeIDs []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_type_id IN (?)", shipmentTypeIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromShipmentFee 通过shipment_fee获取内容 运费
func (obj *_ProductTblMgr) GetFromShipmentFee(shipmentFee int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_fee = ?", shipmentFee).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromShipmentFee 批量唯一主键查找 运费
func (obj *_ProductTblMgr) GetBatchFromShipmentFee(shipmentFees []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_fee IN (?)", shipmentFees).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromContext 通过context获取内容 商品详情
func (obj *_ProductTblMgr) GetFromContext(context string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("context = ?", context).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromContext 批量唯一主键查找 商品详情
func (obj *_ProductTblMgr) GetBatchFromContext(contexts []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("context IN (?)", contexts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVolumeWeight 通过volume_weight获取内容 体积重
func (obj *_ProductTblMgr) GetFromVolumeWeight(volumeWeight int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("volume_weight = ?", volumeWeight).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVolumeWeight 批量唯一主键查找 体积重
func (obj *_ProductTblMgr) GetBatchFromVolumeWeight(volumeWeights []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("volume_weight IN (?)", volumeWeights).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromActualWeight 通过actual_weight获取内容 实际重
func (obj *_ProductTblMgr) GetFromActualWeight(actualWeight int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("actual_weight = ?", actualWeight).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromActualWeight 批量唯一主键查找 实际重
func (obj *_ProductTblMgr) GetBatchFromActualWeight(actualWeights []int64) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("actual_weight IN (?)", actualWeights).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSourceRegion 通过source_region获取内容 货源地名
func (obj *_ProductTblMgr) GetFromSourceRegion(sourceRegion string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("source_region = ?", sourceRegion).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSourceRegion 批量唯一主键查找 货源地名
func (obj *_ProductTblMgr) GetBatchFromSourceRegion(sourceRegions []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("source_region IN (?)", sourceRegions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMaxBuyQty 通过max_buy_qty获取内容 最大购买数量
func (obj *_ProductTblMgr) GetFromMaxBuyQty(maxBuyQty int) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_buy_qty = ?", maxBuyQty).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMaxBuyQty 批量唯一主键查找 最大购买数量
func (obj *_ProductTblMgr) GetBatchFromMaxBuyQty(maxBuyQtys []int) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_buy_qty IN (?)", maxBuyQtys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsSelect 通过is_select获取内容 是否推荐
func (obj *_ProductTblMgr) GetFromIsSelect(isSelect bool) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_select = ?", isSelect).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsSelect 批量唯一主键查找 是否推荐
func (obj *_ProductTblMgr) GetBatchFromIsSelect(isSelects []bool) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_select IN (?)", isSelects).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVaild 通过vaild获取内容 是否有效
func (obj *_ProductTblMgr) GetFromVaild(vaild bool) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild = ?", vaild).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVaild 批量唯一主键查找 是否有效
func (obj *_ProductTblMgr) GetBatchFromVaild(vailds []bool) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild IN (?)", vailds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProductTblMgr) GetFromCreatedBy(createdBy string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProductTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ProductTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_ProductTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProductTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProductTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ProductTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_ProductTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_info_tbl").Where("gpid = ?", results[i].Gpid).Find(&results[i].ProductInfoTbl).Error; err != nil { // 商品附加属性
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductTblMgr) FetchByPrimaryKey(id int64) (result ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByGpid primay or index 获取唯一内容
func (obj *_ProductTblMgr) FetchUniqueByGpid(gpid string) (result ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByNameID primay or index 获取唯一内容
func (obj *_ProductTblMgr) FetchUniqueIndexByNameID(name string, platformID string) (result ProductTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ? AND platform_id = ?", name, platformID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_info_tbl").Where("gpid = ?", result.Gpid).Find(&result.ProductInfoTbl).Error; err != nil { // 商品附加属性
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
