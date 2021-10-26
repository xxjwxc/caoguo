package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CouponTblMgr struct {
	*_BaseMgr
}

// CouponTblMgr open func
func CouponTblMgr(db *gorm.DB) *_CouponTblMgr {
	if db == nil {
		panic(fmt.Errorf("CouponTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CouponTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("coupon_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CouponTblMgr) GetTableName() string {
	return "coupon_tbl"
}

// Get 获取
func (obj *_CouponTblMgr) Get() (result CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CouponTblMgr) Gets() (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CouponTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupName group_name获取 分组名
func (obj *_CouponTblMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// WithTitle title获取 优惠券名字
func (obj *_CouponTblMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithValidity validity获取 有效期(天数)
func (obj *_CouponTblMgr) WithValidity(validity int) Option {
	return optionFunc(func(o *options) { o.query["validity"] = validity })
}

// WithGpid gpid获取 商品优惠券
func (obj *_CouponTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithDenom denom获取 面额
func (obj *_CouponTblMgr) WithDenom(denom int) Option {
	return optionFunc(func(o *options) { o.query["denom"] = denom })
}

// WithCouponType coupon_type获取 优惠券类型(1:全场，2:单个商品)
func (obj *_CouponTblMgr) WithCouponType(couponType int) Option {
	return optionFunc(func(o *options) { o.query["coupon_type"] = couponType })
}

// WithGreatMoney great_money获取 满多少可用
func (obj *_CouponTblMgr) WithGreatMoney(greatMoney int) Option {
	return optionFunc(func(o *options) { o.query["great_money"] = greatMoney })
}

// WithDescrible describle获取 优惠券描述
func (obj *_CouponTblMgr) WithDescrible(describle string) Option {
	return optionFunc(func(o *options) { o.query["describle"] = describle })
}

// WithVaild vaild获取 是否有效
func (obj *_CouponTblMgr) WithVaild(vaild bool) Option {
	return optionFunc(func(o *options) { o.query["vaild"] = vaild })
}

// WithCreatedBy created_by获取 创建者
func (obj *_CouponTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_CouponTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_CouponTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_CouponTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_CouponTblMgr) GetByOption(opts ...Option) (result CouponTbl, err error) {
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
func (obj *_CouponTblMgr) GetByOptions(opts ...Option) (results []*CouponTbl, err error) {
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
func (obj *_CouponTblMgr) GetFromID(id int64) (result CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CouponTblMgr) GetBatchFromID(ids []int64) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容 分组名
func (obj *_CouponTblMgr) GetFromGroupName(groupName string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name = ?", groupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量唯一主键查找 分组名
func (obj *_CouponTblMgr) GetBatchFromGroupName(groupNames []string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name IN (?)", groupNames).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 优惠券名字
func (obj *_CouponTblMgr) GetFromTitle(title string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 优惠券名字
func (obj *_CouponTblMgr) GetBatchFromTitle(titles []string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromValidity 通过validity获取内容 有效期(天数)
func (obj *_CouponTblMgr) GetFromValidity(validity int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validity = ?", validity).Find(&results).Error

	return
}

// GetBatchFromValidity 批量唯一主键查找 有效期(天数)
func (obj *_CouponTblMgr) GetBatchFromValidity(validitys []int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validity IN (?)", validitys).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品优惠券
func (obj *_CouponTblMgr) GetFromGpid(gpid string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品优惠券
func (obj *_CouponTblMgr) GetBatchFromGpid(gpids []string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromDenom 通过denom获取内容 面额
func (obj *_CouponTblMgr) GetFromDenom(denom int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("denom = ?", denom).Find(&results).Error

	return
}

// GetBatchFromDenom 批量唯一主键查找 面额
func (obj *_CouponTblMgr) GetBatchFromDenom(denoms []int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("denom IN (?)", denoms).Find(&results).Error

	return
}

// GetFromCouponType 通过coupon_type获取内容 优惠券类型(1:全场，2:单个商品)
func (obj *_CouponTblMgr) GetFromCouponType(couponType int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_type = ?", couponType).Find(&results).Error

	return
}

// GetBatchFromCouponType 批量唯一主键查找 优惠券类型(1:全场，2:单个商品)
func (obj *_CouponTblMgr) GetBatchFromCouponType(couponTypes []int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_type IN (?)", couponTypes).Find(&results).Error

	return
}

// GetFromGreatMoney 通过great_money获取内容 满多少可用
func (obj *_CouponTblMgr) GetFromGreatMoney(greatMoney int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("great_money = ?", greatMoney).Find(&results).Error

	return
}

// GetBatchFromGreatMoney 批量唯一主键查找 满多少可用
func (obj *_CouponTblMgr) GetBatchFromGreatMoney(greatMoneys []int) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("great_money IN (?)", greatMoneys).Find(&results).Error

	return
}

// GetFromDescrible 通过describle获取内容 优惠券描述
func (obj *_CouponTblMgr) GetFromDescrible(describle string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("describle = ?", describle).Find(&results).Error

	return
}

// GetBatchFromDescrible 批量唯一主键查找 优惠券描述
func (obj *_CouponTblMgr) GetBatchFromDescrible(describles []string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("describle IN (?)", describles).Find(&results).Error

	return
}

// GetFromVaild 通过vaild获取内容 是否有效
func (obj *_CouponTblMgr) GetFromVaild(vaild bool) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild = ?", vaild).Find(&results).Error

	return
}

// GetBatchFromVaild 批量唯一主键查找 是否有效
func (obj *_CouponTblMgr) GetBatchFromVaild(vailds []bool) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild IN (?)", vailds).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_CouponTblMgr) GetFromCreatedBy(createdBy string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_CouponTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_CouponTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_CouponTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_CouponTblMgr) GetFromUpdatedBy(updatedBy string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_CouponTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_CouponTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_CouponTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CouponTblMgr) FetchByPrimaryKey(id int64) (result CouponTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
