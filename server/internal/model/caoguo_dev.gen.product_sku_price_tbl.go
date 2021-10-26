package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductSkuPriceTblMgr struct {
	*_BaseMgr
}

// ProductSkuPriceTblMgr open func
func ProductSkuPriceTblMgr(db *gorm.DB) *_ProductSkuPriceTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductSkuPriceTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductSkuPriceTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_sku_price_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductSkuPriceTblMgr) GetTableName() string {
	return "product_sku_price_tbl"
}

// Get 获取
func (obj *_ProductSkuPriceTblMgr) Get() (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductSkuPriceTblMgr) Gets() (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_ProductSkuPriceTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_ProductSkuPriceTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_ProductSkuPriceTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_ProductSkuPriceTblMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithGpid gpid获取 商品id
func (obj *_ProductSkuPriceTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithSkuList sku_list获取 属性列表
func (obj *_ProductSkuPriceTblMgr) WithSkuList(skuList string) Option {
	return optionFunc(func(o *options) { o.query["sku_list"] = skuList })
}

// WithPremium premium获取 商品单价
func (obj *_ProductSkuPriceTblMgr) WithPremium(premium int64) Option {
	return optionFunc(func(o *options) { o.query["premium"] = premium })
}

// WithDistAmount dist_amount获取 分享收益
func (obj *_ProductSkuPriceTblMgr) WithDistAmount(distAmount int64) Option {
	return optionFunc(func(o *options) { o.query["dist_amount"] = distAmount })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProductSkuPriceTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProductSkuPriceTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithDeletedBy deleted_by获取 删除者
func (obj *_ProductSkuPriceTblMgr) WithDeletedBy(deletedBy string) Option {
	return optionFunc(func(o *options) { o.query["deleted_by"] = deletedBy })
}

// GetByOption 功能选项模式获取
func (obj *_ProductSkuPriceTblMgr) GetByOption(opts ...Option) (result ProductSkuPriceTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductSkuPriceTblMgr) GetByOptions(opts ...Option) (results []*ProductSkuPriceTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_ProductSkuPriceTblMgr) GetFromID(id int64) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 Primary key
func (obj *_ProductSkuPriceTblMgr) GetBatchFromID(ids []int64) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_ProductSkuPriceTblMgr) GetFromCreatedAt(createdAt time.Time) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 created time
func (obj *_ProductSkuPriceTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_ProductSkuPriceTblMgr) GetFromUpdatedAt(updatedAt time.Time) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 updated at
func (obj *_ProductSkuPriceTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_ProductSkuPriceTblMgr) GetFromDeletedAt(deletedAt time.Time) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找 deleted time
func (obj *_ProductSkuPriceTblMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_ProductSkuPriceTblMgr) GetFromGpid(gpid string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_ProductSkuPriceTblMgr) GetBatchFromGpid(gpids []string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromSkuList 通过sku_list获取内容 属性列表
func (obj *_ProductSkuPriceTblMgr) GetFromSkuList(skuList string) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list = ?", skuList).Find(&result).Error

	return
}

// GetBatchFromSkuList 批量唯一主键查找 属性列表
func (obj *_ProductSkuPriceTblMgr) GetBatchFromSkuList(skuLists []string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list IN (?)", skuLists).Find(&results).Error

	return
}

// GetFromPremium 通过premium获取内容 商品单价
func (obj *_ProductSkuPriceTblMgr) GetFromPremium(premium int64) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("premium = ?", premium).Find(&results).Error

	return
}

// GetBatchFromPremium 批量唯一主键查找 商品单价
func (obj *_ProductSkuPriceTblMgr) GetBatchFromPremium(premiums []int64) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("premium IN (?)", premiums).Find(&results).Error

	return
}

// GetFromDistAmount 通过dist_amount获取内容 分享收益
func (obj *_ProductSkuPriceTblMgr) GetFromDistAmount(distAmount int64) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_amount = ?", distAmount).Find(&results).Error

	return
}

// GetBatchFromDistAmount 批量唯一主键查找 分享收益
func (obj *_ProductSkuPriceTblMgr) GetBatchFromDistAmount(distAmounts []int64) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_amount IN (?)", distAmounts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProductSkuPriceTblMgr) GetFromCreatedBy(createdBy string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProductSkuPriceTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProductSkuPriceTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProductSkuPriceTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromDeletedBy 通过deleted_by获取内容 删除者
func (obj *_ProductSkuPriceTblMgr) GetFromDeletedBy(deletedBy string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by = ?", deletedBy).Find(&results).Error

	return
}

// GetBatchFromDeletedBy 批量唯一主键查找 删除者
func (obj *_ProductSkuPriceTblMgr) GetBatchFromDeletedBy(deletedBys []string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by IN (?)", deletedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductSkuPriceTblMgr) FetchByPrimaryKey(id int64) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueBySkuList primay or index 获取唯一内容
func (obj *_ProductSkuPriceTblMgr) FetchUniqueBySkuList(skuList string) (result ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list = ?", skuList).Find(&result).Error

	return
}

// FetchIndexByGpid  获取多个内容
func (obj *_ProductSkuPriceTblMgr) FetchIndexByGpid(gpid string) (results []*ProductSkuPriceTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}
