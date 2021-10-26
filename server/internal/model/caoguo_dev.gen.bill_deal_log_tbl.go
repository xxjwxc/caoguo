package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BillDealLogTblMgr struct {
	*_BaseMgr
}

// BillDealLogTblMgr open func
func BillDealLogTblMgr(db *gorm.DB) *_BillDealLogTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillDealLogTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillDealLogTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_deal_log_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillDealLogTblMgr) GetTableName() string {
	return "bill_deal_log_tbl"
}

// Get 获取
func (obj *_BillDealLogTblMgr) Get() (result BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BillDealLogTblMgr) Gets() (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BillDealLogTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_BillDealLogTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithBillID bill_id获取 账单id
func (obj *_BillDealLogTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithType type获取 操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈
func (obj *_BillDealLogTblMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithContact contact获取 联系方式
func (obj *_BillDealLogTblMgr) WithContact(contact string) Option {
	return optionFunc(func(o *options) { o.query["contact"] = contact })
}

// WithRemak remak获取 备注
func (obj *_BillDealLogTblMgr) WithRemak(remak string) Option {
	return optionFunc(func(o *options) { o.query["remak"] = remak })
}

// WithImgs imgs获取 图片
func (obj *_BillDealLogTblMgr) WithImgs(imgs string) Option {
	return optionFunc(func(o *options) { o.query["imgs"] = imgs })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillDealLogTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BillDealLogTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillDealLogTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_BillDealLogTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_BillDealLogTblMgr) GetByOption(opts ...Option) (result BillDealLogTbl, err error) {
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
func (obj *_BillDealLogTblMgr) GetByOptions(opts ...Option) (results []*BillDealLogTbl, err error) {
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
func (obj *_BillDealLogTblMgr) GetFromID(id int) (result BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_BillDealLogTblMgr) GetBatchFromID(ids []int) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_BillDealLogTblMgr) GetFromUserID(userID string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_BillDealLogTblMgr) GetBatchFromUserID(userIDs []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillDealLogTblMgr) GetFromBillID(billID string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_BillDealLogTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈
func (obj *_BillDealLogTblMgr) GetFromType(_type int) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈
func (obj *_BillDealLogTblMgr) GetBatchFromType(_types []int) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromContact 通过contact获取内容 联系方式
func (obj *_BillDealLogTblMgr) GetFromContact(contact string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contact = ?", contact).Find(&results).Error

	return
}

// GetBatchFromContact 批量唯一主键查找 联系方式
func (obj *_BillDealLogTblMgr) GetBatchFromContact(contacts []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contact IN (?)", contacts).Find(&results).Error

	return
}

// GetFromRemak 通过remak获取内容 备注
func (obj *_BillDealLogTblMgr) GetFromRemak(remak string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remak = ?", remak).Find(&results).Error

	return
}

// GetBatchFromRemak 批量唯一主键查找 备注
func (obj *_BillDealLogTblMgr) GetBatchFromRemak(remaks []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remak IN (?)", remaks).Find(&results).Error

	return
}

// GetFromImgs 通过imgs获取内容 图片
func (obj *_BillDealLogTblMgr) GetFromImgs(imgs string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("imgs = ?", imgs).Find(&results).Error

	return
}

// GetBatchFromImgs 批量唯一主键查找 图片
func (obj *_BillDealLogTblMgr) GetBatchFromImgs(imgss []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("imgs IN (?)", imgss).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillDealLogTblMgr) GetFromCreatedBy(createdBy string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_BillDealLogTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BillDealLogTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_BillDealLogTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillDealLogTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_BillDealLogTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_BillDealLogTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_BillDealLogTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BillDealLogTblMgr) FetchByPrimaryKey(id int) (result BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByBillID  获取多个内容
func (obj *_BillDealLogTblMgr) FetchIndexByBillID(billID string) (results []*BillDealLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}
