package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BillTblMgr struct {
	*_BaseMgr
}

// BillTblMgr open func
func BillTblMgr(db *gorm.DB) *_BillTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillTblMgr) GetTableName() string {
	return "bill_tbl"
}

// Get 获取
func (obj *_BillTblMgr) Get() (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_BillTblMgr) Gets() (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_BillTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_BillTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_BillTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_BillTblMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithUserID user_id获取 用户id
func (obj *_BillTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithBillID bill_id获取 账单id
func (obj *_BillTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithPayPlatform pay_platform获取 支付类型(1:微信支付)
func (obj *_BillTblMgr) WithPayPlatform(payPlatform int) Option {
	return optionFunc(func(o *options) { o.query["pay_platform"] = payPlatform })
}

// WithPayAmount pay_amount获取 支付金额
func (obj *_BillTblMgr) WithPayAmount(payAmount int64) Option {
	return optionFunc(func(o *options) { o.query["pay_amount"] = payAmount })
}

// WithStatus status获取 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func (obj *_BillTblMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDesc desc获取 订单备注
func (obj *_BillTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithDeletedBy deleted_by获取 删除者
func (obj *_BillTblMgr) WithDeletedBy(deletedBy string) Option {
	return optionFunc(func(o *options) { o.query["deleted_by"] = deletedBy })
}

// GetByOption 功能选项模式获取
func (obj *_BillTblMgr) GetByOption(opts ...Option) (result BillTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BillTblMgr) GetByOptions(opts ...Option) (results []*BillTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_BillTblMgr) GetFromID(id int64) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找 Primary key
func (obj *_BillTblMgr) GetBatchFromID(ids []int64) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_BillTblMgr) GetFromCreatedAt(createdAt time.Time) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 created time
func (obj *_BillTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_BillTblMgr) GetFromUpdatedAt(updatedAt time.Time) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 updated at
func (obj *_BillTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_BillTblMgr) GetFromDeletedAt(deletedAt time.Time) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找 deleted time
func (obj *_BillTblMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_BillTblMgr) GetFromUserID(userID string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_BillTblMgr) GetBatchFromUserID(userIDs []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillTblMgr) GetFromBillID(billID string) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_BillTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPayPlatform 通过pay_platform获取内容 支付类型(1:微信支付)
func (obj *_BillTblMgr) GetFromPayPlatform(payPlatform int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pay_platform = ?", payPlatform).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPayPlatform 批量唯一主键查找 支付类型(1:微信支付)
func (obj *_BillTblMgr) GetBatchFromPayPlatform(payPlatforms []int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pay_platform IN (?)", payPlatforms).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPayAmount 通过pay_amount获取内容 支付金额
func (obj *_BillTblMgr) GetFromPayAmount(payAmount int64) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pay_amount = ?", payAmount).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPayAmount 批量唯一主键查找 支付金额
func (obj *_BillTblMgr) GetBatchFromPayAmount(payAmounts []int64) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pay_amount IN (?)", payAmounts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func (obj *_BillTblMgr) GetFromStatus(status int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量唯一主键查找 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func (obj *_BillTblMgr) GetBatchFromStatus(statuss []int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDesc 通过desc获取内容 订单备注
func (obj *_BillTblMgr) GetFromDesc(desc string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDesc 批量唯一主键查找 订单备注
func (obj *_BillTblMgr) GetBatchFromDesc(descs []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillTblMgr) GetFromCreatedBy(createdBy string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_BillTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_BillTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeletedBy 通过deleted_by获取内容 删除者
func (obj *_BillTblMgr) GetFromDeletedBy(deletedBy string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by = ?", deletedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeletedBy 批量唯一主键查找 删除者
func (obj *_BillTblMgr) GetBatchFromDeletedBy(deletedBys []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by IN (?)", deletedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillAddressTbl).Error; err != nil { // 派送信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", results[i].BillID).Find(&results[i].BillOrderTblList).Error; err != nil { // 账单订单列表
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
func (obj *_BillTblMgr) FetchByPrimaryKey(id int64) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByBillID primay or index 获取唯一内容
func (obj *_BillTblMgr) FetchUniqueByBillID(billID string) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("bill_address_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillAddressTbl).Error; err != nil { // 派送信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("bill_order_tbl").Where("bill_id = ?", result.BillID).Find(&result.BillOrderTblList).Error; err != nil { // 账单订单列表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
