package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ShipmentTblMgr struct {
	*_BaseMgr
}

// ShipmentTblMgr open func
func ShipmentTblMgr(db *gorm.DB) *_ShipmentTblMgr {
	if db == nil {
		panic(fmt.Errorf("ShipmentTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ShipmentTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("shipment_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ShipmentTblMgr) GetTableName() string {
	return "shipment_tbl"
}

// Get 获取
func (obj *_ShipmentTblMgr) Get() (result ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("post_tbl").Where("code = ?", result.PostKey).Find(&result.PostTbl).Error; err != nil { // 快递公司信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ShipmentTblMgr) Gets() (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
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
func (obj *_ShipmentTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBillID bill_id获取 账单id
func (obj *_ShipmentTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithShipmentID shipment_id获取 快递单号
func (obj *_ShipmentTblMgr) WithShipmentID(shipmentID string) Option {
	return optionFunc(func(o *options) { o.query["shipment_id"] = shipmentID })
}

// WithPostKey post_key获取 快递id
func (obj *_ShipmentTblMgr) WithPostKey(postKey string) Option {
	return optionFunc(func(o *options) { o.query["post_key"] = postKey })
}

// WithStatus status获取 快递状态  2在途中 3已签收 4 问题件
func (obj *_ShipmentTblMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ShipmentTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ShipmentTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ShipmentTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ShipmentTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ShipmentTblMgr) GetByOption(opts ...Option) (result ShipmentTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("post_tbl").Where("code = ?", result.PostKey).Find(&result.PostTbl).Error; err != nil { // 快递公司信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ShipmentTblMgr) GetByOptions(opts ...Option) (results []*ShipmentTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
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
func (obj *_ShipmentTblMgr) GetFromID(id int) (result ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("post_tbl").Where("code = ?", result.PostKey).Find(&result.PostTbl).Error; err != nil { // 快递公司信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ShipmentTblMgr) GetBatchFromID(ids []int) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_ShipmentTblMgr) GetFromBillID(billID string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_ShipmentTblMgr) GetBatchFromBillID(billIDs []string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromShipmentID 通过shipment_id获取内容 快递单号
func (obj *_ShipmentTblMgr) GetFromShipmentID(shipmentID string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_id = ?", shipmentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromShipmentID 批量唯一主键查找 快递单号
func (obj *_ShipmentTblMgr) GetBatchFromShipmentID(shipmentIDs []string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipment_id IN (?)", shipmentIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPostKey 通过post_key获取内容 快递id
func (obj *_ShipmentTblMgr) GetFromPostKey(postKey string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_key = ?", postKey).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPostKey 批量唯一主键查找 快递id
func (obj *_ShipmentTblMgr) GetBatchFromPostKey(postKeys []string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_key IN (?)", postKeys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容 快递状态  2在途中 3已签收 4 问题件
func (obj *_ShipmentTblMgr) GetFromStatus(status int) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量唯一主键查找 快递状态  2在途中 3已签收 4 问题件
func (obj *_ShipmentTblMgr) GetBatchFromStatus(statuss []int) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ShipmentTblMgr) GetFromCreatedBy(createdBy string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ShipmentTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ShipmentTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_ShipmentTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ShipmentTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ShipmentTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ShipmentTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_ShipmentTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
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
func (obj *_ShipmentTblMgr) FetchByPrimaryKey(id int) (result ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("post_tbl").Where("code = ?", result.PostKey).Find(&result.PostTbl).Error; err != nil { // 快递公司信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByBillShipment primay or index 获取唯一内容
func (obj *_ShipmentTblMgr) FetchUniqueIndexByBillShipment(billID string, shipmentID string) (result ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ? AND shipment_id = ?", billID, shipmentID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("post_tbl").Where("code = ?", result.PostKey).Find(&result.PostTbl).Error; err != nil { // 快递公司信息
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByPostKey  获取多个内容
func (obj *_ShipmentTblMgr) FetchIndexByPostKey(postKey string) (results []*ShipmentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_key = ?", postKey).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("post_tbl").Where("code = ?", results[i].PostKey).Find(&results[i].PostTbl).Error; err != nil { // 快递公司信息
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
