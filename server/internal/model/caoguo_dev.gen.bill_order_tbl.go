package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BillOrderTblMgr struct {
	*_BaseMgr
}

// BillOrderTblMgr open func
func BillOrderTblMgr(db *gorm.DB) *_BillOrderTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillOrderTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillOrderTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_order_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillOrderTblMgr) GetTableName() string {
	return "bill_order_tbl"
}

// Get 获取
func (obj *_BillOrderTblMgr) Get() (result BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BillOrderTblMgr) Gets() (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BillOrderTblMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBillID bill_id获取 账单id
func (obj *_BillOrderTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithGpid gpid获取 商品id
func (obj *_BillOrderTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithName name获取 商品名字
func (obj *_BillOrderTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 商品价格(分)
func (obj *_BillOrderTblMgr) WithPrice(price int64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithIcon icon获取 图标信息
func (obj *_BillOrderTblMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithUserID user_id获取 用户id
func (obj *_BillOrderTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithNumber number获取 数量
func (obj *_BillOrderTblMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithDistAmount dist_amount获取 分享收益
func (obj *_BillOrderTblMgr) WithDistAmount(distAmount int64) Option {
	return optionFunc(func(o *options) { o.query["dist_amount"] = distAmount })
}

// WithSkuList sku_list获取 属性列表
func (obj *_BillOrderTblMgr) WithSkuList(skuList string) Option {
	return optionFunc(func(o *options) { o.query["sku_list"] = skuList })
}

// WithSkuArrt sku_arrt获取 属性描述
func (obj *_BillOrderTblMgr) WithSkuArrt(skuArrt string) Option {
	return optionFunc(func(o *options) { o.query["sku_arrt"] = skuArrt })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillOrderTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BillOrderTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillOrderTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_BillOrderTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_BillOrderTblMgr) GetByOption(opts ...Option) (result BillOrderTbl, err error) {
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
func (obj *_BillOrderTblMgr) GetByOptions(opts ...Option) (results []*BillOrderTbl, err error) {
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
func (obj *_BillOrderTblMgr) GetFromID(id int64) (result BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_BillOrderTblMgr) GetBatchFromID(ids []int64) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillOrderTblMgr) GetFromBillID(billID string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_BillOrderTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_BillOrderTblMgr) GetFromGpid(gpid string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_BillOrderTblMgr) GetBatchFromGpid(gpids []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 商品名字
func (obj *_BillOrderTblMgr) GetFromName(name string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 商品名字
func (obj *_BillOrderTblMgr) GetBatchFromName(names []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价格(分)
func (obj *_BillOrderTblMgr) GetFromPrice(price int64) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 商品价格(分)
func (obj *_BillOrderTblMgr) GetBatchFromPrice(prices []int64) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标信息
func (obj *_BillOrderTblMgr) GetFromIcon(icon string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标信息
func (obj *_BillOrderTblMgr) GetBatchFromIcon(icons []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_BillOrderTblMgr) GetFromUserID(userID string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_BillOrderTblMgr) GetBatchFromUserID(userIDs []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 数量
func (obj *_BillOrderTblMgr) GetFromNumber(number int) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量唯一主键查找 数量
func (obj *_BillOrderTblMgr) GetBatchFromNumber(numbers []int) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number IN (?)", numbers).Find(&results).Error

	return
}

// GetFromDistAmount 通过dist_amount获取内容 分享收益
func (obj *_BillOrderTblMgr) GetFromDistAmount(distAmount int64) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_amount = ?", distAmount).Find(&results).Error

	return
}

// GetBatchFromDistAmount 批量唯一主键查找 分享收益
func (obj *_BillOrderTblMgr) GetBatchFromDistAmount(distAmounts []int64) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dist_amount IN (?)", distAmounts).Find(&results).Error

	return
}

// GetFromSkuList 通过sku_list获取内容 属性列表
func (obj *_BillOrderTblMgr) GetFromSkuList(skuList string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list = ?", skuList).Find(&results).Error

	return
}

// GetBatchFromSkuList 批量唯一主键查找 属性列表
func (obj *_BillOrderTblMgr) GetBatchFromSkuList(skuLists []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_list IN (?)", skuLists).Find(&results).Error

	return
}

// GetFromSkuArrt 通过sku_arrt获取内容 属性描述
func (obj *_BillOrderTblMgr) GetFromSkuArrt(skuArrt string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_arrt = ?", skuArrt).Find(&results).Error

	return
}

// GetBatchFromSkuArrt 批量唯一主键查找 属性描述
func (obj *_BillOrderTblMgr) GetBatchFromSkuArrt(skuArrts []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sku_arrt IN (?)", skuArrts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillOrderTblMgr) GetFromCreatedBy(createdBy string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_BillOrderTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BillOrderTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_BillOrderTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillOrderTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_BillOrderTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_BillOrderTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_BillOrderTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BillOrderTblMgr) FetchByPrimaryKey(id int64) (result BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUserGpid primay or index 获取唯一内容
func (obj *_BillOrderTblMgr) FetchUniqueIndexByUserGpid(billID string, gpid string, userID string, skuList string) (result BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ? AND gpid = ? AND user_id = ? AND sku_list = ?", billID, gpid, userID, skuList).Find(&result).Error

	return
}

// FetchIndexByBillID  获取多个内容
func (obj *_BillOrderTblMgr) FetchIndexByBillID(billID string) (results []*BillOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}
