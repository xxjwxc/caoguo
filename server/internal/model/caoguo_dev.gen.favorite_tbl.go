package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FavoriteTblMgr struct {
	*_BaseMgr
}

// FavoriteTblMgr open func
func FavoriteTblMgr(db *gorm.DB) *_FavoriteTblMgr {
	if db == nil {
		panic(fmt.Errorf("FavoriteTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FavoriteTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("favorite_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FavoriteTblMgr) GetTableName() string {
	return "favorite_tbl"
}

// Get 获取
func (obj *_FavoriteTblMgr) Get() (result FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FavoriteTblMgr) Gets() (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FavoriteTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGpid gpid获取 商品id
func (obj *_FavoriteTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithUserID user_id获取 用户id
func (obj *_FavoriteTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserType user_type获取 用户类型(1:微信)
func (obj *_FavoriteTblMgr) WithUserType(userType int) Option {
	return optionFunc(func(o *options) { o.query["user_type"] = userType })
}

// WithCreatedBy created_by获取 创建者
func (obj *_FavoriteTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_FavoriteTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_FavoriteTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_FavoriteTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_FavoriteTblMgr) GetByOption(opts ...Option) (result FavoriteTbl, err error) {
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
func (obj *_FavoriteTblMgr) GetByOptions(opts ...Option) (results []*FavoriteTbl, err error) {
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
func (obj *_FavoriteTblMgr) GetFromID(id int) (result FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_FavoriteTblMgr) GetBatchFromID(ids []int) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_FavoriteTblMgr) GetFromGpid(gpid string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_FavoriteTblMgr) GetBatchFromGpid(gpids []string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_FavoriteTblMgr) GetFromUserID(userID string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_FavoriteTblMgr) GetBatchFromUserID(userIDs []string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserType 通过user_type获取内容 用户类型(1:微信)
func (obj *_FavoriteTblMgr) GetFromUserType(userType int) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type = ?", userType).Find(&results).Error

	return
}

// GetBatchFromUserType 批量唯一主键查找 用户类型(1:微信)
func (obj *_FavoriteTblMgr) GetBatchFromUserType(userTypes []int) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type IN (?)", userTypes).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_FavoriteTblMgr) GetFromCreatedBy(createdBy string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_FavoriteTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_FavoriteTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_FavoriteTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_FavoriteTblMgr) GetFromUpdatedBy(updatedBy string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_FavoriteTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_FavoriteTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_FavoriteTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_FavoriteTblMgr) FetchByPrimaryKey(id int) (result FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUserGpid primay or index 获取唯一内容
func (obj *_FavoriteTblMgr) FetchUniqueIndexByUserGpid(gpid string, userID string) (result FavoriteTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ? AND user_id = ?", gpid, userID).Find(&result).Error

	return
}
