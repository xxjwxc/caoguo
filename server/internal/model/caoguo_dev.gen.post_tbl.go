package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PostTblMgr struct {
	*_BaseMgr
}

// PostTblMgr open func
func PostTblMgr(db *gorm.DB) *_PostTblMgr {
	if db == nil {
		panic(fmt.Errorf("PostTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PostTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("post_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PostTblMgr) GetTableName() string {
	return "post_tbl"
}

// Get 获取
func (obj *_PostTblMgr) Get() (result PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PostTblMgr) Gets() (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PostTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 快递名称
func (obj *_PostTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 快递id
func (obj *_PostTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithIcon icon获取 快递logo地址
func (obj *_PostTblMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithExpPhone exp_phone获取 快递电话
func (obj *_PostTblMgr) WithExpPhone(expPhone string) Option {
	return optionFunc(func(o *options) { o.query["exp_phone"] = expPhone })
}

// GetByOption 功能选项模式获取
func (obj *_PostTblMgr) GetByOption(opts ...Option) (result PostTbl, err error) {
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
func (obj *_PostTblMgr) GetByOptions(opts ...Option) (results []*PostTbl, err error) {
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
func (obj *_PostTblMgr) GetFromID(id int) (result PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_PostTblMgr) GetBatchFromID(ids []int) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 快递名称
func (obj *_PostTblMgr) GetFromName(name string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 快递名称
func (obj *_PostTblMgr) GetBatchFromName(names []string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 快递id
func (obj *_PostTblMgr) GetFromCode(code string) (result PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&result).Error

	return
}

// GetBatchFromCode 批量唯一主键查找 快递id
func (obj *_PostTblMgr) GetBatchFromCode(codes []string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 快递logo地址
func (obj *_PostTblMgr) GetFromIcon(icon string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 快递logo地址
func (obj *_PostTblMgr) GetBatchFromIcon(icons []string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromExpPhone 通过exp_phone获取内容 快递电话
func (obj *_PostTblMgr) GetFromExpPhone(expPhone string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exp_phone = ?", expPhone).Find(&results).Error

	return
}

// GetBatchFromExpPhone 批量唯一主键查找 快递电话
func (obj *_PostTblMgr) GetBatchFromExpPhone(expPhones []string) (results []*PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exp_phone IN (?)", expPhones).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PostTblMgr) FetchByPrimaryKey(id int) (result PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByKey primay or index 获取唯一内容
func (obj *_PostTblMgr) FetchUniqueByKey(code string) (result PostTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&result).Error

	return
}
