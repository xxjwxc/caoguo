package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BillRefundTblMgr struct {
	*_BaseMgr
}

// BillRefundTblMgr open func
func BillRefundTblMgr(db *gorm.DB) *_BillRefundTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillRefundTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillRefundTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_refund_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillRefundTblMgr) GetTableName() string {
	return "bill_refund_tbl"
}

// Get 获取
func (obj *_BillRefundTblMgr) Get() (result BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BillRefundTblMgr) Gets() (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BillRefundTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_BillRefundTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithBillID bill_id获取 账单id
func (obj *_BillRefundTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithRefundID refund_id获取 退款账单号
func (obj *_BillRefundTblMgr) WithRefundID(refundID string) Option {
	return optionFunc(func(o *options) { o.query["refund_id"] = refundID })
}

// WithRefundFee refund_fee获取 退款金额
func (obj *_BillRefundTblMgr) WithRefundFee(refundFee int64) Option {
	return optionFunc(func(o *options) { o.query["refund_fee"] = refundFee })
}

// WithDesc desc获取 订单备注
func (obj *_BillRefundTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillRefundTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BillRefundTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillRefundTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_BillRefundTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_BillRefundTblMgr) GetByOption(opts ...Option) (result BillRefundTbl, err error) {
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
func (obj *_BillRefundTblMgr) GetByOptions(opts ...Option) (results []*BillRefundTbl, err error) {
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
func (obj *_BillRefundTblMgr) GetFromID(id int) (result BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_BillRefundTblMgr) GetBatchFromID(ids []int) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_BillRefundTblMgr) GetFromUserID(userID string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_BillRefundTblMgr) GetBatchFromUserID(userIDs []string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillRefundTblMgr) GetFromBillID(billID string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_BillRefundTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromRefundID 通过refund_id获取内容 退款账单号
func (obj *_BillRefundTblMgr) GetFromRefundID(refundID string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("refund_id = ?", refundID).Find(&results).Error

	return
}

// GetBatchFromRefundID 批量唯一主键查找 退款账单号
func (obj *_BillRefundTblMgr) GetBatchFromRefundID(refundIDs []string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("refund_id IN (?)", refundIDs).Find(&results).Error

	return
}

// GetFromRefundFee 通过refund_fee获取内容 退款金额
func (obj *_BillRefundTblMgr) GetFromRefundFee(refundFee int64) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("refund_fee = ?", refundFee).Find(&results).Error

	return
}

// GetBatchFromRefundFee 批量唯一主键查找 退款金额
func (obj *_BillRefundTblMgr) GetBatchFromRefundFee(refundFees []int64) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("refund_fee IN (?)", refundFees).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 订单备注
func (obj *_BillRefundTblMgr) GetFromDesc(desc string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 订单备注
func (obj *_BillRefundTblMgr) GetBatchFromDesc(descs []string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillRefundTblMgr) GetFromCreatedBy(createdBy string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_BillRefundTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BillRefundTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_BillRefundTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillRefundTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_BillRefundTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_BillRefundTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_BillRefundTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BillRefundTblMgr) FetchByPrimaryKey(id int) (result BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByBillID  获取多个内容
func (obj *_BillRefundTblMgr) FetchIndexByBillID(billID string) (results []*BillRefundTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}
