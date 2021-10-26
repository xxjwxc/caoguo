package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _DistOrderTblMgr struct {
	*_BaseMgr
}

// DistOrderTblMgr open func
func DistOrderTblMgr(db *gorm.DB) *_DistOrderTblMgr {
	if db == nil {
		panic(fmt.Errorf("DistOrderTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DistOrderTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dist_order_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DistOrderTblMgr) GetTableName() string {
	return "dist_order_tbl"
}

// Get 获取
func (obj *_DistOrderTblMgr) Get() (result DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DistOrderTblMgr) Gets() (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DistOrderTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBillID bill_id获取 账单id
func (obj *_DistOrderTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithUserID user_id获取 分销给的账号
func (obj *_DistOrderTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithGpid gpid获取 商品id
func (obj *_DistOrderTblMgr) WithGpid(gpid string) Option {
	return optionFunc(func(o *options) { o.query["gpid"] = gpid })
}

// WithName name获取 商品名字
func (obj *_DistOrderTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 分销金额
func (obj *_DistOrderTblMgr) WithPrice(price int64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithLevel level获取 分销等级
func (obj *_DistOrderTblMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithTotalPrice total_price获取 总分销金额
func (obj *_DistOrderTblMgr) WithTotalPrice(totalPrice int64) Option {
	return optionFunc(func(o *options) { o.query["total_price"] = totalPrice })
}

// WithStatus status获取 状态(-1:失效,1:待确认,2:可提现，3:已提现)
func (obj *_DistOrderTblMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithNumber number获取 商品数量
func (obj *_DistOrderTblMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_DistOrderTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_DistOrderTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_DistOrderTblMgr) GetByOption(opts ...Option) (result DistOrderTbl, err error) {
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
func (obj *_DistOrderTblMgr) GetByOptions(opts ...Option) (results []*DistOrderTbl, err error) {
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
func (obj *_DistOrderTblMgr) GetFromID(id int) (result DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_DistOrderTblMgr) GetBatchFromID(ids []int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_DistOrderTblMgr) GetFromBillID(billID string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_DistOrderTblMgr) GetBatchFromBillID(billIDs []string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 分销给的账号
func (obj *_DistOrderTblMgr) GetFromUserID(userID string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 分销给的账号
func (obj *_DistOrderTblMgr) GetBatchFromUserID(userIDs []string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromGpid 通过gpid获取内容 商品id
func (obj *_DistOrderTblMgr) GetFromGpid(gpid string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid = ?", gpid).Find(&results).Error

	return
}

// GetBatchFromGpid 批量唯一主键查找 商品id
func (obj *_DistOrderTblMgr) GetBatchFromGpid(gpids []string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gpid IN (?)", gpids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 商品名字
func (obj *_DistOrderTblMgr) GetFromName(name string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 商品名字
func (obj *_DistOrderTblMgr) GetBatchFromName(names []string) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 分销金额
func (obj *_DistOrderTblMgr) GetFromPrice(price int64) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 分销金额
func (obj *_DistOrderTblMgr) GetBatchFromPrice(prices []int64) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 分销等级
func (obj *_DistOrderTblMgr) GetFromLevel(level int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量唯一主键查找 分销等级
func (obj *_DistOrderTblMgr) GetBatchFromLevel(levels []int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level IN (?)", levels).Find(&results).Error

	return
}

// GetFromTotalPrice 通过total_price获取内容 总分销金额
func (obj *_DistOrderTblMgr) GetFromTotalPrice(totalPrice int64) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price = ?", totalPrice).Find(&results).Error

	return
}

// GetBatchFromTotalPrice 批量唯一主键查找 总分销金额
func (obj *_DistOrderTblMgr) GetBatchFromTotalPrice(totalPrices []int64) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price IN (?)", totalPrices).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态(-1:失效,1:待确认,2:可提现，3:已提现)
func (obj *_DistOrderTblMgr) GetFromStatus(status int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态(-1:失效,1:待确认,2:可提现，3:已提现)
func (obj *_DistOrderTblMgr) GetBatchFromStatus(statuss []int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 商品数量
func (obj *_DistOrderTblMgr) GetFromNumber(number int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量唯一主键查找 商品数量
func (obj *_DistOrderTblMgr) GetBatchFromNumber(numbers []int) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number IN (?)", numbers).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_DistOrderTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_DistOrderTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_DistOrderTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_DistOrderTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DistOrderTblMgr) FetchByPrimaryKey(id int) (result DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByBillUserGpidID primay or index 获取唯一内容
func (obj *_DistOrderTblMgr) FetchUniqueIndexByBillUserGpidID(billID string, userID string, gpid string) (result DistOrderTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ? AND user_id = ? AND gpid = ?", billID, userID, gpid).Find(&result).Error

	return
}
