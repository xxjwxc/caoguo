package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CartTblMgr struct {
	*_BaseMgr
}

// CartTblMgr open func
func CartTblMgr(db *gorm.DB) *_CartTblMgr {
	if db == nil {
		panic(fmt.Errorf("CartTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cart_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartTblMgr) GetTableName() string {
	return "cart_tbl"
}

// Get 获取
func (obj *_CartTblMgr) Get() (result CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartTblMgr) Gets() (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CartTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGpid gpid获取 商品id
func (obj *_CartTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithUserID user_id获取 用户id
func (obj *_CartTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserType user_type获取 用户类型(1:微信)
func (obj *_CartTblMgr) WithUserType(userType int) Option {
	return optionFunc(func(o *options) { o.query["user_type"] = userType })
}

// WithNumber number获取 数量
func (obj *_CartTblMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithSkuList sku_list获取 属性列表
func (obj *_CartTblMgr) WithSkuList(skuList string) Option {
	return optionFunc(func(o *options) { o.query["sku_list"] = skuList })
}

// WithCreatedBy created_by获取 创建者
func (obj *_CartTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_CartTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_CartTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_CartTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_CartTblMgr) GetByOption(opts ...Option) (result CartTbl, err error) {
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
func (obj *_CartTblMgr) GetByOptions(opts ...Option) (results []*CartTbl, err error) {
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
func (obj *_CartTblMgr) GetFromID(id int64) (result CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CartTblMgr) GetBatchFromID(ids []int64) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_CartTblMgr) GetFromGpid(gpid string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_CartTblMgr) GetBatchFromGpid(gpids []string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_CartTblMgr) GetFromUserID(userID string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_CartTblMgr) GetBatchFromUserID(userIDs []string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserType 通过user_type获取内容 用户类型(1:微信)
func (obj *_CartTblMgr) GetFromUserType(userType int) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type = ?", userType).Find(&results).Error

	return
}

// GetBatchFromUserType 批量唯一主键查找 用户类型(1:微信)
func (obj *_CartTblMgr) GetBatchFromUserType(userTypes []int) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type IN (?)", userTypes).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 数量
func (obj *_CartTblMgr) GetFromNumber(number int) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量唯一主键查找 数量
func (obj *_CartTblMgr) GetBatchFromNumber(numbers []int) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number IN (?)", numbers).Find(&results).Error

	return
}

// GetFromSkuList 通过sku_list获取内容 属性列表
func (obj *_CartTblMgr) GetFromSkuList(skuList string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list = ?", skuList).Find(&results).Error

	return
}

// GetBatchFromSkuList 批量唯一主键查找 属性列表
func (obj *_CartTblMgr) GetBatchFromSkuList(skuLists []string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list IN (?)", skuLists).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_CartTblMgr) GetFromCreatedBy(createdBy string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_CartTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_CartTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_CartTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_CartTblMgr) GetFromUpdatedBy(updatedBy string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_CartTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_CartTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_CartTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartTblMgr) FetchByPrimaryKey(id int64) (result CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUserGpid primay or index 获取唯一内容
func (obj *_CartTblMgr) FetchUniqueIndexByUserGpid(gpid string, userID string, skuList string) (result CartTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ? AND user_id = ? AND sku_list = ?", gpid, userID, skuList).Find(&result).Error

	return
}
