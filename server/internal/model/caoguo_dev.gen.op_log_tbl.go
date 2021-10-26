package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _OpLogTblMgr struct {
	*_BaseMgr
}

// OpLogTblMgr open func
func OpLogTblMgr(db *gorm.DB) *_OpLogTblMgr {
	if db == nil {
		panic(fmt.Errorf("OpLogTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OpLogTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("op_log_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OpLogTblMgr) GetTableName() string {
	return "op_log_tbl"
}

// Get 获取
func (obj *_OpLogTblMgr) Get() (result OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OpLogTblMgr) Gets() (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_OpLogTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOperator operator获取 操作人
func (obj *_OpLogTblMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

// WithReceive receive获取 接收方
func (obj *_OpLogTblMgr) WithReceive(receive string) Option {
	return optionFunc(func(o *options) { o.query["receive"] = receive })
}

// WithCreateTime create_time获取 操作时间
func (obj *_OpLogTblMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithTopic topic获取 操作主题
func (obj *_OpLogTblMgr) WithTopic(topic string) Option {
	return optionFunc(func(o *options) { o.query["topic"] = topic })
}

// WithBundle bundle获取 子主体
func (obj *_OpLogTblMgr) WithBundle(bundle string) Option {
	return optionFunc(func(o *options) { o.query["bundle"] = bundle })
}

// WithPid pid获取 细分主体
func (obj *_OpLogTblMgr) WithPid(pid string) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithIPAddr ip_addr获取 操作IP地址
func (obj *_OpLogTblMgr) WithIPAddr(ipAddr string) Option {
	return optionFunc(func(o *options) { o.query["ip_addr"] = ipAddr })
}

// WithNum0 num0获取 附加数字1
func (obj *_OpLogTblMgr) WithNum0(num0 int) Option {
	return optionFunc(func(o *options) { o.query["num0"] = num0 })
}

// WithNum1 num1获取 附加数字2
func (obj *_OpLogTblMgr) WithNum1(num1 int) Option {
	return optionFunc(func(o *options) { o.query["num1"] = num1 })
}

// WithNum2 num2获取 附加数字3
func (obj *_OpLogTblMgr) WithNum2(num2 int) Option {
	return optionFunc(func(o *options) { o.query["num2"] = num2 })
}

// WithAttach0 attach0获取 附加字符串1
func (obj *_OpLogTblMgr) WithAttach0(attach0 string) Option {
	return optionFunc(func(o *options) { o.query["attach0"] = attach0 })
}

// WithAttach1 attach1获取 附加字符串2
func (obj *_OpLogTblMgr) WithAttach1(attach1 string) Option {
	return optionFunc(func(o *options) { o.query["attach1"] = attach1 })
}

// WithAttach2 attach2获取 附加字符串3
func (obj *_OpLogTblMgr) WithAttach2(attach2 string) Option {
	return optionFunc(func(o *options) { o.query["attach2"] = attach2 })
}

// GetByOption 功能选项模式获取
func (obj *_OpLogTblMgr) GetByOption(opts ...Option) (result OpLogTbl, err error) {
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
func (obj *_OpLogTblMgr) GetByOptions(opts ...Option) (results []*OpLogTbl, err error) {
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
func (obj *_OpLogTblMgr) GetFromID(id int) (result OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_OpLogTblMgr) GetBatchFromID(ids []int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOperator 通过operator获取内容 操作人
func (obj *_OpLogTblMgr) GetFromOperator(operator string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator = ?", operator).Find(&results).Error

	return
}

// GetBatchFromOperator 批量唯一主键查找 操作人
func (obj *_OpLogTblMgr) GetBatchFromOperator(operators []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator IN (?)", operators).Find(&results).Error

	return
}

// GetFromReceive 通过receive获取内容 接收方
func (obj *_OpLogTblMgr) GetFromReceive(receive string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receive = ?", receive).Find(&results).Error

	return
}

// GetBatchFromReceive 批量唯一主键查找 接收方
func (obj *_OpLogTblMgr) GetBatchFromReceive(receives []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receive IN (?)", receives).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 操作时间
func (obj *_OpLogTblMgr) GetFromCreateTime(createTime time.Time) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 操作时间
func (obj *_OpLogTblMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromTopic 通过topic获取内容 操作主题
func (obj *_OpLogTblMgr) GetFromTopic(topic string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("topic = ?", topic).Find(&results).Error

	return
}

// GetBatchFromTopic 批量唯一主键查找 操作主题
func (obj *_OpLogTblMgr) GetBatchFromTopic(topics []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("topic IN (?)", topics).Find(&results).Error

	return
}

// GetFromBundle 通过bundle获取内容 子主体
func (obj *_OpLogTblMgr) GetFromBundle(bundle string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bundle = ?", bundle).Find(&results).Error

	return
}

// GetBatchFromBundle 批量唯一主键查找 子主体
func (obj *_OpLogTblMgr) GetBatchFromBundle(bundles []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bundle IN (?)", bundles).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 细分主体
func (obj *_OpLogTblMgr) GetFromPid(pid string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pid = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量唯一主键查找 细分主体
func (obj *_OpLogTblMgr) GetBatchFromPid(pids []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pid IN (?)", pids).Find(&results).Error

	return
}

// GetFromIPAddr 通过ip_addr获取内容 操作IP地址
func (obj *_OpLogTblMgr) GetFromIPAddr(ipAddr string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_addr = ?", ipAddr).Find(&results).Error

	return
}

// GetBatchFromIPAddr 批量唯一主键查找 操作IP地址
func (obj *_OpLogTblMgr) GetBatchFromIPAddr(ipAddrs []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_addr IN (?)", ipAddrs).Find(&results).Error

	return
}

// GetFromNum0 通过num0获取内容 附加数字1
func (obj *_OpLogTblMgr) GetFromNum0(num0 int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("num0 = ?", num0).Find(&results).Error

	return
}

// GetBatchFromNum0 批量唯一主键查找 附加数字1
func (obj *_OpLogTblMgr) GetBatchFromNum0(num0s []int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("num0 IN (?)", num0s).Find(&results).Error

	return
}

// GetFromNum1 通过num1获取内容 附加数字2
func (obj *_OpLogTblMgr) GetFromNum1(num1 int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("num1 = ?", num1).Find(&results).Error

	return
}

// GetBatchFromNum1 批量唯一主键查找 附加数字2
func (obj *_OpLogTblMgr) GetBatchFromNum1(num1s []int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("num1 IN (?)", num1s).Find(&results).Error

	return
}

// GetFromNum2 通过num2获取内容 附加数字3
func (obj *_OpLogTblMgr) GetFromNum2(num2 int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("num2 = ?", num2).Find(&results).Error

	return
}

// GetBatchFromNum2 批量唯一主键查找 附加数字3
func (obj *_OpLogTblMgr) GetBatchFromNum2(num2s []int) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("num2 IN (?)", num2s).Find(&results).Error

	return
}

// GetFromAttach0 通过attach0获取内容 附加字符串1
func (obj *_OpLogTblMgr) GetFromAttach0(attach0 string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach0 = ?", attach0).Find(&results).Error

	return
}

// GetBatchFromAttach0 批量唯一主键查找 附加字符串1
func (obj *_OpLogTblMgr) GetBatchFromAttach0(attach0s []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach0 IN (?)", attach0s).Find(&results).Error

	return
}

// GetFromAttach1 通过attach1获取内容 附加字符串2
func (obj *_OpLogTblMgr) GetFromAttach1(attach1 string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach1 = ?", attach1).Find(&results).Error

	return
}

// GetBatchFromAttach1 批量唯一主键查找 附加字符串2
func (obj *_OpLogTblMgr) GetBatchFromAttach1(attach1s []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach1 IN (?)", attach1s).Find(&results).Error

	return
}

// GetFromAttach2 通过attach2获取内容 附加字符串3
func (obj *_OpLogTblMgr) GetFromAttach2(attach2 string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach2 = ?", attach2).Find(&results).Error

	return
}

// GetBatchFromAttach2 批量唯一主键查找 附加字符串3
func (obj *_OpLogTblMgr) GetBatchFromAttach2(attach2s []string) (results []*OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("attach2 IN (?)", attach2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OpLogTblMgr) FetchByPrimaryKey(id int) (result OpLogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
