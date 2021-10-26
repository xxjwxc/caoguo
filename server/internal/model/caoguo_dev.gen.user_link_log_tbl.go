package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserLinkLogTblMgr struct {
	*_BaseMgr
}

// UserLinkLogTblMgr open func
func UserLinkLogTblMgr(db *gorm.DB) *_UserLinkLogTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserLinkLogTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserLinkLogTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_link_log_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserLinkLogTblMgr) GetTableName() string {
	return "user_link_log_tbl"
}

// Get 获取
func (obj *_UserLinkLogTblMgr) Get() (result UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserLinkLogTblMgr) Gets() (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserLinkLogTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOpenID open_id获取 微信用户唯一标识符
func (obj *_UserLinkLogTblMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithLinkOpenID link_open_id获取 关联的微信用户唯一标识符
func (obj *_UserLinkLogTblMgr) WithLinkOpenID(linkOpenID string) Option {
	return optionFunc(func(o *options) { o.query["link_open_id"] = linkOpenID })
}

// WithCreatedBy created_by获取 创建者
func (obj *_UserLinkLogTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_UserLinkLogTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserLinkLogTblMgr) GetByOption(opts ...Option) (result UserLinkLogTbl, err error) {
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
func (obj *_UserLinkLogTblMgr) GetByOptions(opts ...Option) (results []*UserLinkLogTbl, err error) {
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
func (obj *_UserLinkLogTblMgr) GetFromID(id int) (result UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserLinkLogTblMgr) GetBatchFromID(ids []int) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 微信用户唯一标识符
func (obj *_UserLinkLogTblMgr) GetFromOpenID(openID string) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量唯一主键查找 微信用户唯一标识符
func (obj *_UserLinkLogTblMgr) GetBatchFromOpenID(openIDs []string) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromLinkOpenID 通过link_open_id获取内容 关联的微信用户唯一标识符
func (obj *_UserLinkLogTblMgr) GetFromLinkOpenID(linkOpenID string) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_open_id = ?", linkOpenID).Find(&results).Error

	return
}

// GetBatchFromLinkOpenID 批量唯一主键查找 关联的微信用户唯一标识符
func (obj *_UserLinkLogTblMgr) GetBatchFromLinkOpenID(linkOpenIDs []string) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_open_id IN (?)", linkOpenIDs).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_UserLinkLogTblMgr) GetFromCreatedBy(createdBy string) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_UserLinkLogTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_UserLinkLogTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_UserLinkLogTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserLinkLogTblMgr) FetchByPrimaryKey(id int) (result UserLinkLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
