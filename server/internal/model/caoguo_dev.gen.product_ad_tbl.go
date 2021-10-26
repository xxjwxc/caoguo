package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductAdTblMgr struct {
	*_BaseMgr
}

// ProductAdTblMgr open func
func ProductAdTblMgr(db *gorm.DB) *_ProductAdTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAdTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAdTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_ad_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductAdTblMgr) GetTableName() string {
	return "product_ad_tbl"
}

// Get 获取
func (obj *_ProductAdTblMgr) Get() (result ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductAdTblMgr) Gets() (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_ProductAdTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_ProductAdTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_ProductAdTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_ProductAdTblMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithURL url获取 跳转路径
func (obj *_ProductAdTblMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithImg img获取 卡片展示图片
func (obj *_ProductAdTblMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["img"] = img })
}

// WithSortID sort_id获取 排序(越大越前)
func (obj *_ProductAdTblMgr) WithSortID(sortID int) Option {
	return optionFunc(func(o *options) { o.query["sort_id"] = sortID })
}

// WithType type获取 类型(1:轮播图广告，2:类型广告，3:主销广告)
func (obj *_ProductAdTblMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithAttach attach获取 附加属性
func (obj *_ProductAdTblMgr) WithAttach(attach string) Option {
	return optionFunc(func(o *options) { o.query["attach"] = attach })
}

// WithVaild vaild获取 是否有效
func (obj *_ProductAdTblMgr) WithVaild(vaild bool) Option {
	return optionFunc(func(o *options) { o.query["vaild"] = vaild })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProductAdTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProductAdTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithDeletedBy deleted_by获取 删除者
func (obj *_ProductAdTblMgr) WithDeletedBy(deletedBy string) Option {
	return optionFunc(func(o *options) { o.query["deleted_by"] = deletedBy })
}

// GetByOption 功能选项模式获取
func (obj *_ProductAdTblMgr) GetByOption(opts ...Option) (result ProductAdTbl, err error) {
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
func (obj *_ProductAdTblMgr) GetByOptions(opts ...Option) (results []*ProductAdTbl, err error) {
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
func (obj *_ProductAdTblMgr) GetFromID(id int64) (result ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 Primary key
func (obj *_ProductAdTblMgr) GetBatchFromID(ids []int64) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_ProductAdTblMgr) GetFromCreatedAt(createdAt time.Time) (result ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 created time
func (obj *_ProductAdTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_ProductAdTblMgr) GetFromUpdatedAt(updatedAt time.Time) (result ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 updated at
func (obj *_ProductAdTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_ProductAdTblMgr) GetFromDeletedAt(deletedAt time.Time) (result ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找 deleted time
func (obj *_ProductAdTblMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 跳转路径
func (obj *_ProductAdTblMgr) GetFromURL(url string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找 跳转路径
func (obj *_ProductAdTblMgr) GetBatchFromURL(urls []string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromImg 通过img获取内容 卡片展示图片
func (obj *_ProductAdTblMgr) GetFromImg(img string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("img = ?", img).Find(&results).Error

	return
}

// GetBatchFromImg 批量唯一主键查找 卡片展示图片
func (obj *_ProductAdTblMgr) GetBatchFromImg(imgs []string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("img IN (?)", imgs).Find(&results).Error

	return
}

// GetFromSortID 通过sort_id获取内容 排序(越大越前)
func (obj *_ProductAdTblMgr) GetFromSortID(sortID int) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort_id = ?", sortID).Find(&results).Error

	return
}

// GetBatchFromSortID 批量唯一主键查找 排序(越大越前)
func (obj *_ProductAdTblMgr) GetBatchFromSortID(sortIDs []int) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort_id IN (?)", sortIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型(1:轮播图广告，2:类型广告，3:主销广告)
func (obj *_ProductAdTblMgr) GetFromType(_type int) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型(1:轮播图广告，2:类型广告，3:主销广告)
func (obj *_ProductAdTblMgr) GetBatchFromType(_types []int) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromAttach 通过attach获取内容 附加属性
func (obj *_ProductAdTblMgr) GetFromAttach(attach string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach = ?", attach).Find(&results).Error

	return
}

// GetBatchFromAttach 批量唯一主键查找 附加属性
func (obj *_ProductAdTblMgr) GetBatchFromAttach(attachs []string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach IN (?)", attachs).Find(&results).Error

	return
}

// GetFromVaild 通过vaild获取内容 是否有效
func (obj *_ProductAdTblMgr) GetFromVaild(vaild bool) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild = ?", vaild).Find(&results).Error

	return
}

// GetBatchFromVaild 批量唯一主键查找 是否有效
func (obj *_ProductAdTblMgr) GetBatchFromVaild(vailds []bool) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vaild IN (?)", vailds).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProductAdTblMgr) GetFromCreatedBy(createdBy string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProductAdTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProductAdTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProductAdTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromDeletedBy 通过deleted_by获取内容 删除者
func (obj *_ProductAdTblMgr) GetFromDeletedBy(deletedBy string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by = ?", deletedBy).Find(&results).Error

	return
}

// GetBatchFromDeletedBy 批量唯一主键查找 删除者
func (obj *_ProductAdTblMgr) GetBatchFromDeletedBy(deletedBys []string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted_by IN (?)", deletedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductAdTblMgr) FetchByPrimaryKey(id int64) (result ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByGpid  获取多个内容
func (obj *_ProductAdTblMgr) FetchIndexByGpid(url string) (results []*ProductAdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}
