package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EmailNotifyTblMgr struct {
	*_BaseMgr
}

// EmailNotifyTblMgr open func
func EmailNotifyTblMgr(db *gorm.DB) *_EmailNotifyTblMgr {
	if db == nil {
		panic(fmt.Errorf("EmailNotifyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmailNotifyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("email_notify_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EmailNotifyTblMgr) GetTableName() string {
	return "email_notify_tbl"
}

// Get 获取
func (obj *_EmailNotifyTblMgr) Get() (result EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EmailNotifyTblMgr) Gets() (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EmailNotifyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户名
func (obj *_EmailNotifyTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithEmail email获取 邮箱号
func (obj *_EmailNotifyTblMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithCreatedBy created_by获取 创建者
func (obj *_EmailNotifyTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_EmailNotifyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_EmailNotifyTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_EmailNotifyTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_EmailNotifyTblMgr) GetByOption(opts ...Option) (result EmailNotifyTbl, err error) {
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
func (obj *_EmailNotifyTblMgr) GetByOptions(opts ...Option) (results []*EmailNotifyTbl, err error) {
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
func (obj *_EmailNotifyTblMgr) GetFromID(id int) (result EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_EmailNotifyTblMgr) GetBatchFromID(ids []int) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户名
func (obj *_EmailNotifyTblMgr) GetFromUserID(userID string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户名
func (obj *_EmailNotifyTblMgr) GetBatchFromUserID(userIDs []string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱号
func (obj *_EmailNotifyTblMgr) GetFromEmail(email string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱号
func (obj *_EmailNotifyTblMgr) GetBatchFromEmail(emails []string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_EmailNotifyTblMgr) GetFromCreatedBy(createdBy string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_EmailNotifyTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_EmailNotifyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_EmailNotifyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_EmailNotifyTblMgr) GetFromUpdatedBy(updatedBy string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_EmailNotifyTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_EmailNotifyTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_EmailNotifyTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EmailNotifyTblMgr) FetchByPrimaryKey(id int) (result EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByShipmentID  获取多个内容
func (obj *_EmailNotifyTblMgr) FetchIndexByShipmentID(userID string) (results []*EmailNotifyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
