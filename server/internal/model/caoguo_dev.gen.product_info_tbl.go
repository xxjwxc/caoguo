package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductInfoTblMgr struct {
	*_BaseMgr
}

// ProductInfoTblMgr open func
func ProductInfoTblMgr(db *gorm.DB) *_ProductInfoTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProductInfoTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductInfoTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_info_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductInfoTblMgr) GetTableName() string {
	return "product_info_tbl"
}

// Get 获取
func (obj *_ProductInfoTblMgr) Get() (result ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductInfoTblMgr) Gets() (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 自增
func (obj *_ProductInfoTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGpid gpid获取 商品id
func (obj *_ProductInfoTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithQty qty获取 已购买数量(销量)
func (obj *_ProductInfoTblMgr) WithQty(qty int64) Option {
	return optionFunc(func(o *options) { o.query["qty"] = qty })
}

// WithImg img获取 轮播图
func (obj *_ProductInfoTblMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["img"] = img })
}

// WithVideo video获取 视频信息
func (obj *_ProductInfoTblMgr) WithVideo(video string) Option {
	return optionFunc(func(o *options) { o.query["video"] = video })
}

// WithIcon icon获取 商品图标
func (obj *_ProductInfoTblMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithService service获取 商品服务
func (obj *_ProductInfoTblMgr) WithService(service string) Option {
	return optionFunc(func(o *options) { o.query["service"] = service })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProductInfoTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ProductInfoTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProductInfoTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ProductInfoTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ProductInfoTblMgr) GetByOption(opts ...Option) (result ProductInfoTbl, err error) {
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
func (obj *_ProductInfoTblMgr) GetByOptions(opts ...Option) (results []*ProductInfoTbl, err error) {
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

// GetFromID 通过id获取内容 自增
func (obj *_ProductInfoTblMgr) GetFromID(id int64) (result ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 自增
func (obj *_ProductInfoTblMgr) GetBatchFromID(ids []int64) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_ProductInfoTblMgr) GetFromGpid(gpid string) (result ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&result).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_ProductInfoTblMgr) GetBatchFromGpid(gpids []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromQty 通过qty获取内容 已购买数量(销量)
func (obj *_ProductInfoTblMgr) GetFromQty(qty int64) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qty = ?", qty).Find(&results).Error

	return
}

// GetBatchFromQty 批量唯一主键查找 已购买数量(销量)
func (obj *_ProductInfoTblMgr) GetBatchFromQty(qtys []int64) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qty IN (?)", qtys).Find(&results).Error

	return
}

// GetFromImg 通过img获取内容 轮播图
func (obj *_ProductInfoTblMgr) GetFromImg(img string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("img = ?", img).Find(&results).Error

	return
}

// GetBatchFromImg 批量唯一主键查找 轮播图
func (obj *_ProductInfoTblMgr) GetBatchFromImg(imgs []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("img IN (?)", imgs).Find(&results).Error

	return
}

// GetFromVideo 通过video获取内容 视频信息
func (obj *_ProductInfoTblMgr) GetFromVideo(video string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("video = ?", video).Find(&results).Error

	return
}

// GetBatchFromVideo 批量唯一主键查找 视频信息
func (obj *_ProductInfoTblMgr) GetBatchFromVideo(videos []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("video IN (?)", videos).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 商品图标
func (obj *_ProductInfoTblMgr) GetFromIcon(icon string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 商品图标
func (obj *_ProductInfoTblMgr) GetBatchFromIcon(icons []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromService 通过service获取内容 商品服务
func (obj *_ProductInfoTblMgr) GetFromService(service string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("service = ?", service).Find(&results).Error

	return
}

// GetBatchFromService 批量唯一主键查找 商品服务
func (obj *_ProductInfoTblMgr) GetBatchFromService(services []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("service IN (?)", services).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProductInfoTblMgr) GetFromCreatedBy(createdBy string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProductInfoTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ProductInfoTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_ProductInfoTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProductInfoTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProductInfoTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ProductInfoTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_ProductInfoTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductInfoTblMgr) FetchByPrimaryKey(id int64) (result ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByGpid primay or index 获取唯一内容
func (obj *_ProductInfoTblMgr) FetchUniqueByGpid(gpid string) (result ProductInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&result).Error

	return
}
