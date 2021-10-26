package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _DistUserinfoTblMgr struct {
	*_BaseMgr
}

// DistUserinfoTblMgr open func
func DistUserinfoTblMgr(db *gorm.DB) *_DistUserinfoTblMgr {
	if db == nil {
		panic(fmt.Errorf("DistUserinfoTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DistUserinfoTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dist_userinfo_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DistUserinfoTblMgr) GetTableName() string {
	return "dist_userinfo_tbl"
}

// Get 获取
func (obj *_DistUserinfoTblMgr) Get() (result DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DistUserinfoTblMgr) Gets() (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DistUserinfoTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 子节点
func (obj *_DistUserinfoTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithParntID parnt_id获取 父节点
func (obj *_DistUserinfoTblMgr) WithParntID(parntID string) Option {
	return optionFunc(func(o *options) { o.query["parnt_id"] = parntID })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_DistUserinfoTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_DistUserinfoTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_DistUserinfoTblMgr) GetByOption(opts ...Option) (result DistUserinfoTbl, err error) {
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
func (obj *_DistUserinfoTblMgr) GetByOptions(opts ...Option) (results []*DistUserinfoTbl, err error) {
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
func (obj *_DistUserinfoTblMgr) GetFromID(id int) (result DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_DistUserinfoTblMgr) GetBatchFromID(ids []int) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 子节点
func (obj *_DistUserinfoTblMgr) GetFromUserID(userID string) (result DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 子节点
func (obj *_DistUserinfoTblMgr) GetBatchFromUserID(userIDs []string) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromParntID 通过parnt_id获取内容 父节点
func (obj *_DistUserinfoTblMgr) GetFromParntID(parntID string) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parnt_id = ?", parntID).Find(&results).Error

	return
}

// GetBatchFromParntID 批量唯一主键查找 父节点
func (obj *_DistUserinfoTblMgr) GetBatchFromParntID(parntIDs []string) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parnt_id IN (?)", parntIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_DistUserinfoTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_DistUserinfoTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_DistUserinfoTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_DistUserinfoTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DistUserinfoTblMgr) FetchByPrimaryKey(id int) (result DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUserid primay or index 获取唯一内容
func (obj *_DistUserinfoTblMgr) FetchUniqueByUserid(userID string) (result DistUserinfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}
