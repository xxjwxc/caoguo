package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserLinkTblMgr struct {
	*_BaseMgr
}

// UserLinkTblMgr open func
func UserLinkTblMgr(db *gorm.DB) *_UserLinkTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserLinkTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserLinkTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_link_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserLinkTblMgr) GetTableName() string {
	return "user_link_tbl"
}

// Get 获取
func (obj *_UserLinkTblMgr) Get() (result UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserLinkTblMgr) Gets() (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserLinkTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 被邀请人
func (obj *_UserLinkTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithParentUserID parent_user_id获取 邀请人
func (obj *_UserLinkTblMgr) WithParentUserID(parentUserID string) Option {
	return optionFunc(func(o *options) { o.query["parent_user_id"] = parentUserID })
}

// WithCreatedBy created_by获取 创建者
func (obj *_UserLinkTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UserLinkTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_UserLinkTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_UserLinkTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserLinkTblMgr) GetByOption(opts ...Option) (result UserLinkTbl, err error) {
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
func (obj *_UserLinkTblMgr) GetByOptions(opts ...Option) (results []*UserLinkTbl, err error) {
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
func (obj *_UserLinkTblMgr) GetFromID(id int) (result UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserLinkTblMgr) GetBatchFromID(ids []int) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 被邀请人
func (obj *_UserLinkTblMgr) GetFromUserID(userID string) (result UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 被邀请人
func (obj *_UserLinkTblMgr) GetBatchFromUserID(userIDs []string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromParentUserID 通过parent_user_id获取内容 邀请人
func (obj *_UserLinkTblMgr) GetFromParentUserID(parentUserID string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_user_id = ?", parentUserID).Find(&results).Error

	return
}

// GetBatchFromParentUserID 批量唯一主键查找 邀请人
func (obj *_UserLinkTblMgr) GetBatchFromParentUserID(parentUserIDs []string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_user_id IN (?)", parentUserIDs).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_UserLinkTblMgr) GetFromCreatedBy(createdBy string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_UserLinkTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UserLinkTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_UserLinkTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_UserLinkTblMgr) GetFromUpdatedBy(updatedBy string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_UserLinkTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_UserLinkTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_UserLinkTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserLinkTblMgr) FetchByPrimaryKey(id int) (result UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOpenid primay or index 获取唯一内容
func (obj *_UserLinkTblMgr) FetchUniqueByOpenid(userID string) (result UserLinkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}
