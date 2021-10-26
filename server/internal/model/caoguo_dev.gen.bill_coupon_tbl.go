package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BillCouponTblMgr struct {
	*_BaseMgr
}

// BillCouponTblMgr open func
func BillCouponTblMgr(db *gorm.DB) *_BillCouponTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillCouponTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillCouponTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_coupon_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillCouponTblMgr) GetTableName() string {
	return "bill_coupon_tbl"
}

// Get 获取
func (obj *_BillCouponTblMgr) Get() (result BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BillCouponTblMgr) Gets() (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BillCouponTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBillID bill_id获取 账单id
func (obj *_BillCouponTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithCouponID coupon_id获取 我领取的优惠券id
func (obj *_BillCouponTblMgr) WithCouponID(couponID int64) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithTitle title获取 优惠券名字
func (obj *_BillCouponTblMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDenom denom获取 面额
func (obj *_BillCouponTblMgr) WithDenom(denom int) Option {
	return optionFunc(func(o *options) { o.query["denom"] = denom })
}

// WithCouponType coupon_type获取 优惠券类型(1:全场，2:单个商品)
func (obj *_BillCouponTblMgr) WithCouponType(couponType int) Option {
	return optionFunc(func(o *options) { o.query["coupon_type"] = couponType })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillCouponTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BillCouponTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillCouponTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_BillCouponTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_BillCouponTblMgr) GetByOption(opts ...Option) (result BillCouponTbl, err error) {
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
func (obj *_BillCouponTblMgr) GetByOptions(opts ...Option) (results []*BillCouponTbl, err error) {
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
func (obj *_BillCouponTblMgr) GetFromID(id int64) (result BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_BillCouponTblMgr) GetBatchFromID(ids []int64) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillCouponTblMgr) GetFromBillID(billID string) (result BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&result).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_BillCouponTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromCouponID 通过coupon_id获取内容 我领取的优惠券id
func (obj *_BillCouponTblMgr) GetFromCouponID(couponID int64) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_id = ?", couponID).Find(&results).Error

	return
}

// GetBatchFromCouponID 批量唯一主键查找 我领取的优惠券id
func (obj *_BillCouponTblMgr) GetBatchFromCouponID(couponIDs []int64) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_id IN (?)", couponIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 优惠券名字
func (obj *_BillCouponTblMgr) GetFromTitle(title string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 优惠券名字
func (obj *_BillCouponTblMgr) GetBatchFromTitle(titles []string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromDenom 通过denom获取内容 面额
func (obj *_BillCouponTblMgr) GetFromDenom(denom int) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("denom = ?", denom).Find(&results).Error

	return
}

// GetBatchFromDenom 批量唯一主键查找 面额
func (obj *_BillCouponTblMgr) GetBatchFromDenom(denoms []int) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("denom IN (?)", denoms).Find(&results).Error

	return
}

// GetFromCouponType 通过coupon_type获取内容 优惠券类型(1:全场，2:单个商品)
func (obj *_BillCouponTblMgr) GetFromCouponType(couponType int) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_type = ?", couponType).Find(&results).Error

	return
}

// GetBatchFromCouponType 批量唯一主键查找 优惠券类型(1:全场，2:单个商品)
func (obj *_BillCouponTblMgr) GetBatchFromCouponType(couponTypes []int) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_type IN (?)", couponTypes).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillCouponTblMgr) GetFromCreatedBy(createdBy string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_BillCouponTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BillCouponTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_BillCouponTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillCouponTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_BillCouponTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_BillCouponTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_BillCouponTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BillCouponTblMgr) FetchByPrimaryKey(id int64) (result BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByBillID primay or index 获取唯一内容
func (obj *_BillCouponTblMgr) FetchUniqueByBillID(billID string) (result BillCouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&result).Error

	return
}
