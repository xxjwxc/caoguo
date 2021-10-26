package model

import (
	"gorm.io/gorm"
	"time"
)

/******sql******
CREATE TABLE `address_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(255) DEFAULT NULL COMMENT '[@gormt default:''123456'';]手机号',
  `address_name` varchar(1024) DEFAULT NULL COMMENT '地址名称',
  `address` varchar(1024) DEFAULT NULL COMMENT '详细地址',
  `area` varchar(1024) DEFAULT NULL COMMENT '单元门牌号',
  `default` tinyint(1) DEFAULT NULL COMMENT '是否是默认值',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 COMMENT='用户地址列表'
******sql******/
// AddressTbl 用户地址列表
type AddressTbl struct {
	ID          int64     `gorm:"primaryKey"`
	UserID      string    `gorm:"index:user_id"` // 用户id
	Name        string    // 用户名
	Mobile      string    // [@gormt default:'123456';]手机号
	AddressName string    // 地址名称
	Address     string    // 详细地址
	Area        string    // 单元门牌号
	Default     bool      // 是否是默认值
	CreatedBy   string    `gorm:"default:''"` // 创建者
	CreatedAt   time.Time // 创建时间
	UpdatedBy   string    `gorm:"default:''"` // 更新者
	UpdatedAt   time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *AddressTbl) TableName() string {
	return "address_tbl"
}

// AddressTblColumns get sql column name.获取数据库列名
var AddressTblColumns = struct {
	ID          string
	UserID      string
	Name        string
	Mobile      string
	AddressName string
	Address     string
	Area        string
	Default     string
	CreatedBy   string
	CreatedAt   string
	UpdatedBy   string
	UpdatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	Name:        "name",
	Mobile:      "mobile",
	AddressName: "address_name",
	Address:     "address",
	Area:        "area",
	Default:     "default",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
}

/******sql******
CREATE TABLE `bill_address_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addr_id` bigint(20) DEFAULT NULL COMMENT '地址id',
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `bill_id` varchar(32) DEFAULT NULL COMMENT '账单id',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(255) DEFAULT NULL COMMENT '手机号',
  `address_name` varchar(1024) DEFAULT NULL COMMENT '地址名称',
  `address` varchar(1024) DEFAULT NULL COMMENT '详细地址',
  `area` varchar(1024) DEFAULT NULL COMMENT '单元门牌号',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bill_id` (`bill_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8mb4 COMMENT='派送信息'
******sql******/
// BillAddressTbl 派送信息
type BillAddressTbl struct {
	ID          int64     `gorm:"primaryKey"` // 主键
	AddrID      int64     // 地址id
	UserID      string    // 用户id
	BillID      string    `gorm:"unique"` // 账单id
	Name        string    // 用户名
	Mobile      string    // 手机号
	AddressName string    // 地址名称
	Address     string    // 详细地址
	Area        string    // 单元门牌号
	CreatedBy   string    `gorm:"default:''"` // 创建者
	CreatedAt   time.Time // 创建时间
	UpdatedBy   string    `gorm:"default:''"` // 更新者
	UpdatedAt   time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *BillAddressTbl) TableName() string {
	return "bill_address_tbl"
}

// BillAddressTblColumns get sql column name.获取数据库列名
var BillAddressTblColumns = struct {
	ID          string
	AddrID      string
	UserID      string
	BillID      string
	Name        string
	Mobile      string
	AddressName string
	Address     string
	Area        string
	CreatedBy   string
	CreatedAt   string
	UpdatedBy   string
	UpdatedAt   string
}{
	ID:          "id",
	AddrID:      "addr_id",
	UserID:      "user_id",
	BillID:      "bill_id",
	Name:        "name",
	Mobile:      "mobile",
	AddressName: "address_name",
	Address:     "address",
	Area:        "area",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
}

/******sql******
CREATE TABLE `bill_coupon_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) DEFAULT NULL COMMENT '账单id',
  `coupon_id` bigint(20) NOT NULL COMMENT '我领取的优惠券id',
  `title` varchar(255) NOT NULL COMMENT '优惠券名字',
  `denom` int(11) DEFAULT NULL COMMENT '面额',
  `coupon_type` int(11) DEFAULT NULL COMMENT '优惠券类型(1:全场，2:单个商品)',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bill_id` (`bill_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COMMENT='优惠券列表池'
******sql******/
// BillCouponTbl 优惠券列表池
type BillCouponTbl struct {
	ID         int64     `gorm:"primaryKey"`
	BillID     string    `gorm:"unique"` // 账单id
	CouponID   int64     // 我领取的优惠券id
	Title      string    // 优惠券名字
	Denom      int       // 面额
	CouponType int       // 优惠券类型(1:全场，2:单个商品)
	CreatedBy  string    `gorm:"default:''"` // 创建者
	CreatedAt  time.Time // 创建时间
	UpdatedBy  string    `gorm:"default:''"` // 更新者
	UpdatedAt  time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *BillCouponTbl) TableName() string {
	return "bill_coupon_tbl"
}

