package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CouponMyTblMgr struct {
	*_BaseMgr
}

// CouponMyTblMgr open func
func CouponMyTblMgr(db *gorm.DB) *_CouponMyTblMgr {
	if db == nil {
		panic(fmt.Errorf("CouponMyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CouponMyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("coupon_my_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CouponMyTblMgr) GetTableName() string {
	return "coupon_my_tbl"
}

// Get 获取
func (obj *_CouponMyTblMgr) Get() (result CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("coupon_tbl").Where("id = ?", result.CouponID).Find(&result.CouponTbl).Error; err != nil { // 优惠券列表池
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_CouponMyTblMgr) Gets() (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CouponMyTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_CouponMyTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCouponID coupon_id获取 优惠券id
func (obj *_CouponMyTblMgr) WithCouponID(couponID int64) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithExpiresTime expires_time获取 过期时间
func (obj *_CouponMyTblMgr) WithExpiresTime(expiresTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires_time"] = expiresTime })
}

// WithStatus status获取 当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)
func (obj *_CouponMyTblMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedBy created_by获取 创建者
func (obj *_CouponMyTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_CouponMyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_CouponMyTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_CouponMyTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_CouponMyTblMgr) GetByOption(opts ...Option) (result CouponMyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("coupon_tbl").Where("id = ?", result.CouponID).Find(&result.CouponTbl).Error; err != nil { // 优惠券列表池
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CouponMyTblMgr) GetByOptions(opts ...Option) (results []*CouponMyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_CouponMyTblMgr) GetFromID(id int64) (result CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("coupon_tbl").Where("id = ?", result.CouponID).Find(&result.CouponTbl).Error; err != nil { // 优惠券列表池
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CouponMyTblMgr) GetBatchFromID(ids []int64) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_CouponMyTblMgr) GetFromUserID(userID string) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_CouponMyTblMgr) GetBatchFromUserID(userIDs []string) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCouponID 通过coupon_id获取内容 优惠券id
func (obj *_CouponMyTblMgr) GetFromCouponID(couponID int64) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_id = ?", couponID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCouponID 批量唯一主键查找 优惠券id
func (obj *_CouponMyTblMgr) GetBatchFromCouponID(couponIDs []int64) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_id IN (?)", couponIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromExpiresTime 通过expires_time获取内容 过期时间
func (obj *_CouponMyTblMgr) GetFromExpiresTime(expiresTime time.Time) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expires_time = ?", expiresTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromExpiresTime 批量唯一主键查找 过期时间
func (obj *_CouponMyTblMgr) GetBatchFromExpiresTime(expiresTimes []time.Time) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expires_time IN (?)", expiresTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容 当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)
func (obj *_CouponMyTblMgr) GetFromStatus(status int) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量唯一主键查找 当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)
func (obj *_CouponMyTblMgr) GetBatchFromStatus(statuss []int) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_CouponMyTblMgr) GetFromCreatedBy(createdBy string) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_CouponMyTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_CouponMyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_CouponMyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_CouponMyTblMgr) GetFromUpdatedBy(updatedBy string) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_CouponMyTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_CouponMyTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_CouponMyTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CouponMyTblMgr) FetchByPrimaryKey(id int64) (result CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("coupon_tbl").Where("id = ?", result.CouponID).Find(&result.CouponTbl).Error; err != nil { // 优惠券列表池
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByCouponID  获取多个内容
func (obj *_CouponMyTblMgr) FetchIndexByCouponID(couponID int64) (results []*CouponMyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("coupon_id = ?", couponID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("coupon_tbl").Where("id = ?", results[i].CouponID).Find(&results[i].CouponTbl).Error; err != nil { // 优惠券列表池
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
