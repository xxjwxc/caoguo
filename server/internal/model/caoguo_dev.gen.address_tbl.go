package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AddressTblMgr struct {
	*_BaseMgr
}

// AddressTblMgr open func
func AddressTblMgr(db *gorm.DB) *_AddressTblMgr {
	if db == nil {
		panic(fmt.Errorf("AddressTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AddressTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("address_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AddressTblMgr) GetTableName() string {
	return "address_tbl"
}

// Get 获取
func (obj *_AddressTblMgr) Get() (result AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AddressTblMgr) Gets() (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AddressTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_AddressTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithName name获取 用户名
func (obj *_AddressTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMobile mobile获取 [@gormt default:'123456';]手机号
func (obj *_AddressTblMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithAddressName address_name获取 地址名称
func (obj *_AddressTblMgr) WithAddressName(addressName string) Option {
	return optionFunc(func(o *options) { o.query["address_name"] = addressName })
}

// WithAddress address获取 详细地址
func (obj *_AddressTblMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithArea area获取 单元门牌号
func (obj *_AddressTblMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// WithDefault default获取 是否是默认值
func (obj *_AddressTblMgr) WithDefault(_default bool) Option {
	return optionFunc(func(o *options) { o.query["default"] = _default })
}

// WithCreatedBy created_by获取 创建者
func (obj *_AddressTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AddressTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_AddressTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_AddressTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_AddressTblMgr) GetByOption(opts ...Option) (result AddressTbl, err error) {
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
func (obj *_AddressTblMgr) GetByOptions(opts ...Option) (results []*AddressTbl, err error) {
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
func (obj *_AddressTblMgr) GetFromID(id int64) (result AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_AddressTblMgr) GetBatchFromID(ids []int64) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_AddressTblMgr) GetFromUserID(userID string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_AddressTblMgr) GetBatchFromUserID(userIDs []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 用户名
func (obj *_AddressTblMgr) GetFromName(name string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 用户名
func (obj *_AddressTblMgr) GetBatchFromName(names []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 [@gormt default:'123456';]手机号
func (obj *_AddressTblMgr) GetFromMobile(mobile string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 [@gormt default:'123456';]手机号
func (obj *_AddressTblMgr) GetBatchFromMobile(mobiles []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromAddressName 通过address_name获取内容 地址名称
func (obj *_AddressTblMgr) GetFromAddressName(addressName string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address_name = ?", addressName).Find(&results).Error

	return
}

// GetBatchFromAddressName 批量唯一主键查找 地址名称
func (obj *_AddressTblMgr) GetBatchFromAddressName(addressNames []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address_name IN (?)", addressNames).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 详细地址
func (obj *_AddressTblMgr) GetFromAddress(address string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量唯一主键查找 详细地址
func (obj *_AddressTblMgr) GetBatchFromAddress(addresss []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address IN (?)", addresss).Find(&results).Error

	return
}

// GetFromArea 通过area获取内容 单元门牌号
func (obj *_AddressTblMgr) GetFromArea(area string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("area = ?", area).Find(&results).Error

	return
}

// GetBatchFromArea 批量唯一主键查找 单元门牌号
func (obj *_AddressTblMgr) GetBatchFromArea(areas []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("area IN (?)", areas).Find(&results).Error

	return
}

// GetFromDefault 通过default获取内容 是否是默认值
func (obj *_AddressTblMgr) GetFromDefault(_default bool) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default = ?", _default).Find(&results).Error

	return
}

// GetBatchFromDefault 批量唯一主键查找 是否是默认值
func (obj *_AddressTblMgr) GetBatchFromDefault(_defaults []bool) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default IN (?)", _defaults).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_AddressTblMgr) GetFromCreatedBy(createdBy string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_AddressTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AddressTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_AddressTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_AddressTblMgr) GetFromUpdatedBy(updatedBy string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_AddressTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_AddressTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_AddressTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AddressTblMgr) FetchByPrimaryKey(id int64) (result AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_AddressTblMgr) FetchIndexByUserID(userID string) (results []*AddressTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