// BillCouponTblColumns get sql column name.获取数据库列名
var BillCouponTblColumns = struct {
	ID         string
	BillID     string
	CouponID   string
	Title      string
	Denom      string
	CouponType string
	CreatedBy  string
	CreatedAt  string
	UpdatedBy  string
	UpdatedAt  string
}{
	ID:         "id",
	BillID:     "bill_id",
	CouponID:   "coupon_id",
	Title:      "title",
	Denom:      "denom",
	CouponType: "coupon_type",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

/******sql******
CREATE TABLE `bill_deal_log_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) NOT NULL COMMENT '账单id',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈',
  `contact` varchar(64) NOT NULL DEFAULT '0' COMMENT '联系方式',
  `remak` varchar(1024) DEFAULT NULL COMMENT '备注',
  `imgs` text COMMENT '图片',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `bill_id` (`bill_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COMMENT='账单操作日志'
******sql******/
// BillDealLogTbl 账单操作日志
type BillDealLogTbl struct {
	ID        int       `gorm:"primaryKey"`
	UserID    string    `gorm:"default:0"`     // 用户id
	BillID    string    `gorm:"index:bill_id"` // 账单id
	Type      int       `gorm:"default:0"`     // 操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈
	Contact   string    `gorm:"default:0"`     // 联系方式
	Remak     string    // 备注
	Imgs      string    // 图片
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *BillDealLogTbl) TableName() string {
	return "bill_deal_log_tbl"
}

// BillDealLogTblColumns get sql column name.获取数据库列名
var BillDealLogTblColumns = struct {
	ID        string
	UserID    string
	BillID    string
	Type      string
	Contact   string
	Remak     string
	Imgs      string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	BillID:    "bill_id",
	Type:      "type",
	Contact:   "contact",
	Remak:     "remak",
	Imgs:      "imgs",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `bill_order_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) DEFAULT NULL COMMENT '账单id',
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `name` varchar(255) NOT NULL COMMENT '商品名字',
  `price` bigint(20) NOT NULL COMMENT '商品价格(分)',
  `icon` text NOT NULL COMMENT '图标信息',
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `number` int(11) NOT NULL COMMENT '数量',
  `dist_amount` bigint(20) DEFAULT NULL COMMENT '分享收益',
  `sku_list` varchar(255) DEFAULT NULL COMMENT '属性列表',
  `sku_arrt` varchar(255) DEFAULT NULL COMMENT '属性描述',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_gpid` (`gpid`,`user_id`,`sku_list`,`bill_id`) USING BTREE,
  KEY `bill_id` (`bill_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8mb4 COMMENT='账单订单列表'
******sql******/
// BillOrderTbl 账单订单列表
type BillOrderTbl struct {
	ID         int64     `gorm:"primaryKey"`
	BillID     string    `gorm:"uniqueIndex:user_gpid;index:bill_id"` // 账单id
	Gpid       string    `gorm:"uniqueIndex:user_gpid"`               // 商品id
	Name       string    // 商品名字
	Price      int64     // 商品价格(分)
	Icon       string    // 图标信息
	UserID     string    `gorm:"uniqueIndex:user_gpid"` // 用户id
	Number     int       // 数量
	DistAmount int64     // 分享收益
	SkuList    string    `gorm:"uniqueIndex:user_gpid"` // 属性列表
	SkuArrt    string    // 属性描述
	CreatedBy  string    `gorm:"default:''"` // 创建者
	CreatedAt  time.Time // 创建时间
	UpdatedBy  string    `gorm:"default:''"` // 更新者
	UpdatedAt  time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *BillOrderTbl) TableName() string {
	return "bill_order_tbl"
}

// BillOrderTblColumns get sql column name.获取数据库列名
var BillOrderTblColumns = struct {
	ID         string
	BillID     string
	Gpid       string
	Name       string
	Price      string
	Icon       string
	UserID     string
	Number     string
	DistAmount string
	SkuList    string
	SkuArrt    string
	CreatedBy  string
	CreatedAt  string
	UpdatedBy  string
	UpdatedAt  string
}{
	ID:         "id",
	BillID:     "bill_id",
	Gpid:       "gpid",
	Name:       "name",
	Price:      "price",
	Icon:       "icon",
	UserID:     "user_id",
	Number:     "number",
	DistAmount: "dist_amount",
	SkuList:    "sku_list",
	SkuArrt:    "sku_arrt",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

/******sql******
CREATE TABLE `bill_refund_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) NOT NULL COMMENT '账单id',
  `refund_id` varchar(32) NOT NULL DEFAULT '0' COMMENT '退款账单号',
  `refund_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '退款金额',
  `desc` varchar(1024) DEFAULT NULL COMMENT '订单备注',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `bill_id` (`bill_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='用户退款信息'
******sql******/
// BillRefundTbl 用户退款信息
type BillRefundTbl struct {
	ID        int       `gorm:"primaryKey"`
	UserID    string    `gorm:"default:0"`     // 用户id
	BillID    string    `gorm:"index:bill_id"` // 账单id
	RefundID  string    `gorm:"default:0"`     // 退款账单号
	RefundFee int64     `gorm:"default:0"`     // 退款金额
	Desc      string    // 订单备注
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *BillRefundTbl) TableName() string {
	return "bill_refund_tbl"
}

// BillRefundTblColumns get sql column name.获取数据库列名
var BillRefundTblColumns = struct {
	ID        string
	UserID    string
	BillID    string
	RefundID  string
	RefundFee string
	Desc      string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	BillID:    "bill_id",
	RefundID:  "refund_id",
	RefundFee: "refund_fee",
	Desc:      "desc",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `bill_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) NOT NULL COMMENT '账单id',
  `pay_platform` int(11) NOT NULL DEFAULT '0' COMMENT '支付类型(1:微信支付)',
  `pay_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付金额',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)',
  `desc` varchar(1024) DEFAULT NULL COMMENT '订单备注',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bill_id` (`bill_id`) USING BTREE,
  CONSTRAINT `bill_tbl_ibfk_1` FOREIGN KEY (`bill_id`) REFERENCES `bill_address_tbl` (`bill_id`),
  CONSTRAINT `bill_tbl_ibfk_2` FOREIGN KEY (`bill_id`) REFERENCES `bill_order_tbl` (`bill_id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COMMENT='用户账单信息'
******sql******/
// BillTbl 用户账单信息
type BillTbl struct {
	gorm.Model
	UserID           string         `gorm:"default:0"`                                 // 用户id
	BillID           string         `gorm:"unique"`                                    // 账单id
	BillAddressTbl   BillAddressTbl `gorm:"joinForeignKey:bill_id;foreignKey:bill_id"` // 派送信息
	BillOrderTblList []BillOrderTbl `gorm:"joinForeignKey:bill_id;foreignKey:bill_id"` // 账单订单列表
	PayPlatform      int            `gorm:"default:0"`                                 // 支付类型(1:微信支付)
	PayAmount        int64          `gorm:"default:0"`                                 // 支付金额
	Status           int            `gorm:"default:0"`                                 // 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
	Desc             string         // 订单备注
	CreatedBy        string         `gorm:"default:''"` // 创建者
	UpdatedBy        string         `gorm:"default:''"` // 更新者
	DeletedBy        string         `gorm:"default:''"` // 删除者
}

// TableName get sql table name.获取数据库表名
func (m *BillTbl) TableName() string {
	return "bill_tbl"
}

// BillTblColumns get sql column name.获取数据库列名
var BillTblColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	UserID      string
	BillID      string
	PayPlatform string
	PayAmount   string
	Status      string
	Desc        string
	CreatedBy   string
	UpdatedBy   string
	DeletedBy   string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	UserID:      "user_id",
	BillID:      "bill_id",
	PayPlatform: "pay_platform",
	PayAmount:   "pay_amount",
	Status:      "status",
	Desc:        "desc",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	DeletedBy:   "deleted_by",
}

/******sql******
CREATE TABLE `cart_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `user_type` int(11) DEFAULT NULL COMMENT '用户类型(1:微信)',
  `number` int(11) DEFAULT NULL COMMENT '数量',
  `sku_list` varchar(255) DEFAULT NULL COMMENT '属性列表',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_gpid` (`gpid`,`user_id`,`sku_list`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 COMMENT='购物车列表'
******sql******/
// CartTbl 购物车列表
type CartTbl struct {
	ID        int64     `gorm:"primaryKey"`
	Gpid      string    `gorm:"uniqueIndex:user_gpid"` // 商品id
	UserID    string    `gorm:"uniqueIndex:user_gpid"` // 用户id
	UserType  int       // 用户类型(1:微信)
	Number    int       // 数量
	SkuList   string    `gorm:"uniqueIndex:user_gpid"` // 属性列表
	CreatedBy string    `gorm:"default:''"`            // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *CartTbl) TableName() string {
	return "cart_tbl"
}

// CartTblColumns get sql column name.获取数据库列名
var CartTblColumns = struct {
	ID        string
	Gpid      string
	UserID    string
	UserType  string
	Number    string
	SkuList   string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	Gpid:      "gpid",
	UserID:    "user_id",
	UserType:  "user_type",
	Number:    "number",
	SkuList:   "sku_list",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `cart_tmp_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `user_type` int(11) DEFAULT NULL COMMENT '用户类型(1:微信)',
  `number` int(11) DEFAULT NULL COMMENT '数量',
  `sku_list` varchar(255) DEFAULT NULL COMMENT '属性列表',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=150 DEFAULT CHARSET=utf8mb4 COMMENT='虚拟购物车列表'
******sql******/
// CartTmpTbl 虚拟购物车列表
type CartTmpTbl struct {
	ID        int64     `gorm:"primaryKey"`
	Gpid      string    // 商品id
	UserID    string    `gorm:"unique"` // 用户id
	UserType  int       // 用户类型(1:微信)
	Number    int       // 数量
	SkuList   string    // 属性列表
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *CartTmpTbl) TableName() string {
	return "cart_tmp_tbl"
}

// CartTmpTblColumns get sql column name.获取数据库列名
var CartTmpTblColumns = struct {
	ID        string
	Gpid      string
	UserID    string
	UserType  string
	Number    string
	SkuList   string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	Gpid:      "gpid",
	UserID:    "user_id",
	UserType:  "user_type",
	Number:    "number",
	SkuList:   "sku_list",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `coupon_my_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `coupon_id` bigint(20) NOT NULL COMMENT '优惠券id',
  `expires_time` timestamp NULL DEFAULT NULL COMMENT '过期时间',
  `status` int(11) DEFAULT NULL COMMENT '当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `coupon_id` (`coupon_id`) USING BTREE,
  CONSTRAINT `coupon_my_tbl_ibfk_1` FOREIGN KEY (`coupon_id`) REFERENCES `coupon_tbl` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=483 DEFAULT CHARSET=utf8mb4 COMMENT='优惠券列表池'
******sql******/
// CouponMyTbl 优惠券列表池
type CouponMyTbl struct {
	ID          int64     `gorm:"primaryKey"`
	UserID      string    // 用户id
	CouponID    int64     `gorm:"index:coupon_id"`                        // 优惠券id
	CouponTbl   CouponTbl `gorm:"joinForeignKey:coupon_id;foreignKey:id"` // 优惠券列表池
	ExpiresTime time.Time // 过期时间
	Status      int       // 当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)
	CreatedBy   string    `gorm:"default:''"` // 创建者
	CreatedAt   time.Time // 创建时间
	UpdatedBy   string    `gorm:"default:''"` // 更新者
	UpdatedAt   time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *CouponMyTbl) TableName() string {
	return "coupon_my_tbl"
}

// CouponMyTblColumns get sql column name.获取数据库列名
var CouponMyTblColumns = struct {
	ID          string
	UserID      string
	CouponID    string
	ExpiresTime string
	Status      string
	CreatedBy   string
	CreatedAt   string
	UpdatedBy   string
	UpdatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	CouponID:    "coupon_id",
	ExpiresTime: "expires_time",
	Status:      "status",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
}

/******sql******
CREATE TABLE `coupon_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `group_name` varchar(32) NOT NULL COMMENT '分组名',
  `title` varchar(255) NOT NULL COMMENT '优惠券名字',
  `validity` int(11) DEFAULT NULL COMMENT '有效期(天数)',
  `gpid` varchar(32) DEFAULT NULL COMMENT '商品优惠券',
  `denom` int(11) DEFAULT NULL COMMENT '面额',
  `coupon_type` int(11) DEFAULT NULL COMMENT '优惠券类型(1:全场，2:单个商品)',
  `great_money` int(11) DEFAULT NULL COMMENT '满多少可用',
  `describle` varchar(255) DEFAULT NULL COMMENT '优惠券描述',
  `vaild` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='优惠券列表池'
******sql******/
// CouponTbl 优惠券列表池
type CouponTbl struct {
	ID         int64     `gorm:"primaryKey"`
	GroupName  string    // 分组名
	Title      string    // 优惠券名字
	Validity   int       // 有效期(天数)
	Gpid       string    // 商品优惠券
	Denom      int       // 面额
	CouponType int       // 优惠券类型(1:全场，2:单个商品)
	GreatMoney int       // 满多少可用
	Describle  string    // 优惠券描述
	Vaild      bool      `gorm:"default:1"`  // 是否有效
	CreatedBy  string    `gorm:"default:''"` // 创建者
	CreatedAt  time.Time // 创建时间
	UpdatedBy  string    `gorm:"default:''"` // 更新者
	UpdatedAt  time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *CouponTbl) TableName() string {
	return "coupon_tbl"
}

// CouponTblColumns get sql column name.获取数据库列名
var CouponTblColumns = struct {
	ID         string
	GroupName  string
	Title      string
	Validity   string
	Gpid       string
	Denom      string
	CouponType string
	GreatMoney string
	Describle  string
	Vaild      string
	CreatedBy  string
	CreatedAt  string
	UpdatedBy  string
	UpdatedAt  string
}{
	ID:         "id",
	GroupName:  "group_name",
	Title:      "title",
	Validity:   "validity",
	Gpid:       "gpid",
	Denom:      "denom",
	CouponType: "coupon_type",
	GreatMoney: "great_money",
	Describle:  "describle",
	Vaild:      "vaild",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

/******sql******
CREATE TABLE `dist_order_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) NOT NULL COMMENT '账单id',
  `user_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '分销给的账号',
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `name` varchar(255) DEFAULT NULL COMMENT '商品名字',
  `price` bigint(20) NOT NULL COMMENT '分销金额',
  `level` int(255) DEFAULT NULL COMMENT '分销等级',
  `total_price` bigint(20) DEFAULT NULL COMMENT '总分销金额',
  `status` int(11) DEFAULT NULL COMMENT '状态(-1:失效,1:待确认,2:可提现，3:已提现)',
  `number` int(11) DEFAULT NULL COMMENT '商品数量',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bill_user_gpid_id` (`bill_id`,`user_id`,`gpid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分销详细账单'
******sql******/
// DistOrderTbl 分销详细账单
type DistOrderTbl struct {
	ID         int       `gorm:"primaryKey"`
	BillID     string    `gorm:"uniqueIndex:bill_user_gpid_id"` // 账单id
	UserID     string    `gorm:"uniqueIndex:bill_user_gpid_id"` // 分销给的账号
	Gpid       string    `gorm:"uniqueIndex:bill_user_gpid_id"` // 商品id
	Name       string    // 商品名字
	Price      int64     // 分销金额
	Level      int       // 分销等级
	TotalPrice int64     // 总分销金额
	Status     int       // 状态(-1:失效,1:待确认,2:可提现，3:已提现)
	Number     int       // 商品数量
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *DistOrderTbl) TableName() string {
	return "dist_order_tbl"
}

// DistOrderTblColumns get sql column name.获取数据库列名
var DistOrderTblColumns = struct {
	ID         string
	BillID     string
	UserID     string
	Gpid       string
	Name       string
	Price      string
	Level      string
	TotalPrice string
	Status     string
	Number     string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	BillID:     "bill_id",
	UserID:     "user_id",
	Gpid:       "gpid",
	Name:       "name",
	Price:      "price",
	Level:      "level",
	TotalPrice: "total_price",
	Status:     "status",
	Number:     "number",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

/******sql******
CREATE TABLE `dist_userinfo_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '子节点',
  `parnt_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '父节点',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `userid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信用户信息'
******sql******/
// DistUserinfoTbl 微信用户信息
type DistUserinfoTbl struct {
	ID        int       `gorm:"primaryKey"`
	UserID    string    `gorm:"unique"` // 子节点
	ParntID   string    // 父节点
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *DistUserinfoTbl) TableName() string {
	return "dist_userinfo_tbl"
}

// DistUserinfoTblColumns get sql column name.获取数据库列名
var DistUserinfoTblColumns = struct {
	ID        string
	UserID    string
	ParntID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	ParntID:   "parnt_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `email_notify_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户名',
  `email` varchar(512) DEFAULT NULL COMMENT '邮箱号',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `shipment_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='运输信息明细'
******sql******/
// EmailNotifyTbl 运输信息明细
type EmailNotifyTbl struct {
	ID        int       `gorm:"primaryKey"`
	UserID    string    `gorm:"index:shipment_id"` // 用户名
	Email     string    // 邮箱号
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *EmailNotifyTbl) TableName() string {
	return "email_notify_tbl"
}

// EmailNotifyTblColumns get sql column name.获取数据库列名
var EmailNotifyTblColumns = struct {
	ID        string
	UserID    string
	Email     string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Email:     "email",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `favorite_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `user_id` varchar(255) NOT NULL COMMENT '用户id',
  `user_type` int(11) NOT NULL COMMENT '用户类型(1:微信)',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_gpid` (`gpid`,`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COMMENT='收藏列表'
******sql******/
// FavoriteTbl 收藏列表
type FavoriteTbl struct {
	ID        int       `gorm:"primaryKey"`
	Gpid      string    `gorm:"uniqueIndex:user_gpid"` // 商品id
	UserID    string    `gorm:"uniqueIndex:user_gpid"` // 用户id
	UserType  int       // 用户类型(1:微信)
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FavoriteTbl) TableName() string {
	return "favorite_tbl"
}

// FavoriteTblColumns get sql column name.获取数据库列名
var FavoriteTblColumns = struct {
	ID        string
	Gpid      string
	UserID    string
	UserType  string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	Gpid:      "gpid",
	UserID:    "user_id",
	UserType:  "user_type",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `logistics_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shipment_id` varchar(255) DEFAULT NULL COMMENT '快递单号',
  `times` datetime DEFAULT NULL COMMENT '时间',
  `context` varchar(512) DEFAULT NULL COMMENT '描述信息',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `shipment_id` (`shipment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1698 DEFAULT CHARSET=utf8mb4 COMMENT='运输信息明细'
******sql******/
// LogisticsTbl 运输信息明细
type LogisticsTbl struct {
	ID         int       `gorm:"primaryKey"`
	ShipmentID string    `gorm:"index:shipment_id"` // 快递单号
	Times      time.Time // 时间
	Context    string    // 描述信息
	CreatedBy  string    `gorm:"default:''"` // 创建者
	CreatedAt  time.Time // 创建时间
	UpdatedBy  string    `gorm:"default:''"` // 更新者
	UpdatedAt  time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *LogisticsTbl) TableName() string {
	return "logistics_tbl"
}

// LogisticsTblColumns get sql column name.获取数据库列名
var LogisticsTblColumns = struct {
	ID         string
	ShipmentID string
	Times      string
	Context    string
	CreatedBy  string
	CreatedAt  string
	UpdatedBy  string
	UpdatedAt  string
}{
	ID:         "id",
	ShipmentID: "shipment_id",
	Times:      "times",
	Context:    "context",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

/******sql******
CREATE TABLE `op_log_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operator` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '操作人',
  `receive` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '接收方',
  `create_time` datetime DEFAULT NULL COMMENT '操作时间',
  `topic` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '操作主题',
  `bundle` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '子主体',
  `pid` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '细分主体',
  `ip_addr` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '操作IP地址',
  `num0` int(11) DEFAULT NULL COMMENT '附加数字1',
  `num1` int(11) DEFAULT NULL COMMENT '附加数字2',
  `num2` int(11) DEFAULT NULL COMMENT '附加数字3',
  `attach0` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '附加字符串1',
  `attach1` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '附加字符串2',
  `attach2` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '附加字符串3',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表'
******sql******/
// OpLogTbl 操作日志表
type OpLogTbl struct {
	ID         int       `gorm:"primaryKey"`
	Operator   string    // 操作人
	Receive    string    // 接收方
	CreateTime time.Time // 操作时间
	Topic      string    // 操作主题
	Bundle     string    // 子主体
	Pid        string    // 细分主体
	IPAddr     string    // 操作IP地址
	Num0       int       // 附加数字1
	Num1       int       // 附加数字2
	Num2       int       // 附加数字3
	Attach0    string    // 附加字符串1
	Attach1    string    // 附加字符串2
	Attach2    string    // 附加字符串3
}

// TableName get sql table name.获取数据库表名
func (m *OpLogTbl) TableName() string {
	return "op_log_tbl"
}

// OpLogTblColumns get sql column name.获取数据库列名
var OpLogTblColumns = struct {
	ID         string
	Operator   string
	Receive    string
	CreateTime string
	Topic      string
	Bundle     string
	Pid        string
	IPAddr     string
	Num0       string
	Num1       string
	Num2       string
	Attach0    string
	Attach1    string
	Attach2    string
}{
	ID:         "id",
	Operator:   "operator",
	Receive:    "receive",
	CreateTime: "create_time",
	Topic:      "topic",
	Bundle:     "bundle",
	Pid:        "pid",
	IPAddr:     "ip_addr",
	Num0:       "num0",
	Num1:       "num1",
	Num2:       "num2",
	Attach0:    "attach0",
	Attach1:    "attach1",
	Attach2:    "attach2",
}

/******sql******
CREATE TABLE `post_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '快递名称',
  `code` varchar(255) DEFAULT NULL COMMENT '快递id',
  `icon` varchar(255) DEFAULT NULL COMMENT '快递logo地址',
  `exp_phone` varchar(255) DEFAULT NULL COMMENT '快递电话',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`code`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='快递公司信息'
******sql******/
// PostTbl 快递公司信息
type PostTbl struct {
	ID       int    `gorm:"primaryKey"`
	Name     string // 快递名称
	Code     string `gorm:"unique"` // 快递id
	Icon     string // 快递logo地址
	ExpPhone string // 快递电话
}

// TableName get sql table name.获取数据库表名
func (m *PostTbl) TableName() string {
	return "post_tbl"
}

// PostTblColumns get sql column name.获取数据库列名
var PostTblColumns = struct {
	ID       string
	Name     string
	Code     string
	Icon     string
	ExpPhone string
}{
	ID:       "id",
	Name:     "name",
	Code:     "code",
	Icon:     "icon",
	ExpPhone: "exp_phone",
}

/******sql******
CREATE TABLE `product_ad_group_tbl` (
  `id` int(11) NOT NULL,
  `main_gpid` varchar(32) DEFAULT NULL COMMENT '主商品id\n',
  `sub_gpid` varchar(32) DEFAULT NULL COMMENT '附加商品id',
  `vaild` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效',
  `sort_id` int(11) NOT NULL COMMENT '排序(越大越前)',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `main_sub_gpid` (`main_gpid`,`sub_gpid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页精选列表'
******sql******/
// ProductAdGroupTbl 首页精选列表
type ProductAdGroupTbl struct {
	ID        int       `gorm:"primaryKey"`
	MainGpid  string    `gorm:"uniqueIndex:main_sub_gpid"` // 主商品id,
	SubGpid   string    `gorm:"uniqueIndex:main_sub_gpid"` // 附加商品id
	Vaild     bool      `gorm:"default:1"`                 // 是否有效
	SortID    int       // 排序(越大越前)
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductAdGroupTbl) TableName() string {
	return "product_ad_group_tbl"
}

// ProductAdGroupTblColumns get sql column name.获取数据库列名
var ProductAdGroupTblColumns = struct {
	ID        string
	MainGpid  string
	SubGpid   string
	Vaild     string
	SortID    string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	MainGpid:  "main_gpid",
	SubGpid:   "sub_gpid",
	Vaild:     "vaild",
	SortID:    "sort_id",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `product_ad_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(512) DEFAULT NULL COMMENT '跳转路径',
  `img` varchar(255) DEFAULT NULL COMMENT '卡片展示图片',
  `sort_id` int(11) NOT NULL DEFAULT '0' COMMENT '排序(越大越前)',
  `type` int(11) DEFAULT NULL COMMENT '类型(1:轮播图广告，2:类型广告，3:主销广告)',
  `attach` varchar(255) DEFAULT NULL COMMENT '附加属性',
  `vaild` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `gpid` (`url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='广告轮播图'
******sql******/
// ProductAdTbl 广告轮播图
type ProductAdTbl struct {
	gorm.Model
	URL       string `gorm:"index:gpid"` // 跳转路径
	Img       string // 卡片展示图片
	SortID    int    `gorm:"default:0"` // 排序(越大越前)
	Type      int    // 类型(1:轮播图广告，2:类型广告，3:主销广告)
	Attach    string // 附加属性
	Vaild     bool   `gorm:"default:1"`  // 是否有效
	CreatedBy string `gorm:"default:''"` // 创建者
	UpdatedBy string `gorm:"default:''"` // 更新者
	DeletedBy string `gorm:"default:''"` // 删除者
}

// TableName get sql table name.获取数据库表名
func (m *ProductAdTbl) TableName() string {
	return "product_ad_tbl"
}

// ProductAdTblColumns get sql column name.获取数据库列名
var ProductAdTblColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	URL       string
	Img       string
	SortID    string
	Type      string
	Attach    string
	Vaild     string
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	URL:       "url",
	Img:       "img",
	SortID:    "sort_id",
	Type:      "type",
	Attach:    "attach",
	Vaild:     "vaild",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	DeletedBy: "deleted_by",
}

/******sql******
CREATE TABLE `product_info_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增',
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `qty` bigint(20) NOT NULL COMMENT '已购买数量(销量)',
  `img` text NOT NULL COMMENT '轮播图',
  `video` text NOT NULL COMMENT '视频信息',
  `icon` text NOT NULL COMMENT '商品图标',
  `service` text COMMENT '商品服务',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `gpid` (`gpid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COMMENT='商品附加属性'
******sql******/
// ProductInfoTbl 商品附加属性
type ProductInfoTbl struct {
	ID        int64     `gorm:"primaryKey"` // 自增
	Gpid      string    `gorm:"unique"`     // 商品id
	Qty       int64     // 已购买数量(销量)
	Img       string    // 轮播图
	Video     string    // 视频信息
	Icon      string    // 商品图标
	Service   string    // 商品服务
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductInfoTbl) TableName() string {
	return "product_info_tbl"
}

// ProductInfoTblColumns get sql column name.获取数据库列名
var ProductInfoTblColumns = struct {
	ID        string
	Gpid      string
	Qty       string
	Img       string
	Video     string
	Icon      string
	Service   string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	Gpid:      "gpid",
	Qty:       "qty",
	Img:       "img",
	Video:     "video",
	Icon:      "icon",
	Service:   "service",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `product_sku_price_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `sku_list` varchar(32) NOT NULL COMMENT '属性列表',
  `premium` bigint(20) DEFAULT NULL COMMENT '商品单价',
  `dist_amount` bigint(20) DEFAULT NULL COMMENT '分享收益',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `sku_list` (`sku_list`),
  KEY `gpid` (`gpid`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COMMENT='产品sku 属性溢价表'
******sql******/
// ProductSkuPriceTbl 产品sku 属性溢价表
type ProductSkuPriceTbl struct {
	gorm.Model
	Gpid       string `gorm:"index:gpid"` // 商品id
	SkuList    string `gorm:"unique"`     // 属性列表
	Premium    int64  // 商品单价
	DistAmount int64  // 分享收益
	CreatedBy  string `gorm:"default:''"` // 创建者
	UpdatedBy  string `gorm:"default:''"` // 更新者
	DeletedBy  string `gorm:"default:''"` // 删除者
}

// TableName get sql table name.获取数据库表名
func (m *ProductSkuPriceTbl) TableName() string {
	return "product_sku_price_tbl"
}

// ProductSkuPriceTblColumns get sql column name.获取数据库列名
var ProductSkuPriceTblColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	Gpid       string
	SkuList    string
	Premium    string
	DistAmount string
	CreatedBy  string
	UpdatedBy  string
	DeletedBy  string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	Gpid:       "gpid",
	SkuList:    "sku_list",
	Premium:    "premium",
	DistAmount: "dist_amount",
	CreatedBy:  "created_by",
	UpdatedBy:  "updated_by",
	DeletedBy:  "deleted_by",
}

/******sql******
CREATE TABLE `product_sku_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `type_name` varchar(255) DEFAULT NULL COMMENT '标签类型',
  `tag_name` varchar(255) DEFAULT NULL COMMENT '标签名字',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `gpid_tag` (`gpid`,`tag_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='产品sku 属性'
******sql******/
// ProductSkuTbl 产品sku 属性
type ProductSkuTbl struct {
	gorm.Model
	Gpid      string `gorm:"uniqueIndex:gpid_tag"` // 商品id
	TypeName  string // 标签类型
	TagName   string `gorm:"uniqueIndex:gpid_tag"` // 标签名字
	CreatedBy string `gorm:"default:''"`           // 创建者
	UpdatedBy string `gorm:"default:''"`           // 更新者
	DeletedBy string `gorm:"default:''"`           // 删除者
}

// TableName get sql table name.获取数据库表名
func (m *ProductSkuTbl) TableName() string {
	return "product_sku_tbl"
}

// ProductSkuTblColumns get sql column name.获取数据库列名
var ProductSkuTblColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Gpid      string
	TypeName  string
	TagName   string
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Gpid:      "gpid",
	TypeName:  "type_name",
	TagName:   "tag_name",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	DeletedBy: "deleted_by",
}

/******sql******
CREATE TABLE `product_tbl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增',
  `gpid` varchar(32) NOT NULL COMMENT '商品id',
  `gtid` varchar(255) NOT NULL COMMENT '商品分类id',
  `name` varchar(255) NOT NULL COMMENT '商品名字',
  `price` bigint(20) NOT NULL COMMENT '商品价格(分)',
  `original_price` bigint(20) NOT NULL COMMENT '商品原始价格(分)',
  `dist_amount` bigint(20) DEFAULT NULL COMMENT '分享收益',
  `platform_id` varchar(255) NOT NULL COMMENT '商铺客户id',
  `qty` bigint(20) NOT NULL COMMENT '数量(库存)',
  `shipment_type_id` bigint(20) NOT NULL COMMENT '运输方式类型',
  `shipment_fee` bigint(20) NOT NULL COMMENT '运费',
  `context` text NOT NULL COMMENT '商品详情',
  `volume_weight` bigint(20) NOT NULL DEFAULT '0' COMMENT '体积重',
  `actual_weight` bigint(20) NOT NULL DEFAULT '0' COMMENT '实际重',
  `source_region` varchar(255) DEFAULT '0' COMMENT '货源地名',
  `max_buy_qty` int(11) NOT NULL COMMENT '最大购买数量',
  `is_select` tinyint(1) DEFAULT NULL COMMENT '是否推荐',
  `vaild` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name_id` (`name`,`platform_id`) USING BTREE COMMENT '商品名索引',
  UNIQUE KEY `gpid` (`gpid`) USING BTREE COMMENT '商品索引',
  CONSTRAINT `product_tbl_ibfk_1` FOREIGN KEY (`gpid`) REFERENCES `product_info_tbl` (`gpid`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COMMENT='商品信息'
******sql******/
// ProductTbl 商品信息
type ProductTbl struct {
	ID             int64          `gorm:"primaryKey"`                          // 自增
	Gpid           string         `gorm:"unique"`                              // 商品id
	ProductInfoTbl ProductInfoTbl `gorm:"joinForeignKey:gpid;foreignKey:gpid"` // 商品附加属性
	Gtid           string         // 商品分类id
	Name           string         `gorm:"uniqueIndex:name_id"` // 商品名字
	Price          int64          // 商品价格(分)
	OriginalPrice  int64          // 商品原始价格(分)
	DistAmount     int64          // 分享收益
	PlatformID     string         `gorm:"uniqueIndex:name_id"` // 商铺客户id
	Qty            int64          // 数量(库存)
	ShipmentTypeID int64          // 运输方式类型
	ShipmentFee    int64          // 运费
	Context        string         // 商品详情
	VolumeWeight   int64          `gorm:"default:0"` // 体积重
	ActualWeight   int64          `gorm:"default:0"` // 实际重
	SourceRegion   string         `gorm:"default:0"` // 货源地名
	MaxBuyQty      int            // 最大购买数量
	IsSelect       bool           // 是否推荐
	Vaild          bool           `gorm:"default:1"`  // 是否有效
	CreatedBy      string         `gorm:"default:''"` // 创建者
	CreatedAt      time.Time      // 创建时间
	UpdatedBy      string         `gorm:"default:''"` // 更新者
	UpdatedAt      time.Time      // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductTbl) TableName() string {
	return "product_tbl"
}

// ProductTblColumns get sql column name.获取数据库列名
var ProductTblColumns = struct {
	ID             string
	Gpid           string
	Gtid           string
	Name           string
	Price          string
	OriginalPrice  string
	DistAmount     string
	PlatformID     string
	Qty            string
	ShipmentTypeID string
	ShipmentFee    string
	Context        string
	VolumeWeight   string
	ActualWeight   string
	SourceRegion   string
	MaxBuyQty      string
	IsSelect       string
	Vaild          string
	CreatedBy      string
	CreatedAt      string
	UpdatedBy      string
	UpdatedAt      string
}{
	ID:             "id",
	Gpid:           "gpid",
	Gtid:           "gtid",
	Name:           "name",
	Price:          "price",
	OriginalPrice:  "original_price",
	DistAmount:     "dist_amount",
	PlatformID:     "platform_id",
	Qty:            "qty",
	ShipmentTypeID: "shipment_type_id",
	ShipmentFee:    "shipment_fee",
	Context:        "context",
	VolumeWeight:   "volume_weight",
	ActualWeight:   "actual_weight",
	SourceRegion:   "source_region",
	MaxBuyQty:      "max_buy_qty",
	IsSelect:       "is_select",
	Vaild:          "vaild",
	CreatedBy:      "created_by",
	CreatedAt:      "created_at",
	UpdatedBy:      "updated_by",
	UpdatedAt:      "updated_at",
}

/******sql******
CREATE TABLE `product_type_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gtid` varchar(32) DEFAULT NULL COMMENT '产品类型id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `gtid` (`gtid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='产品类型'
******sql******/
// ProductTypeTbl 产品类型
type ProductTypeTbl struct {
	ID   int    `gorm:"primaryKey"`
	Gtid string `gorm:"unique"` // 产品类型id
}

// TableName get sql table name.获取数据库表名
func (m *ProductTypeTbl) TableName() string {
	return "product_type_tbl"
}

// ProductTypeTblColumns get sql column name.获取数据库列名
var ProductTypeTblColumns = struct {
	ID   string
	Gtid string
}{
	ID:   "id",
	Gtid: "gtid",
}

/******sql******
CREATE TABLE `proposal_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) NOT NULL COMMENT '账单id',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '操作：-1：取消，1：申请售后，',
  `contact` varchar(64) NOT NULL DEFAULT '0' COMMENT '联系方式',
  `remak` varchar(1024) DEFAULT NULL COMMENT '备注',
  `imgs` text COMMENT '图片',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `bill_id` (`bill_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='建议与反馈'
******sql******/
// ProposalTbl 建议与反馈
type ProposalTbl struct {
	ID        int       `gorm:"primaryKey"`
	UserID    string    `gorm:"default:0"`     // 用户id
	BillID    string    `gorm:"index:bill_id"` // 账单id
	Type      int       `gorm:"default:0"`     // 操作：-1：取消，1：申请售后，
	Contact   string    `gorm:"default:0"`     // 联系方式
	Remak     string    // 备注
	Imgs      string    // 图片
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProposalTbl) TableName() string {
	return "proposal_tbl"
}

// ProposalTblColumns get sql column name.获取数据库列名
var ProposalTblColumns = struct {
	ID        string
	UserID    string
	BillID    string
	Type      string
	Contact   string
	Remak     string
	Imgs      string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	BillID:    "bill_id",
	Type:      "type",
	Contact:   "contact",
	Remak:     "remak",
	Imgs:      "imgs",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `shipment_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) NOT NULL COMMENT '账单id',
  `shipment_id` varchar(255) DEFAULT NULL COMMENT '快递单号',
  `post_key` varchar(255) DEFAULT NULL COMMENT '快递id',
  `status` int(11) DEFAULT NULL COMMENT '快递状态  2在途中 3已签收 4 问题件',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `bill_shipment` (`bill_id`,`shipment_id`) USING BTREE,
  KEY `post_key` (`post_key`),
  CONSTRAINT `shipment_tbl_ibfk_1` FOREIGN KEY (`post_key`) REFERENCES `post_tbl` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=169 DEFAULT CHARSET=utf8mb4 COMMENT='用户快递单信息'
******sql******/
// ShipmentTbl 用户快递单信息
type ShipmentTbl struct {
	ID         int       `gorm:"primaryKey"`
	BillID     string    `gorm:"uniqueIndex:bill_shipment"`               // 账单id
	ShipmentID string    `gorm:"uniqueIndex:bill_shipment"`               // 快递单号
	PostKey    string    `gorm:"index:post_key"`                          // 快递id
	PostTbl    PostTbl   `gorm:"joinForeignKey:post_key;foreignKey:code"` // 快递公司信息
	Status     int       // 快递状态  2在途中 3已签收 4 问题件
	CreatedBy  string    `gorm:"default:''"` // 创建者
	CreatedAt  time.Time // 创建时间
	UpdatedBy  string    `gorm:"default:''"` // 更新者
	UpdatedAt  time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ShipmentTbl) TableName() string {
	return "shipment_tbl"
}

// ShipmentTblColumns get sql column name.获取数据库列名
var ShipmentTblColumns = struct {
	ID         string
	BillID     string
	ShipmentID string
	PostKey    string
	Status     string
	CreatedBy  string
	CreatedAt  string
	UpdatedBy  string
	UpdatedAt  string
}{
	ID:         "id",
	BillID:     "bill_id",
	ShipmentID: "shipment_id",
	PostKey:    "post_key",
	Status:     "status",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

/******sql******
CREATE TABLE `user_link_log_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '微信用户唯一标识符',
  `link_open_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '关联的微信用户唯一标识符',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信用户信息'
******sql******/
// UserLinkLogTbl 微信用户信息
type UserLinkLogTbl struct {
	ID         int       `gorm:"primaryKey"`
	OpenID     string    // 微信用户唯一标识符
	LinkOpenID string    // 关联的微信用户唯一标识符
	CreatedBy  string    `gorm:"default:''"` // 创建者
	CreatedAt  time.Time // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *UserLinkLogTbl) TableName() string {
	return "user_link_log_tbl"
}

// UserLinkLogTblColumns get sql column name.获取数据库列名
var UserLinkLogTblColumns = struct {
	ID         string
	OpenID     string
	LinkOpenID string
	CreatedBy  string
	CreatedAt  string
}{
	ID:         "id",
	OpenID:     "open_id",
	LinkOpenID: "link_open_id",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
}

/******sql******
CREATE TABLE `user_link_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '被邀请人',
  `parent_user_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '邀请人',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `openid` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信用户信息'
******sql******/
// UserLinkTbl 微信用户信息
type UserLinkTbl struct {
	ID           int       `gorm:"primaryKey"`
	UserID       string    `gorm:"unique"` // 被邀请人
	ParentUserID string    // 邀请人
	CreatedBy    string    `gorm:"default:''"` // 创建者
	CreatedAt    time.Time // 创建时间
	UpdatedBy    string    `gorm:"default:''"` // 更新者
	UpdatedAt    time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserLinkTbl) TableName() string {
	return "user_link_tbl"
}

// UserLinkTblColumns get sql column name.获取数据库列名
var UserLinkTblColumns = struct {
	ID           string
	UserID       string
	ParentUserID string
	CreatedBy    string
	CreatedAt    string
	UpdatedBy    string
	UpdatedAt    string
}{
	ID:           "id",
	UserID:       "user_id",
	ParentUserID: "parent_user_id",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	UpdatedBy:    "updated_by",
	UpdatedAt:    "updated_at",
}

/******sql******
CREATE TABLE `user_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '微信用户唯一标识符',
  `vip` tinyint(1) DEFAULT NULL COMMENT '是否vip',
  `dist_vip` tinyint(1) DEFAULT NULL COMMENT '是否分销vip',
  `balance` bigint(20) DEFAULT NULL COMMENT '用户余额',
  `points` bigint(20) DEFAULT NULL COMMENT '用户积分',
  `created_by` varchar(255) DEFAULT '' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) DEFAULT '' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `openid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COMMENT='微信用户信息'
******sql******/
// UserTbl 微信用户信息
type UserTbl struct {
	ID        int       `gorm:"primaryKey"`
	UserID    string    `gorm:"unique"` // 微信用户唯一标识符
	Vip       bool      // 是否vip
	DistVip   bool      // 是否分销vip
	Balance   int64     // 用户余额
	Points    int64     // 用户积分
	CreatedBy string    `gorm:"default:''"` // 创建者
	CreatedAt time.Time // 创建时间
	UpdatedBy string    `gorm:"default:''"` // 更新者
	UpdatedAt time.Time // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserTbl) TableName() string {
	return "user_tbl"
}

// UserTblColumns get sql column name.获取数据库列名
var UserTblColumns = struct {
	ID        string
	UserID    string
	Vip       string
	DistVip   string
	Balance   string
	Points    string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Vip:       "vip",
	DistVip:   "dist_vip",
	Balance:   "balance",
	Points:    "points",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

/******sql******
CREATE TABLE `wx_userinfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '微信用户唯一标识符\n\n微信用户唯一标识符',
  `nick_name` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户昵称',
  `avatar_url` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。',
  `gender` int(11) DEFAULT NULL COMMENT '用户的性别，值为1时是男性，值为2时是女性，值为0时是未知',
  `city` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户所在城市',
  `province` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户所在省份',
  `country` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户所在国家',
  `language` varbinary(16) DEFAULT NULL COMMENT '用户的语言，简体中文为zh_CN',
  `phone` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户绑定的手机号',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `openid` (`openid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1212 DEFAULT CHARSET=utf8mb4 COMMENT='微信用户信息'
******sql******/
// WxUserinfo 微信用户信息
type WxUserinfo struct {
	ID        int    `gorm:"primaryKey"`
	Openid    string `gorm:"unique"` // 微信用户唯一标识符,,微信用户唯一标识符
	NickName  string // 用户昵称
	AvatarURL string // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Gender    int    // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City      string // 用户所在城市
	Province  string // 用户所在省份
	Country   string // 用户所在国家
	Language  []byte // 用户的语言，简体中文为zh_CN
	Phone     string // 用户绑定的手机号
}

// TableName get sql table name.获取数据库表名
func (m *WxUserinfo) TableName() string {
	return "wx_userinfo"
}

// WxUserinfoColumns get sql column name.获取数据库列名
var WxUserinfoColumns = struct {
	ID        string
	Openid    string
	NickName  string
	AvatarURL string
	Gender    string
	City      string
	Province  string
	Country   string
	Language  string
	Phone     string
}{
	ID:        "id",
	Openid:    "openid",
	NickName:  "nick_name",
	AvatarURL: "avatar_url",
	Gender:    "gender",
	City:      "city",
	Province:  "province",
	Country:   "country",
	Language:  "language",
	Phone:     "phone",
}
