package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CartTmpTblMgr struct {
	*_BaseMgr
}

// CartTmpTblMgr open func
func CartTmpTblMgr(db *gorm.DB) *_CartTmpTblMgr {
	if db == nil {
		panic(fmt.Errorf("CartTmpTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartTmpTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cart_tmp_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartTmpTblMgr) GetTableName() string {
	return "cart_tmp_tbl"
}

// Get 获取
func (obj *_CartTmpTblMgr) Get() (result CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartTmpTblMgr) Gets() (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CartTmpTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGpid gpid获取 商品id
func (obj *_CartTmpTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithUserID user_id获取 用户id
func (obj *_CartTmpTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserType user_type获取 用户类型(1:微信)
func (obj *_CartTmpTblMgr) WithUserType(userType int) Option {
	return optionFunc(func(o *options) { o.query["user_type"] = userType })
}

// WithNumber number获取 数量
func (obj *_CartTmpTblMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithSkuList sku_list获取 属性列表
func (obj *_CartTmpTblMgr) WithSkuList(skuList string) Option {
	return optionFunc(func(o *options) { o.query["sku_list"] = skuList })
}

// WithCreatedBy created_by获取 创建者
func (obj *_CartTmpTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_CartTmpTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_CartTmpTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_CartTmpTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_CartTmpTblMgr) GetByOption(opts ...Option) (result CartTmpTbl, err error) {
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
func (obj *_CartTmpTblMgr) GetByOptions(opts ...Option) (results []*CartTmpTbl, err error) {
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

// GetFromID 通过id获取内容
func (obj *_CartTmpTblMgr) GetFromID(id int64) (result CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CartTmpTblMgr) GetBatchFromID(ids []int64) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_CartTmpTblMgr) GetFromGpid(gpid string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_CartTmpTblMgr) GetBatchFromGpid(gpids []string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_CartTmpTblMgr) GetFromUserID(userID string) (result CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_CartTmpTblMgr) GetBatchFromUserID(userIDs []string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserType 通过user_type获取内容 用户类型(1:微信)
func (obj *_CartTmpTblMgr) GetFromUserType(userType int) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type = ?", userType).Find(&results).Error

	return
}

// GetBatchFromUserType 批量唯一主键查找 用户类型(1:微信)
func (obj *_CartTmpTblMgr) GetBatchFromUserType(userTypes []int) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type IN (?)", userTypes).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 数量
func (obj *_CartTmpTblMgr) GetFromNumber(number int) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量唯一主键查找 数量
func (obj *_CartTmpTblMgr) GetBatchFromNumber(numbers []int) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number IN (?)", numbers).Find(&results).Error

	return
}

// GetFromSkuList 通过sku_list获取内容 属性列表
func (obj *_CartTmpTblMgr) GetFromSkuList(skuList string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list = ?", skuList).Find(&results).Error

	return
}

// GetBatchFromSkuList 批量唯一主键查找 属性列表
func (obj *_CartTmpTblMgr) GetBatchFromSkuList(skuLists []string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list IN (?)", skuLists).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_CartTmpTblMgr) GetFromCreatedBy(createdBy string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_CartTmpTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_CartTmpTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_CartTmpTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_CartTmpTblMgr) GetFromUpdatedBy(updatedBy string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_CartTmpTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_CartTmpTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_CartTmpTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartTmpTblMgr) FetchByPrimaryKey(id int64) (result CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUserID primay or index 获取唯一内容
func (obj *_CartTmpTblMgr) FetchUniqueByUserID(userID string) (result CartTmpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}
