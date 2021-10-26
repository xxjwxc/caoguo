package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserTblMgr struct {
	*_BaseMgr
}

// UserTblMgr open func
func UserTblMgr(db *gorm.DB) *_UserTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserTblMgr) GetTableName() string {
	return "user_tbl"
}

// Get 获取
func (obj *_UserTblMgr) Get() (result UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserTblMgr) Gets() (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 微信用户唯一标识符
func (obj *_UserTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithVip vip获取 是否vip
func (obj *_UserTblMgr) WithVip(vip bool) Option {
	return optionFunc(func(o *options) { o.query["vip"] = vip })
}

// WithDistVip dist_vip获取 是否分销vip
func (obj *_UserTblMgr) WithDistVip(distVip bool) Option {
	return optionFunc(func(o *options) { o.query["dist_vip"] = distVip })
}

// WithBalance balance获取 用户余额
func (obj *_UserTblMgr) WithBalance(balance int64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithPoints points获取 用户积分
func (obj *_UserTblMgr) WithPoints(points int64) Option {
	return optionFunc(func(o *options) { o.query["points"] = points })
}

// WithCreatedBy created_by获取 创建者
func (obj *_UserTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UserTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_UserTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_UserTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserTblMgr) GetByOption(opts ...Option) (result UserTbl, err error) {
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
func (obj *_UserTblMgr) GetByOptions(opts ...Option) (results []*UserTbl, err error) {
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
func (obj *_UserTblMgr) GetFromID(id int) (result UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserTblMgr) GetBatchFromID(ids []int) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 微信用户唯一标识符
func (obj *_UserTblMgr) GetFromUserID(userID string) (result UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 微信用户唯一标识符
func (obj *_UserTblMgr) GetBatchFromUserID(userIDs []string) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromVip 通过vip获取内容 是否vip
func (obj *_UserTblMgr) GetFromVip(vip bool) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vip = ?", vip).Find(&results).Error

	return
}

// GetBatchFromVip 批量唯一主键查找 是否vip
func (obj *_UserTblMgr) GetBatchFromVip(vips []bool) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vip IN (?)", vips).Find(&results).Error

	return
}

// GetFromDistVip 通过dist_vip获取内容 是否分销vip
func (obj *_UserTblMgr) GetFromDistVip(distVip bool) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_vip = ?", distVip).Find(&results).Error

	return
}

// GetBatchFromDistVip 批量唯一主键查找 是否分销vip
func (obj *_UserTblMgr) GetBatchFromDistVip(distVips []bool) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_vip IN (?)", distVips).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 用户余额
func (obj *_UserTblMgr) GetFromBalance(balance int64) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("balance = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量唯一主键查找 用户余额
func (obj *_UserTblMgr) GetBatchFromBalance(balances []int64) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("balance IN (?)", balances).Find(&results).Error

	return
}

// GetFromPoints 通过points获取内容 用户积分
func (obj *_UserTblMgr) GetFromPoints(points int64) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("points = ?", points).Find(&results).Error

	return
}

// GetBatchFromPoints 批量唯一主键查找 用户积分
func (obj *_UserTblMgr) GetBatchFromPoints(pointss []int64) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("points IN (?)", pointss).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_UserTblMgr) GetFromCreatedBy(createdBy string) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_UserTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UserTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_UserTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_UserTblMgr) GetFromUpdatedBy(updatedBy string) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_UserTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_UserTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_UserTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserTblMgr) FetchByPrimaryKey(id int) (result UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOpenid primay or index 获取唯一内容
func (obj *_UserTblMgr) FetchUniqueByOpenid(userID string) (result UserTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}
