package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductSkuTblMgr struct {
	*_BaseMgr
}

// ProductSkuTblMgr open func
func ProductSkuTblMgr(db *gorm.DB) *_ProductSkuTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductSkuTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductSkuTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_sku_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductSkuTblMgr) GetTableName() string {
	return "product_sku_tbl"
}

// Get 获取
func (obj *_ProductSkuTblMgr) Get() (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductSkuTblMgr) Gets() (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_ProductSkuTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_ProductSkuTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_ProductSkuTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_ProductSkuTblMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithGpid gpid获取 商品id
func (obj *_ProductSkuTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithTypeName type_name获取 标签类型
func (obj *_ProductSkuTblMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithTagName tag_name获取 标签名字
func (obj *_ProductSkuTblMgr) WithTagName(tagName string) Option {
	return optionFunc(func(o *options) { o.query["tag_name"] = tagName })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProductSkuTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProductSkuTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithDeletedBy deleted_by获取 删除者
func (obj *_ProductSkuTblMgr) WithDeletedBy(deletedBy string) Option {
	return optionFunc(func(o *options) { o.query["deleted_by"] = deletedBy })
}

// GetByOption 功能选项模式获取
func (obj *_ProductSkuTblMgr) GetByOption(opts ...Option) (result ProductSkuTbl, err error) {
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
func (obj *_ProductSkuTblMgr) GetByOptions(opts ...Option) (results []*ProductSkuTbl, err error) {
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

// GetFromID 通过id获取内容 Primary key
func (obj *_ProductSkuTblMgr) GetFromID(id int64) (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 Primary key
func (obj *_ProductSkuTblMgr) GetBatchFromID(ids []int64) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_ProductSkuTblMgr) GetFromCreatedAt(createdAt time.Time) (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 created time
func (obj *_ProductSkuTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_ProductSkuTblMgr) GetFromUpdatedAt(updatedAt time.Time) (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 updated at
func (obj *_ProductSkuTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_ProductSkuTblMgr) GetFromDeletedAt(deletedAt time.Time) (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找 deleted time
func (obj *_ProductSkuTblMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_ProductSkuTblMgr) GetFromGpid(gpid string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_ProductSkuTblMgr) GetBatchFromGpid(gpids []string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容 标签类型
func (obj *_ProductSkuTblMgr) GetFromTypeName(typeName string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_name = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量唯一主键查找 标签类型
func (obj *_ProductSkuTblMgr) GetBatchFromTypeName(typeNames []string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_name IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromTagName 通过tag_name获取内容 标签名字
func (obj *_ProductSkuTblMgr) GetFromTagName(tagName string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_name = ?", tagName).Find(&results).Error

	return
}

// GetBatchFromTagName 批量唯一主键查找 标签名字
func (obj *_ProductSkuTblMgr) GetBatchFromTagName(tagNames []string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_name IN (?)", tagNames).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProductSkuTblMgr) GetFromCreatedBy(createdBy string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProductSkuTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProductSkuTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProductSkuTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromDeletedBy 通过deleted_by获取内容 删除者
func (obj *_ProductSkuTblMgr) GetFromDeletedBy(deletedBy string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by = ?", deletedBy).Find(&results).Error

	return
}

// GetBatchFromDeletedBy 批量唯一主键查找 删除者
func (obj *_ProductSkuTblMgr) GetBatchFromDeletedBy(deletedBys []string) (results []*ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by IN (?)", deletedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductSkuTblMgr) FetchByPrimaryKey(id int64) (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByGpidTag primay or index 获取唯一内容
func (obj *_ProductSkuTblMgr) FetchUniqueIndexByGpidTag(gpid string, tagName string) (result ProductSkuTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ? AND tag_name = ?", gpid, tagName).Find(&result).Error

	return
}
