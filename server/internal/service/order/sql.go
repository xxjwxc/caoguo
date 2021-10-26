package order

import "time"

// BillOrderTbl 账单订单列表
type BillOrderTbl struct {
	ID        int64     `gorm:"primary_key"`
	BillID    string    `gorm:"unique_index:user_gpid;index:bill_id"` // 账单id
	Gpid      string    `gorm:"unique_index:user_gpid"`               // 商品id
	Name      string    // 商品名字
	Price     int64     // 商品价格(分)
	Icon      string    // 图标信息
	UserID    string    `gorm:"unique_index:user_gpid"` // 用户id
	Number    int       // 数量
	SkuList   string    `gorm:"unique_index:user_gpid"` // 属性列表
	SkuArrt   string    // 属性描述
	CreatedBy string    // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    // 更新者
	UpdatedAt time.Time // 更新时间
	Status    int       // 状态
	CouponID  int       // 优惠券
	Title     string    // 优惠券名称
}
