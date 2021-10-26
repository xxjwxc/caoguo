package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProductTypeTblMgr struct {
	*_BaseMgr
}

// ProductTypeTblMgr open func
func ProductTypeTblMgr(db *gorm.DB) *_ProductTypeTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductTypeTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductTypeTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_type_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductTypeTblMgr) GetTableName() string {
	return "product_type_tbl"
}

// Get 获取
func (obj *_ProductTypeTblMgr) Get() (result ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductTypeTblMgr) Gets() (results []*ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProductTypeTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGtid gtid获取 产品类型id
func (obj *_ProductTypeTblMgr) WithGtid(gtid string) Option {
	return optionFunc(func(o *options) { o.query["gtid"] = gtid })
}

// GetByOption 功能选项模式获取
func (obj *_ProductTypeTblMgr) GetByOption(opts ...Option) (result ProductTypeTbl, err error) {
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
func (obj *_ProductTypeTblMgr) GetByOptions(opts ...Option) (results []*ProductTypeTbl, err error) {
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
func (obj *_ProductTypeTblMgr) GetFromID(id int) (result ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ProductTypeTblMgr) GetBatchFromID(ids []int) (results []*ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGtid 通过gtid获取内容 产品类型id
func (obj *_ProductTypeTblMgr) GetFromGtid(gtid string) (result ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gtid = ?", gtid).Find(&result).Error

	return
}

// GetBatchFromGtid 批量唯一主键查找 产品类型id
func (obj *_ProductTypeTblMgr) GetBatchFromGtid(gtids []string) (results []*ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gtid IN (?)", gtids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductTypeTblMgr) FetchByPrimaryKey(id int) (result ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByGtid primay or index 获取唯一内容
func (obj *_ProductTypeTblMgr) FetchUniqueByGtid(gtid string) (result ProductTypeTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gtid = ?", gtid).Find(&result).Error

	return
}
