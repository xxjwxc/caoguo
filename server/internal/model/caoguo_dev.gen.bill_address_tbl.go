package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BillAddressTblMgr struct {
	*_BaseMgr
}

// BillAddressTblMgr open func
func BillAddressTblMgr(db *gorm.DB) *_BillAddressTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillAddressTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillAddressTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_address_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillAddressTblMgr) GetTableName() string {
	return "bill_address_tbl"
}

// Get 获取
func (obj *_BillAddressTblMgr) Get() (result BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BillAddressTblMgr) Gets() (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_BillAddressTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAddrID addr_id获取 地址id
func (obj *_BillAddressTblMgr) WithAddrID(addrID int64) Option {
	return optionFunc(func(o *options) { o.query["addr_id"] = addrID })
}

// WithUserID user_id获取 用户id
func (obj *_BillAddressTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithBillID bill_id获取 账单id
func (obj *_BillAddressTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithName name获取 用户名
func (obj *_BillAddressTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMobile mobile获取 手机号
func (obj *_BillAddressTblMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithAddressName address_name获取 地址名称
func (obj *_BillAddressTblMgr) WithAddressName(addressName string) Option {
	return optionFunc(func(o *options) { o.query["address_name"] = addressName })
}

// WithAddress address获取 详细地址
func (obj *_BillAddressTblMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithArea area获取 单元门牌号
func (obj *_BillAddressTblMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillAddressTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BillAddressTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillAddressTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_BillAddressTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_BillAddressTblMgr) GetByOption(opts ...Option) (result BillAddressTbl, err error) {
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
func (obj *_BillAddressTblMgr) GetByOptions(opts ...Option) (results []*BillAddressTbl, err error) {
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

// GetFromID 通过id获取内容 主键
func (obj *_BillAddressTblMgr) GetFromID(id int64) (result BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_BillAddressTblMgr) GetBatchFromID(ids []int64) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAddrID 通过addr_id获取内容 地址id
func (obj *_BillAddressTblMgr) GetFromAddrID(addrID int64) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("addr_id = ?", addrID).Find(&results).Error

	return
}

// GetBatchFromAddrID 批量唯一主键查找 地址id
func (obj *_BillAddressTblMgr) GetBatchFromAddrID(addrIDs []int64) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("addr_id IN (?)", addrIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_BillAddressTblMgr) GetFromUserID(userID string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_BillAddressTblMgr) GetBatchFromUserID(userIDs []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillAddressTblMgr) GetFromBillID(billID string) (result BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&result).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_BillAddressTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 用户名
func (obj *_BillAddressTblMgr) GetFromName(name string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 用户名
func (obj *_BillAddressTblMgr) GetBatchFromName(names []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *_BillAddressTblMgr) GetFromMobile(mobile string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 手机号
func (obj *_BillAddressTblMgr) GetBatchFromMobile(mobiles []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromAddressName 通过address_name获取内容 地址名称
func (obj *_BillAddressTblMgr) GetFromAddressName(addressName string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address_name = ?", addressName).Find(&results).Error

	return
}

// GetBatchFromAddressName 批量唯一主键查找 地址名称
func (obj *_BillAddressTblMgr) GetBatchFromAddressName(addressNames []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address_name IN (?)", addressNames).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 详细地址
func (obj *_BillAddressTblMgr) GetFromAddress(address string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量唯一主键查找 详细地址
func (obj *_BillAddressTblMgr) GetBatchFromAddress(addresss []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address IN (?)", addresss).Find(&results).Error

	return
}

// GetFromArea 通过area获取内容 单元门牌号
func (obj *_BillAddressTblMgr) GetFromArea(area string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("area = ?", area).Find(&results).Error

	return
}

// GetBatchFromArea 批量唯一主键查找 单元门牌号
func (obj *_BillAddressTblMgr) GetBatchFromArea(areas []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("area IN (?)", areas).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillAddressTblMgr) GetFromCreatedBy(createdBy string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_BillAddressTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BillAddressTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_BillAddressTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillAddressTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_BillAddressTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_BillAddressTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_BillAddressTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BillAddressTblMgr) FetchByPrimaryKey(id int64) (result BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByBillID primay or index 获取唯一内容
func (obj *_BillAddressTblMgr) FetchUniqueByBillID(billID string) (result BillAddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&result).Error

	return
}
