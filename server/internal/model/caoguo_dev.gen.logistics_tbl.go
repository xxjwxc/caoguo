package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LogisticsTblMgr struct {
	*_BaseMgr
}

// LogisticsTblMgr open func
func LogisticsTblMgr(db *gorm.DB) *_LogisticsTblMgr {
	if db == nil {
		panic(fmt.Errorf("LogisticsTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LogisticsTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("logistics_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LogisticsTblMgr) GetTableName() string {
	return "logistics_tbl"
}

// Get 获取
func (obj *_LogisticsTblMgr) Get() (result LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LogisticsTblMgr) Gets() (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LogisticsTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShipmentID shipment_id获取 快递单号
func (obj *_LogisticsTblMgr) WithShipmentID(shipmentID string) Option {
	return optionFunc(func(o *options) { o.query["shipment_id"] = shipmentID })
}

// WithTimes times获取 时间
func (obj *_LogisticsTblMgr) WithTimes(times time.Time) Option {
	return optionFunc(func(o *options) { o.query["times"] = times })
}

// WithContext context获取 描述信息
func (obj *_LogisticsTblMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["context"] = context })
}

// WithCreatedBy created_by获取 创建者
func (obj *_LogisticsTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_LogisticsTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_LogisticsTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_LogisticsTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_LogisticsTblMgr) GetByOption(opts ...Option) (result LogisticsTbl, err error) {
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
func (obj *_LogisticsTblMgr) GetByOptions(opts ...Option) (results []*LogisticsTbl, err error) {
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
func (obj *_LogisticsTblMgr) GetFromID(id int) (result LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_LogisticsTblMgr) GetBatchFromID(ids []int) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromShipmentID 通过shipment_id获取内容 快递单号
func (obj *_LogisticsTblMgr) GetFromShipmentID(shipmentID string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_id = ?", shipmentID).Find(&results).Error

	return
}

// GetBatchFromShipmentID 批量唯一主键查找 快递单号
func (obj *_LogisticsTblMgr) GetBatchFromShipmentID(shipmentIDs []string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_id IN (?)", shipmentIDs).Find(&results).Error

	return
}

// GetFromTimes 通过times获取内容 时间
func (obj *_LogisticsTblMgr) GetFromTimes(times time.Time) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("times = ?", times).Find(&results).Error

	return
}

// GetBatchFromTimes 批量唯一主键查找 时间
func (obj *_LogisticsTblMgr) GetBatchFromTimes(timess []time.Time) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("times IN (?)", timess).Find(&results).Error

	return
}

// GetFromContext 通过context获取内容 描述信息
func (obj *_LogisticsTblMgr) GetFromContext(context string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("context = ?", context).Find(&results).Error

	return
}

// GetBatchFromContext 批量唯一主键查找 描述信息
func (obj *_LogisticsTblMgr) GetBatchFromContext(contexts []string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("context IN (?)", contexts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_LogisticsTblMgr) GetFromCreatedBy(createdBy string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_LogisticsTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_LogisticsTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_LogisticsTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_LogisticsTblMgr) GetFromUpdatedBy(updatedBy string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_LogisticsTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_LogisticsTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_LogisticsTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LogisticsTblMgr) FetchByPrimaryKey(id int) (result LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByShipmentID  获取多个内容
func (obj *_LogisticsTblMgr) FetchIndexByShipmentID(shipmentID string) (results []*LogisticsTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_id = ?", shipmentID).Find(&results).Error

	return
}
