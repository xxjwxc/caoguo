package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProposalTblMgr struct {
	*_BaseMgr
}

// ProposalTblMgr open func
func ProposalTblMgr(db *gorm.DB) *_ProposalTblMgr {
	if db == nil {
		panic(fmt.Errorf("ProposalTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProposalTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("proposal_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProposalTblMgr) GetTableName() string {
	return "proposal_tbl"
}

// Get 获取
func (obj *_ProposalTblMgr) Get() (result ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProposalTblMgr) Gets() (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProposalTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_ProposalTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithBillID bill_id获取 账单id
func (obj *_ProposalTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithType type获取 操作：-1：取消，1：申请售后，
func (obj *_ProposalTblMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithContact contact获取 联系方式
func (obj *_ProposalTblMgr) WithContact(contact string) Option {
	return optionFunc(func(o *options) { o.query["contact"] = contact })
}

// WithRemak remak获取 备注
func (obj *_ProposalTblMgr) WithRemak(remak string) Option {
	return optionFunc(func(o *options) { o.query["remak"] = remak })
}

// WithImgs imgs获取 图片
func (obj *_ProposalTblMgr) WithImgs(imgs string) Option {
	return optionFunc(func(o *options) { o.query["imgs"] = imgs })
}

// WithCreatedBy created_by获取 创建者
func (obj *_ProposalTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ProposalTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_ProposalTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ProposalTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_ProposalTblMgr) GetByOption(opts ...Option) (result ProposalTbl, err error) {
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
func (obj *_ProposalTblMgr) GetByOptions(opts ...Option) (results []*ProposalTbl, err error) {
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
func (obj *_ProposalTblMgr) GetFromID(id int) (result ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ProposalTblMgr) GetBatchFromID(ids []int) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_ProposalTblMgr) GetFromUserID(userID string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_ProposalTblMgr) GetBatchFromUserID(userIDs []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_ProposalTblMgr) GetFromBillID(billID string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量唯一主键查找 账单id
func (obj *_ProposalTblMgr) GetBatchFromBillID(billIDs []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 操作：-1：取消，1：申请售后，
func (obj *_ProposalTblMgr) GetFromType(_type int) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 操作：-1：取消，1：申请售后，
func (obj *_ProposalTblMgr) GetBatchFromType(_types []int) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromContact 通过contact获取内容 联系方式
func (obj *_ProposalTblMgr) GetFromContact(contact string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contact = ?", contact).Find(&results).Error

	return
}

// GetBatchFromContact 批量唯一主键查找 联系方式
func (obj *_ProposalTblMgr) GetBatchFromContact(contacts []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contact IN (?)", contacts).Find(&results).Error

	return
}

// GetFromRemak 通过remak获取内容 备注
func (obj *_ProposalTblMgr) GetFromRemak(remak string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remak = ?", remak).Find(&results).Error

	return
}

// GetBatchFromRemak 批量唯一主键查找 备注
func (obj *_ProposalTblMgr) GetBatchFromRemak(remaks []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remak IN (?)", remaks).Find(&results).Error

	return
}

// GetFromImgs 通过imgs获取内容 图片
func (obj *_ProposalTblMgr) GetFromImgs(imgs string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("imgs = ?", imgs).Find(&results).Error

	return
}

// GetBatchFromImgs 批量唯一主键查找 图片
func (obj *_ProposalTblMgr) GetBatchFromImgs(imgss []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("imgs IN (?)", imgss).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_ProposalTblMgr) GetFromCreatedBy(createdBy string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找 创建者
func (obj *_ProposalTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ProposalTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_ProposalTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_ProposalTblMgr) GetFromUpdatedBy(updatedBy string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量唯一主键查找 更新者
func (obj *_ProposalTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_by IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ProposalTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_ProposalTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProposalTblMgr) FetchByPrimaryKey(id int) (result ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByBillID  获取多个内容
func (obj *_ProposalTblMgr) FetchIndexByBillID(billID string) (results []*ProposalTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bill_id = ?", billID).Find(&results).Error

	return
}
