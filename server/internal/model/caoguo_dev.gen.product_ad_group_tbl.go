package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductAdGroupTblMgr struct {
	*_BaseMgr
}

// ProductAdGroupTblMgr open func
func ProductAdGroupTblMgr(db *gorm.DB) *_ProductAdGroupTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAdGroupTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAdGroupTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_ad_group_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductAdGroupTblMgr) GetTableName() string {
	return "product_ad_group_tbl"
}

// Get 获取
func (obj *_ProductAdGroupTblMgr) Get() (result ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductAdGroupTblMgr) Gets() (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProductAdGroupTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMainGpid main_gpid获取 主商品id
//
func (obj *_ProductAdGroupTblMgr) WithMainGpid(mainGpid string) Option {
	return optionFunc(func(o *options) { o.query["main_gpid"] = mainGpid })
}

// WithSubGpid sub_gpid获取 附加商品id
func (obj *_ProductAdGroupTblMgr) WithSubGpid(subGpid string) Option {
	return optionFunc(func(o *options) { o.query["sub_gpid"] = subGpid })
}

// WithVaild vaild获取 是否有效
func (obj *_ProductAdGroupTblMgr) WithVaild(vaild bool) Option {
	return optionFunc(func(o *options) { o.query["vaild"] = vaild })
}

// WithSortID sort_id获取 排序(越大越前)
func (obj *_ProductAdGroupTblMgr) WithSortID(sortID int) Option {
	return optionFunc(func(o *options) { o.query["sort_id"] = sortID })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProductAdGroupTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ProductAdGroupTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProductAdGroupTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ProductAdGroupTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ProductAdGroupTblMgr) GetByOption(opts ...Option) (result ProductAdGroupTbl, err error) {
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
func (obj *_ProductAdGroupTblMgr) GetByOptions(opts ...Option) (results []*ProductAdGroupTbl, err error) {
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
func (obj *_ProductAdGroupTblMgr) GetFromID(id int) (result ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ProductAdGroupTblMgr) GetBatchFromID(ids []int) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMainGpid 通过main_gpid获取内容 主商品id
//
func (obj *_ProductAdGroupTblMgr) GetFromMainGpid(mainGpid string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main_gpid = ?", mainGpid).Find(&results).Error

	return
}

// GetBatchFromMainGpid 批量唯一主键查找 主商品id
//
func (obj *_ProductAdGroupTblMgr) GetBatchFromMainGpid(mainGpids []string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main_gpid IN (?)", mainGpids).Find(&results).Error

	return
}

// GetFromSubGpid 通过sub_gpid获取内容 附加商品id
func (obj *_ProductAdGroupTblMgr) GetFromSubGpid(subGpid string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sub_gpid = ?", subGpid).Find(&results).Error

	return
}

// GetBatchFromSubGpid 批量唯一主键查找 附加商品id
func (obj *_ProductAdGroupTblMgr) GetBatchFromSubGpid(subGpids []string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sub_gpid IN (?)", subGpids).Find(&results).Error

	return
}

// GetFromVaild 通过vaild获取内容 是否有效
func (obj *_ProductAdGroupTblMgr) GetFromVaild(vaild bool) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild = ?", vaild).Find(&results).Error

	return
}

// GetBatchFromVaild 批量唯一主键查找 是否有效
func (obj *_ProductAdGroupTblMgr) GetBatchFromVaild(vailds []bool) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild IN (?)", vailds).Find(&results).Error

	return
}

// GetFromSortID 通过sort_id获取内容 排序(越大越前)
func (obj *_ProductAdGroupTblMgr) GetFromSortID(sortID int) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort_id = ?", sortID).Find(&results).Error

	return
}

// GetBatchFromSortID 批量唯一主键查找 排序(越大越前)
func (obj *_ProductAdGroupTblMgr) GetBatchFromSortID(sortIDs []int) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort_id IN (?)", sortIDs).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProductAdGroupTblMgr) GetFromCreatedBy(createdBy string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProductAdGroupTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ProductAdGroupTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_ProductAdGroupTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProductAdGroupTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProductAdGroupTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ProductAdGroupTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_ProductAdGroupTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductAdGroupTblMgr) FetchByPrimaryKey(id int) (result ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByMainSubGpid primay or index 获取唯一内容
func (obj *_ProductAdGroupTblMgr) FetchUniqueIndexByMainSubGpid(mainGpid string, subGpid string) (result ProductAdGroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main_gpid = ? AND sub_gpid = ?", mainGpid, subGpid).Find(&result).Error

	return
}
