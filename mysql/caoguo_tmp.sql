/*
 Navicat Premium Data Transfer

 Source Server         : 175.24.103.30
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 175.24.103.30:3306
 Source Schema         : caoguo_tmp

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 26/10/2021 14:26:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address_tbl
-- ----------------------------
DROP TABLE IF EXISTS `address_tbl`;
CREATE TABLE `address_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '[@gormt default:\'123456\';]手机号',
  `address_name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址名称',
  `address` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '详细地址',
  `area` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单元门牌号',
  `default` tinyint(1) NULL DEFAULT NULL COMMENT '是否是默认值',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户地址列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of address_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for bill_address_tbl
-- ----------------------------
DROP TABLE IF EXISTS `bill_address_tbl`;
CREATE TABLE `bill_address_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addr_id` bigint(20) NULL DEFAULT NULL COMMENT '地址id',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账单id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `address_name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址名称',
  `address` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '详细地址',
  `area` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单元门牌号',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `bill_id`(`bill_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 156 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '派送信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill_address_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for bill_coupon_tbl
-- ----------------------------
DROP TABLE IF EXISTS `bill_coupon_tbl`;
CREATE TABLE `bill_coupon_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账单id',
  `coupon_id` bigint(20) NOT NULL COMMENT '我领取的优惠券id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '优惠券名字',
  `denom` int(11) NULL DEFAULT NULL COMMENT '面额',
  `coupon_type` int(11) NULL DEFAULT NULL COMMENT '优惠券类型(1:全场，2:单个商品)',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `bill_id`(`bill_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '优惠券列表池' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill_coupon_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for bill_deal_log_tbl
-- ----------------------------
DROP TABLE IF EXISTS `bill_deal_log_tbl`;
CREATE TABLE `bill_deal_log_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账单id',
  `type` int(11) NOT NULL DEFAULT 0 COMMENT '操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈',
  `contact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '联系方式',
  `remak` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `imgs` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图片',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `bill_id`(`bill_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '账单操作日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill_deal_log_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for bill_order_tbl
-- ----------------------------
DROP TABLE IF EXISTS `bill_order_tbl`;
CREATE TABLE `bill_order_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账单id',
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品名字',
  `price` bigint(20) NOT NULL COMMENT '商品价格(分)',
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标信息',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `number` int(11) NOT NULL COMMENT '数量',
  `dist_amount` bigint(20) NULL DEFAULT NULL COMMENT '分享收益',
  `sku_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性列表',
  `sku_arrt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性描述',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_gpid`(`gpid`, `user_id`, `sku_list`, `bill_id`) USING BTREE,
  INDEX `bill_id`(`bill_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 187 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '账单订单列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill_order_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for bill_refund_tbl
-- ----------------------------
DROP TABLE IF EXISTS `bill_refund_tbl`;
CREATE TABLE `bill_refund_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账单id',
  `refund_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '退款账单号',
  `refund_fee` bigint(20) NOT NULL DEFAULT 0 COMMENT '退款金额',
  `desc` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '订单备注',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `bill_id`(`bill_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户退款信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill_refund_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for bill_tbl
-- ----------------------------
DROP TABLE IF EXISTS `bill_tbl`;
CREATE TABLE `bill_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账单id',
  `pay_platform` int(11) NOT NULL DEFAULT 0 COMMENT '支付类型(1:微信支付)',
  `pay_amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '支付金额',
  `status` int(11) NOT NULL DEFAULT 0 COMMENT '支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)',
  `desc` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '订单备注',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `bill_id`(`bill_id`) USING BTREE,
  CONSTRAINT `bill_tbl_ibfk_1` FOREIGN KEY (`bill_id`) REFERENCES `bill_address_tbl` (`bill_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `bill_tbl_ibfk_2` FOREIGN KEY (`bill_id`) REFERENCES `bill_order_tbl` (`bill_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户账单信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for cart_tbl
-- ----------------------------
DROP TABLE IF EXISTS `cart_tbl`;
CREATE TABLE `cart_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `user_type` int(11) NULL DEFAULT NULL COMMENT '用户类型(1:微信)',
  `number` int(11) NULL DEFAULT NULL COMMENT '数量',
  `sku_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性列表',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_gpid`(`gpid`, `user_id`, `sku_list`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '购物车列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cart_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for cart_tmp_tbl
-- ----------------------------
DROP TABLE IF EXISTS `cart_tmp_tbl`;
CREATE TABLE `cart_tmp_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `user_type` int(11) NULL DEFAULT NULL COMMENT '用户类型(1:微信)',
  `number` int(11) NULL DEFAULT NULL COMMENT '数量',
  `sku_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '属性列表',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '虚拟购物车列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cart_tmp_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for coupon_my_tbl
-- ----------------------------
DROP TABLE IF EXISTS `coupon_my_tbl`;
CREATE TABLE `coupon_my_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `coupon_id` bigint(20) NOT NULL COMMENT '优惠券id',
  `expires_time` timestamp(0) NULL DEFAULT NULL COMMENT '过期时间',
  `status` int(11) NULL DEFAULT NULL COMMENT '当前优惠券状态(1:未使用(有效),2:已使用,-1:已过期)',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `coupon_id`(`coupon_id`) USING BTREE,
  CONSTRAINT `coupon_my_tbl_ibfk_1` FOREIGN KEY (`coupon_id`) REFERENCES `coupon_tbl` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '优惠券列表池' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of coupon_my_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for coupon_tbl
-- ----------------------------
DROP TABLE IF EXISTS `coupon_tbl`;
CREATE TABLE `coupon_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `group_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分组名',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '优惠券名字',
  `validity` int(11) NULL DEFAULT NULL COMMENT '有效期(天数)',
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品优惠券',
  `denom` int(11) NULL DEFAULT NULL COMMENT '面额',
  `coupon_type` int(11) NULL DEFAULT NULL COMMENT '优惠券类型(1:全场，2:单个商品)',
  `great_money` int(11) NULL DEFAULT NULL COMMENT '满多少可用',
  `describle` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '优惠券描述',
  `vaild` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否有效',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '优惠券列表池' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of coupon_tbl
-- ----------------------------
INSERT INTO `coupon_tbl` VALUES (1, '新人优惠券', '新用户5元优惠券', 30, '', 500, 1, 500, '新用户专享优惠券', 0, '', NULL, '', '2020-10-07 13:50:19');
INSERT INTO `coupon_tbl` VALUES (2, '新人优惠券', '新用户10元优惠券', 30, '', 1000, 1, 10000, '新用户专享优惠券', 0, '', NULL, '', '2020-11-25 18:16:41');
INSERT INTO `coupon_tbl` VALUES (3, '商品优惠券', '猕猴桃限时5元优惠券', 30, '1280145890979352576', 500, 2, 5000, '特殊商品优惠券', 1, '', NULL, '', '2020-09-30 15:38:25');
INSERT INTO `coupon_tbl` VALUES (4, '商品优惠券', '猕猴桃限时12优惠券', 30, '1280145890979352576', 1200, 2, 10000, '特殊商品优惠券', 1, '', NULL, '', '2020-10-19 18:29:58');
INSERT INTO `coupon_tbl` VALUES (5, '商品优惠券', '猕猴桃限时30优惠券', 30, '1280145890979352576', 3000, 2, 20000, '特殊商品优惠券', 1, '', NULL, '', '2020-10-08 12:34:04');
INSERT INTO `coupon_tbl` VALUES (6, '类别优惠券', '全场水果优惠券', 30, NULL, 999, 2, 1000, '特殊类别优惠券', 0, '', NULL, '', NULL);
INSERT INTO `coupon_tbl` VALUES (7, '运费优惠券', '全场运费优惠券', 30, NULL, 999, 3, 1000, '运费优惠券', 0, '', NULL, '', NULL);

-- ----------------------------
-- Table structure for dist_order_tbl
-- ----------------------------
DROP TABLE IF EXISTS `dist_order_tbl`;
CREATE TABLE `dist_order_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账单id',
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分销给的账号',
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品名字',
  `price` bigint(20) NOT NULL COMMENT '分销金额',
  `level` int(255) NULL DEFAULT NULL COMMENT '分销等级',
  `total_price` bigint(20) NULL DEFAULT NULL COMMENT '总分销金额',
  `status` int(11) NULL DEFAULT NULL COMMENT '状态(-1:失效,1:待确认,2:可提现，3:已提现)',
  `number` int(11) NULL DEFAULT NULL COMMENT '商品数量',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `bill_user_gpid_id`(`bill_id`, `user_id`, `gpid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分销详细账单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dist_order_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for dist_userinfo_tbl
-- ----------------------------
DROP TABLE IF EXISTS `dist_userinfo_tbl`;
CREATE TABLE `dist_userinfo_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '子节点',
  `parnt_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '父节点',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `userid`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信用户信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dist_userinfo_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for email_notify_tbl
-- ----------------------------
DROP TABLE IF EXISTS `email_notify_tbl`;
CREATE TABLE `email_notify_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `email` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱号',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `shipment_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '运输信息明细' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of email_notify_tbl
-- ----------------------------
INSERT INTO `email_notify_tbl` VALUES (1, NULL, '346944475@qq.com', '', '2020-08-28 11:55:00', '', NULL);
INSERT INTO `email_notify_tbl` VALUES (2, NULL, '359208695@qq.com', '', '2020-08-28 11:58:03', '', NULL);
INSERT INTO `email_notify_tbl` VALUES (3, '13795896867', '545308274@qq.com', '', '2020-09-15 18:31:52', '', NULL);
INSERT INTO `email_notify_tbl` VALUES (4, '787940823', '787940823@qq.com', '', '2020-09-16 17:43:20', '', NULL);
INSERT INTO `email_notify_tbl` VALUES (5, '13520697936', 'zhanghao@iqcheng.com', '', '2020-11-25 18:39:23', '', NULL);

-- ----------------------------
-- Table structure for favorite_tbl
-- ----------------------------
DROP TABLE IF EXISTS `favorite_tbl`;
CREATE TABLE `favorite_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `user_type` int(11) NOT NULL COMMENT '用户类型(1:微信)',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_gpid`(`gpid`, `user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of favorite_tbl
-- ----------------------------
INSERT INTO `favorite_tbl` VALUES (6, '1280145890979352576', 'oXebs4jQXYNezUHMBQLgGPUlnw0M', 1, '', '2020-09-03 02:41:51', '', '2020-09-03 02:41:51');
INSERT INTO `favorite_tbl` VALUES (7, '1280146196534398976', 'oXebs4jQXYNezUHMBQLgGPUlnw0M', 1, '', '2020-09-03 02:42:00', '', '2020-09-03 02:42:00');
INSERT INTO `favorite_tbl` VALUES (8, '1280146196534398976', 'oXebs4pxYa5VGoCY3K_uhgABE10c', 1, '', '2020-09-04 14:23:56', '', '2020-09-04 14:23:56');
INSERT INTO `favorite_tbl` VALUES (9, '1280145890979352576', 'oXebs4inTXsT6H9tNfR4OnpZcqG8', 1, '', '2020-09-04 15:06:32', '', '2020-09-04 15:06:32');
INSERT INTO `favorite_tbl` VALUES (11, '1280146196534398976', 'oXebs4jVLgvk486pQSlm04Ja_Qog', 1, '', '2020-09-05 00:19:56', '', '2020-09-05 00:19:56');
INSERT INTO `favorite_tbl` VALUES (12, '1280145890979352576', 'oXebs4slYHHeDtvlCjkmn-AkNCqE', 1, '', '2020-09-05 10:00:05', '', '2020-09-05 10:00:05');
INSERT INTO `favorite_tbl` VALUES (13, '1280146196534398976', 'oXebs4slYHHeDtvlCjkmn-AkNCqE', 1, '', '2020-09-05 10:00:14', '', '2020-09-05 10:00:14');
INSERT INTO `favorite_tbl` VALUES (15, '1280146196534398976', 'oXebs4sgZVyjzzznqClpHjBuIV7U', 1, '', '2020-09-06 23:47:41', '', '2020-09-06 23:47:41');
INSERT INTO `favorite_tbl` VALUES (16, '1280145890979352576', 'oXebs4lMMJz4_gdzcBSM5_m97cCM', 1, '', '2020-09-07 00:31:37', '', '2020-09-07 00:31:37');
INSERT INTO `favorite_tbl` VALUES (17, '1280146196534398976', 'oXebs4lMMJz4_gdzcBSM5_m97cCM', 1, '', '2020-09-07 00:31:46', '', '2020-09-07 00:31:46');
INSERT INTO `favorite_tbl` VALUES (18, '1280145890979352576', 'oXebs4t0GxU93i319WkKEVxB5AyU', 1, '', '2020-09-08 02:06:00', '', '2020-09-08 02:06:00');
INSERT INTO `favorite_tbl` VALUES (19, '1303017018819088384', 'oXebs4i_j7ps4SRIiEIxiD-QsVnM', 1, '', '2020-09-08 11:10:58', '', '2020-09-08 11:10:58');
INSERT INTO `favorite_tbl` VALUES (22, '1280145890979352576', 'oXebs4jVLgvk486pQSlm04Ja_Qog', 1, '', '2020-09-10 00:46:58', '', '2020-09-10 00:46:58');
INSERT INTO `favorite_tbl` VALUES (23, '1280145890979352576', 'oXebs4qOVVUf8SkuDFjEItJk9fEQ', 1, '', '2020-09-10 10:41:48', '', '2020-09-10 10:41:48');
INSERT INTO `favorite_tbl` VALUES (26, '1280145890979352576', 'oXebs4sgZVyjzzznqClpHjBuIV7U', 1, '', '2020-09-12 00:46:00', '', '2020-09-12 00:46:00');
INSERT INTO `favorite_tbl` VALUES (28, '1280145890979352576', 'oXebs4mR1hZEj6XTAWCISS5gYSIo', 1, '', '2020-09-23 15:34:58', '', '2020-09-23 15:34:58');
INSERT INTO `favorite_tbl` VALUES (29, '1331221900822581248', 'oXebs4gacVDm5wGLjaKAN4WiIY1o', 1, '', '2020-12-07 21:10:29', '', '2020-12-07 21:10:29');

-- ----------------------------
-- Table structure for logistics_tbl
-- ----------------------------
DROP TABLE IF EXISTS `logistics_tbl`;
CREATE TABLE `logistics_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shipment_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递单号',
  `times` datetime(0) NULL DEFAULT NULL COMMENT '时间',
  `context` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述信息',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `shipment_id`(`shipment_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1698 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '运输信息明细' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of logistics_tbl
-- ----------------------------
INSERT INTO `logistics_tbl` VALUES (2, 'YT4765645114700', '2020-09-06 08:13:25', '添加快递单', 'admin', '2020-09-06 08:13:25', '', '2020-09-06 08:13:25');
INSERT INTO `logistics_tbl` VALUES (3, 'YT4766119173574', '2020-09-06 10:48:20', '添加快递单', 'admin', '2020-09-06 10:48:20', '', '2020-09-06 10:48:20');
INSERT INTO `logistics_tbl` VALUES (4, 'JDVD00934676229-1-1-', '2020-09-06 11:30:23', '添加快递单', 'admin', '2020-09-06 11:30:23', '', '2020-09-06 11:30:23');
INSERT INTO `logistics_tbl` VALUES (5, 'JDVD00934692194', '2020-09-06 11:31:17', '添加快递单', 'admin', '2020-09-06 11:31:17', '', '2020-09-06 11:31:17');
INSERT INTO `logistics_tbl` VALUES (6, 'JDVD00934679583-1-1-', '2020-09-06 11:36:03', '添加快递单', 'admin', '2020-09-06 11:36:03', '', '2020-09-06 11:36:03');
INSERT INTO `logistics_tbl` VALUES (7, 'JDVD00934685489-1-1-', '2020-09-06 11:36:44', '添加快递单', 'admin', '2020-09-06 11:36:44', '', '2020-09-06 11:36:44');
INSERT INTO `logistics_tbl` VALUES (8, 'JDVD00934659318-1-1-', '2020-09-06 11:38:58', '添加快递单', 'admin', '2020-09-06 11:38:58', '', '2020-09-06 11:38:58');
INSERT INTO `logistics_tbl` VALUES (9, 'JDVD00934683078', '2020-09-06 11:39:44', '添加快递单', 'admin', '2020-09-06 11:39:44', '', '2020-09-06 11:39:44');
INSERT INTO `logistics_tbl` VALUES (10, 'JDVD00936221520', '2020-09-06 12:25:17', '添加快递单', 'admin', '2020-09-06 12:25:17', '', '2020-09-06 12:25:17');
INSERT INTO `logistics_tbl` VALUES (11, 'JDVD00936221539', '2020-09-06 12:26:34', '添加快递单', 'admin', '2020-09-06 12:26:34', '', '2020-09-06 12:26:34');
INSERT INTO `logistics_tbl` VALUES (12, 'JDVD00936215570', '2020-09-06 12:27:42', '添加快递单', 'admin', '2020-09-06 12:27:42', '', '2020-09-06 12:27:42');
INSERT INTO `logistics_tbl` VALUES (13, 'JDVD00936201478', '2020-09-06 12:35:33', '添加快递单', 'admin', '2020-09-06 12:35:33', '', '2020-09-06 12:35:33');
INSERT INTO `logistics_tbl` VALUES (14, 'JDVD00936221512', '2020-09-06 12:37:26', '添加快递单', 'admin', '2020-09-06 12:37:26', '', '2020-09-06 12:37:26');
INSERT INTO `logistics_tbl` VALUES (15, 'JDVD00936215561', '2020-09-06 12:41:13', '添加快递单', 'admin', '2020-09-06 12:41:13', '', '2020-09-06 12:41:13');
INSERT INTO `logistics_tbl` VALUES (16, 'JDVD00936213420', '2020-09-06 12:42:33', '添加快递单', 'admin', '2020-09-06 12:42:33', '', '2020-09-06 12:42:33');
INSERT INTO `logistics_tbl` VALUES (17, 'JDVD00936215561', '2020-09-06 11:53:38', '揽收任务已分配给邓洁', '', '2020-09-06 15:50:17', '', '2020-09-06 15:50:17');
INSERT INTO `logistics_tbl` VALUES (18, 'JDVD00936215561', '2020-09-06 14:48:43', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 15:50:17', '', '2020-09-06 15:50:17');
INSERT INTO `logistics_tbl` VALUES (19, 'JDVD00936215561', '2020-09-06 14:48:43', '京东快递 已收取快件', '', '2020-09-06 15:50:17', '', '2020-09-06 15:50:17');
INSERT INTO `logistics_tbl` VALUES (20, 'JDVD00936213420', '2020-09-06 11:53:39', '揽收任务已分配给邓洁', '', '2020-09-06 15:50:36', '', '2020-09-06 15:50:36');
INSERT INTO `logistics_tbl` VALUES (21, 'JDVD00936213420', '2020-09-06 14:48:38', '京东快递 已收取快件', '', '2020-09-06 15:50:36', '', '2020-09-06 15:50:36');
INSERT INTO `logistics_tbl` VALUES (22, 'JDVD00936213420', '2020-09-06 14:48:38', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 15:50:36', '', '2020-09-06 15:50:36');
INSERT INTO `logistics_tbl` VALUES (23, 'JDVD00934683078', '2020-09-05 17:04:50', '揽收任务已分配给邓洁', '', '2020-09-06 15:57:50', '', '2020-09-06 15:57:50');
INSERT INTO `logistics_tbl` VALUES (24, 'JDVD00934683078', '2020-09-06 14:45:41', '京东快递 已收取快件', '', '2020-09-06 15:57:50', '', '2020-09-06 15:57:50');
INSERT INTO `logistics_tbl` VALUES (25, 'JDVD00934683078', '2020-09-06 14:45:41', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 15:57:50', '', '2020-09-06 15:57:50');
INSERT INTO `logistics_tbl` VALUES (26, 'JDVD00934692194', '2020-09-05 17:04:51', '揽收任务已分配给邓洁', '', '2020-09-06 15:58:03', '', '2020-09-06 15:58:03');
INSERT INTO `logistics_tbl` VALUES (27, 'JDVD00934692194', '2020-09-06 14:45:51', '京东快递 已收取快件', '', '2020-09-06 15:58:03', '', '2020-09-06 15:58:03');
INSERT INTO `logistics_tbl` VALUES (28, 'JDVD00934692194', '2020-09-06 14:45:51', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 15:58:03', '', '2020-09-06 15:58:03');
INSERT INTO `logistics_tbl` VALUES (29, 'YT4767034499495', '2020-09-06 16:49:02', '添加快递单', 'admin', '2020-09-06 16:49:02', '', '2020-09-06 16:49:02');
INSERT INTO `logistics_tbl` VALUES (30, 'YT4767033731540', '2020-09-06 16:49:18', '添加快递单', 'admin', '2020-09-06 16:49:18', '', '2020-09-06 16:49:18');
INSERT INTO `logistics_tbl` VALUES (31, 'YT4767033732744', '2020-09-06 16:49:43', '添加快递单', 'admin', '2020-09-06 16:49:43', '', '2020-09-06 16:49:43');
INSERT INTO `logistics_tbl` VALUES (32, 'YT4767030836799', '2020-09-06 16:49:54', '添加快递单', 'admin', '2020-09-06 16:49:54', '', '2020-09-06 16:49:54');
INSERT INTO `logistics_tbl` VALUES (33, '773054585661279', '2020-09-06 17:53:11', '添加快递单', 'admin', '2020-09-06 17:53:11', '', '2020-09-06 17:53:11');
INSERT INTO `logistics_tbl` VALUES (35, '773054585245568', '2020-09-06 18:39:28', '添加快递单', 'admin', '2020-09-06 18:39:28', '', '2020-09-06 18:39:28');
INSERT INTO `logistics_tbl` VALUES (36, '773054585245568', '2020-09-06 17:50:36', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 18:39:35', '', '2020-09-06 18:39:35');
INSERT INTO `logistics_tbl` VALUES (37, 'YT4767034499495', '2020-09-06 17:07:52', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-06 18:45:29', '', '2020-09-06 18:45:29');
INSERT INTO `logistics_tbl` VALUES (38, 'YT4767033731540', '2020-09-06 17:07:57', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-06 18:45:59', '', '2020-09-06 18:45:59');
INSERT INTO `logistics_tbl` VALUES (39, 'YT4767030836799', '2020-09-06 17:07:47', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-06 18:46:05', '', '2020-09-06 18:46:05');
INSERT INTO `logistics_tbl` VALUES (40, 'YT4767033732744', '2020-09-06 17:08:00', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-06 18:46:10', '', '2020-09-06 18:46:10');
INSERT INTO `logistics_tbl` VALUES (41, 'YT4765645114700', '2020-09-06 17:02:41', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-06 18:46:37', '', '2020-09-06 18:46:37');
INSERT INTO `logistics_tbl` VALUES (42, 'JDVD00934683078', '2020-09-06 17:44:32', '您的快件已到达【广元苍溪营业部】', '', '2020-09-06 18:46:54', '', '2020-09-06 18:46:54');
INSERT INTO `logistics_tbl` VALUES (43, 'JDVD00934683078', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-06 18:46:54', '', '2020-09-06 18:46:54');
INSERT INTO `logistics_tbl` VALUES (44, 'JDVD00934683078', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-06 18:46:54', '', '2020-09-06 18:46:54');
INSERT INTO `logistics_tbl` VALUES (45, 'YT4766119173574', '2020-09-06 17:02:38', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-06 18:47:02', '', '2020-09-06 18:47:02');
INSERT INTO `logistics_tbl` VALUES (46, '773054585661279', '2020-09-06 17:50:08', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 18:47:07', '', '2020-09-06 18:47:07');
INSERT INTO `logistics_tbl` VALUES (47, '773054585661279', '2020-09-06 18:34:41', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 18:47:07', '', '2020-09-06 18:47:07');
INSERT INTO `logistics_tbl` VALUES (48, 'JDVD00936221539', '2020-09-06 11:53:40', '揽收任务已分配给邓洁', '', '2020-09-06 18:47:21', '', '2020-09-06 18:47:21');
INSERT INTO `logistics_tbl` VALUES (49, 'JDVD00936221539', '2020-09-06 14:47:49', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 18:47:21', '', '2020-09-06 18:47:21');
INSERT INTO `logistics_tbl` VALUES (50, 'JDVD00936221539', '2020-09-06 14:47:49', '京东快递 已收取快件', '', '2020-09-06 18:47:21', '', '2020-09-06 18:47:21');
INSERT INTO `logistics_tbl` VALUES (51, 'JDVD00936221520', '2020-09-06 11:53:39', '揽收任务已分配给邓洁', '', '2020-09-06 18:47:27', '', '2020-09-06 18:47:27');
INSERT INTO `logistics_tbl` VALUES (52, 'JDVD00936221520', '2020-09-06 14:47:53', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 18:47:27', '', '2020-09-06 18:47:27');
INSERT INTO `logistics_tbl` VALUES (53, 'JDVD00936221520', '2020-09-06 14:47:53', '京东快递 已收取快件', '', '2020-09-06 18:47:27', '', '2020-09-06 18:47:27');
INSERT INTO `logistics_tbl` VALUES (54, 'JDVD00936221520', '2020-09-06 17:44:29', '您的快件已到达【广元苍溪营业部】', '', '2020-09-06 18:47:27', '', '2020-09-06 18:47:27');
INSERT INTO `logistics_tbl` VALUES (55, 'JDVD00936221520', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-06 18:47:27', '', '2020-09-06 18:47:27');
INSERT INTO `logistics_tbl` VALUES (56, 'JDVD00936221520', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-06 18:47:27', '', '2020-09-06 18:47:27');
INSERT INTO `logistics_tbl` VALUES (57, 'JDVD00936221512', '2020-09-06 11:53:39', '揽收任务已分配给邓洁', '', '2020-09-06 18:47:37', '', '2020-09-06 18:47:37');
INSERT INTO `logistics_tbl` VALUES (58, 'JDVD00936221512', '2020-09-06 14:48:26', '京东快递 已收取快件', '', '2020-09-06 18:47:37', '', '2020-09-06 18:47:37');
INSERT INTO `logistics_tbl` VALUES (59, 'JDVD00936221512', '2020-09-06 14:48:26', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 18:47:37', '', '2020-09-06 18:47:37');
INSERT INTO `logistics_tbl` VALUES (60, 'JDVD00936221512', '2020-09-06 17:44:35', '您的快件已到达【广元苍溪营业部】', '', '2020-09-06 18:47:37', '', '2020-09-06 18:47:37');
INSERT INTO `logistics_tbl` VALUES (61, 'JDVD00936221512', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-06 18:47:37', '', '2020-09-06 18:47:37');
INSERT INTO `logistics_tbl` VALUES (62, 'JDVD00936221512', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-06 18:47:37', '', '2020-09-06 18:47:37');
INSERT INTO `logistics_tbl` VALUES (63, 'JDVD00936201478', '2020-09-06 11:53:39', '揽收任务已分配给邓洁', '', '2020-09-06 18:47:49', '', '2020-09-06 18:47:49');
INSERT INTO `logistics_tbl` VALUES (64, 'JDVD00936201478', '2020-09-06 14:47:43', '京东快递 已收取快件', '', '2020-09-06 18:47:49', '', '2020-09-06 18:47:49');
INSERT INTO `logistics_tbl` VALUES (65, 'JDVD00936201478', '2020-09-06 14:47:43', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 18:47:49', '', '2020-09-06 18:47:49');
INSERT INTO `logistics_tbl` VALUES (66, 'JDVD00936201478', '2020-09-06 17:44:28', '您的快件已到达【广元苍溪营业部】', '', '2020-09-06 18:47:49', '', '2020-09-06 18:47:49');
INSERT INTO `logistics_tbl` VALUES (67, 'JDVD00936201478', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-06 18:47:49', '', '2020-09-06 18:47:49');
INSERT INTO `logistics_tbl` VALUES (68, 'JDVD00936201478', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-06 18:47:49', '', '2020-09-06 18:47:49');
INSERT INTO `logistics_tbl` VALUES (69, '773054585660092', '2020-09-06 23:29:01', '添加快递单', 'admin', '2020-09-06 23:29:01', '', '2020-09-06 23:29:01');
INSERT INTO `logistics_tbl` VALUES (70, '773054585660134', '2020-09-06 23:29:29', '添加快递单', 'admin', '2020-09-06 23:29:29', '', '2020-09-06 23:29:29');
INSERT INTO `logistics_tbl` VALUES (71, '773054585733964', '2020-09-06 23:30:07', '添加快递单', 'admin', '2020-09-06 23:30:07', '', '2020-09-06 23:30:07');
INSERT INTO `logistics_tbl` VALUES (72, '773054583991096', '2020-09-06 23:30:27', '添加快递单', 'admin', '2020-09-06 23:30:27', '', '2020-09-06 23:30:27');
INSERT INTO `logistics_tbl` VALUES (73, '773054585812874', '2020-09-06 23:31:58', '添加快递单', 'admin', '2020-09-06 23:31:58', '', '2020-09-06 23:31:58');
INSERT INTO `logistics_tbl` VALUES (74, '773054585660663', '2020-09-06 23:32:17', '添加快递单', 'admin', '2020-09-06 23:32:17', '', '2020-09-06 23:32:17');
INSERT INTO `logistics_tbl` VALUES (75, '773054585734185', '2020-09-06 23:32:33', '添加快递单', 'admin', '2020-09-06 23:32:33', '', '2020-09-06 23:32:33');
INSERT INTO `logistics_tbl` VALUES (76, '773054585245386', '2020-09-06 23:33:31', '添加快递单', 'admin', '2020-09-06 23:33:31', '', '2020-09-06 23:33:31');
INSERT INTO `logistics_tbl` VALUES (77, '773054585582825', '2020-09-06 23:33:58', '添加快递单', 'admin', '2020-09-06 23:33:58', '', '2020-09-06 23:33:58');
INSERT INTO `logistics_tbl` VALUES (78, '773054583991235', '2020-09-06 23:35:12', '添加快递单', 'admin', '2020-09-06 23:35:12', '', '2020-09-06 23:35:12');
INSERT INTO `logistics_tbl` VALUES (79, '773054585583004', '2020-09-06 23:37:05', '添加快递单', 'admin', '2020-09-06 23:37:05', '', '2020-09-06 23:37:05');
INSERT INTO `logistics_tbl` VALUES (80, '773054585734702', '2020-09-06 23:38:01', '添加快递单', 'admin', '2020-09-06 23:38:01', '', '2020-09-06 23:38:01');
INSERT INTO `logistics_tbl` VALUES (81, '773054585734878', '2020-09-06 23:40:17', '添加快递单', 'admin', '2020-09-06 23:40:17', '', '2020-09-06 23:40:17');
INSERT INTO `logistics_tbl` VALUES (82, '773054585734933', '2020-09-06 23:40:49', '添加快递单', 'admin', '2020-09-06 23:40:49', '', '2020-09-06 23:40:49');
INSERT INTO `logistics_tbl` VALUES (83, '773054585510778', '2020-09-06 23:41:35', '添加快递单', 'admin', '2020-09-06 23:41:35', '', '2020-09-06 23:41:35');
INSERT INTO `logistics_tbl` VALUES (84, '773054585661840', '2020-09-06 23:42:22', '添加快递单', 'admin', '2020-09-06 23:42:22', '', '2020-09-06 23:42:22');
INSERT INTO `logistics_tbl` VALUES (85, '773054585813745', '2020-09-06 23:42:54', '添加快递单', 'admin', '2020-09-06 23:42:54', '', '2020-09-06 23:42:54');
INSERT INTO `logistics_tbl` VALUES (86, '773054584544493', '2020-09-06 23:43:09', '添加快递单', 'admin', '2020-09-06 23:43:09', '', '2020-09-06 23:43:09');
INSERT INTO `logistics_tbl` VALUES (87, '773054583991345', '2020-09-06 23:43:21', '添加快递单', 'admin', '2020-09-06 23:43:21', '', '2020-09-06 23:43:21');
INSERT INTO `logistics_tbl` VALUES (88, 'JDVD00936221520', '2020-09-06 22:15:24', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-06 23:49:09', '', '2020-09-06 23:49:09');
INSERT INTO `logistics_tbl` VALUES (89, 'JDVD00936221520', '2020-09-06 22:21:26', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-06 23:49:09', '', '2020-09-06 23:49:09');
INSERT INTO `logistics_tbl` VALUES (90, 'JDVD00936221520', '2020-09-06 23:05:51', '您的快件由【成都祥福分拣中心】准备发往【南充分拣场】', '', '2020-09-06 23:49:09', '', '2020-09-06 23:49:09');
INSERT INTO `logistics_tbl` VALUES (91, 'JDVD00936215570', '2020-09-06 11:53:40', '揽收任务已分配给邓洁', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (92, 'JDVD00936215570', '2020-09-06 14:47:46', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (93, 'JDVD00936215570', '2020-09-06 14:47:46', '京东快递 已收取快件', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (94, 'JDVD00936215570', '2020-09-06 17:44:29', '您的快件已到达【广元苍溪营业部】', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (95, 'JDVD00936215570', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (96, 'JDVD00936215570', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (97, 'JDVD00936215570', '2020-09-06 22:15:14', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (98, 'JDVD00936215570', '2020-09-06 22:19:18', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (99, 'JDVD00936215570', '2020-09-06 22:19:23', '您的快件由【成都祥福分拣中心】准备发往【成都青白江分拣中心】', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (100, 'JDVD00936215570', '2020-09-06 23:12:38', '您的快件已发车', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (101, 'JDVD00936201478', '2020-09-06 22:15:37', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (102, 'JDVD00936201478', '2020-09-06 22:20:02', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (103, 'JDVD00936201478', '2020-09-06 22:20:07', '您的快件由【成都祥福分拣中心】准备发往【成都青白江分拣中心】', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (104, 'JDVD00936201478', '2020-09-06 22:50:50', '您的快件已发车', '', '2020-09-06 23:49:10', '', '2020-09-06 23:49:10');
INSERT INTO `logistics_tbl` VALUES (105, 'JDVD00936215561', '2020-09-06 17:44:29', '您的快件已到达【广元苍溪营业部】', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (106, 'JDVD00936215561', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (107, 'JDVD00936215561', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (108, 'JDVD00936215561', '2020-09-06 22:15:06', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (109, 'JDVD00936215561', '2020-09-06 22:43:08', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (110, 'JDVD00936215561', '2020-09-06 23:10:40', '您的快件由【成都祥福分拣中心】准备发往【芜湖分拣场】', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (111, '773054585661279', '2020-09-06 20:46:43', '快件已到达【四川南充转运中心】扫描员是【李雪】', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (112, '773054585661279', '2020-09-06 20:57:03', '快件由【四川南充转运中心】发往【安徽合肥转运中心】', '', '2020-09-06 23:49:11', '', '2020-09-06 23:49:11');
INSERT INTO `logistics_tbl` VALUES (113, '773054585583004', '2020-09-06 17:50:50', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (114, '773054585583004', '2020-09-06 18:33:54', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (115, '773054585583004', '2020-09-06 20:45:10', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (116, '773054585583004', '2020-09-06 20:57:47', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (117, '773054585734702', '2020-09-06 17:50:35', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (118, '773054585734702', '2020-09-06 18:31:15', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (119, '773054585734702', '2020-09-06 20:42:36', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (120, '773054585734702', '2020-09-06 20:57:07', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-06 23:49:12', '', '2020-09-06 23:49:12');
INSERT INTO `logistics_tbl` VALUES (121, '773054585734878', '2020-09-06 17:50:48', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (122, '773054585734878', '2020-09-06 18:34:20', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (123, '773054585734878', '2020-09-06 21:04:27', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (124, '773054585734878', '2020-09-06 21:09:41', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (125, '773054585734933', '2020-09-06 17:50:51', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (126, '773054585734933', '2020-09-06 18:30:41', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (127, '773054585734933', '2020-09-06 21:01:29', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (128, '773054585734933', '2020-09-06 21:11:38', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (129, '773054585510778', '2020-09-06 17:50:02', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (130, '773054585510778', '2020-09-06 18:31:54', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (131, '773054585510778', '2020-09-06 20:55:01', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (132, '773054585510778', '2020-09-06 21:09:50', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (133, '773054585661840', '2020-09-06 17:50:58', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (134, '773054585661840', '2020-09-06 18:41:50', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (135, '773054585661840', '2020-09-06 21:01:52', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (136, '773054585661840', '2020-09-06 21:09:46', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (137, '773054585813745', '2020-09-06 17:50:19', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (138, '773054585813745', '2020-09-06 18:32:10', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (139, '773054585813745', '2020-09-06 21:25:27', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (140, '773054585813745', '2020-09-06 21:33:05', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (141, '773054584544493', '2020-09-06 17:49:41', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (142, '773054584544493', '2020-09-06 18:36:31', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (143, '773054584544493', '2020-09-06 21:14:47', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (144, '773054584544493', '2020-09-06 21:28:01', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (145, '773054583991345', '2020-09-06 17:50:32', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (146, '773054583991345', '2020-09-06 18:32:59', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (147, '773054583991345', '2020-09-06 20:55:01', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (148, '773054583991345', '2020-09-06 21:09:40', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-06 23:49:13', '', '2020-09-06 23:49:13');
INSERT INTO `logistics_tbl` VALUES (149, 'JDVD00934692194', '2020-09-06 17:44:33', '您的快件已到达【广元苍溪营业部】', '', '2020-09-07 00:33:13', '', '2020-09-07 00:33:13');
INSERT INTO `logistics_tbl` VALUES (150, 'JDVD00934692194', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-07 00:33:13', '', '2020-09-07 00:33:13');
INSERT INTO `logistics_tbl` VALUES (151, 'JDVD00934692194', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-07 00:33:13', '', '2020-09-07 00:33:13');
INSERT INTO `logistics_tbl` VALUES (152, 'JDVD00934692194', '2020-09-06 22:15:29', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 00:33:13', '', '2020-09-07 00:33:13');
INSERT INTO `logistics_tbl` VALUES (153, 'JDVD00934692194', '2020-09-06 22:20:55', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 00:33:13', '', '2020-09-07 00:33:13');
INSERT INTO `logistics_tbl` VALUES (154, 'JDVD00934692194', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 00:33:13', '', '2020-09-07 00:33:13');
INSERT INTO `logistics_tbl` VALUES (155, 'JDVD00936221539', '2020-09-06 22:15:22', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (156, 'JDVD00936221539', '2020-09-06 22:20:50', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (157, 'JDVD00936221539', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (158, 'JDVD00936221512', '2020-09-06 22:14:03', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (159, 'JDVD00936221512', '2020-09-06 22:19:08', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (160, 'JDVD00936221512', '2020-09-06 22:19:13', '您的快件由【成都祥福分拣中心】准备发往【苏州昆山分拣中心】', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (161, 'JDVD00936221512', '2020-09-06 23:07:15', '您的快件已发车', '', '2020-09-07 00:33:15', '', '2020-09-07 00:33:15');
INSERT INTO `logistics_tbl` VALUES (162, 'JDVD00936213420', '2020-09-06 17:44:29', '您的快件已到达【广元苍溪营业部】', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (163, 'JDVD00936213420', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (164, 'JDVD00936213420', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (165, 'JDVD00936213420', '2020-09-06 22:15:26', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (166, 'JDVD00936213420', '2020-09-06 22:19:41', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (167, 'JDVD00936213420', '2020-09-06 22:19:46', '您的快件由【成都祥福分拣中心】准备发往【成都青白江分拣中心】', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (168, 'JDVD00936213420', '2020-09-06 23:12:38', '您的快件已发车', '', '2020-09-07 00:33:16', '', '2020-09-07 00:33:16');
INSERT INTO `logistics_tbl` VALUES (169, 'YT4768219922650', '2020-09-07 11:02:08', '添加快递单', 'admin', '2020-09-07 11:02:08', '', '2020-09-07 11:02:08');
INSERT INTO `logistics_tbl` VALUES (170, 'YT4768216385233', '2020-09-07 11:03:09', '添加快递单', 'admin', '2020-09-07 11:03:09', '', '2020-09-07 11:03:09');
INSERT INTO `logistics_tbl` VALUES (171, 'YT4768216384701', '2020-09-07 11:04:22', '添加快递单', 'admin', '2020-09-07 11:04:22', '', '2020-09-07 11:04:22');
INSERT INTO `logistics_tbl` VALUES (172, 'YT4768219923622', '2020-09-07 11:04:38', '添加快递单', 'admin', '2020-09-07 11:04:38', '', '2020-09-07 11:04:38');
INSERT INTO `logistics_tbl` VALUES (173, '773054585734702', '2020-09-07 04:51:38', '快件已到达【四川成都转运中心】扫描员是【出港捷信顺胥树芳】', '', '2020-09-07 11:27:44', '', '2020-09-07 11:27:44');
INSERT INTO `logistics_tbl` VALUES (174, '773054585734702', '2020-09-07 05:04:35', '快件由【四川成都转运中心】发往【四川成都锦绣三圣公司】', '', '2020-09-07 11:27:44', '', '2020-09-07 11:27:44');
INSERT INTO `logistics_tbl` VALUES (175, '773054585734702', '2020-09-07 07:22:09', '快件已到达【四川成都锦绣三圣公司】扫描员是【李燕】', '', '2020-09-07 11:27:44', '', '2020-09-07 11:27:44');
INSERT INTO `logistics_tbl` VALUES (176, '773054585734702', '2020-09-07 08:29:46', '【四川成都锦绣三圣公司】的派件员【李豪邦(17366936919)】正在为您派件，如有疑问请联系派件员，联系电话【17366936919】', '', '2020-09-07 11:27:44', '', '2020-09-07 11:27:44');
INSERT INTO `logistics_tbl` VALUES (177, '773054585734702', '2020-09-07 10:52:13', '快件已暂存至成都东御佲家店菜鸟驿站，如有疑问请联系13688356435', '', '2020-09-07 11:27:44', '', '2020-09-07 11:27:44');
INSERT INTO `logistics_tbl` VALUES (178, '773054585583004', '2020-09-07 04:47:16', '快件已到达【四川成都转运中心】扫描员是【进港李鑫杰】', '', '2020-09-07 11:27:56', '', '2020-09-07 11:27:56');
INSERT INTO `logistics_tbl` VALUES (179, '773054585583004', '2020-09-07 04:58:24', '快件由【四川成都转运中心】发往【四川安岳公司】', '', '2020-09-07 11:27:56', '', '2020-09-07 11:27:56');
INSERT INTO `logistics_tbl` VALUES (180, '773054585583004', '2020-09-07 10:42:55', '快件已到达【四川安岳公司】扫描员是【李豪】', '', '2020-09-07 11:27:56', '', '2020-09-07 11:27:56');
INSERT INTO `logistics_tbl` VALUES (181, '773054585583004', '2020-09-07 10:48:44', '【四川安岳公司】的派件员【李豪(18882501863)】正在为您派件，如有疑问请联系派件员，联系电话【18882501863】', '', '2020-09-07 11:27:56', '', '2020-09-07 11:27:56');
INSERT INTO `logistics_tbl` VALUES (182, 'JDVD00936215561', '2020-09-07 01:55:13', '您的快件已发车', '', '2020-09-07 11:41:59', '', '2020-09-07 11:41:59');
INSERT INTO `logistics_tbl` VALUES (183, '773054585734878', '2020-09-07 07:22:57', '快件已到达【重庆转运中心】扫描员是【B班HW组朱成元】', '', '2020-09-07 12:41:53', '', '2020-09-07 12:41:53');
INSERT INTO `logistics_tbl` VALUES (184, '773054585734878', '2020-09-07 07:34:16', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:41:53', '', '2020-09-07 12:41:53');
INSERT INTO `logistics_tbl` VALUES (185, '773054583991345', '2020-09-07 03:22:46', '快件已到达【重庆转运中心】扫描员是【B班BD组组长陈邦秀】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (186, '773054583991345', '2020-09-07 03:34:44', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (187, '773054583991345', '2020-09-07 08:34:03', '快件已到达【重庆机场公司】扫描员是【周达】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (188, '773054584544493', '2020-09-07 03:11:36', '快件已到达【重庆转运中心】扫描员是【B班BD组组长陈邦秀】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (189, '773054584544493', '2020-09-07 03:22:55', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (190, '773054585734933', '2020-09-07 03:18:40', '快件已到达【重庆转运中心】扫描员是【B班BD组组长陈邦秀】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (191, '773054585734933', '2020-09-07 03:28:39', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (192, '773054585734933', '2020-09-07 08:29:16', '快件已到达【重庆机场公司】扫描员是【周达】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (193, '773054585813745', '2020-09-07 03:06:30', '快件已到达【重庆转运中心】扫描员是【B班BD组组长陈邦秀】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (194, '773054585813745', '2020-09-07 03:18:17', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (195, '773054585813745', '2020-09-07 08:39:49', '快件已到达【重庆机场公司】扫描员是【周达】', '', '2020-09-07 12:42:02', '', '2020-09-07 12:42:02');
INSERT INTO `logistics_tbl` VALUES (196, 'JDVD00934683078', '2020-09-06 22:14:36', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (197, 'JDVD00934683078', '2020-09-06 22:20:05', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (198, 'JDVD00934683078', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (199, 'JDVD00934683078', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (200, 'JDVD00934683078', '2020-09-07 10:43:27', '您的快件已到达【重庆分拣中心】', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (201, 'JDVD00934683078', '2020-09-07 10:54:06', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (202, 'JDVD00934683078', '2020-09-07 10:54:11', '您的快件由【重庆分拣中心】准备发往【重庆渝北营业部】', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (203, 'JDVD00934683078', '2020-09-07 11:24:43', '您的快件已发车', '', '2020-09-07 12:42:06', '', '2020-09-07 12:42:06');
INSERT INTO `logistics_tbl` VALUES (204, '773054585510778', '2020-09-07 03:23:00', '快件已到达【重庆转运中心】扫描员是【B班BD组组长陈邦秀】', '', '2020-09-07 12:42:46', '', '2020-09-07 12:42:46');
INSERT INTO `logistics_tbl` VALUES (205, '773054585510778', '2020-09-07 03:33:56', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:42:46', '', '2020-09-07 12:42:46');
INSERT INTO `logistics_tbl` VALUES (206, '773054585661840', '2020-09-07 03:29:42', '快件已到达【重庆转运中心】扫描员是【B班BD组组长陈邦秀】', '', '2020-09-07 12:42:46', '', '2020-09-07 12:42:46');
INSERT INTO `logistics_tbl` VALUES (207, '773054585661840', '2020-09-07 03:36:22', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-07 12:42:46', '', '2020-09-07 12:42:46');
INSERT INTO `logistics_tbl` VALUES (208, '773054585661840', '2020-09-07 08:30:39', '快件已到达【重庆机场公司】扫描员是【周达】', '', '2020-09-07 12:42:46', '', '2020-09-07 12:42:46');
INSERT INTO `logistics_tbl` VALUES (209, 'YT4768493934859', '2020-09-07 12:58:45', '添加快递单', 'admin', '2020-09-07 12:58:45', '', '2020-09-07 12:58:45');
INSERT INTO `logistics_tbl` VALUES (210, 'YT4768496031904', '2020-09-07 12:59:07', '添加快递单', 'admin', '2020-09-07 12:59:07', '', '2020-09-07 12:59:07');
INSERT INTO `logistics_tbl` VALUES (211, 'YT4768496704095', '2020-09-07 12:59:27', '添加快递单', 'admin', '2020-09-07 12:59:27', '', '2020-09-07 12:59:27');
INSERT INTO `logistics_tbl` VALUES (212, 'YT4768497494319', '2020-09-07 12:59:37', '添加快递单', 'admin', '2020-09-07 12:59:37', '', '2020-09-07 12:59:37');
INSERT INTO `logistics_tbl` VALUES (213, 'YT4768219923622', '2020-09-07 13:36:45', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 14:47:00', '', '2020-09-07 14:47:00');
INSERT INTO `logistics_tbl` VALUES (214, 'YT4768216384701', '2020-09-07 13:36:51', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 14:47:06', '', '2020-09-07 14:47:06');
INSERT INTO `logistics_tbl` VALUES (215, 'YT4768216385233', '2020-09-07 13:36:57', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 14:47:06', '', '2020-09-07 14:47:06');
INSERT INTO `logistics_tbl` VALUES (216, '773054739511845', '2020-09-07 15:05:40', '添加快递单', 'admin', '2020-09-07 15:05:40', '', '2020-09-07 15:05:40');
INSERT INTO `logistics_tbl` VALUES (217, '773054741020685', '2020-09-07 15:06:05', '添加快递单', 'admin', '2020-09-07 15:06:05', '', '2020-09-07 15:06:05');
INSERT INTO `logistics_tbl` VALUES (218, '773054740940829', '2020-09-07 15:19:48', '添加快递单', 'admin', '2020-09-07 15:19:48', '', '2020-09-07 15:19:48');
INSERT INTO `logistics_tbl` VALUES (219, '773054585510778', '2020-09-07 12:25:22', '快件已到达【重庆机场公司】扫描员是【暨华片区】', '', '2020-09-07 16:26:05', '', '2020-09-07 16:26:05');
INSERT INTO `logistics_tbl` VALUES (220, '773054585510778', '2020-09-07 12:34:17', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-07 16:26:05', '', '2020-09-07 16:26:05');
INSERT INTO `logistics_tbl` VALUES (221, '773054585661840', '2020-09-07 12:34:30', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-07 16:26:05', '', '2020-09-07 16:26:05');
INSERT INTO `logistics_tbl` VALUES (222, 'JDVD00934676229', '2020-09-05 17:04:51', '揽收任务已分配给邓洁', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (223, 'JDVD00934676229', '2020-09-06 14:45:30', '京东快递 已收取快件', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (224, 'JDVD00934676229', '2020-09-06 14:45:30', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (225, 'JDVD00934676229', '2020-09-06 17:44:40', '您的快件已到达【广元苍溪营业部】', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (226, 'JDVD00934676229', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (227, 'JDVD00934676229', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (228, 'JDVD00934676229', '2020-09-06 22:13:07', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (229, 'JDVD00934676229', '2020-09-06 22:18:28', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (230, 'JDVD00934676229', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (231, 'JDVD00934676229', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (232, 'JDVD00934676229', '2020-09-07 10:44:19', '您的快件已到达【重庆分拣中心】', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (233, 'JDVD00934676229', '2020-09-07 10:54:17', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (234, 'JDVD00934676229', '2020-09-07 10:54:22', '您的快件由【重庆分拣中心】准备发往【重庆渝北营业部】', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (235, 'JDVD00934676229', '2020-09-07 11:24:43', '您的快件已发车', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (236, 'JDVD00934676229', '2020-09-07 13:22:12', '您的快件已到达【重庆渝北营业部】', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (237, 'JDVD00934676229', '2020-09-07 13:40:15', '您的快件正在派送中，请您准备签收（快递员：温鹏，联系电话：17815289082）', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (238, 'JDVD00934676229', '2020-09-07 15:49:42', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-07 16:26:14', '', '2020-09-07 16:26:14');
INSERT INTO `logistics_tbl` VALUES (239, 'JDVD00934679583', '2020-09-05 17:04:51', '揽收任务已分配给邓洁', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (240, 'JDVD00934679583', '2020-09-06 14:45:20', '京东快递 已收取快件', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (241, 'JDVD00934679583', '2020-09-06 14:45:20', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (242, 'JDVD00934679583', '2020-09-06 22:15:14', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (243, 'JDVD00934679583', '2020-09-06 22:19:43', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (244, 'JDVD00934679583', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (245, 'JDVD00934679583', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (246, 'JDVD00934679583', '2020-09-07 10:43:28', '您的快件已到达【重庆分拣中心】', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (247, 'JDVD00934679583', '2020-09-07 10:54:10', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (248, 'JDVD00934679583', '2020-09-07 10:54:15', '您的快件由【重庆分拣中心】准备发往【重庆渝北营业部】', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (249, 'JDVD00934679583', '2020-09-07 11:24:43', '您的快件已发车', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (250, 'JDVD00934679583', '2020-09-07 13:21:40', '您的快件已到达【重庆渝北营业部】', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (251, 'JDVD00934679583', '2020-09-07 13:40:37', '您的快件正在派送中，请您准备签收（快递员：温鹏，联系电话：17815289082）', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (252, 'JDVD00934679583', '2020-09-07 15:49:42', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-07 16:26:34', '', '2020-09-07 16:26:34');
INSERT INTO `logistics_tbl` VALUES (253, 'JDVD00934685489', '2020-09-05 17:04:50', '揽收任务已分配给邓洁', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (254, 'JDVD00934685489', '2020-09-06 14:45:46', '京东快递 已收取快件', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (255, 'JDVD00934685489', '2020-09-06 14:45:46', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (256, 'JDVD00934685489', '2020-09-06 17:44:29', '您的快件已到达【广元苍溪营业部】', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (257, 'JDVD00934685489', '2020-09-06 17:51:17', '快递司机收箱', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (258, 'JDVD00934685489', '2020-09-06 17:51:22', '您的快件已发车', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (259, 'JDVD00934685489', '2020-09-06 22:15:08', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (260, 'JDVD00934685489', '2020-09-06 22:20:28', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (261, 'JDVD00934685489', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (262, 'JDVD00934685489', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (263, 'JDVD00934685489', '2020-09-07 12:36:29', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (264, 'JDVD00934685489', '2020-09-07 14:03:16', '您的快件由【重庆分拣中心】准备发往【重庆渝北营业部】', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (265, 'JDVD00934685489', '2020-09-07 14:12:12', '您的快件已发车', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (266, 'JDVD00934685489', '2020-09-07 15:51:20', '您的快件已到达【重庆渝北营业部】', '', '2020-09-07 16:26:35', '', '2020-09-07 16:26:35');
INSERT INTO `logistics_tbl` VALUES (267, 'JDVD00934659318', '2020-09-05 17:04:50', '揽收任务已分配给邓洁', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (268, 'JDVD00934659318', '2020-09-06 14:44:59', '京东快递 已收取快件', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (269, 'JDVD00934659318', '2020-09-06 14:45:00', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (270, 'JDVD00934659318', '2020-09-06 22:14:39', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (271, 'JDVD00934659318', '2020-09-06 22:19:50', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (272, 'JDVD00934659318', '2020-09-06 22:33:54', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (273, 'JDVD00934659318', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (274, 'JDVD00934659318', '2020-09-07 10:53:41', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (275, 'JDVD00934659318', '2020-09-07 10:53:46', '您的快件由【重庆分拣中心】准备发往【重庆渝北营业部】', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (276, 'JDVD00934659318', '2020-09-07 11:24:43', '您的快件已发车', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (277, 'JDVD00934659318', '2020-09-07 13:21:58', '您的快件已到达【重庆渝北营业部】', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (278, 'JDVD00934659318', '2020-09-07 13:40:07', '您的快件正在派送中，请您准备签收（快递员：温鹏，联系电话：17815289082）', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (279, 'JDVD00934659318', '2020-09-07 15:49:42', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-07 16:26:40', '', '2020-09-07 16:26:40');
INSERT INTO `logistics_tbl` VALUES (280, '773054583991345', '2020-09-07 12:34:42', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-07 16:26:43', '', '2020-09-07 16:26:43');
INSERT INTO `logistics_tbl` VALUES (281, '773054584544493', '2020-09-07 12:24:18', '快件已到达【重庆机场公司】扫描员是【暨华片区】', '', '2020-09-07 16:26:44', '', '2020-09-07 16:26:44');
INSERT INTO `logistics_tbl` VALUES (282, '773054584544493', '2020-09-07 12:29:32', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-07 16:26:44', '', '2020-09-07 16:26:44');
INSERT INTO `logistics_tbl` VALUES (283, '773054585734933', '2020-09-07 12:29:19', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-07 16:26:44', '', '2020-09-07 16:26:44');
INSERT INTO `logistics_tbl` VALUES (284, '773054585813745', '2020-09-07 12:29:26', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-07 16:26:44', '', '2020-09-07 16:26:44');
INSERT INTO `logistics_tbl` VALUES (285, '773054585245386', '2020-09-07 17:59:43', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:37', '', '2020-09-07 18:27:37');
INSERT INTO `logistics_tbl` VALUES (286, '773054585245386', '2020-09-07 18:21:04', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:37', '', '2020-09-07 18:27:37');
INSERT INTO `logistics_tbl` VALUES (287, '773054585660663', '2020-09-07 17:59:44', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:37', '', '2020-09-07 18:27:37');
INSERT INTO `logistics_tbl` VALUES (288, '773054585660663', '2020-09-07 18:15:54', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:37', '', '2020-09-07 18:27:37');
INSERT INTO `logistics_tbl` VALUES (289, '773054585734185', '2020-09-07 17:59:45', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:38', '', '2020-09-07 18:27:38');
INSERT INTO `logistics_tbl` VALUES (290, '773054585734185', '2020-09-07 18:16:11', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:38', '', '2020-09-07 18:27:38');
INSERT INTO `logistics_tbl` VALUES (291, '773054585812874', '2020-09-07 17:59:43', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:38', '', '2020-09-07 18:27:38');
INSERT INTO `logistics_tbl` VALUES (292, '773054585812874', '2020-09-07 18:24:05', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:38', '', '2020-09-07 18:27:38');
INSERT INTO `logistics_tbl` VALUES (293, '773054583991096', '2020-09-07 17:59:42', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (294, '773054583991096', '2020-09-07 18:25:42', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (295, '773054585660092', '2020-09-07 17:59:37', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (296, '773054585660092', '2020-09-07 18:17:44', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (297, '773054585660134', '2020-09-07 17:59:39', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (298, '773054585660134', '2020-09-07 18:16:46', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (299, '773054585733964', '2020-09-07 17:59:40', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (300, '773054585733964', '2020-09-07 18:21:52', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 18:27:45', '', '2020-09-07 18:27:45');
INSERT INTO `logistics_tbl` VALUES (301, 'YT4768493934859', '2020-09-07 13:37:26', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 22:01:05', '', '2020-09-07 22:01:05');
INSERT INTO `logistics_tbl` VALUES (302, 'YT4768493934859', '2020-09-07 21:07:30', '【南充转运中心公司】 已收入', '', '2020-09-07 22:01:05', '', '2020-09-07 22:01:05');
INSERT INTO `logistics_tbl` VALUES (303, 'YT4768493934859', '2020-09-07 21:35:06', '【南充转运中心】 已发出 下一站 【四川省广安市公司】', '', '2020-09-07 22:01:05', '', '2020-09-07 22:01:05');
INSERT INTO `logistics_tbl` VALUES (304, 'YT4768496031904', '2020-09-07 13:37:32', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 22:01:05', '', '2020-09-07 22:01:05');
INSERT INTO `logistics_tbl` VALUES (305, 'YT4768496031904', '2020-09-07 21:10:51', '【南充转运中心公司】 已收入', '', '2020-09-07 22:01:05', '', '2020-09-07 22:01:05');
INSERT INTO `logistics_tbl` VALUES (306, 'YT4768496031904', '2020-09-07 21:34:55', '【南充转运中心】 已发出 下一站 【四川省广安市公司】', '', '2020-09-07 22:01:05', '', '2020-09-07 22:01:05');
INSERT INTO `logistics_tbl` VALUES (307, '773054583991345', '2020-09-07 19:29:05', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-07 22:01:17', '', '2020-09-07 22:01:17');
INSERT INTO `logistics_tbl` VALUES (308, '773054584544493', '2020-09-07 19:29:05', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-07 22:01:17', '', '2020-09-07 22:01:17');
INSERT INTO `logistics_tbl` VALUES (309, '773054585734933', '2020-09-07 19:29:04', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-07 22:01:18', '', '2020-09-07 22:01:18');
INSERT INTO `logistics_tbl` VALUES (310, '773054585813745', '2020-09-07 19:29:05', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-07 22:01:18', '', '2020-09-07 22:01:18');
INSERT INTO `logistics_tbl` VALUES (311, 'JDVD00934685489', '2020-09-07 16:49:07', '您的快件正在派送中，请您准备签收（快递员：温鹏，联系电话：17815289082）', '', '2020-09-07 22:01:42', '', '2020-09-07 22:01:42');
INSERT INTO `logistics_tbl` VALUES (312, 'JDVD00934685489', '2020-09-07 18:11:47', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-07 22:01:42', '', '2020-09-07 22:01:42');
INSERT INTO `logistics_tbl` VALUES (313, '773054585734702', '2020-09-07 18:58:44', '已签收，签收人凭取货码签收。', '', '2020-09-07 22:02:01', '', '2020-09-07 22:02:01');
INSERT INTO `logistics_tbl` VALUES (314, '773054740940829', '2020-09-07 16:33:50', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 22:02:11', '', '2020-09-07 22:02:11');
INSERT INTO `logistics_tbl` VALUES (315, '773054740940829', '2020-09-07 18:20:41', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 22:02:11', '', '2020-09-07 22:02:11');
INSERT INTO `logistics_tbl` VALUES (316, 'YT4768496704095', '2020-09-07 13:37:40', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 22:02:11', '', '2020-09-07 22:02:11');
INSERT INTO `logistics_tbl` VALUES (317, 'YT4768496704095', '2020-09-07 21:07:37', '【南充转运中心公司】 已收入', '', '2020-09-07 22:02:11', '', '2020-09-07 22:02:11');
INSERT INTO `logistics_tbl` VALUES (318, 'YT4768496704095', '2020-09-07 21:34:49', '【南充转运中心】 已发出 下一站 【四川省广安市公司】', '', '2020-09-07 22:02:11', '', '2020-09-07 22:02:11');
INSERT INTO `logistics_tbl` VALUES (319, 'YT4768497494319', '2020-09-07 13:37:37', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-07 22:02:12', '', '2020-09-07 22:02:12');
INSERT INTO `logistics_tbl` VALUES (320, 'YT4768497494319', '2020-09-07 21:14:45', '【南充转运中心公司】 已收入', '', '2020-09-07 22:02:12', '', '2020-09-07 22:02:12');
INSERT INTO `logistics_tbl` VALUES (321, 'YT4768497494319', '2020-09-07 21:34:46', '【南充转运中心】 已发出 下一站 【四川省广安市公司】', '', '2020-09-07 22:02:12', '', '2020-09-07 22:02:12');
INSERT INTO `logistics_tbl` VALUES (322, '773054741020685', '2020-09-07 17:54:15', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-07 22:02:21', '', '2020-09-07 22:02:21');
INSERT INTO `logistics_tbl` VALUES (323, '773054741020685', '2020-09-07 18:18:56', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-07 22:02:21', '', '2020-09-07 22:02:21');
INSERT INTO `logistics_tbl` VALUES (324, 'JDVD00936221539', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (325, 'JDVD00936221539', '2020-09-07 11:03:20', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (326, 'JDVD00936221539', '2020-09-07 11:03:25', '您的快件由【重庆分拣中心】准备发往【重庆恒大同景国际京东星配站】', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (327, 'JDVD00936221539', '2020-09-07 11:24:43', '您的快件已发车', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (328, 'JDVD00936221539', '2020-09-07 12:17:01', '您的快件已到达【重庆恒大同景国际京东星配站】', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (329, 'JDVD00936221539', '2020-09-07 14:39:05', '您的快件已到达【重庆奥园越时代京东星配站】', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (330, 'JDVD00936221539', '2020-09-07 14:39:39', '您的快件正在派送中，请您准备签收（快递员：吴力辉，联系电话：15223378355）', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (331, 'JDVD00936221539', '2020-09-07 18:42:14', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 00:05:05', '', '2020-09-08 00:05:05');
INSERT INTO `logistics_tbl` VALUES (332, '773054741020685', '2020-09-07 21:52:11', '快件已到达【四川南充转运中心】扫描员是【顶扫五】', '', '2020-09-08 00:39:16', '', '2020-09-08 00:39:16');
INSERT INTO `logistics_tbl` VALUES (333, '773054741020685', '2020-09-07 22:06:07', '快件由【四川南充转运中心】发往【四川自贡转运中心】', '', '2020-09-08 00:39:16', '', '2020-09-08 00:39:16');
INSERT INTO `logistics_tbl` VALUES (334, 'YT4768219922650', '2020-09-07 13:36:12', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-08 00:39:22', '', '2020-09-08 00:39:22');
INSERT INTO `logistics_tbl` VALUES (335, 'YT4768219922650', '2020-09-07 21:07:13', '【南充转运中心公司】 已收入', '', '2020-09-08 00:39:22', '', '2020-09-08 00:39:22');
INSERT INTO `logistics_tbl` VALUES (336, 'YT4768219922650', '2020-09-07 21:32:02', '【南充转运中心】 已发出 下一站 【自贡转运中心公司】', '', '2020-09-08 00:39:22', '', '2020-09-08 00:39:22');
INSERT INTO `logistics_tbl` VALUES (337, '773054585734878', '2020-09-07 13:47:25', '快件已到达【重庆机场公司】扫描员是【暨华片区】', '', '2020-09-08 09:25:10', '', '2020-09-08 09:25:10');
INSERT INTO `logistics_tbl` VALUES (338, '773054585734878', '2020-09-07 13:55:22', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-08 09:25:10', '', '2020-09-08 09:25:10');
INSERT INTO `logistics_tbl` VALUES (339, '773054585734878', '2020-09-07 19:29:04', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 09:25:10', '', '2020-09-08 09:25:10');
INSERT INTO `logistics_tbl` VALUES (340, '773054583991096', '2020-09-07 21:48:33', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (341, '773054583991096', '2020-09-07 22:02:24', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (342, '773054583991096', '2020-09-08 03:00:26', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (343, '773054583991096', '2020-09-08 03:14:18', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (344, '773054583991096', '2020-09-08 07:20:38', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (345, '773054583991096', '2020-09-08 08:31:55', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (346, '773054583991096', '2020-09-08 08:38:29', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (347, '773054583991096', '2020-09-08 10:44:13', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (348, '773054585660092', '2020-09-07 21:51:48', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (349, '773054585660092', '2020-09-07 22:03:03', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (350, '773054585660092', '2020-09-08 03:03:48', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (351, '773054585660092', '2020-09-08 03:14:40', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (352, '773054585660092', '2020-09-08 07:24:40', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (353, '773054585660092', '2020-09-08 08:31:46', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (354, '773054585660092', '2020-09-08 08:38:35', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (355, '773054585660092', '2020-09-08 10:42:41', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 11:20:23', '', '2020-09-08 11:20:23');
INSERT INTO `logistics_tbl` VALUES (356, '773054585660134', '2020-09-07 21:53:06', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (357, '773054585660134', '2020-09-07 22:01:59', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (358, '773054585660134', '2020-09-08 03:04:39', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (359, '773054585660134', '2020-09-08 03:12:06', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (360, '773054585660134', '2020-09-08 07:15:45', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (361, '773054585660134', '2020-09-08 08:31:07', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (362, '773054585660134', '2020-09-08 08:38:42', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (363, '773054585660134', '2020-09-08 10:42:52', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (364, '773054585733964', '2020-09-07 21:43:22', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (365, '773054585733964', '2020-09-07 21:51:08', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (366, '773054585733964', '2020-09-08 04:44:42', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (367, '773054585733964', '2020-09-08 04:55:54', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (368, '773054585733964', '2020-09-08 06:58:21', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (369, '773054585733964', '2020-09-08 08:30:24', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (370, '773054585733964', '2020-09-08 08:38:38', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (371, '773054585733964', '2020-09-08 10:44:37', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 11:20:24', '', '2020-09-08 11:20:24');
INSERT INTO `logistics_tbl` VALUES (372, 'YT4770457089144', '2020-09-08 11:38:35', '添加快递单', 'admin', '2020-09-08 11:38:35', '', '2020-09-08 11:38:35');
INSERT INTO `logistics_tbl` VALUES (373, 'YT4770460243366', '2020-09-08 11:38:45', '添加快递单', 'admin', '2020-09-08 11:38:45', '', '2020-09-08 11:38:45');
INSERT INTO `logistics_tbl` VALUES (374, '773054585245386', '2020-09-07 21:59:43', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (375, '773054585245386', '2020-09-07 22:07:45', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (376, '773054585245386', '2020-09-08 02:54:54', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (377, '773054585245386', '2020-09-08 03:07:00', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (378, '773054585245386', '2020-09-08 07:22:04', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (379, '773054585245386', '2020-09-08 08:32:02', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (380, '773054585245386', '2020-09-08 08:38:47', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (381, '773054585245386', '2020-09-08 10:49:26', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (382, '773054585660663', '2020-09-07 21:43:45', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (383, '773054585660663', '2020-09-07 21:56:56', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (384, '773054585660663', '2020-09-08 04:04:01', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (385, '773054585660663', '2020-09-08 04:15:22', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (386, '773054585660663', '2020-09-08 07:15:55', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (387, '773054585660663', '2020-09-08 08:32:50', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (388, '773054585660663', '2020-09-08 08:38:40', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (389, '773054585660663', '2020-09-08 10:50:18', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (390, '773054585734185', '2020-09-07 21:48:48', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (391, '773054585734185', '2020-09-07 21:56:50', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (392, '773054585734185', '2020-09-08 03:10:04', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (393, '773054585734185', '2020-09-08 03:17:10', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (394, '773054585734185', '2020-09-08 07:16:14', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (395, '773054585734185', '2020-09-08 08:33:02', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (396, '773054585734185', '2020-09-08 08:38:45', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (397, '773054585734185', '2020-09-08 10:50:04', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 12:51:19', '', '2020-09-08 12:51:19');
INSERT INTO `logistics_tbl` VALUES (398, '773054585812874', '2020-09-07 21:56:37', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (399, '773054585812874', '2020-09-07 22:01:51', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (400, '773054585812874', '2020-09-08 03:04:57', '快件已到达【四川成都转运中心】扫描员是【进港夜班兴燚恒】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (401, '773054585812874', '2020-09-08 03:15:18', '快件由【四川成都转运中心】发往【四川成都华阳公司】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (402, '773054585812874', '2020-09-08 07:13:58', '快件已到达【四川成都华阳公司】扫描员是【罗成】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (403, '773054585812874', '2020-09-08 08:32:05', '快件已到达【四川成都华阳公司】扫描员是【凯华丽景承包区】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (404, '773054585812874', '2020-09-08 08:38:32', '【四川成都华阳公司】的派件员【凯华丽景承包区(17635727438)】正在为您派件，如有疑问请联系派件员，联系电话【17635727438】', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (405, '773054585812874', '2020-09-08 10:49:45', '已签收，签收(17635727438)人是【本人】,如有疑问请联系:17635727438,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 12:51:20', '', '2020-09-08 12:51:20');
INSERT INTO `logistics_tbl` VALUES (406, 'JDVD00936221520', '2020-09-07 01:25:04', '您的快件已发车', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (407, 'JDVD00936221520', '2020-09-07 05:33:12', '您的快件已到达【南充分拣场】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (408, 'JDVD00936221520', '2020-09-07 05:34:31', '您的快件在【南充分拣场】分拣完成', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (409, 'JDVD00936221520', '2020-09-07 05:34:36', '您的快件由【南充分拣场】准备发往【广安城南营业部】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (410, 'JDVD00936221520', '2020-09-07 07:18:35', '您的快件已发车', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (411, 'JDVD00936221520', '2020-09-07 09:06:48', '您的快件已到达【广安城南营业部】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (412, 'JDVD00936221520', '2020-09-07 09:31:05', '您的快件正在派送中，请您准备签收（快递员：周世策，联系电话：18780151907）', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (413, 'JDVD00936221520', '2020-09-07 11:13:09', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (414, 'JDVD00936215570', '2020-09-06 23:58:05', '您的快件已到达【成都青白江分拣中心】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (415, 'JDVD00936215570', '2020-09-07 00:02:11', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (416, 'JDVD00936215570', '2020-09-07 00:02:16', '您的快件由【成都青白江分拣中心】准备发往【昆明分拣中心】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (417, 'JDVD00936215570', '2020-09-07 02:04:10', '您的快件已发车', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (418, 'JDVD00936215570', '2020-09-07 18:04:53', '您的快件已到达【昆明分拣中心】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (419, 'JDVD00936215570', '2020-09-07 19:31:32', '您的快件在【昆明分拣中心】分拣完成', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (420, 'JDVD00936215570', '2020-09-07 19:31:37', '您的快件由【昆明分拣中心】准备发往【昆明关上营业部】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (421, 'JDVD00936215570', '2020-09-08 05:36:40', '您的快件已发车', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (422, 'JDVD00936215570', '2020-09-08 07:47:27', '您的快件在【昆明关上营业部】收货完成', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (423, 'JDVD00936215570', '2020-09-08 07:47:28', '您的快件已到达【昆明关上营业部】', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (424, 'JDVD00936215570', '2020-09-08 08:09:23', '您的快件正在派送中，请您准备签收（快递员：邹启志，联系电话：13669792793）', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (425, 'JDVD00936215570', '2020-09-08 10:46:08', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 13:59:55', '', '2020-09-08 13:59:55');
INSERT INTO `logistics_tbl` VALUES (426, 'JDVD00936201478', '2020-09-07 00:56:12', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (427, 'JDVD00936201478', '2020-09-07 04:49:23', '您的快件由【成都青白江分拣中心】准备发往【成都大源营业部】', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (428, 'JDVD00936201478', '2020-09-07 05:49:48', '您的快件已发车', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (429, 'JDVD00936201478', '2020-09-07 07:16:31', '您的快件在【成都大源营业部】收货完成', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (430, 'JDVD00936201478', '2020-09-07 07:16:32', '您的快件已到达【成都大源营业部】', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (431, 'JDVD00936201478', '2020-09-07 08:02:38', '您的快件正在派送中，请您准备签收（快递员：王锐，联系电话：13408593075）', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (432, 'JDVD00936201478', '2020-09-07 12:37:36', '您的快件派送不成功，原因【联系不上客户】，预计下次派送时间2020-09-07 18:00', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (433, 'JDVD00936201478', '2020-09-07 14:19:22', '您的快件正在派送中，请您准备签收（快递员：王锐，联系电话：13408593075）', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (434, 'JDVD00936201478', '2020-09-07 19:46:29', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (435, 'JDVD00936215561', '2020-09-08 12:03:04', '您的快件已到达【芜湖分拣场】', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (436, 'JDVD00936215561', '2020-09-08 12:04:06', '您的快件在【芜湖分拣场】分拣完成', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (437, 'JDVD00936215561', '2020-09-08 12:04:11', '您的快件由【芜湖分拣场】准备发往【宣城宁国营业部】', '', '2020-09-08 13:59:56', '', '2020-09-08 13:59:56');
INSERT INTO `logistics_tbl` VALUES (438, '773054585661279', '2020-09-08 03:25:52', '快件已到达【安徽合肥转运中心】扫描员是【李群】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (439, '773054585661279', '2020-09-08 03:39:19', '快件由【安徽合肥转运中心】发往【安徽芜湖转运中心】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (440, '773054585661279', '2020-09-08 07:59:40', '快件已到达【安徽芜湖转运中心】扫描员是【合肥高扫1】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (441, '773054585661279', '2020-09-08 08:11:17', '快件由【安徽芜湖转运中心】发往【安徽马鞍山公司】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (442, '773054585661279', '2020-09-08 11:13:08', '快件由【安徽马鞍山公司】发往【安徽马鞍山映翠分部】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (443, '773054585582825', '2020-09-07 17:59:46', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (444, '773054585582825', '2020-09-07 18:24:26', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (445, '773054585582825', '2020-09-07 21:47:14', '快件已到达【四川南充转运中心】扫描员是【临时工】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (446, '773054585582825', '2020-09-07 21:55:23', '快件由【四川南充转运中心】发往【上海转运中心】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (447, '773054583991235', '2020-09-07 17:59:48', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (448, '773054583991235', '2020-09-07 18:15:39', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (449, '773054583991235', '2020-09-08 00:57:16', '快件已到达【四川南充转运中心】扫描员是【临时工】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (450, '773054583991235', '2020-09-08 01:02:17', '快件由【四川南充转运中心】发往【四川广安公司】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (451, '773054583991235', '2020-09-08 08:03:31', '快件已到达【四川广安公司】扫描员是【西溪新城片区】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (452, '773054583991235', '2020-09-08 08:09:57', '【四川广安公司】的派件员【西溪新城片区(18808280635)】正在为您派件，如有疑问请联系派件员，联系电话【18808280635】', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (453, '773054583991235', '2020-09-08 09:43:15', '已签收，签(18808280635)收人是【门卫代签】,如有疑问请联系:18808280635,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (454, '773054585583004', '2020-09-07 13:14:48', '已(18882501863)签收，签收人是【盛世首座门卫】,如有疑问请联系:18882501863,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (455, '773054585510778', '2020-09-07 19:29:05', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 13:59:57', '', '2020-09-08 13:59:57');
INSERT INTO `logistics_tbl` VALUES (456, '773054585661840', '2020-09-07 19:29:05', '已签收(15123280019)，签收人是【朋友代签】,如有疑问请联系:15123280019,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (457, 'YT4768219922650', '2020-09-08 07:50:47', '【自贡转运中心公司】 已收入', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (458, 'YT4768219922650', '2020-09-08 08:17:04', '【自贡转运中心】 已发出 下一站 【四川省泸州市公司】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (459, 'YT4768219922650', '2020-09-08 13:24:03', '【四川省泸州市公司】 已收入', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (460, 'YT4768219922650', '2020-09-08 13:28:54', '【四川省泸州市】 已发出 下一站 【四川省泸州市江阳区万象汇公司】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (461, 'YT4768493934859', '2020-09-08 07:01:53', '【四川省广安市公司】 已收入', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (462, 'YT4768493934859', '2020-09-08 07:58:46', '【四川省广安市公司】 派件中  派件人: 余翔 电话 18808280065  如有疑问，请联系：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (463, 'YT4768493934859', '2020-09-08 11:08:47', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18808280065，投诉电话：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (464, 'YT4768496031904', '2020-09-08 07:01:59', '【四川省广安市公司】 已收入', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (465, 'YT4768496031904', '2020-09-08 07:37:27', '【四川省广安市公司】 派件中  派件人: 余翔 电话 18808280065  如有疑问，请联系：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (466, 'YT4768496031904', '2020-09-08 11:08:42', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18808280065，投诉电话：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (467, 'YT4768496704095', '2020-09-08 07:01:49', '【四川省广安市公司】 已收入', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (468, 'YT4768496704095', '2020-09-08 07:25:51', '【四川省广安市公司】 派件中  派件人: 贺兴华 电话 18808280085  如有疑问，请联系：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (469, 'YT4768496704095', '2020-09-08 11:46:47', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18808280085，投诉电话：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (470, 'YT4768497494319', '2020-09-08 07:01:53', '【四川省广安市公司】 已收入', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (471, 'YT4768497494319', '2020-09-08 07:41:35', '【四川省广安市公司】 派件中  派件人: 贺兴华 电话 18808280085  如有疑问，请联系：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (472, 'YT4768497494319', '2020-09-08 11:46:45', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18808280085，投诉电话：0826-2345387', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (473, '773054739511845', '2020-09-07 17:54:30', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (474, '773054739511845', '2020-09-07 18:17:45', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (475, '773054739511845', '2020-09-08 00:53:36', '快件已到达【四川南充转运中心】扫描员是【马海蓉】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (476, '773054739511845', '2020-09-08 01:01:59', '快件由【四川南充转运中心】发往【四川仪陇公司】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (477, '773054741020685', '2020-09-08 03:26:14', '快件已到达【四川自贡转运中心】扫描员是【进朱友明】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (478, '773054741020685', '2020-09-08 03:32:37', '快件由【四川自贡转运中心】发往【四川泸州公司】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (479, '773054741020685', '2020-09-08 07:41:42', '快件已到达【四川泸州公司】扫描员是【到件扫描】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (480, '773054741020685', '2020-09-08 08:55:16', '快件由【四川泸州公司】发往【四川泸州佳乐世纪城】', '', '2020-09-08 13:59:58', '', '2020-09-08 13:59:58');
INSERT INTO `logistics_tbl` VALUES (481, '773054740940829', '2020-09-08 00:49:46', '快件已到达【四川南充转运中心】扫描员是【临时工】', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (482, '773054740940829', '2020-09-08 01:04:00', '快件由【四川南充转运中心】发往【四川广安公司】', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (483, '773054740940829', '2020-09-08 07:52:57', '快件已到达【四川广安公司】扫描员是【申通到件机器人】', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (484, '773054740940829', '2020-09-08 08:13:27', '快件已到达【四川广安公司】扫描员是【康大军】', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (485, '773054740940829', '2020-09-08 08:18:44', '【四川广安公司】的派件员【康大军(18808281880)】正在为您派件，如有疑问请联系派件员，联系电话【18808281880】', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (486, '773054740940829', '2020-09-08 10:10:42', '已签(18808281880)收，签收人是【朋友/同事代签】,如有疑问请联系:18808281880,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (487, 'YT4770457089144', '2020-09-08 13:49:16', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (488, 'YT4770460243366', '2020-09-08 13:49:12', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-08 13:59:59', '', '2020-09-08 13:59:59');
INSERT INTO `logistics_tbl` VALUES (489, '773054741020685', '2020-09-08 14:36:49', '快件已到达【四川泸州佳乐世纪城】扫描员是【世纪城邬磊】', '', '2020-09-08 15:29:38', '', '2020-09-08 15:29:38');
INSERT INTO `logistics_tbl` VALUES (490, '773054741020685', '2020-09-08 14:42:43', '【四川泸州佳乐世纪城】的派件员【武国强(18783007955)】正在为您派件，如有疑问请联系派件员，联系电话【18783007955】', '', '2020-09-08 15:29:38', '', '2020-09-08 15:29:38');
INSERT INTO `logistics_tbl` VALUES (491, '773054906506161', '2020-09-08 15:56:26', '添加快递单', 'admin', '2020-09-08 15:56:26', '', '2020-09-08 15:56:26');
INSERT INTO `logistics_tbl` VALUES (492, '773054904868479', '2020-09-08 15:56:39', '添加快递单', 'admin', '2020-09-08 15:56:39', '', '2020-09-08 15:56:39');
INSERT INTO `logistics_tbl` VALUES (493, '773054904635511', '2020-09-08 15:56:57', '添加快递单', 'admin', '2020-09-08 15:56:57', '', '2020-09-08 15:56:57');
INSERT INTO `logistics_tbl` VALUES (494, '773054906075645', '2020-09-08 15:58:09', '添加快递单', 'admin', '2020-09-08 15:58:09', '', '2020-09-08 15:58:09');
INSERT INTO `logistics_tbl` VALUES (495, '773054906427998', '2020-09-08 15:58:15', '添加快递单', 'admin', '2020-09-08 15:58:15', '', '2020-09-08 15:58:15');
INSERT INTO `logistics_tbl` VALUES (496, '773054741020685', '2020-09-08 19:40:57', '已签(18783007955)收，签收人是【前台代签】,如有疑问请联系:18783007955,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 19:48:54', '', '2020-09-08 19:48:54');
INSERT INTO `logistics_tbl` VALUES (497, '773054739511845', '2020-09-08 14:09:59', '快件已到达【四川仪陇公司】扫描员是【高扫】', '', '2020-09-08 19:49:55', '', '2020-09-08 19:49:55');
INSERT INTO `logistics_tbl` VALUES (498, '773054739511845', '2020-09-08 14:10:14', '快件已到达【四川仪陇金城镇服务点】扫描员是【建设路营业部】', '', '2020-09-08 19:49:55', '', '2020-09-08 19:49:55');
INSERT INTO `logistics_tbl` VALUES (499, '773054739511845', '2020-09-08 14:10:29', '【仪陇金城镇建设路共配站S】小件员【罗豹】正在为您派件，如有疑问请联系小件员，联系电话【18380774878】', '', '2020-09-08 19:49:55', '', '2020-09-08 19:49:55');
INSERT INTO `logistics_tbl` VALUES (500, '773054739511845', '2020-09-08 14:10:59', '快件已被【仪陇金城镇建设路共配站S】【建设路交警大队公交旁】站点代收，请及时取件。有问题请联系【18881787239】', '', '2020-09-08 19:49:55', '', '2020-09-08 19:49:55');
INSERT INTO `logistics_tbl` VALUES (501, '773054585661279', '2020-09-08 14:56:32', '快件已到达【安徽马鞍山映翠分部】扫描员是【欧尚分部】', '', '2020-09-08 19:50:13', '', '2020-09-08 19:50:13');
INSERT INTO `logistics_tbl` VALUES (502, '773054585661279', '2020-09-08 15:03:11', '【安徽马鞍山映翠分部】的派件员【马钢花园快递之家(15715553677)】正在为您派件，如有疑问请联系派件员，联系电话【15715553677】', '', '2020-09-08 19:50:13', '', '2020-09-08 19:50:13');
INSERT INTO `logistics_tbl` VALUES (503, '773054585661279', '2020-09-08 17:20:53', '已签收，签收人(15715553677)是【钢城花园快递之家】,如有疑问请联系:15715553677,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-09-08 19:50:13', '', '2020-09-08 19:50:13');
INSERT INTO `logistics_tbl` VALUES (504, 'JDVD00936215561', '2020-09-08 14:03:12', '您的快件已发车', '', '2020-09-08 19:50:44', '', '2020-09-08 19:50:44');
INSERT INTO `logistics_tbl` VALUES (505, 'JDVD00934692194', '2020-09-07 02:54:32', '您的快件已发车', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (506, 'JDVD00934692194', '2020-09-07 10:43:17', '您的快件已到达【重庆分拣中心】', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (507, 'JDVD00934692194', '2020-09-07 10:53:38', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (508, 'JDVD00934692194', '2020-09-07 10:53:43', '您的快件由【重庆分拣中心】准备发往【重庆渝北营业部】', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (509, 'JDVD00934692194', '2020-09-07 11:24:43', '您的快件已发车', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (510, 'JDVD00934692194', '2020-09-07 13:21:13', '您的快件已到达【重庆渝北营业部】', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (511, 'JDVD00934692194', '2020-09-07 13:40:31', '您的快件正在派送中，请您准备签收（快递员：温鹏，联系电话：17815289082）', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (512, 'JDVD00934692194', '2020-09-07 15:49:42', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 20:35:15', '', '2020-09-08 20:35:15');
INSERT INTO `logistics_tbl` VALUES (513, 'JDVD00936221512', '2020-09-08 07:25:52', '您的快件已到达【苏州昆山分拣中心】', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (514, 'JDVD00936221512', '2020-09-08 07:25:52', '苏州昆山分拣中心分拣中心已收箱，箱号JDVD00936221512-1-1-', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (515, 'JDVD00936221512', '2020-09-08 09:58:20', '您的快件在【苏州昆山分拣中心】分拣完成', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (516, 'JDVD00936221512', '2020-09-08 09:58:25', '您的快件由【苏州昆山分拣中心】准备发往【苏州淀山分拣中心】', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (517, 'JDVD00936221512', '2020-09-08 10:31:50', '您的快件已发车', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (518, 'JDVD00936221512', '2020-09-08 11:12:49', '您的快件已到达【苏州淀山分拣中心】', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (519, 'JDVD00936221512', '2020-09-08 11:20:40', '您的快件在【苏州淀山分拣中心】分拣完成', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (520, 'JDVD00936221512', '2020-09-08 11:20:45', '您的快件由【苏州淀山分拣中心】准备发往【上海真沪营业部】', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (521, 'JDVD00936221512', '2020-09-08 13:28:46', '您的快件已发车', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (522, 'JDVD00936221512', '2020-09-08 15:32:03', '您的快件在【上海真沪营业部】收货完成', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (523, 'JDVD00936221512', '2020-09-08 15:32:04', '您的快件已到达【上海真沪营业部】', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (524, 'JDVD00936221512', '2020-09-08 16:05:43', '您的快件正在派送中，请您准备签收（快递员：田始通，联系电话：15165886690）', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (525, 'JDVD00936221512', '2020-09-08 16:41:51', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 22:11:56', '', '2020-09-08 22:11:56');
INSERT INTO `logistics_tbl` VALUES (526, 'JDVD00936213420', '2020-09-07 01:24:02', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (527, 'JDVD00936213420', '2020-09-07 05:04:13', '您的快件由【成都青白江分拣中心】准备发往【成都滨河营业部】', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (528, 'JDVD00936213420', '2020-09-07 05:49:13', '您的快件已发车', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (529, 'JDVD00936213420', '2020-09-07 07:21:00', '您的快件在【成都滨河营业部】收货完成', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (530, 'JDVD00936213420', '2020-09-07 07:21:01', '您的快件已到达【成都滨河营业部】', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (531, 'JDVD00936213420', '2020-09-07 08:05:40', '您的快件正在派送中，请您准备签收（快递员：毛磊，联系电话：15183233267）', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (532, 'JDVD00936213420', '2020-09-07 11:08:18', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-08 22:11:58', '', '2020-09-08 22:11:58');
INSERT INTO `logistics_tbl` VALUES (533, 'YT4767033732744', '2020-09-06 21:32:06', '南充转运中心 已收入', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (534, 'YT4767033732744', '2020-09-06 21:58:11', '南充转运中心 已发出 下一站  成都转运中心', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (535, 'YT4767033732744', '2020-09-07 17:26:50', '成都转运中心 已发出 下一站  贵阳转运中心', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (536, 'YT4767033732744', '2020-09-08 08:19:45', '贵阳转运中心 已收入', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (537, 'YT4767033732744', '2020-09-08 08:42:46', '贵阳转运中心 已发出 下一站  贵州省贵阳市观山湖区三部', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (538, 'YT4767033732744', '2020-09-08 14:40:45', '贵州省贵阳市观山湖区三部 已收入', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (539, 'YT4767033732744', '2020-09-08 15:22:37', '贵州省贵阳市观山湖区三部 派件中 派件人: 芶明兴 电话 15284662011', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (540, 'YT4767033732744', '2020-09-08 16:38:08', '快件已暂存至贵阳会展城B区东座店菜鸟驿站，如有疑问请联系19185000221', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (541, 'YT4767033732744', '2020-09-08 18:05:52', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务', '', '2020-09-09 13:01:14', '', '2020-09-09 13:01:14');
INSERT INTO `logistics_tbl` VALUES (542, 'YT4765645114700', '2020-09-06 21:41:53', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (543, 'YT4765645114700', '2020-09-06 22:02:38', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (544, 'YT4765645114700', '2020-09-07 04:13:34', '【成都转运中心公司】 已收入', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (545, 'YT4765645114700', '2020-09-07 04:38:23', '【成都转运中心】 已发出 下一站 【四川省成都市双流区华阳二部公司】', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (546, 'YT4765645114700', '2020-09-07 06:47:50', '【四川省成都市双流区华阳二部公司】 已收入', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (547, 'YT4765645114700', '2020-09-07 07:00:19', '【四川省成都市双流区华阳二部公司】 派件中  派件人: 何显雄 电话 18081120174  如有疑问，请联系：18521172098', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (548, 'YT4765645114700', '2020-09-07 08:09:38', '快件已由油建新区活动中心旁丰巢柜丰巢柜代收，取件码已发送，请及时取件。', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (549, 'YT4765645114700', '2020-09-07 20:50:50', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18081120174，投诉电话：18521172098', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (550, 'YT4766119173574', '2020-09-06 21:23:40', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (551, 'YT4766119173574', '2020-09-06 21:52:10', '【南充转运中心】 已发出 下一站 【四川省广安市公司】', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (552, 'YT4766119173574', '2020-09-07 06:58:55', '【四川省广安市公司】 已收入', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (553, 'YT4766119173574', '2020-09-07 08:04:24', '【四川省广安市公司】 派件中  派件人: 杨海军 电话 19961159176  如有疑问，请联系：0826-2345387', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (554, 'YT4766119173574', '2020-09-07 11:18:55', '客户签收人: 家人 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：19961159176，投诉电话：0826-2345387', '', '2020-09-09 13:01:26', '', '2020-09-09 13:01:26');
INSERT INTO `logistics_tbl` VALUES (555, 'JDVD00934683078', '2020-09-07 13:21:50', '您的快件已到达【重庆渝北营业部】', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (556, 'JDVD00934683078', '2020-09-07 13:40:41', '您的快件正在派送中，请您准备签收（快递员：温鹏，联系电话：17815289082）', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (557, 'JDVD00934683078', '2020-09-07 15:49:42', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (558, 'YT4767034499495', '2020-09-06 21:32:12', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (559, 'YT4767034499495', '2020-09-06 21:54:14', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (560, 'YT4767034499495', '2020-09-07 03:02:45', '【重庆转运中心公司】 已收入', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (561, 'YT4767034499495', '2020-09-07 03:32:41', '【重庆转运中心】 已发出 下一站 【重庆市江北区公司】', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (562, 'YT4767034499495', '2020-09-07 07:23:08', '【重庆市江北区公司】 已收入', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (563, 'YT4767034499495', '2020-09-07 09:51:51', '【重庆市江北区公司】 派件中  派件人: 黎钱勇 电话 17205126668  如有疑问，请联系：023-67955558', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (564, 'YT4767034499495', '2020-09-07 11:10:42', '快件已由重庆银行总行停车场入口旁丰巢柜丰巢柜代收，取件码已发送，请及时取件。', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (565, 'YT4767034499495', '2020-09-08 06:55:27', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：17265126668，投诉电话：023-67955558', '', '2020-09-09 13:01:27', '', '2020-09-09 13:01:27');
INSERT INTO `logistics_tbl` VALUES (566, 'YT4767033731540', '2020-09-06 21:27:49', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (567, 'YT4767033731540', '2020-09-06 21:54:24', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (568, 'YT4767033731540', '2020-09-07 03:54:56', '【重庆转运中心公司】 已收入', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (569, 'YT4767033731540', '2020-09-07 04:22:41', '【重庆转运中心】 已发出 下一站 【重庆市渝北区人和公司】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (570, 'YT4767033731540', '2020-09-07 06:56:36', '【重庆市渝北区人和公司】 已收入', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (571, 'YT4767033731540', '2020-09-07 07:41:11', '【重庆市渝北区人和公司】 派件中  派件人: 廖渝洲 电话 18521170280  如有疑问，请联系：023-62593882', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (572, 'YT4767033731540', '2020-09-07 16:09:17', '快件已由涉外商务区B1栋负一楼电梯口丰巢柜丰巢柜代收，取件码已发送，请及时取件。', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (573, 'YT4767033731540', '2020-09-07 21:20:04', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170280，投诉电话：023-62593882', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (574, 'YT4767030836799', '2020-09-06 21:25:11', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (575, 'YT4767030836799', '2020-09-06 21:54:20', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (576, 'YT4767030836799', '2020-09-07 02:40:05', '【重庆转运中心公司】 已收入', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (577, 'YT4767030836799', '2020-09-07 03:03:05', '【重庆转运中心】 已发出 下一站 【重庆市大学城公司】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (578, 'YT4767030836799', '2020-09-07 06:34:31', '【重庆市大学城公司】 已收入', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (579, 'YT4767030836799', '2020-09-07 06:37:44', '【重庆市大学城公司】 派件中  派件人: 圆通张玉强 电话 15023679871 ', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (580, 'YT4767030836799', '2020-09-07 16:30:27', '客户签收人: 前台 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：15023679871', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (581, '773054585245568', '2020-09-06 18:37:56', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (582, '773054585245568', '2020-09-06 20:47:08', '快件已到达【四川南充转运中心】扫描员是【兰玉风】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (583, '773054585245568', '2020-09-06 20:57:09', '快件由【四川南充转运中心】发往【四川成都转运中心】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (584, '773054585245568', '2020-09-07 04:59:33', '快件已到达【四川成都转运中心】扫描员是【出港龙仁冯伟伦】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (585, '773054585245568', '2020-09-07 05:08:31', '快件由【四川成都转运中心】发往【四川成都双丰公司】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (586, '773054585245568', '2020-09-07 09:03:01', '快件已到达【四川成都双丰公司】扫描员是【陈敏】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (587, '773054585245568', '2020-09-07 09:09:57', '快件已到达【四川成都双丰公司】扫描员是【李键】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (588, '773054585245568', '2020-09-07 09:10:19', '【四川成都双丰公司】的派件员【陈敏(15828255260)】正在为您派件，如有疑问请联系派件员，联系电话【15828255260】', '', '2020-09-09 13:01:28', '', '2020-09-09 13:01:28');
INSERT INTO `logistics_tbl` VALUES (589, '773054585245568', '2020-09-07 10:11:56', '快递已被【富有收件宝快递柜】成都市武侯区龙腾东路7号华馨名屋中庭富友快递柜代收，请及时取件。有问题请联系18200227625', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (590, '773054585245568', '2020-09-07 18:43:48', '已签收，签收人凭取货码签收', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (591, '773054585582825', '2020-09-09 06:23:27', '快件已到达【上海转运中心】扫描员是【宋永宣】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (592, '773054585582825', '2020-09-09 06:30:45', '快件由【上海转运中心】发往【上海宝山转运中心】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (593, '773054585582825', '2020-09-09 09:25:55', '快件已到达【上海宝山转运中心】扫描员是【汪波】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (594, '773054585582825', '2020-09-09 09:37:10', '快件由【上海宝山转运中心】发往【上海宝山大场公司】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (595, '773054585582825', '2020-09-09 10:25:55', '快件已到达【上海宝山大场公司】扫描员是【大场扫描】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (596, 'YT4768216385233', '2020-09-07 22:27:58', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (597, 'YT4768216385233', '2020-09-07 22:52:57', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (598, 'YT4768216385233', '2020-09-08 03:33:18', '【重庆转运中心公司】 已收入', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (599, 'YT4768216385233', '2020-09-08 03:53:27', '【重庆转运中心】 已发出 下一站 【潍坊转运中心公司】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (600, 'YT4768216385233', '2020-09-09 07:17:28', '【潍坊转运中心公司】 已收入', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (601, 'YT4768216385233', '2020-09-09 07:45:12', '【潍坊转运中心】 已发出 下一站 【青岛转运中心公司】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (602, 'YT4768216385233', '2020-09-09 11:34:18', '【青岛转运中心公司】 已收入', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (603, 'YT4768216385233', '2020-09-09 11:55:42', '【青岛转运中心】 已发出 下一站 【山东省青岛市崂山区公司】', '', '2020-09-09 13:01:29', '', '2020-09-09 13:01:29');
INSERT INTO `logistics_tbl` VALUES (604, 'YT4768216384701', '2020-09-07 22:25:51', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (605, 'YT4768216384701', '2020-09-07 22:55:15', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (606, 'YT4768216384701', '2020-09-08 03:13:21', '【重庆转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (607, 'YT4768216384701', '2020-09-08 03:35:49', '【重庆转运中心】 已发出 下一站 【潍坊转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (608, 'YT4768216384701', '2020-09-09 07:22:24', '【潍坊转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (609, 'YT4768216384701', '2020-09-09 07:47:17', '【潍坊转运中心】 已发出 下一站 【青岛转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (610, 'YT4768216384701', '2020-09-09 11:39:58', '【青岛转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (611, 'YT4768216384701', '2020-09-09 12:02:21', '【青岛转运中心】 已发出 下一站 【山东省青岛市崂山区公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (612, 'YT4768219923622', '2020-09-07 22:28:36', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (613, 'YT4768219923622', '2020-09-07 22:53:44', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (614, 'YT4768219923622', '2020-09-08 03:31:30', '【重庆转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (615, 'YT4768219923622', '2020-09-08 03:53:38', '【重庆转运中心】 已发出 下一站 【潍坊转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (616, 'YT4768219923622', '2020-09-09 07:23:43', '【潍坊转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (617, 'YT4768219923622', '2020-09-09 07:48:55', '【潍坊转运中心】 已发出 下一站 【青岛转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (618, 'YT4768219923622', '2020-09-09 11:18:39', '【青岛转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (619, 'YT4768219923622', '2020-09-09 11:39:12', '【青岛转运中心】 已发出 下一站 【山东省青岛市李沧区李村公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (620, 'YT4770457089144', '2020-09-08 22:32:54', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (621, 'YT4770457089144', '2020-09-08 22:57:59', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (622, 'YT4770457089144', '2020-09-09 02:39:18', '【成都转运中心公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (623, 'YT4770457089144', '2020-09-09 03:06:30', '【成都转运中心】 已发出 下一站 【四川省成都市金府顺公司】', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (624, 'YT4770457089144', '2020-09-09 07:25:09', '【四川省成都市金府顺公司】 已收入', '', '2020-09-09 13:01:30', '', '2020-09-09 13:01:30');
INSERT INTO `logistics_tbl` VALUES (625, 'YT4770457089144', '2020-09-09 07:49:58', '【四川省成都市金府顺公司】 派件中  派件人: 贺川 电话 18521871490  如有疑问，请联系：028-65092999', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (626, 'YT4770457089144', '2020-09-09 10:31:42', '圆通合作点【兔喜快递超市】快件已到达丽都路538号3栋2单元架空层驿站,联系电话17360048900', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (627, 'YT4770460243366', '2020-09-08 21:17:54', '【南充转运中心公司】 已收入', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (628, 'YT4770460243366', '2020-09-08 21:38:08', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (629, 'YT4770460243366', '2020-09-09 03:21:59', '【成都转运中心公司】 已收入', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (630, 'YT4770460243366', '2020-09-09 03:44:01', '【成都转运中心】 已发出 下一站 【四川省成都市金府顺公司】', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (631, 'YT4770460243366', '2020-09-09 07:17:14', '【四川省成都市金府顺公司】 已收入', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (632, 'YT4770460243366', '2020-09-09 07:42:58', '【四川省成都市金府顺公司】 派件中  派件人: 贺川 电话 18521871490  如有疑问，请联系：028-65092999', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (633, 'YT4770460243366', '2020-09-09 10:31:42', '圆通合作点【兔喜快递超市】快件已到达丽都路538号3栋2单元架空层驿站,联系电话17360048900', '', '2020-09-09 13:01:31', '', '2020-09-09 13:01:31');
INSERT INTO `logistics_tbl` VALUES (634, 'YT4773416999820', '2020-09-09 13:47:07', '添加快递单', 'admin', '2020-09-09 13:47:07', '', '2020-09-09 13:47:07');
INSERT INTO `logistics_tbl` VALUES (635, 'YT4773415789541', '2020-09-09 13:47:18', '添加快递单', 'admin', '2020-09-09 13:47:18', '', '2020-09-09 13:47:18');
INSERT INTO `logistics_tbl` VALUES (636, 'YT4773419880126', '2020-09-09 13:47:32', '添加快递单', 'admin', '2020-09-09 13:47:32', '', '2020-09-09 13:47:32');
INSERT INTO `logistics_tbl` VALUES (637, 'YT4773417016273', '2020-09-09 13:47:40', '添加快递单', 'admin', '2020-09-09 13:47:40', '', '2020-09-09 13:47:40');
INSERT INTO `logistics_tbl` VALUES (638, 'YT4773419154635', '2020-09-09 13:47:54', '添加快递单', 'admin', '2020-09-09 13:47:54', '', '2020-09-09 13:47:54');
INSERT INTO `logistics_tbl` VALUES (639, 'YT4773417656935', '2020-09-09 13:48:01', '添加快递单', 'admin', '2020-09-09 13:48:01', '', '2020-09-09 13:48:01');
INSERT INTO `logistics_tbl` VALUES (640, 'YT4773418439870', '2020-09-09 13:48:12', '添加快递单', 'admin', '2020-09-09 13:48:12', '', '2020-09-09 13:48:12');
INSERT INTO `logistics_tbl` VALUES (641, 'YT4773419910786', '2020-09-09 13:48:17', '添加快递单', 'admin', '2020-09-09 13:48:17', '', '2020-09-09 13:48:17');
INSERT INTO `logistics_tbl` VALUES (642, 'YT4773419902199', '2020-09-09 13:48:23', '添加快递单', 'admin', '2020-09-09 13:48:23', '', '2020-09-09 13:48:23');
INSERT INTO `logistics_tbl` VALUES (643, 'YT4773418461886', '2020-09-09 13:48:35', '添加快递单', 'admin', '2020-09-09 13:48:35', '', '2020-09-09 13:48:35');
INSERT INTO `logistics_tbl` VALUES (644, 'YT4773418460044', '2020-09-09 13:48:40', '添加快递单', 'admin', '2020-09-09 13:48:40', '', '2020-09-09 13:48:40');
INSERT INTO `logistics_tbl` VALUES (645, 'YT4773415806150', '2020-09-09 13:48:52', '添加快递单', 'admin', '2020-09-09 13:48:52', '', '2020-09-09 13:48:52');
INSERT INTO `logistics_tbl` VALUES (646, 'YT4773415801755', '2020-09-09 13:48:57', '添加快递单', 'admin', '2020-09-09 13:48:57', '', '2020-09-09 13:48:57');
INSERT INTO `logistics_tbl` VALUES (647, 'YT4773416451220', '2020-09-09 13:49:03', '添加快递单', 'admin', '2020-09-09 13:49:03', '', '2020-09-09 13:49:03');
INSERT INTO `logistics_tbl` VALUES (648, 'YT4773419169365', '2020-09-09 13:49:08', '添加快递单', 'admin', '2020-09-09 13:49:08', '', '2020-09-09 13:49:08');
INSERT INTO `logistics_tbl` VALUES (649, 'YT4773419890573', '2020-09-09 13:49:17', '添加快递单', 'admin', '2020-09-09 13:49:17', '', '2020-09-09 13:49:17');
INSERT INTO `logistics_tbl` VALUES (650, 'YT4773419172800', '2020-09-09 13:49:22', '添加快递单', 'admin', '2020-09-09 13:49:22', '', '2020-09-09 13:49:22');
INSERT INTO `logistics_tbl` VALUES (651, 'YT4773417677734', '2020-09-09 13:49:27', '添加快递单', 'admin', '2020-09-09 13:49:27', '', '2020-09-09 13:49:27');
INSERT INTO `logistics_tbl` VALUES (652, 'YT4773417665418', '2020-09-09 13:49:31', '添加快递单', 'admin', '2020-09-09 13:49:31', '', '2020-09-09 13:49:31');
INSERT INTO `logistics_tbl` VALUES (653, 'YT4773419884144', '2020-09-09 13:49:36', '添加快递单', 'admin', '2020-09-09 13:49:36', '', '2020-09-09 13:49:36');
INSERT INTO `logistics_tbl` VALUES (654, 'YT4768219923622', '2020-09-09 13:09:37', '【山东省青岛市李沧区宝龙广场分部公司】 已收入', '', '2020-09-09 14:43:43', '', '2020-09-09 14:43:43');
INSERT INTO `logistics_tbl` VALUES (655, 'YT4768219923622', '2020-09-09 13:46:53', '【山东省青岛市李沧区宝龙广场分部公司】 派件中  派件人: 于洋 电话 17669795331  如有疑问，请联系：18521846242', '', '2020-09-09 14:43:43', '', '2020-09-09 14:43:43');
INSERT INTO `logistics_tbl` VALUES (656, 'YT4768219923622', '2020-09-09 14:14:14', '快件已暂存至青岛海尔鼎世华府B区10号楼店菜鸟驿站，如有疑问请联系13165077850', '', '2020-09-09 14:43:43', '', '2020-09-09 14:43:43');
INSERT INTO `logistics_tbl` VALUES (657, 'YT4768216384701', '2020-09-09 14:18:41', '【山东省青岛市崂山区公司】 派件中  派件人: 刘祥南 电话 15336485120  如有疑问，请联系：0532-80992611', '', '2020-09-09 14:43:51', '', '2020-09-09 14:43:51');
INSERT INTO `logistics_tbl` VALUES (658, 'YT4768216385233', '2020-09-09 14:05:11', '【山东省青岛市崂山区公司】 派件中  派件人: 刘祥南 电话 15336485120  如有疑问，请联系：0532-80992611', '', '2020-09-09 14:43:51', '', '2020-09-09 14:43:51');
INSERT INTO `logistics_tbl` VALUES (659, 'YT4774201651604', '2020-09-09 16:31:26', '添加快递单', 'admin', '2020-09-09 16:31:26', '', '2020-09-09 16:31:26');
INSERT INTO `logistics_tbl` VALUES (660, 'JDVD00936215561', '2020-09-09 10:14:26', '您的快件正在派送中，请您准备签收（快递员：贾华林，联系电话：18010830550）', '', '2020-09-09 18:38:25', '', '2020-09-09 18:38:25');
INSERT INTO `logistics_tbl` VALUES (661, 'JDVD00936215561', '2020-09-09 12:34:54', '您的快件已由亲属代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-09 18:38:25', '', '2020-09-09 18:38:25');
INSERT INTO `logistics_tbl` VALUES (662, '773054904635511', '2020-09-09 17:07:39', '【四川苍溪公司】的收件员【办公室1(18784946469)】已收件', '', '2020-09-09 22:11:02', '', '2020-09-09 22:11:02');
INSERT INTO `logistics_tbl` VALUES (663, '773054904635511', '2020-09-09 18:31:56', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-09 22:11:02', '', '2020-09-09 22:11:02');
INSERT INTO `logistics_tbl` VALUES (664, '773054906506161', '2020-09-09 17:07:43', '【四川苍溪公司】的收件员【办公室1(18784946469)】已收件', '', '2020-09-09 22:11:02', '', '2020-09-09 22:11:02');
INSERT INTO `logistics_tbl` VALUES (665, '773054906506161', '2020-09-09 18:30:27', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-09 22:11:02', '', '2020-09-09 22:11:02');
INSERT INTO `logistics_tbl` VALUES (666, '773054904868479', '2020-09-09 17:07:42', '【四川苍溪公司】的收件员【办公室1(18784946469)】已收件', '', '2020-09-09 22:11:21', '', '2020-09-09 22:11:21');
INSERT INTO `logistics_tbl` VALUES (667, '773054904868479', '2020-09-09 18:41:07', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-09 22:11:21', '', '2020-09-09 22:11:21');
INSERT INTO `logistics_tbl` VALUES (668, '773054585582825', '2020-09-09 14:28:08', '快件已到达【上海宝山大场公司】扫描员是【王锦凯】', '', '2020-09-09 22:34:15', '', '2020-09-09 22:34:15');
INSERT INTO `logistics_tbl` VALUES (669, '773054585582825', '2020-09-09 14:38:08', '【上海宝山大场公司】的派件员【王锦凯(13564768019)】正在为您派件，如有疑问请联系派件员，联系电话【13564768019】', '', '2020-09-09 22:34:15', '', '2020-09-09 22:34:15');
INSERT INTO `logistics_tbl` VALUES (670, '773054585582825', '2020-09-09 17:10:16', '快递已被【丰巢智能柜】上海市宝山区经纬学府阳光家园22号楼2号柜代收，请及时取件。有问题请联系13564768019', '', '2020-09-09 22:34:16', '', '2020-09-09 22:34:16');
INSERT INTO `logistics_tbl` VALUES (671, 'YT4768216385233', '2020-09-09 19:09:34', '客户签收人: 快递超市 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：15336485120，投诉电话：0532-80992611', '', '2020-09-09 22:34:17', '', '2020-09-09 22:34:17');
INSERT INTO `logistics_tbl` VALUES (672, 'YT4768216384701', '2020-09-09 19:09:37', '客户签收人: 快递超市 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：15336485120，投诉电话：0532-80992611', '', '2020-09-09 22:34:18', '', '2020-09-09 22:34:18');
INSERT INTO `logistics_tbl` VALUES (673, 'YT4768219923622', '2020-09-09 17:30:52', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：17669795331，投诉电话：18521846242', '', '2020-09-09 22:34:18', '', '2020-09-09 22:34:18');
INSERT INTO `logistics_tbl` VALUES (674, '773054906075645', '2020-09-09 17:01:33', '【四川苍溪公司】的收件员【办公室1(18784946469)】已收件', '', '2020-09-09 22:34:19', '', '2020-09-09 22:34:19');
INSERT INTO `logistics_tbl` VALUES (675, '773054906075645', '2020-09-09 18:41:08', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-09 22:34:20', '', '2020-09-09 22:34:20');
INSERT INTO `logistics_tbl` VALUES (676, '773054906427998', '2020-09-09 17:01:37', '【四川苍溪公司】的收件员【办公室1(18784946469)】已收件', '', '2020-09-09 22:34:20', '', '2020-09-09 22:34:20');
INSERT INTO `logistics_tbl` VALUES (677, '773054906427998', '2020-09-09 18:35:12', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-09 22:34:20', '', '2020-09-09 22:34:20');
INSERT INTO `logistics_tbl` VALUES (678, 'YT4773416999820', '2020-09-09 16:22:14', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:20', '', '2020-09-09 22:34:20');
INSERT INTO `logistics_tbl` VALUES (679, 'YT4773415789541', '2020-09-09 16:25:19', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (680, 'YT4773415789541', '2020-09-09 21:09:37', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (681, 'YT4773415789541', '2020-09-09 21:37:04', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (682, 'YT4773419880126', '2020-09-09 16:25:56', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (683, 'YT4773419880126', '2020-09-09 21:09:56', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (684, 'YT4773419880126', '2020-09-09 21:38:19', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (685, 'YT4773417016273', '2020-09-09 16:22:16', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (686, 'YT4773417016273', '2020-09-09 21:21:27', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (687, 'YT4773417016273', '2020-09-09 21:41:28', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:21', '', '2020-09-09 22:34:21');
INSERT INTO `logistics_tbl` VALUES (688, 'YT4773419154635', '2020-09-09 16:25:12', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (689, 'YT4773419154635', '2020-09-09 21:07:42', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (690, 'YT4773419154635', '2020-09-09 21:34:44', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (691, 'YT4773417656935', '2020-09-09 16:26:02', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (692, 'YT4773417656935', '2020-09-09 21:06:13', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (693, 'YT4773417656935', '2020-09-09 21:35:44', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (694, 'YT4773418439870', '2020-09-09 16:25:43', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (695, 'YT4773418439870', '2020-09-09 21:09:48', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (696, 'YT4773418439870', '2020-09-09 21:38:22', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:22', '', '2020-09-09 22:34:22');
INSERT INTO `logistics_tbl` VALUES (697, 'YT4773419910786', '2020-09-09 16:25:32', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (698, 'YT4773419910786', '2020-09-09 21:14:14', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (699, 'YT4773419910786', '2020-09-09 21:37:22', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (700, 'YT4773419902199', '2020-09-09 16:25:41', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (701, 'YT4773419902199', '2020-09-09 21:07:19', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (702, 'YT4773419902199', '2020-09-09 21:36:30', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (703, 'YT4773418461886', '2020-09-09 16:25:39', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:23', '', '2020-09-09 22:34:23');
INSERT INTO `logistics_tbl` VALUES (704, 'YT4773418460044', '2020-09-09 16:25:27', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:24', '', '2020-09-09 22:34:24');
INSERT INTO `logistics_tbl` VALUES (705, 'YT4773418460044', '2020-09-09 21:06:41', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:24', '', '2020-09-09 22:34:24');
INSERT INTO `logistics_tbl` VALUES (706, 'YT4773418460044', '2020-09-09 21:36:34', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:24', '', '2020-09-09 22:34:24');
INSERT INTO `logistics_tbl` VALUES (707, 'YT4773415806150', '2020-09-09 16:25:51', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:24', '', '2020-09-09 22:34:24');
INSERT INTO `logistics_tbl` VALUES (708, 'YT4773415806150', '2020-09-09 21:16:33', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:24', '', '2020-09-09 22:34:24');
INSERT INTO `logistics_tbl` VALUES (709, 'YT4773415806150', '2020-09-09 21:36:51', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:24', '', '2020-09-09 22:34:24');
INSERT INTO `logistics_tbl` VALUES (710, 'YT4773415801755', '2020-09-09 16:23:15', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (711, 'YT4773415801755', '2020-09-09 21:12:01', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (712, 'YT4773415801755', '2020-09-09 21:39:58', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (713, 'YT4773416451220', '2020-09-09 16:23:21', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (714, 'YT4773416451220', '2020-09-09 21:13:26', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (715, 'YT4773416451220', '2020-09-09 21:40:11', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (716, 'YT4773419169365', '2020-09-09 16:25:55', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (717, 'YT4773419169365', '2020-09-09 21:07:25', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (718, 'YT4773419169365', '2020-09-09 21:37:14', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:25', '', '2020-09-09 22:34:25');
INSERT INTO `logistics_tbl` VALUES (719, 'YT4773419890573', '2020-09-09 16:25:48', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:26', '', '2020-09-09 22:34:26');
INSERT INTO `logistics_tbl` VALUES (720, 'YT4773419890573', '2020-09-09 21:13:33', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:26', '', '2020-09-09 22:34:26');
INSERT INTO `logistics_tbl` VALUES (721, 'YT4773419890573', '2020-09-09 21:36:51', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:26', '', '2020-09-09 22:34:26');
INSERT INTO `logistics_tbl` VALUES (722, 'YT4773419172800', '2020-09-09 16:25:45', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (723, 'YT4773419172800', '2020-09-09 21:06:36', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (724, 'YT4773419172800', '2020-09-09 21:36:06', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (725, 'YT4773417677734', '2020-09-09 16:22:02', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (726, 'YT4773417677734', '2020-09-09 21:22:14', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (727, 'YT4773417677734', '2020-09-09 21:43:15', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (728, 'YT4773417665418', '2020-09-09 16:23:16', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (729, 'YT4773417665418', '2020-09-09 21:12:16', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (730, 'YT4773417665418', '2020-09-09 21:33:23', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (731, 'YT4773419884144', '2020-09-09 16:22:59', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (732, 'YT4773419884144', '2020-09-09 21:20:34', '【南充转运中心公司】 已收入', '', '2020-09-09 22:34:27', '', '2020-09-09 22:34:27');
INSERT INTO `logistics_tbl` VALUES (733, 'YT4773419884144', '2020-09-09 21:40:37', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-09 22:34:28', '', '2020-09-09 22:34:28');
INSERT INTO `logistics_tbl` VALUES (734, 'YT4774201651604', '2020-09-09 17:57:55', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-09 22:34:28', '', '2020-09-09 22:34:28');
INSERT INTO `logistics_tbl` VALUES (735, '773054906506161', '2020-09-09 22:42:30', '快件已到达【四川南充转运中心】扫描员是【临时工】', '', '2020-09-09 23:16:26', '', '2020-09-09 23:16:26');
INSERT INTO `logistics_tbl` VALUES (736, '773054906506161', '2020-09-09 22:54:30', '快件由【四川南充转运中心】发往【上海转运中心】', '', '2020-09-09 23:16:26', '', '2020-09-09 23:16:26');
INSERT INTO `logistics_tbl` VALUES (737, '773054904868479', '2020-09-09 22:44:36', '快件已到达【四川南充转运中心】扫描员是【顶扫五】', '', '2020-09-09 23:16:26', '', '2020-09-09 23:16:26');
INSERT INTO `logistics_tbl` VALUES (738, '773054904868479', '2020-09-09 22:52:18', '快件由【四川南充转运中心】发往【四川自贡转运中心】', '', '2020-09-09 23:16:26', '', '2020-09-09 23:16:26');
INSERT INTO `logistics_tbl` VALUES (739, '773054906075645', '2020-09-09 22:46:08', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-09 23:16:27', '', '2020-09-09 23:16:27');
INSERT INTO `logistics_tbl` VALUES (740, '773054906075645', '2020-09-09 22:51:28', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-09 23:16:27', '', '2020-09-09 23:16:27');
INSERT INTO `logistics_tbl` VALUES (741, '773054906427998', '2020-09-09 22:49:09', '快件已到达【四川南充转运中心】扫描员是【余群】', '', '2020-09-09 23:16:27', '', '2020-09-09 23:16:27');
INSERT INTO `logistics_tbl` VALUES (742, '773054906427998', '2020-09-09 22:55:12', '快件由【四川南充转运中心】发往【重庆转运中心】', '', '2020-09-09 23:16:27', '', '2020-09-09 23:16:27');
INSERT INTO `logistics_tbl` VALUES (743, 'YT4773416999820', '2020-09-09 22:21:00', '南充转运中心 已收入', '', '2020-09-09 23:16:27', '', '2020-09-09 23:16:27');
INSERT INTO `logistics_tbl` VALUES (744, 'YT4773416999820', '2020-09-09 22:41:46', '南充转运中心 已发出 下一站  重庆转运中心', '', '2020-09-09 23:16:27', '', '2020-09-09 23:16:27');
INSERT INTO `logistics_tbl` VALUES (745, 'YT4775134367550', '2020-09-10 09:01:00', '添加快递单', 'admin', '2020-09-10 09:01:00', '', '2020-09-10 09:01:00');
INSERT INTO `logistics_tbl` VALUES (746, 'YT4775131545884', '2020-09-10 09:14:45', '添加快递单', 'admin', '2020-09-10 09:14:45', '', '2020-09-10 09:14:45');
INSERT INTO `logistics_tbl` VALUES (747, 'YT4775132833220', '2020-09-10 09:15:03', '添加快递单', 'admin', '2020-09-10 09:15:03', '', '2020-09-10 09:15:03');
INSERT INTO `logistics_tbl` VALUES (748, 'YT4775132837766', '2020-09-10 09:15:23', '添加快递单', 'admin', '2020-09-10 09:15:23', '', '2020-09-10 09:15:23');
INSERT INTO `logistics_tbl` VALUES (749, 'YT4775135853366', '2020-09-10 09:15:46', '添加快递单', 'admin', '2020-09-10 09:15:46', '', '2020-09-10 09:15:46');
INSERT INTO `logistics_tbl` VALUES (750, 'YT4775135848765', '2020-09-10 09:16:20', '添加快递单', 'admin', '2020-09-10 09:16:20', '', '2020-09-10 09:16:20');
INSERT INTO `logistics_tbl` VALUES (751, 'YT4775134358570', '2020-09-10 09:17:54', '添加快递单', 'admin', '2020-09-10 09:17:54', '', '2020-09-10 09:17:54');
INSERT INTO `logistics_tbl` VALUES (752, 'YT4775135096174', '2020-09-10 09:18:08', '添加快递单', 'admin', '2020-09-10 09:18:08', '', '2020-09-10 09:18:08');
INSERT INTO `logistics_tbl` VALUES (753, 'YT4775135098761', '2020-09-10 09:18:16', '添加快递单', 'admin', '2020-09-10 09:18:16', '', '2020-09-10 09:18:16');
INSERT INTO `logistics_tbl` VALUES (754, 'YT4774201651604', '2020-09-10 01:26:38', '【南充转运中心公司】 已收入', '', '2020-09-10 10:29:21', '', '2020-09-10 10:29:21');
INSERT INTO `logistics_tbl` VALUES (755, 'YT4774201651604', '2020-09-10 01:49:34', '【南充转运中心】 已发出 下一站 【上海转运中心公司】', '', '2020-09-10 10:29:21', '', '2020-09-10 10:29:21');
INSERT INTO `logistics_tbl` VALUES (756, 'YT4770460243366', '2020-09-10 10:33:48', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521871490，投诉电话：028-65092999', '', '2020-09-10 12:09:39', '', '2020-09-10 12:09:39');
INSERT INTO `logistics_tbl` VALUES (757, 'YT4775134367550', '2020-09-10 13:49:18', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 15:31:50', '', '2020-09-10 15:31:50');
INSERT INTO `logistics_tbl` VALUES (758, '773054906427998', '2020-09-10 04:53:08', '快件已到达【重庆转运中心】扫描员是【A班AC组组长陈俊琼】', '', '2020-09-10 15:31:56', '', '2020-09-10 15:31:56');
INSERT INTO `logistics_tbl` VALUES (759, '773054906427998', '2020-09-10 04:58:11', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-10 15:31:56', '', '2020-09-10 15:31:56');
INSERT INTO `logistics_tbl` VALUES (760, '773054906427998', '2020-09-10 13:28:01', '快件已到达【重庆机场公司】扫描员是【周达】', '', '2020-09-10 15:31:56', '', '2020-09-10 15:31:56');
INSERT INTO `logistics_tbl` VALUES (761, '773054906427998', '2020-09-10 13:32:27', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-10 15:31:56', '', '2020-09-10 15:31:56');
INSERT INTO `logistics_tbl` VALUES (762, '773054904868479', '2020-09-10 03:43:37', '快件已到达【四川自贡转运中心】扫描员是【银敦英】', '', '2020-09-10 15:32:14', '', '2020-09-10 15:32:14');
INSERT INTO `logistics_tbl` VALUES (763, '773054904868479', '2020-09-10 03:51:50', '快件由【四川自贡转运中心】发往【广东深圳转运中心】', '', '2020-09-10 15:32:14', '', '2020-09-10 15:32:14');
INSERT INTO `logistics_tbl` VALUES (764, 'YT4773415789541', '2020-09-10 03:28:31', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (765, 'YT4773415789541', '2020-09-10 03:49:31', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (766, 'YT4773415789541', '2020-09-10 08:09:20', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (767, 'YT4773415789541', '2020-09-10 11:17:56', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (768, 'YT4773415801755', '2020-09-10 02:39:09', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (769, 'YT4773415801755', '2020-09-10 03:01:59', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (770, 'YT4773415801755', '2020-09-10 08:06:38', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (771, 'YT4773415801755', '2020-09-10 11:19:19', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:13', '', '2020-09-10 15:33:13');
INSERT INTO `logistics_tbl` VALUES (772, 'YT4773415806150', '2020-09-10 03:23:47', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (773, 'YT4773415806150', '2020-09-10 03:48:30', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (774, 'YT4773415806150', '2020-09-10 08:08:39', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (775, 'YT4773415806150', '2020-09-10 11:18:24', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (776, 'YT4773416451220', '2020-09-10 03:22:50', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (777, 'YT4773416451220', '2020-09-10 03:47:09', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (778, 'YT4773416451220', '2020-09-10 08:00:08', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (779, 'YT4773416451220', '2020-09-10 11:18:47', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (780, 'YT4773416999820', '2020-09-10 05:17:46', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (781, 'YT4773416999820', '2020-09-10 05:44:34', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (782, 'YT4773417016273', '2020-09-10 03:47:58', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (783, 'YT4773417016273', '2020-09-10 04:17:09', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (784, 'YT4773417016273', '2020-09-10 08:45:35', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (785, 'YT4773417016273', '2020-09-10 11:17:32', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:14', '', '2020-09-10 15:33:14');
INSERT INTO `logistics_tbl` VALUES (786, 'YT4773417656935', '2020-09-10 05:13:45', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (787, 'YT4773417656935', '2020-09-10 05:39:16', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (788, 'YT4773417665418', '2020-09-10 03:51:45', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (789, 'YT4773417665418', '2020-09-10 04:13:27', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (790, 'YT4773417665418', '2020-09-10 08:32:55', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (791, 'YT4773417665418', '2020-09-10 11:19:34', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (792, 'YT4773417677734', '2020-09-10 04:08:00', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (793, 'YT4773417677734', '2020-09-10 04:37:22', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (794, 'YT4773418439870', '2020-09-10 05:02:56', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (795, 'YT4773418439870', '2020-09-10 05:25:02', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (796, 'YT4773418439870', '2020-09-10 08:16:48', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (797, 'YT4773418439870', '2020-09-10 11:17:28', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (798, 'YT4773418460044', '2020-09-10 03:40:49', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (799, 'YT4773418460044', '2020-09-10 04:08:39', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (800, 'YT4773418460044', '2020-09-10 08:43:02', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (801, 'YT4773418460044', '2020-09-10 11:17:48', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:15', '', '2020-09-10 15:33:15');
INSERT INTO `logistics_tbl` VALUES (802, 'YT4773419154635', '2020-09-10 03:31:12', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (803, 'YT4773419154635', '2020-09-10 03:52:42', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (804, 'YT4773419154635', '2020-09-10 08:17:23', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (805, 'YT4773419154635', '2020-09-10 11:19:39', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (806, 'YT4773419169365', '2020-09-10 04:05:25', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (807, 'YT4773419169365', '2020-09-10 04:35:08', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (808, 'YT4773419172800', '2020-09-10 03:37:34', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (809, 'YT4773419172800', '2020-09-10 04:02:58', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (810, 'YT4773419172800', '2020-09-10 08:46:57', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (811, 'YT4773419172800', '2020-09-10 11:21:51', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (812, 'YT4773419880126', '2020-09-10 02:48:17', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (813, 'YT4773419880126', '2020-09-10 03:12:37', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (814, 'YT4773419880126', '2020-09-10 08:24:08', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (815, 'YT4773419880126', '2020-09-10 11:18:11', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (816, 'YT4773419884144', '2020-09-10 04:27:43', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (817, 'YT4773419884144', '2020-09-10 04:49:12', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (818, 'YT4773419890573', '2020-09-10 05:18:30', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (819, 'YT4773419890573', '2020-09-10 05:42:24', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (820, 'YT4773419890573', '2020-09-10 14:33:53', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:16', '', '2020-09-10 15:33:16');
INSERT INTO `logistics_tbl` VALUES (821, 'YT4773419902199', '2020-09-10 05:17:07', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (822, 'YT4773419902199', '2020-09-10 05:41:58', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (823, 'YT4773419902199', '2020-09-10 14:21:51', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (824, 'YT4773419910786', '2020-09-10 02:48:09', '【重庆转运中心公司】 已收入', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (825, 'YT4773419910786', '2020-09-10 03:11:26', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (826, 'YT4773419910786', '2020-09-10 08:24:13', '【重庆市巴南区公司】 已收入', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (827, 'YT4773419910786', '2020-09-10 11:17:23', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-10 15:33:17', '', '2020-09-10 15:33:17');
INSERT INTO `logistics_tbl` VALUES (828, 'YT4775134358570', '2020-09-10 13:48:32', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 15:33:30', '', '2020-09-10 15:33:30');
INSERT INTO `logistics_tbl` VALUES (829, 'YT4775132833220', '2020-09-10 13:49:27', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:22:57', '', '2020-09-10 18:22:57');
INSERT INTO `logistics_tbl` VALUES (830, 'YT4775131545884', '2020-09-10 13:48:29', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:23:11', '', '2020-09-10 18:23:11');
INSERT INTO `logistics_tbl` VALUES (831, 'YT4775132837766', '2020-09-10 13:49:14', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:23:17', '', '2020-09-10 18:23:17');
INSERT INTO `logistics_tbl` VALUES (832, 'YT4775135096174', '2020-09-10 13:49:07', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:23:21', '', '2020-09-10 18:23:21');
INSERT INTO `logistics_tbl` VALUES (833, 'YT4775135098761', '2020-09-10 13:48:35', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:24:07', '', '2020-09-10 18:24:07');
INSERT INTO `logistics_tbl` VALUES (834, 'YT4773417656935', '2020-09-10 15:58:19', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 17823750906  如有疑问，请联系：023-62593622', '', '2020-09-10 18:41:35', '', '2020-09-10 18:41:35');
INSERT INTO `logistics_tbl` VALUES (835, 'YT4773419902199', '2020-09-10 15:56:56', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 17823750906  如有疑问，请联系：023-62593622', '', '2020-09-10 18:41:36', '', '2020-09-10 18:41:36');
INSERT INTO `logistics_tbl` VALUES (836, 'YT4773419890573', '2020-09-10 15:58:14', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 17823750906  如有疑问，请联系：023-62593622', '', '2020-09-10 18:41:36', '', '2020-09-10 18:41:36');
INSERT INTO `logistics_tbl` VALUES (837, 'YT4775135853366', '2020-09-10 13:49:21', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:41:38', '', '2020-09-10 18:41:38');
INSERT INTO `logistics_tbl` VALUES (838, 'YT4775135848765', '2020-09-10 13:49:25', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-10 18:41:38', '', '2020-09-10 18:41:38');
INSERT INTO `logistics_tbl` VALUES (839, '773054906075645', '2020-09-10 18:03:48', '快件已到达【重庆转运中心】扫描员是【A班AC组组长陈俊琼】', '', '2020-09-10 19:07:24', '', '2020-09-10 19:07:24');
INSERT INTO `logistics_tbl` VALUES (840, '773054906075645', '2020-09-10 18:13:52', '快件由【重庆转运中心】发往【重庆机场公司】', '', '2020-09-10 19:07:24', '', '2020-09-10 19:07:24');
INSERT INTO `logistics_tbl` VALUES (841, '773054906427998', '2020-09-10 18:51:05', '快件已暂存至重庆渝北寰泰市场店菜鸟驿站，如有疑问请联系18883282919', '', '2020-09-10 19:07:25', '', '2020-09-10 19:07:25');
INSERT INTO `logistics_tbl` VALUES (842, '773054904868479', '2020-09-11 03:17:03', '快件已到达【广东深圳转运中心】扫描员是【1408】', '', '2020-09-11 07:41:43', '', '2020-09-11 07:41:43');
INSERT INTO `logistics_tbl` VALUES (843, '773054904868479', '2020-09-11 03:29:57', '快件由【广东深圳转运中心】发往【广东惠州转运中心】', '', '2020-09-11 07:41:43', '', '2020-09-11 07:41:43');
INSERT INTO `logistics_tbl` VALUES (844, 'YT4774201651604', '2020-09-11 11:28:44', '【上海转运中心公司】 已收入', '', '2020-09-11 12:17:26', '', '2020-09-11 12:17:26');
INSERT INTO `logistics_tbl` VALUES (845, 'YT4774201651604', '2020-09-11 11:56:03', '【上海转运中心】 已发出 下一站 【浦东转运中心公司】', '', '2020-09-11 12:17:26', '', '2020-09-11 12:17:26');
INSERT INTO `logistics_tbl` VALUES (846, 'YT4775134358570', '2020-09-10 21:20:38', '【南充转运中心公司】 已收入', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (847, 'YT4775134358570', '2020-09-10 21:41:57', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (848, 'YT4775134358570', '2020-09-11 03:41:03', '【重庆转运中心公司】 已收入', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (849, 'YT4775134358570', '2020-09-11 04:03:15', '【重庆转运中心】 已发出 下一站 【重庆市渝北区汽博中心公司】', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (850, 'YT4775134358570', '2020-09-11 09:44:52', '【重庆市渝北区汽博中心公司】 派件中  派件人: 张敏 电话 18559477228  如有疑问，请联系：023-62593882', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (851, 'YT4775134358570', '2020-09-11 10:27:21', '快件已由壹间公寓后门空地丰巢柜丰巢柜代收，取件码已发送，请及时取件。', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (852, 'YT4775134358570', '2020-09-11 12:11:00', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18559477228，投诉电话：023-62593882', '', '2020-09-11 13:11:00', '', '2020-09-11 13:11:00');
INSERT INTO `logistics_tbl` VALUES (853, 'YT4775135098761', '2020-09-10 21:12:58', '【南充转运中心公司】 已收入', '', '2020-09-11 13:11:18', '', '2020-09-11 13:11:18');
INSERT INTO `logistics_tbl` VALUES (854, 'YT4775135098761', '2020-09-10 21:42:04', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-11 13:11:18', '', '2020-09-11 13:11:18');
INSERT INTO `logistics_tbl` VALUES (855, 'YT4775135098761', '2020-09-11 02:11:02', '【重庆转运中心公司】 已收入', '', '2020-09-11 13:11:18', '', '2020-09-11 13:11:18');
INSERT INTO `logistics_tbl` VALUES (856, 'YT4775135098761', '2020-09-11 02:32:12', '【重庆转运中心】 已发出 下一站 【北京转运中心公司】', '', '2020-09-11 13:11:18', '', '2020-09-11 13:11:18');
INSERT INTO `logistics_tbl` VALUES (857, 'YT4775135096174', '2020-09-10 21:15:41', '【南充转运中心公司】 已收入', '', '2020-09-11 13:11:25', '', '2020-09-11 13:11:25');
INSERT INTO `logistics_tbl` VALUES (858, 'YT4775135096174', '2020-09-10 21:40:44', '【南充转运中心】 已发出 下一站 【自贡转运中心公司】', '', '2020-09-11 13:11:25', '', '2020-09-11 13:11:25');
INSERT INTO `logistics_tbl` VALUES (859, 'YT4775135096174', '2020-09-11 06:03:05', '【自贡转运中心公司】 已收入', '', '2020-09-11 13:11:25', '', '2020-09-11 13:11:25');
INSERT INTO `logistics_tbl` VALUES (860, 'YT4775135096174', '2020-09-11 06:25:22', '【自贡转运中心】 已发出 下一站 【四川省自贡市公司】', '', '2020-09-11 13:11:25', '', '2020-09-11 13:11:25');
INSERT INTO `logistics_tbl` VALUES (861, 'YT4775135096174', '2020-09-11 12:46:08', '【四川省自贡市公司】 已收入', '', '2020-09-11 13:11:25', '', '2020-09-11 13:11:25');
INSERT INTO `logistics_tbl` VALUES (862, 'YT4775135096174', '2020-09-11 12:54:48', '【四川省自贡市公司】 派件中  派件人: 王开立 电话 18882016637  如有疑问，请联系：0813-8102333', '', '2020-09-11 13:11:25', '', '2020-09-11 13:11:25');
INSERT INTO `logistics_tbl` VALUES (863, 'YT4775132837766', '2020-09-10 21:20:27', '【南充转运中心公司】 已收入', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (864, 'YT4775132837766', '2020-09-10 21:41:26', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (865, 'YT4775132837766', '2020-09-11 03:46:23', '【成都转运中心公司】 已收入', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (866, 'YT4775132837766', '2020-09-11 04:08:25', '【成都转运中心】 已发出 下一站 【四川省成都市双流县天府中和公司】', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (867, 'YT4775132837766', '2020-09-11 07:30:08', '【四川省成都市双流县天府中和公司】 已收入', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (868, 'YT4775132837766', '2020-09-11 08:15:17', '【四川省成都市双流县天府中和公司】 派件中  派件人: 柯全友 电话 18402854424  如有疑问，请联系：18521153451', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (869, 'YT4775132837766', '2020-09-11 09:43:01', '快件已由大有智慧广场左侧丰巢柜丰巢柜代收，取件码已发送，请及时取件。', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (870, 'YT4775132837766', '2020-09-11 12:47:27', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18402854424，投诉电话：18521153451', '', '2020-09-11 13:11:32', '', '2020-09-11 13:11:32');
INSERT INTO `logistics_tbl` VALUES (871, 'YT4775131545884', '2020-09-10 21:16:22', '【南充转运中心公司】 已收入', '', '2020-09-11 13:11:43', '', '2020-09-11 13:11:43');
INSERT INTO `logistics_tbl` VALUES (872, 'YT4775131545884', '2020-09-10 21:41:55', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-11 13:11:43', '', '2020-09-11 13:11:43');
INSERT INTO `logistics_tbl` VALUES (873, 'YT4775131545884', '2020-09-11 01:40:58', '【重庆转运中心公司】 已收入', '', '2020-09-11 13:11:43', '', '2020-09-11 13:11:43');
INSERT INTO `logistics_tbl` VALUES (874, 'YT4775131545884', '2020-09-11 02:04:09', '【重庆转运中心】 已发出 下一站 【南昌转运中心公司】', '', '2020-09-11 13:11:43', '', '2020-09-11 13:11:43');
INSERT INTO `logistics_tbl` VALUES (875, 'YT4775135853366', '2020-09-10 21:16:28', '【南充转运中心公司】 已收入', '', '2020-09-11 13:11:47', '', '2020-09-11 13:11:47');
INSERT INTO `logistics_tbl` VALUES (876, 'YT4775135853366', '2020-09-10 21:40:48', '【南充转运中心】 已发出 下一站 【自贡转运中心公司】', '', '2020-09-11 13:11:47', '', '2020-09-11 13:11:47');
INSERT INTO `logistics_tbl` VALUES (877, 'YT4775135853366', '2020-09-11 05:34:24', '【自贡转运中心公司】 已收入', '', '2020-09-11 13:11:47', '', '2020-09-11 13:11:47');
INSERT INTO `logistics_tbl` VALUES (878, 'YT4775135853366', '2020-09-11 05:58:00', '【自贡转运中心】 已发出 下一站 【四川省自贡市公司】', '', '2020-09-11 13:11:47', '', '2020-09-11 13:11:47');
INSERT INTO `logistics_tbl` VALUES (879, 'YT4775134367550', '2020-09-10 21:16:18', '【南充转运中心公司】 已收入', '', '2020-09-11 15:25:58', '', '2020-09-11 15:25:58');
INSERT INTO `logistics_tbl` VALUES (880, 'YT4775134367550', '2020-09-10 21:41:22', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-11 15:25:58', '', '2020-09-11 15:25:58');
INSERT INTO `logistics_tbl` VALUES (881, 'YT4775134367550', '2020-09-11 03:54:42', '【成都转运中心公司】 已收入', '', '2020-09-11 15:25:58', '', '2020-09-11 15:25:58');
INSERT INTO `logistics_tbl` VALUES (882, 'YT4775134367550', '2020-09-11 04:21:35', '【成都转运中心】 已发出 下一站 【四川省成都市武侯区二部公司】', '', '2020-09-11 15:25:58', '', '2020-09-11 15:25:58');
INSERT INTO `logistics_tbl` VALUES (883, 'YT4768219922650', '2020-09-11 15:23:27', '【四川省泸州市江阳区万象汇公司】 派件中  派件人: 杨先兵 电话 18181865511 ', '', '2020-09-11 16:32:25', '', '2020-09-11 16:32:25');
INSERT INTO `logistics_tbl` VALUES (884, 'YT4775135096174', '2020-09-11 16:31:20', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18882016637，投诉电话：0813-8102333', '', '2020-09-11 18:34:49', '', '2020-09-11 18:34:49');
INSERT INTO `logistics_tbl` VALUES (885, '773054904868479', '2020-09-11 08:19:34', '快件已到达【广东惠州转运中心】扫描员是【惠阳】', '', '2020-09-11 19:24:59', '', '2020-09-11 19:24:59');
INSERT INTO `logistics_tbl` VALUES (886, '773054904868479', '2020-09-11 08:34:22', '快件由【广东惠州转运中心】发往【广东惠阳公司】', '', '2020-09-11 19:24:59', '', '2020-09-11 19:24:59');
INSERT INTO `logistics_tbl` VALUES (887, '773054904868479', '2020-09-11 14:16:59', '快件已到达【广东惠阳公司】扫描员是【郝小波】', '', '2020-09-11 19:24:59', '', '2020-09-11 19:24:59');
INSERT INTO `logistics_tbl` VALUES (888, '773054904868479', '2020-09-11 18:10:27', '【广东惠阳公司】的派件员【赵斌(13891159702)】正在为您派件，如有疑问请联系派件员，联系电话【13891159702】', '', '2020-09-11 19:24:59', '', '2020-09-11 19:24:59');
INSERT INTO `logistics_tbl` VALUES (889, '773054906506161', '2020-09-11 08:05:44', '快件已到达【上海转运中心】扫描员是【金家群】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (890, '773054906506161', '2020-09-11 08:06:58', '快件已到达【上海青浦重固集散中心】扫描员是【扫描员3】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (891, '773054906506161', '2020-09-11 08:13:55', '快件由【上海转运中心】发往【上海青浦重固集散中心】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (892, '773054906506161', '2020-09-11 08:15:05', '快件由【上海青浦重固集散中心】发往【上海嘉定集散中心】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (893, '773054906506161', '2020-09-11 12:00:11', '快件已到达【上海嘉定集散中心】扫描员是【青浦组进港01】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (894, '773054906506161', '2020-09-11 12:06:36', '快件由【上海嘉定集散中心】发往【上海青浦徐泾国展公司】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (895, '773054906506161', '2020-09-11 13:16:53', '快件已到达【上海青浦徐泾国展公司】扫描员是【徐泾青浦国展公司】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (896, '773054906506161', '2020-09-11 15:18:03', '【上海青浦徐泾国展公司】的派件员【章鹏(13916534983)】正在为您派件，如有疑问请联系派件员，联系电话【13916534983】', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (897, '773054906506161', '2020-09-11 15:29:35', '快件已暂存至上海诸光路1355西郊家园店菜鸟驿站，如有疑问请联系13916564913', '', '2020-09-11 19:25:18', '', '2020-09-11 19:25:18');
INSERT INTO `logistics_tbl` VALUES (898, 'YT4779927240766', '2020-09-11 21:54:42', '添加快递单', 'admin', '2020-09-11 21:54:42', '', '2020-09-11 21:54:42');
INSERT INTO `logistics_tbl` VALUES (899, 'YT4779924412828 ', '2020-09-11 21:55:46', '添加快递单', 'admin', '2020-09-11 21:55:46', '', '2020-09-11 21:55:46');
INSERT INTO `logistics_tbl` VALUES (900, 'YT4779925047925', '2020-09-11 21:56:07', '添加快递单', 'admin', '2020-09-11 21:56:07', '', '2020-09-11 21:56:07');
INSERT INTO `logistics_tbl` VALUES (901, 'YT4775135853366', '2020-09-11 13:06:10', '【四川省自贡市公司】 已收入', '', '2020-09-11 23:05:35', '', '2020-09-11 23:05:35');
INSERT INTO `logistics_tbl` VALUES (902, 'YT4775135853366', '2020-09-11 13:18:41', '【四川省自贡市公司】 派件中  派件人: 徐亮 电话 18808221182  如有疑问，请联系：0813-8102333', '', '2020-09-11 23:05:35', '', '2020-09-11 23:05:35');
INSERT INTO `logistics_tbl` VALUES (903, 'YT4775135853366', '2020-09-11 15:41:53', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18808221182，投诉电话：0813-8102333', '', '2020-09-11 23:05:35', '', '2020-09-11 23:05:35');
INSERT INTO `logistics_tbl` VALUES (904, 'YT4775135848765', '2020-09-10 21:19:57', '【南充转运中心公司】 已收入', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (905, 'YT4775135848765', '2020-09-10 21:40:25', '【南充转运中心】 已发出 下一站 【自贡转运中心公司】', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (906, 'YT4775135848765', '2020-09-11 05:34:40', '【自贡转运中心公司】 已收入', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (907, 'YT4775135848765', '2020-09-11 05:58:42', '【自贡转运中心】 已发出 下一站 【四川省自贡市公司】', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (908, 'YT4775135848765', '2020-09-11 13:24:53', '【四川省自贡市公司】 已收入', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (909, 'YT4775135848765', '2020-09-11 13:32:19', '【四川省自贡市公司】 派件中  派件人: 肖伟 电话 15700650208  如有疑问，请联系：0813-8102333', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (910, 'YT4775135848765', '2020-09-11 19:36:12', '客户签收人: 门卫 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：15700650208，投诉电话：0813-8102333', '', '2020-09-11 23:05:59', '', '2020-09-11 23:05:59');
INSERT INTO `logistics_tbl` VALUES (911, 'YT4768219922650', '2020-09-11 21:30:27', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18181865511', '', '2020-09-12 00:23:26', '', '2020-09-12 00:23:26');
INSERT INTO `logistics_tbl` VALUES (912, '773054904868479', '2020-09-11 20:23:00', '已(13891159702)签收，签收人是【代收点代签收】,如有疑问请联系:13891159702,如您未收到此快递，请拨打投诉电话：0752-3399110! ', '', '2020-09-12 00:47:41', '', '2020-09-12 00:47:41');
INSERT INTO `logistics_tbl` VALUES (913, '773054906075645', '2020-09-11 07:50:15', '快件已到达【重庆机场公司】扫描员是【周达】', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (914, '773054906075645', '2020-09-11 08:29:46', '【重庆机场公司】的派件员【暨华片区(15123280019)】正在为您派件，如有疑问请联系派件员，联系电话【15123280019】', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (915, '773054906075645', '2020-09-11 11:42:47', '快件已暂存至重庆渝北寰泰市场店菜鸟驿站，如有疑问请联系18883282919', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (916, '773054906075645', '2020-09-11 20:06:54', '已签收，签收人凭取货码签收。', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (917, '773054906427998', '2020-09-11 20:06:37', '已签收，签收人凭取货码签收。', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (918, 'YT4773416999820', '2020-09-11 11:53:43', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (919, 'YT4773416999820', '2020-09-11 21:35:19', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (920, 'YT4773415789541', '2020-09-10 21:20:42', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (921, 'YT4773419880126', '2020-09-10 21:20:49', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (922, 'YT4773417016273', '2020-09-10 21:20:36', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:42', '', '2020-09-12 00:47:42');
INSERT INTO `logistics_tbl` VALUES (923, 'YT4773419154635', '2020-09-10 21:20:40', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (924, 'YT4773417656935', '2020-09-10 20:41:00', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (925, 'YT4773418439870', '2020-09-10 21:20:27', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (926, 'YT4773419910786', '2020-09-10 21:20:44', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (927, 'YT4773419902199', '2020-09-10 20:28:38', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (928, 'YT4773418461886', '2020-09-11 22:14:31', '【深圳转运中心公司】 已收入', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (929, 'YT4773418461886', '2020-09-11 22:44:12', '【深圳转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (930, 'YT4773418460044', '2020-09-10 21:20:46', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:43', '', '2020-09-12 00:47:43');
INSERT INTO `logistics_tbl` VALUES (931, 'YT4773415806150', '2020-09-10 21:20:38', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (932, 'YT4773415801755', '2020-09-10 21:20:34', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (933, 'YT4773416451220', '2020-09-10 21:20:25', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (934, 'YT4773419169365', '2020-09-11 09:35:37', '【重庆市巴南区公司】 已收入', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (935, 'YT4773419169365', '2020-09-11 11:56:00', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (936, 'YT4773419169365', '2020-09-11 21:35:26', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (937, 'YT4773419890573', '2020-09-10 20:41:22', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (938, 'YT4773419172800', '2020-09-10 21:20:29', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (939, 'YT4773417677734', '2020-09-11 09:13:24', '【重庆市巴南区公司】 已收入', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (940, 'YT4773417677734', '2020-09-11 11:53:57', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (941, 'YT4773417677734', '2020-09-11 21:35:34', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:44', '', '2020-09-12 00:47:44');
INSERT INTO `logistics_tbl` VALUES (942, 'YT4773417665418', '2020-09-10 21:20:32', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (943, 'YT4773419884144', '2020-09-11 09:24:49', '【重庆市巴南区公司】 已收入', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (944, 'YT4773419884144', '2020-09-11 11:55:07', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (945, 'YT4773419884144', '2020-09-11 21:35:45', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (946, 'YT4774201651604', '2020-09-11 16:09:53', '【浦东转运中心公司】 已收入', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (947, 'YT4774201651604', '2020-09-11 16:33:00', '【浦东转运中心】 已发出 下一站 【上海市浦东新区川沙唐镇分部公司】', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (948, 'YT4775134367550', '2020-09-11 16:24:53', '【四川省成都市武侯区二部公司】 已收入', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (949, 'YT4775134367550', '2020-09-11 16:41:29', '【四川省成都市武侯区二部公司】 派件中  派件人: 李波 电话 18581852005  如有疑问，请联系：028-82009168', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (950, 'YT4775134367550', '2020-09-11 18:10:52', '圆通合作点【兔喜快递超市】快件已到达科华街6号附14号驿站,联系电话13438345158', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (951, 'YT4775132833220', '2020-09-10 21:19:43', '【南充转运中心公司】 已收入', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (952, 'YT4775132833220', '2020-09-10 21:40:37', '【南充转运中心】 已发出 下一站 【自贡转运中心公司】', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (953, 'YT4775132833220', '2020-09-11 05:28:38', '【自贡转运中心公司】 已收入', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (954, 'YT4775132833220', '2020-09-11 05:54:19', '【自贡转运中心】 已发出 下一站 【四川省自贡市公司】', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (955, 'YT4775132833220', '2020-09-11 13:25:58', '【四川省自贡市公司】 已收入', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (956, 'YT4775132833220', '2020-09-11 13:29:03', '【四川省自贡市公司】 派件中  派件人: 郭玲 电话 13350102735  如有疑问，请联系：0813-8102333', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (957, 'YT4775132833220', '2020-09-11 18:10:41', '快件已到达创兴城小区18栋1楼附1号妈妈驿站,联系电话0813-8233256', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (958, 'YT4775132833220', '2020-09-11 18:10:42', '创兴城小区18栋1楼附1号妈妈驿站已发出自提短信,请上门自提,联系电话0813-8233256', '', '2020-09-12 00:47:45', '', '2020-09-12 00:47:45');
INSERT INTO `logistics_tbl` VALUES (959, 'YT4775131545884', '2020-09-12 01:25:07', '【南昌转运中心公司】 已收入', '', '2020-09-12 14:19:34', '', '2020-09-12 14:19:34');
INSERT INTO `logistics_tbl` VALUES (960, 'YT4775131545884', '2020-09-12 01:50:52', '【南昌转运中心】 已发出 下一站 【赣州转运中心公司】', '', '2020-09-12 14:19:34', '', '2020-09-12 14:19:34');
INSERT INTO `logistics_tbl` VALUES (961, 'YT4775131545884', '2020-09-12 09:20:25', '【赣州转运中心公司】 已收入', '', '2020-09-12 14:19:34', '', '2020-09-12 14:19:34');
INSERT INTO `logistics_tbl` VALUES (962, 'YT4775131545884', '2020-09-12 09:41:08', '【赣州转运中心】 已发出 下一站 【江西省赣州市信丰县公司】', '', '2020-09-12 14:19:34', '', '2020-09-12 14:19:34');
INSERT INTO `logistics_tbl` VALUES (963, 'YT4779927240766', '2020-09-12 17:50:01', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-12 19:18:17', '', '2020-09-12 19:18:17');
INSERT INTO `logistics_tbl` VALUES (964, 'YT4779927240766', '2020-09-12 18:45:12', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-12 19:18:17', '', '2020-09-12 19:18:17');
INSERT INTO `logistics_tbl` VALUES (965, 'YT4775131545884', '2020-09-12 14:23:33', '【江西省赣州市信丰县公司】 派件中  派件人: 李太阳 电话 18170131068  如有疑问，请联系：0797-3332060', '', '2020-09-12 19:32:15', '', '2020-09-12 19:32:15');
INSERT INTO `logistics_tbl` VALUES (966, 'YT4775135098761', '2020-09-12 19:03:55', '【北京转运中心公司】 已收入', '', '2020-09-12 19:32:37', '', '2020-09-12 19:32:37');
INSERT INTO `logistics_tbl` VALUES (967, 'YT4775135098761', '2020-09-12 19:28:17', '【北京转运中心】 已发出 下一站 【北京市怀柔区城区公司】', '', '2020-09-12 19:32:37', '', '2020-09-12 19:32:37');
INSERT INTO `logistics_tbl` VALUES (968, 'YT4775132833220', '2020-09-12 19:01:11', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：13350102735，投诉电话：0813-8102333', '', '2020-09-12 19:49:02', '', '2020-09-12 19:49:02');
INSERT INTO `logistics_tbl` VALUES (969, 'YT4770457089144', '2020-09-12 19:22:08', '客户签收人: 他人代收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521871490，投诉电话：028-65092999', '', '2020-09-12 19:49:12', '', '2020-09-12 19:49:12');
INSERT INTO `logistics_tbl` VALUES (970, 'YT4774201651604', '2020-09-12 08:48:15', '【上海市浦东新区川沙唐镇分部公司】 已收入', '', '2020-09-12 20:00:05', '', '2020-09-12 20:00:05');
INSERT INTO `logistics_tbl` VALUES (971, 'YT4774201651604', '2020-09-12 08:48:44', '【上海市浦东新区川沙唐镇分部公司】 派件中  派件人: 陈志军 电话 17721499456  如有疑问，请联系：021-33831367', '', '2020-09-12 20:00:05', '', '2020-09-12 20:00:05');
INSERT INTO `logistics_tbl` VALUES (972, 'YT4774201651604', '2020-09-12 11:50:36', '快件已到达齐爱路471号邮佳驾校驿站妈妈驿站,联系电话17721499456', '', '2020-09-12 20:00:05', '', '2020-09-12 20:00:05');
INSERT INTO `logistics_tbl` VALUES (973, 'YT4774201651604', '2020-09-12 11:50:36', '齐爱路471号邮佳驾校驿站妈妈驿站已发出自提短信,请上门自提,联系电话17721499456', '', '2020-09-12 20:00:05', '', '2020-09-12 20:00:05');
INSERT INTO `logistics_tbl` VALUES (974, 'YT4775131545884', '2020-09-12 19:39:15', '快件已暂存至赣州市信丰县双钻名汇5栋11号店菜鸟驿站，如有疑问请联系18170131068', '', '2020-09-12 20:00:05', '', '2020-09-12 20:00:05');
INSERT INTO `logistics_tbl` VALUES (975, 'YT4775131545884', '2020-09-12 19:39:21', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务', '', '2020-09-12 20:00:05', '', '2020-09-12 20:00:05');
INSERT INTO `logistics_tbl` VALUES (976, 'YT4779924412828', '2020-09-12 17:51:04', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-12 20:00:06', '', '2020-09-12 20:00:06');
INSERT INTO `logistics_tbl` VALUES (977, 'YT4779924412828', '2020-09-12 18:45:12', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-12 20:00:06', '', '2020-09-12 20:00:06');
INSERT INTO `logistics_tbl` VALUES (978, 'YT4779925047925', '2020-09-12 17:50:38', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-12 20:00:06', '', '2020-09-12 20:00:06');
INSERT INTO `logistics_tbl` VALUES (979, 'YT4779925047925', '2020-09-12 18:45:12', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-12 20:00:06', '', '2020-09-12 20:00:06');
INSERT INTO `logistics_tbl` VALUES (980, 'YT4779927240766', '2020-09-12 20:55:21', '【南充转运中心公司】 已收入', '', '2020-09-13 08:12:00', '', '2020-09-13 08:12:00');
INSERT INTO `logistics_tbl` VALUES (981, 'YT4779927240766', '2020-09-12 21:20:16', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-13 08:12:00', '', '2020-09-13 08:12:00');
INSERT INTO `logistics_tbl` VALUES (982, 'YT4779927240766', '2020-09-13 01:53:21', '【重庆转运中心公司】 已收入', '', '2020-09-13 08:12:00', '', '2020-09-13 08:12:00');
INSERT INTO `logistics_tbl` VALUES (983, 'YT4779927240766', '2020-09-13 02:15:20', '【重庆转运中心】 已发出 下一站 【重庆市渝北区加州公司】', '', '2020-09-13 08:12:00', '', '2020-09-13 08:12:00');
INSERT INTO `logistics_tbl` VALUES (984, 'YT4779927240766', '2020-09-13 07:42:03', '【重庆市渝北区加州公司】 已收入', '', '2020-09-13 08:12:00', '', '2020-09-13 08:12:00');
INSERT INTO `logistics_tbl` VALUES (985, 'YT4779927240766', '2020-09-13 07:44:24', '【重庆市渝北区加州公司】 派件中  派件人: 潘建华 电话 18521170812  如有疑问，请联系：023-63026620', '', '2020-09-13 08:12:00', '', '2020-09-13 08:12:00');
INSERT INTO `logistics_tbl` VALUES (986, 'YT4779927240766', '2020-09-13 09:21:14', '快件已暂存至重庆首创鸿恩三期6栋店菜鸟驿站，如有疑问请联系15802368609', '', '2020-09-13 10:57:55', '', '2020-09-13 10:57:55');
INSERT INTO `logistics_tbl` VALUES (987, 'YT4775135098761', '2020-09-13 06:41:12', '【北京市怀柔区城区公司】 已收入', '', '2020-09-13 18:26:55', '', '2020-09-13 18:26:55');
INSERT INTO `logistics_tbl` VALUES (988, 'YT4775135098761', '2020-09-13 06:53:53', '【北京市怀柔区城区公司】 派件中  派件人: 张丽云 电话 15701303908 。 圆通快递小哥每天已测体温，请放心收寄快递 如有疑问，请联系：010-60689907', '', '2020-09-13 18:26:55', '', '2020-09-13 18:26:55');
INSERT INTO `logistics_tbl` VALUES (989, 'YT4775135098761', '2020-09-13 09:03:28', '客户签收人: 门卫15702303908 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：15701303908，投诉电话：010-60689907。疫情期间圆通每天对网点多次消毒，快递小哥每天测量体温，佩戴口罩', '', '2020-09-13 18:26:55', '', '2020-09-13 18:26:55');
INSERT INTO `logistics_tbl` VALUES (990, 'YT4779927240766', '2020-09-13 19:26:46', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170812，投诉电话：023-63026620', '', '2020-09-13 19:34:50', '', '2020-09-13 19:34:50');
INSERT INTO `logistics_tbl` VALUES (991, 'YT4779925047925', '2020-09-12 20:49:42', '【南充转运中心公司】 已收入', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (992, 'YT4779925047925', '2020-09-12 21:14:55', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (993, 'YT4779925047925', '2020-09-13 02:30:39', '【成都转运中心公司】 已收入', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (994, 'YT4779925047925', '2020-09-13 02:53:41', '【成都转运中心】 已发出 下一站 【四川省成都市高新区三部公司】', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (995, 'YT4779925047925', '2020-09-13 07:06:43', '【四川省成都市高新区三部公司】 已收入', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (996, 'YT4779925047925', '2020-09-13 08:53:19', '【四川省成都市高新区三部公司】 派件中  派件人: 余林杰 电话 18521172504  如有疑问，请联系：18483625902', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (997, 'YT4779925047925', '2020-09-13 13:22:44', '客户签收人: 公司件周一派送 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521172504，投诉电话：18483625902', '', '2020-09-13 19:35:07', '', '2020-09-13 19:35:07');
INSERT INTO `logistics_tbl` VALUES (998, 'YT4779924412828', '2020-09-12 20:58:39', '【南充转运中心公司】 已收入', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (999, 'YT4779924412828', '2020-09-12 21:19:29', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1000, 'YT4779924412828', '2020-09-13 02:47:44', '【成都转运中心公司】 已收入', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1001, 'YT4779924412828', '2020-09-13 03:11:51', '【成都转运中心】 已发出 下一站 【四川省成都市高新区中和一部公司】', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1002, 'YT4779924412828', '2020-09-13 06:39:46', '【四川省成都市高新区中和一部公司】 已收入', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1003, 'YT4779924412828', '2020-09-13 07:58:57', '【四川省成都市高新区中和一部公司】 派件中  派件人: 朱永良 电话 13558826956  如有疑问，请联系：18521178676', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1004, 'YT4779924412828', '2020-09-13 10:57:57', '圆通合作点【喵站】快件已到达御庭上郡大门右侧157号科鲜森生鲜超市驿站,联系电话17311222077', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1005, 'YT4779924412828', '2020-09-13 14:38:09', '客户签收人: 无 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：13558826956，投诉电话：18521178676', '', '2020-09-13 19:35:15', '', '2020-09-13 19:35:15');
INSERT INTO `logistics_tbl` VALUES (1006, 'YT4773417656935', '2020-09-11 11:12:35', '重庆市巴南区 派件中 派件人: 杨成勇 电话 18521170626', '', '2020-09-13 19:35:31', '', '2020-09-13 19:35:31');
INSERT INTO `logistics_tbl` VALUES (1007, 'YT4773417656935', '2020-09-11 21:35:00', '重庆市巴南区 签收失败,原因:企事业单位放假', '', '2020-09-13 19:35:31', '', '2020-09-13 19:35:31');
INSERT INTO `logistics_tbl` VALUES (1008, 'YT4773418461886', '2020-09-13 03:39:04', '【重庆转运中心公司】 已收入', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1009, 'YT4773418461886', '2020-09-13 04:08:32', '【重庆转运中心】 已发出 下一站 【重庆市巴南区公司】', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1010, 'YT4773418461886', '2020-09-13 07:39:46', '【重庆市巴南区公司】 已收入', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1011, 'YT4773418461886', '2020-09-13 09:27:28', '【重庆市巴南区公司】 派件中  派件人: 杨成勇 电话 18521170626  如有疑问，请联系：023-62593622', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1012, 'YT4773419890573', '2020-09-11 11:12:36', '重庆市巴南区 派件中 派件人: 杨成勇 电话 18521170626', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1013, 'YT4773419890573', '2020-09-11 21:34:41', '重庆市巴南区 签收失败,原因:企事业单位放假', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1014, 'YT4773419902199', '2020-09-11 11:17:57', '重庆市巴南区 派件中 派件人: 杨成勇 电话 18521170626', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1015, 'YT4773419902199', '2020-09-11 21:35:10', '重庆市巴南区 签收失败,原因:企事业单位放假', '', '2020-09-13 19:35:32', '', '2020-09-13 19:35:32');
INSERT INTO `logistics_tbl` VALUES (1016, '773054904635511', '2020-09-13 17:33:00', '【四川苍溪公司】的收件员【梁婷(13219900508)】已收件', '', '2020-09-13 19:35:47', '', '2020-09-13 19:35:47');
INSERT INTO `logistics_tbl` VALUES (1017, '773054904635511', '2020-09-13 18:16:59', '快件由【四川苍溪公司】发往【四川南充转运中心】', '', '2020-09-13 19:35:47', '', '2020-09-13 19:35:47');
INSERT INTO `logistics_tbl` VALUES (1018, '773054906506161', '2020-09-13 11:39:23', '已签收，签收人凭取货码签收。', '', '2020-09-13 19:35:47', '', '2020-09-13 19:35:47');
INSERT INTO `logistics_tbl` VALUES (1019, 'YT4785137386427', '2020-09-14 11:17:12', '添加快递单', 'admin', '2020-09-14 11:17:12', '', '2020-09-14 11:17:12');
INSERT INTO `logistics_tbl` VALUES (1020, '773054904635511', '2020-09-13 22:11:03', '快件已到达【四川南充转运中心】扫描员是【许彩莲】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1021, '773054904635511', '2020-09-13 22:23:51', '快件由【四川南充转运中心】发往【四川仪陇公司】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1022, '773054904635511', '2020-09-14 15:36:03', '快件由【四川仪陇公司】发往【四川仪陇永光乡服务点】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1023, '773054904635511', '2020-09-14 15:44:26', '快件已到达【四川仪陇公司】扫描员是【高扫】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1024, '773054904635511', '2020-09-14 15:44:41', '快件已到达【四川仪陇永光乡服务点】扫描员是【仪陇永光乡服务点】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1025, '773054904635511', '2020-09-14 15:44:56', '【仪陇永光乡共配站S】小件员【蒋梅】正在为您派件，如有疑问请联系小件员，联系电话【13458278739】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1026, '773054904635511', '2020-09-14 15:45:26', '快件已被【仪陇永光乡共配站S】【文化街127号】站点代收，请及时取件。有问题请联系【13458278739】', '', '2020-09-14 16:10:25', '', '2020-09-14 16:10:25');
INSERT INTO `logistics_tbl` VALUES (1027, 'YT4773418461886', '2020-09-13 21:13:28', '【重庆市巴南区公司】 失败签收录入 杨成勇 （18521170626）', '', '2020-09-14 20:00:00', '', '2020-09-14 20:00:00');
INSERT INTO `logistics_tbl` VALUES (1028, 'YT4774201651604', '2020-09-14 16:13:20', '客户签收人: 邮件收发章 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521100638，投诉电话：021-33831367', '', '2020-09-14 20:00:01', '', '2020-09-14 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1029, 'YT4775134367550', '2020-09-13 20:29:36', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18581852005，投诉电话：028-82009168', '', '2020-09-14 20:00:01', '', '2020-09-14 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1030, 'YT4785137386427', '2020-09-14 15:42:43', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-14 20:00:01', '', '2020-09-14 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1031, 'YT4785137386427', '2020-09-14 17:37:13', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-14 20:00:01', '', '2020-09-14 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1032, 'YT4773418461886', '2020-09-14 21:42:07', '客户签收人: 刘玉 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521170626，投诉电话：023-62593622', '', '2020-09-14 22:04:15', '', '2020-09-14 22:04:15');
INSERT INTO `logistics_tbl` VALUES (1033, 'YT4785137386427', '2020-09-14 20:54:06', '【南充转运中心公司】 已收入', '', '2020-09-14 22:04:54', '', '2020-09-14 22:04:54');
INSERT INTO `logistics_tbl` VALUES (1034, 'YT4785137386427', '2020-09-14 21:18:37', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-14 22:04:54', '', '2020-09-14 22:04:54');
INSERT INTO `logistics_tbl` VALUES (1035, 'YT4785137386427', '2020-09-15 03:35:22', '【成都转运中心公司】 已收入', '', '2020-09-15 11:51:06', '', '2020-09-15 11:51:06');
INSERT INTO `logistics_tbl` VALUES (1036, 'YT4785137386427', '2020-09-15 04:04:50', '【成都转运中心】 已发出 下一站 【四川省成都市高新西区公司】', '', '2020-09-15 11:51:06', '', '2020-09-15 11:51:06');
INSERT INTO `logistics_tbl` VALUES (1037, 'YT4785137386427', '2020-09-15 07:38:42', '【四川省成都市高新西区公司】 已收入', '', '2020-09-15 11:51:06', '', '2020-09-15 11:51:06');
INSERT INTO `logistics_tbl` VALUES (1038, 'YT4785137386427', '2020-09-15 07:52:46', '【四川省成都市高新西区公司】 派件中  派件人: 杨雪梅 电话 13558768119  如有疑问，请联系：028-87876326', '', '2020-09-15 11:51:06', '', '2020-09-15 11:51:06');
INSERT INTO `logistics_tbl` VALUES (1039, 'YT4785137386427', '2020-09-15 10:19:11', '快件已由四川大学锦城学院一车棚菜鸟驿站代收，请及时取件，如有疑问请联系13558768119', '', '2020-09-15 11:51:06', '', '2020-09-15 11:51:06');
INSERT INTO `logistics_tbl` VALUES (1040, 'YT4788197734601', '2020-09-15 16:37:28', '添加快递单', 'admin', '2020-09-15 16:37:28', '', '2020-09-15 16:37:28');
INSERT INTO `logistics_tbl` VALUES (1041, 'YT4788195739310', '2020-09-15 16:37:42', '添加快递单', 'admin', '2020-09-15 16:37:42', '', '2020-09-15 16:37:42');
INSERT INTO `logistics_tbl` VALUES (1042, 'YT4788197145362', '2020-09-15 16:41:02', '添加快递单', 'admin', '2020-09-15 16:41:02', '', '2020-09-15 16:41:02');
INSERT INTO `logistics_tbl` VALUES (1043, 'YT4788893009381', '2020-09-15 16:41:22', '添加快递单', 'admin', '2020-09-15 16:41:22', '', '2020-09-15 16:41:22');
INSERT INTO `logistics_tbl` VALUES (1044, 'YT4788893009381', '2020-09-15 17:42:20', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-15 18:06:47', '', '2020-09-15 18:06:47');
INSERT INTO `logistics_tbl` VALUES (1045, 'YT4788893009381', '2020-09-15 18:02:02', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-15 18:06:47', '', '2020-09-15 18:06:47');
INSERT INTO `logistics_tbl` VALUES (1046, 'YT4788195739310', '2020-09-15 17:56:23', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-15 18:06:56', '', '2020-09-15 18:06:56');
INSERT INTO `logistics_tbl` VALUES (1047, 'YT4788195739310', '2020-09-15 18:02:02', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-15 18:06:56', '', '2020-09-15 18:06:56');
INSERT INTO `logistics_tbl` VALUES (1048, 'YT4788197734601', '2020-09-15 17:56:25', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-15 18:06:56', '', '2020-09-15 18:06:56');
INSERT INTO `logistics_tbl` VALUES (1049, 'YT4788197734601', '2020-09-15 18:02:02', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-15 18:06:56', '', '2020-09-15 18:06:56');
INSERT INTO `logistics_tbl` VALUES (1050, 'YT4788197145362', '2020-09-15 17:56:42', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-15 18:07:02', '', '2020-09-15 18:07:02');
INSERT INTO `logistics_tbl` VALUES (1051, 'YT4788197145362', '2020-09-15 18:02:02', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-15 18:07:02', '', '2020-09-15 18:07:02');
INSERT INTO `logistics_tbl` VALUES (1052, 'YT4788893009381', '2020-09-15 20:18:07', '【南充转运中心公司】 已收入', '', '2020-09-15 23:21:01', '', '2020-09-15 23:21:01');
INSERT INTO `logistics_tbl` VALUES (1053, 'YT4788893009381', '2020-09-15 20:43:36', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-09-15 23:21:01', '', '2020-09-15 23:21:01');
INSERT INTO `logistics_tbl` VALUES (1054, 'YT4788195739310', '2020-09-15 23:11:26', '【南充转运中心公司】 已收入', '', '2020-09-16 13:49:59', '', '2020-09-16 13:49:59');
INSERT INTO `logistics_tbl` VALUES (1055, 'YT4788195739310', '2020-09-15 23:36:42', '【南充转运中心】 已发出 下一站 【上海转运中心公司】', '', '2020-09-16 13:49:59', '', '2020-09-16 13:49:59');
INSERT INTO `logistics_tbl` VALUES (1056, 'YT4788197734601', '2020-09-15 23:08:11', '【南充转运中心公司】 已收入', '', '2020-09-16 13:49:59', '', '2020-09-16 13:49:59');
INSERT INTO `logistics_tbl` VALUES (1057, 'YT4788197734601', '2020-09-15 23:36:47', '【南充转运中心】 已发出 下一站 【上海转运中心公司】', '', '2020-09-16 13:49:59', '', '2020-09-16 13:49:59');
INSERT INTO `logistics_tbl` VALUES (1058, 'YT4788197145362', '2020-09-15 23:26:13', '【南充转运中心公司】 已收入', '', '2020-09-16 13:50:04', '', '2020-09-16 13:50:04');
INSERT INTO `logistics_tbl` VALUES (1059, 'YT4788197145362', '2020-09-15 23:48:40', '【南充转运中心】 已发出 下一站 【上海转运中心公司】', '', '2020-09-16 13:50:04', '', '2020-09-16 13:50:04');
INSERT INTO `logistics_tbl` VALUES (1060, 'YT4788893009381', '2020-09-16 01:02:25', '【成都转运中心公司】 已收入', '', '2020-09-16 17:39:10', '', '2020-09-16 17:39:10');
INSERT INTO `logistics_tbl` VALUES (1061, 'YT4788893009381', '2020-09-16 01:30:00', '【成都转运中心】 已发出 下一站 【四川省成都市青羊区一部公司】', '', '2020-09-16 17:39:10', '', '2020-09-16 17:39:10');
INSERT INTO `logistics_tbl` VALUES (1062, 'YT4788893009381', '2020-09-16 07:16:28', '【四川省成都市青羊区一部公司】 已收入', '', '2020-09-16 17:39:10', '', '2020-09-16 17:39:10');
INSERT INTO `logistics_tbl` VALUES (1063, 'YT4788893009381', '2020-09-16 08:04:43', '【四川省成都市青羊区一部公司】 派件中  派件人: 严超成 电话 18280307624  如有疑问，请联系：028-87713449', '', '2020-09-16 17:39:10', '', '2020-09-16 17:39:10');
INSERT INTO `logistics_tbl` VALUES (1064, 'YT4788893009381', '2020-09-16 11:38:57', '快件已暂存至成都橙花风景店菜鸟驿站，如有疑问请联系18328362508', '', '2020-09-16 17:39:10', '', '2020-09-16 17:39:10');
INSERT INTO `logistics_tbl` VALUES (1065, 'YT4788893009381', '2020-09-16 12:44:00', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18280307624，投诉电话：028-87713449', '', '2020-09-16 17:39:10', '', '2020-09-16 17:39:10');
INSERT INTO `logistics_tbl` VALUES (1066, 'YT4791786040000', '2020-09-16 19:15:55', '添加快递单', 'admin', '2020-09-16 19:15:55', '', '2020-09-16 19:15:55');
INSERT INTO `logistics_tbl` VALUES (1067, 'YT4791007099990', '2020-09-16 19:16:23', '添加快递单', 'admin', '2020-09-16 19:16:23', '', '2020-09-16 19:16:23');
INSERT INTO `logistics_tbl` VALUES (1068, 'YT4791894898331', '2020-09-17 07:41:22', '添加快递单', 'admin', '2020-09-17 07:41:22', '', '2020-09-17 07:41:22');
INSERT INTO `logistics_tbl` VALUES (1069, 'YT4791895530363', '2020-09-17 07:53:24', '添加快递单', 'admin', '2020-09-17 07:53:24', '', '2020-09-17 07:53:24');
INSERT INTO `logistics_tbl` VALUES (1070, 'YT4791879822002', '2020-09-17 07:54:23', '添加快递单', 'admin', '2020-09-17 07:54:23', '', '2020-09-17 07:54:23');
INSERT INTO `logistics_tbl` VALUES (1071, 'YT4791898277719', '2020-09-17 07:57:02', '添加快递单', 'admin', '2020-09-17 07:57:02', '', '2020-09-17 07:57:02');
INSERT INTO `logistics_tbl` VALUES (1072, 'YT4791911088503', '2020-09-17 08:02:40', '添加快递单', 'admin', '2020-09-17 08:02:40', '', '2020-09-17 08:02:40');
INSERT INTO `logistics_tbl` VALUES (1073, 'YT4791894898331', '2020-09-17 11:07:17', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-17 12:05:47', '', '2020-09-17 12:05:47');
INSERT INTO `logistics_tbl` VALUES (1074, 'YT4791894898331', '2020-09-17 11:07:27', '【四川省直营市场部】 已发出 下一站 【成都转运中心公司】', '', '2020-09-17 12:05:47', '', '2020-09-17 12:05:47');
INSERT INTO `logistics_tbl` VALUES (1075, 'YT4791911088503', '2020-09-17 11:07:15', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-17 12:06:14', '', '2020-09-17 12:06:14');
INSERT INTO `logistics_tbl` VALUES (1076, 'YT4791911088503', '2020-09-17 11:07:25', '【四川省直营市场部】 已发出 下一站 【成都转运中心公司】', '', '2020-09-17 12:06:14', '', '2020-09-17 12:06:14');
INSERT INTO `logistics_tbl` VALUES (1077, 'YT4788195739310', '2020-09-17 10:10:08', '【上海转运中心公司】 已收入', '', '2020-09-17 12:06:34', '', '2020-09-17 12:06:34');
INSERT INTO `logistics_tbl` VALUES (1078, 'YT4788195739310', '2020-09-17 10:31:20', '【上海转运中心】 已发出 下一站 【浦东转运中心公司】', '', '2020-09-17 12:06:34', '', '2020-09-17 12:06:34');
INSERT INTO `logistics_tbl` VALUES (1079, 'YT4788197734601', '2020-09-17 09:58:32', '【上海转运中心公司】 已收入', '', '2020-09-17 12:06:35', '', '2020-09-17 12:06:35');
INSERT INTO `logistics_tbl` VALUES (1080, 'YT4788197734601', '2020-09-17 10:28:18', '【上海转运中心】 已发出 下一站 【浦东转运中心公司】', '', '2020-09-17 12:06:35', '', '2020-09-17 12:06:35');
INSERT INTO `logistics_tbl` VALUES (1081, 'YT4788197145362', '2020-09-17 09:48:24', '【上海转运中心公司】 已收入', '', '2020-09-17 12:06:58', '', '2020-09-17 12:06:58');
INSERT INTO `logistics_tbl` VALUES (1082, 'YT4788197145362', '2020-09-17 10:14:37', '【上海转运中心】 已发出 下一站 【上海宝山城配中心公司】', '', '2020-09-17 12:06:58', '', '2020-09-17 12:06:58');
INSERT INTO `logistics_tbl` VALUES (1083, 'YT4793111423033', '2020-09-17 14:09:02', '添加快递单', 'admin', '2020-09-17 14:09:02', '', '2020-09-17 14:09:02');
INSERT INTO `logistics_tbl` VALUES (1084, 'YT4793107685946', '2020-09-17 14:09:19', '添加快递单', 'admin', '2020-09-17 14:09:19', '', '2020-09-17 14:09:19');
INSERT INTO `logistics_tbl` VALUES (1085, 'YT4788195739310', '2020-09-17 13:10:14', '【浦东转运中心公司】 已收入', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1086, 'YT4788195739310', '2020-09-17 13:30:53', '【浦东转运中心】 已发出 下一站 【上海市浦东机场公司】', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1087, 'YT4788195739310', '2020-09-17 14:53:34', '【上海市浦东机场公司】 已收入', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1088, 'YT4788195739310', '2020-09-17 14:58:03', '【上海市浦东机场公司】 派件中  派件人: 饶园丁 电话 17317214698  如有疑问，请联系：023-49589172', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1089, 'YT4788197734601', '2020-09-17 13:14:01', '【浦东转运中心公司】 已收入', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1090, 'YT4788197734601', '2020-09-17 13:40:41', '【浦东转运中心】 已发出 下一站 【上海市浦东机场公司】', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1091, 'YT4788197734601', '2020-09-17 15:09:18', '【上海市浦东机场公司】 已收入', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1092, 'YT4788197734601', '2020-09-17 15:11:16', '【上海市浦东机场公司】 派件中  派件人: 饶园丁 电话 17317214698  如有疑问，请联系：023-49589172', '', '2020-09-17 16:41:53', '', '2020-09-17 16:41:53');
INSERT INTO `logistics_tbl` VALUES (1093, 'YT4788197145362', '2020-09-17 11:55:40', '【上海宝山城配中心公司】 已收入', '', '2020-09-17 16:42:03', '', '2020-09-17 16:42:03');
INSERT INTO `logistics_tbl` VALUES (1094, 'YT4788197145362', '2020-09-17 12:17:45', '【上海宝山城配中心】 已发出 下一站 【上海市杨浦区复旦公司】', '', '2020-09-17 16:42:03', '', '2020-09-17 16:42:03');
INSERT INTO `logistics_tbl` VALUES (1095, 'YT4788197145362', '2020-09-17 15:02:43', '【上海市杨浦区复旦公司】 已收入', '', '2020-09-17 16:42:03', '', '2020-09-17 16:42:03');
INSERT INTO `logistics_tbl` VALUES (1096, 'YT4788197145362', '2020-09-17 15:14:30', '【上海市杨浦区复旦公司】 派件中  派件人: 赵向勇 电话 18521100053  如有疑问，请联系：021-65682009', '', '2020-09-17 16:42:03', '', '2020-09-17 16:42:03');
INSERT INTO `logistics_tbl` VALUES (1097, 'YT4793111423033', '2020-09-17 14:51:50', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-17 17:15:34', '', '2020-09-17 17:15:34');
INSERT INTO `logistics_tbl` VALUES (1098, 'YT4788197734601', '2020-09-17 18:59:37', '圆通合作点【兔喜快递超市】快件已到达508弄73号104室驿站,联系电话17802118959', '', '2020-09-17 20:00:01', '', '2020-09-17 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1099, 'YT4788195739310', '2020-09-17 18:59:37', '圆通合作点【兔喜快递超市】快件已到达508弄73号104室驿站,联系电话17802118959', '', '2020-09-17 20:00:02', '', '2020-09-17 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1100, 'YT4788197145362', '2020-09-17 18:01:32', '客户签收人: 2号楼过道 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521100053，投诉电话：021-65682009', '', '2020-09-17 20:00:03', '', '2020-09-17 20:00:03');
INSERT INTO `logistics_tbl` VALUES (1101, 'YT4791786040000', '2020-09-17 14:54:22', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-17 20:00:03', '', '2020-09-17 20:00:03');
INSERT INTO `logistics_tbl` VALUES (1102, 'YT4791786040000', '2020-09-17 17:57:20', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-17 20:00:03', '', '2020-09-17 20:00:03');
INSERT INTO `logistics_tbl` VALUES (1103, 'YT4791894898331', '2020-09-17 19:01:50', '【成都转运中心公司】 已收入', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1104, 'YT4791894898331', '2020-09-17 19:28:01', '【成都转运中心】 已发出 下一站 【福州转运中心公司】', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1105, 'YT4791895530363', '2020-09-17 11:07:17', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1106, 'YT4791895530363', '2020-09-17 11:07:27', '【四川省直营市场部】 已发出 下一站 【成都转运中心公司】', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1107, 'YT4791879822002', '2020-09-17 11:07:17', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1108, 'YT4791879822002', '2020-09-17 11:07:27', '【四川省直营市场部】 已发出 下一站 【成都转运中心公司】', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1109, 'YT4791879822002', '2020-09-17 18:52:04', '【成都转运中心公司】 已收入', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1110, 'YT4791879822002', '2020-09-17 19:13:43', '【成都转运中心】 已发出 下一站 【肃宁转运中心公司】', '', '2020-09-17 20:00:04', '', '2020-09-17 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1111, 'YT4791898277719', '2020-09-17 11:07:17', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1112, 'YT4791898277719', '2020-09-17 11:07:27', '【四川省直营市场部】 已发出 下一站 【成都转运中心公司】', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1113, 'YT4791898277719', '2020-09-17 19:12:08', '【成都转运中心公司】 已收入', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1114, 'YT4791898277719', '2020-09-17 19:32:11', '【成都转运中心】 已发出 下一站 【肃宁转运中心公司】', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1115, 'YT4793111423033', '2020-09-17 17:57:20', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1116, 'YT4793107685946', '2020-09-17 14:51:41', '【四川省广元市苍溪县公司】 已收件 取件人: 杨俐 (13183559799)', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1117, 'YT4793107685946', '2020-09-17 17:57:20', '【四川省广元市苍溪县】 已发出 下一站 【南充转运中心公司】', '', '2020-09-17 20:00:05', '', '2020-09-17 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1118, 'YT4793111423033', '2020-09-17 21:15:21', '【南充转运中心公司】 已收入', '', '2020-09-18 13:43:40', '', '2020-09-18 13:43:40');
INSERT INTO `logistics_tbl` VALUES (1119, 'YT4793111423033', '2020-09-17 21:39:28', '【南充转运中心】 已发出 下一站 【深圳转运中心公司】', '', '2020-09-18 13:43:40', '', '2020-09-18 13:43:40');
INSERT INTO `logistics_tbl` VALUES (1120, 'YT4791786040000', '2020-09-17 21:29:02', '【南充转运中心公司】 已收入', '', '2020-09-18 20:00:02', '', '2020-09-18 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1121, 'YT4791786040000', '2020-09-17 21:51:40', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-18 20:00:02', '', '2020-09-18 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1122, 'YT4791786040000', '2020-09-18 02:57:18', '【重庆转运中心公司】 已收入', '', '2020-09-18 20:00:02', '', '2020-09-18 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1123, 'YT4791786040000', '2020-09-18 03:21:07', '【重庆转运中心】 已发出 下一站 【三明转运中心公司】', '', '2020-09-18 20:00:02', '', '2020-09-18 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1124, 'YT4793107685946', '2020-09-17 22:39:01', '【南充转运中心公司】 已收入', '', '2020-09-18 20:00:05', '', '2020-09-18 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1125, 'YT4793107685946', '2020-09-17 23:01:19', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-18 20:00:05', '', '2020-09-18 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1126, 'YT4793107685946', '2020-09-18 03:00:25', '【重庆转运中心公司】 已收入', '', '2020-09-18 20:00:05', '', '2020-09-18 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1127, 'YT4793107685946', '2020-09-18 03:27:46', '【重庆转运中心】 已发出 下一站 【合肥转运中心公司】', '', '2020-09-18 20:00:05', '', '2020-09-18 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1128, ' YT4795520163866', '2020-09-18 21:36:41', '添加快递单', 'admin', '2020-09-18 21:36:41', '', '2020-09-18 21:36:41');
INSERT INTO `logistics_tbl` VALUES (1129, 'YT4795518645450', '2020-09-18 21:37:35', '添加快递单', 'admin', '2020-09-18 21:37:35', '', '2020-09-18 21:37:35');
INSERT INTO `logistics_tbl` VALUES (1130, 'YT4788197734601', '2020-09-18 22:22:09', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：17317214698，投诉电话：023-49589172', '', '2020-09-19 20:00:00', '', '2020-09-19 20:00:00');
INSERT INTO `logistics_tbl` VALUES (1131, 'YT4788195739310', '2020-09-18 22:22:09', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：17317214698，投诉电话：023-49589172', '', '2020-09-19 20:00:01', '', '2020-09-19 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1132, 'YT4791786040000', '2020-09-19 04:41:18', '【三明转运中心公司】 已收入', '', '2020-09-19 20:00:02', '', '2020-09-19 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1133, 'YT4791786040000', '2020-09-19 05:02:49', '【三明转运中心】 已发出 下一站 【福州转运中心公司】', '', '2020-09-19 20:00:02', '', '2020-09-19 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1134, 'YT4791786040000', '2020-09-19 11:32:53', '【福州转运中心公司】 已收入', '', '2020-09-19 20:00:02', '', '2020-09-19 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1135, 'YT4791786040000', '2020-09-19 11:55:03', '【福州转运中心】 已发出 下一站 【福建省福州市晋安区西园公司】', '', '2020-09-19 20:00:02', '', '2020-09-19 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1136, 'YT4791786040000', '2020-09-19 16:13:49', '【福建省福州市晋安区西园分部公司】 派件中  派件人: 廖启佳 电话 18350034440  如有疑问，请联系：18059129309', '', '2020-09-19 20:00:02', '', '2020-09-19 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1137, 'YT4791786040000', '2020-09-19 17:43:26', '圆通合作点【小栈】快件已到达新店镇磐石新城二区2号楼109驿站,联系电话18350034440', '', '2020-09-19 20:00:02', '', '2020-09-19 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1138, 'YT4791894898331', '2020-09-19 13:41:09', '【福州转运中心公司】 已收入', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1139, 'YT4791894898331', '2020-09-19 14:01:22', '【福州转运中心】 已发出 下一站 【福建省福州市晋安区西园公司】', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1140, 'YT4791894898331', '2020-09-19 16:05:14', '【福建省福州市晋安区西园分部公司】 派件中  派件人: 廖启佳 电话 18350034440  如有疑问，请联系：18059129309', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1141, 'YT4791894898331', '2020-09-19 17:43:26', '圆通合作点【小栈】快件已到达新店镇磐石新城二区2号楼109驿站,联系电话18350034440', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1142, 'YT4791895530363', '2020-09-19 01:53:41', '【成都转运中心公司】 已收入', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1143, 'YT4791895530363', '2020-09-19 02:19:10', '【成都转运中心】 已发出 下一站 【四川省成都市双流区华阳二部公司】', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1144, 'YT4791895530363', '2020-09-19 08:11:41', '【四川省成都市双流区华阳二部公司】 已收入', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1145, 'YT4791895530363', '2020-09-19 08:23:29', '【四川省成都市双流区华阳二部公司】 派件中  派件人: 邓林 电话 13693433523  如有疑问，请联系：18521172098', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1146, 'YT4791895530363', '2020-09-19 12:14:36', '快件已暂存至成都宏达世纪锦城店菜鸟驿站，如有疑问请联系13330947776', '', '2020-09-19 20:00:05', '', '2020-09-19 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1147, 'YT4791879822002', '2020-09-19 08:46:48', '【肃宁转运中心公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1148, 'YT4791879822002', '2020-09-19 09:10:23', '【肃宁转运中心】 已发出 下一站 【河北省沧州市任丘市公司】', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1149, 'YT4791879822002', '2020-09-19 12:10:52', '【河北省沧州市任丘市公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1150, 'YT4791879822002', '2020-09-19 14:35:27', '【河北省沧州市任丘市公司】 派件中  派件人: 王忆 电话 18521188770  如有疑问，请联系：0317-2215988', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1151, 'YT4791879822002', '2020-09-19 15:17:37', '客户签收人: 大堂保安 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521188770，投诉电话：0317-2215988', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1152, 'YT4791898277719', '2020-09-19 08:35:25', '【肃宁转运中心公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1153, 'YT4791898277719', '2020-09-19 08:59:35', '【肃宁转运中心】 已发出 下一站 【河北省沧州市任丘市公司】', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1154, 'YT4791898277719', '2020-09-19 12:12:39', '【河北省沧州市任丘市公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1155, 'YT4791898277719', '2020-09-19 12:24:39', '【河北省沧州市任丘市公司】 派件中  派件人: 赵健 电话 15931704117  如有疑问，请联系：0317-2215988', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1156, 'YT4791898277719', '2020-09-19 16:58:01', '快件已到达长安花园大门北侧豪豪平价超市妈妈驿站,联系电话13292700597', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1157, 'YT4791898277719', '2020-09-19 16:58:02', '长安花园大门北侧豪豪平价超市妈妈驿站已发出自提短信,请上门自提,联系电话13292700597', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1158, 'YT4791911088503', '2020-09-19 01:39:06', '【成都转运中心公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1159, 'YT4791911088503', '2020-09-19 02:00:51', '【成都转运中心】 已发出 下一站 【南充转运中心公司】', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1160, 'YT4791911088503', '2020-09-19 08:24:04', '【南充转运中心公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1161, 'YT4791911088503', '2020-09-19 08:49:42', '【南充转运中心】 已发出 下一站 【四川省广元市苍溪县公司】', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1162, 'YT4793107685946', '2020-09-19 02:54:10', '【合肥转运中心公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1163, 'YT4793107685946', '2020-09-19 03:17:36', '【合肥转运中心】 已发出 下一站 【安徽省马鞍山市公司】', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1164, 'YT4793107685946', '2020-09-19 09:49:25', '【安徽省马鞍山市公司】 已收入', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1165, 'YT4793107685946', '2020-09-19 14:48:32', '【安徽省马鞍山市公司】 派件中  派件人: 李晓林 电话 18895589919  如有疑问，请联系：0555-2458977', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1166, 'YT4793107685946', '2020-09-19 15:17:24', '圆通合作点【小兵驿站】快件已到达钢城浴室旁马鞍山快递裹裹生活驿站驿站,联系电话18895589915', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1167, 'YT4793107685946', '2020-09-19 15:17:27', '钢城浴室旁马鞍山快递裹裹生活驿站小兵驿站已发出自提通知，请上门自提，联系电话18895589915', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1168, 'YT4793107685946', '2020-09-19 16:41:31', '客户签收人: 谢晓宇 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18895589919，投诉电话：0555-2458977', '', '2020-09-19 20:00:06', '', '2020-09-19 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1169, 'JDVD00976222190', '2020-09-20 09:03:44', '添加快递单', 'admin', '2020-09-20 09:03:44', '', '2020-09-20 09:03:44');
INSERT INTO `logistics_tbl` VALUES (1170, 'JDVD00976222190', '2020-09-20 08:54:33', '揽收任务已分配给寇方正', '', '2020-09-20 09:03:51', '', '2020-09-20 09:03:51');
INSERT INTO `logistics_tbl` VALUES (1171, 'JDVD00976193380', '2020-09-20 09:05:44', '添加快递单', 'admin', '2020-09-20 09:05:44', '', '2020-09-20 09:05:44');
INSERT INTO `logistics_tbl` VALUES (1172, 'JDVD00976193399', '2020-09-20 09:06:07', '添加快递单', 'admin', '2020-09-20 09:06:07', '', '2020-09-20 09:06:07');
INSERT INTO `logistics_tbl` VALUES (1173, 'JDVD00976185082', '2020-09-20 09:06:42', '添加快递单', 'admin', '2020-09-20 09:06:42', '', '2020-09-20 09:06:42');
INSERT INTO `logistics_tbl` VALUES (1174, 'YT4791911088503', '2020-09-20 08:44:35', '【四川省广元市苍溪县公司】 已收入', '', '2020-09-20 09:06:52', '', '2020-09-20 09:06:52');
INSERT INTO `logistics_tbl` VALUES (1175, 'JDVD00976242970', '2020-09-20 09:07:58', '添加快递单', 'admin', '2020-09-20 09:07:58', '', '2020-09-20 09:07:58');
INSERT INTO `logistics_tbl` VALUES (1176, 'JDVD00976193380', '2020-09-20 08:54:33', '揽收任务已分配给寇方正', '', '2020-09-20 09:09:09', '', '2020-09-20 09:09:09');
INSERT INTO `logistics_tbl` VALUES (1177, 'JDVD00976193399', '2020-09-20 08:54:33', '揽收任务已分配给寇方正', '', '2020-09-20 09:09:16', '', '2020-09-20 09:09:16');
INSERT INTO `logistics_tbl` VALUES (1178, 'JDVD00976242970', '2020-09-20 08:54:33', '揽收任务已分配给寇方正', '', '2020-09-20 09:10:05', '', '2020-09-20 09:10:05');
INSERT INTO `logistics_tbl` VALUES (1179, 'YT4791894898331', '2020-09-19 20:56:44', '客户签收人: 磐石新城二区韵达快递超市 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18350034440，投诉电话：18059129309', '', '2020-09-20 09:10:28', '', '2020-09-20 09:10:28');
INSERT INTO `logistics_tbl` VALUES (1180, 'JDVD00976185082', '2020-09-20 08:54:33', '揽收任务已分配给寇方正', '', '2020-09-20 09:37:20', '', '2020-09-20 09:37:20');
INSERT INTO `logistics_tbl` VALUES (1181, 'JDVD00973573239', '2020-09-20 09:43:46', '添加快递单', 'admin', '2020-09-20 09:43:46', '', '2020-09-20 09:43:46');
INSERT INTO `logistics_tbl` VALUES (1182, 'JDVD00973573239', '2020-09-19 17:51:30', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1183, 'JDVD00973573239', '2020-09-19 17:51:30', '京东快递 已收取快件', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1184, 'JDVD00973573239', '2020-09-19 17:51:31', '您的快件已到达【广元苍溪营业部】', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1185, 'JDVD00973573239', '2020-09-19 19:00:40', '快递司机收箱', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1186, 'JDVD00973573239', '2020-09-19 19:00:45', '您的快件已发车', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1187, 'JDVD00973573239', '2020-09-20 01:04:00', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1188, 'JDVD00973573239', '2020-09-20 01:07:48', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1189, 'JDVD00973573239', '2020-09-20 01:07:53', '您的快件由【成都祥福分拣中心】准备发往【成都青白江分拣中心】', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1190, 'JDVD00973573239', '2020-09-20 01:25:06', '您的快件已发车', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1191, 'JDVD00973573239', '2020-09-20 01:55:03', '您的快件已到达【成都青白江分拣中心】', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1192, 'JDVD00973573239', '2020-09-20 02:08:51', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1193, 'JDVD00973573239', '2020-09-20 05:12:21', '您的快件由【成都青白江分拣中心】准备发往【成都经开营业部】', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1194, 'JDVD00973573239', '2020-09-20 05:33:53', '您的快件已发车', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1195, 'JDVD00973573239', '2020-09-20 07:15:01', '您的快件在【成都经开营业部】收货完成', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1196, 'JDVD00973573239', '2020-09-20 07:15:02', '您的快件已到达【成都经开营业部】', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1197, 'JDVD00973573239', '2020-09-20 08:12:44', '您的快件正在派送中，请您准备签收（快递员：宋强，联系电话：18215546285）', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1198, 'JDVD00973573239', '2020-09-20 10:54:20', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-20 10:54:46', '', '2020-09-20 10:54:46');
INSERT INTO `logistics_tbl` VALUES (1199, 'JDVD00977312945', '2020-09-20 13:32:10', '添加快递单', 'admin', '2020-09-20 13:32:10', '', '2020-09-20 13:32:10');
INSERT INTO `logistics_tbl` VALUES (1200, 'JDVD00977342825', '2020-09-20 13:32:33', '添加快递单', 'admin', '2020-09-20 13:32:33', '', '2020-09-20 13:32:33');
INSERT INTO `logistics_tbl` VALUES (1201, 'JDVD00977342825', '2020-09-20 13:21:29', '揽收任务已分配给寇方正', '', '2020-09-20 13:32:39', '', '2020-09-20 13:32:39');
INSERT INTO `logistics_tbl` VALUES (1202, 'JDVD00977342825', '2020-09-20 13:21:29', '揽收任务已分配给寇方正', '', '2020-09-20 13:32:39', '', '2020-09-20 13:32:39');
INSERT INTO `logistics_tbl` VALUES (1203, 'YT4793111423033', '2020-09-20 04:25:00', '【深圳转运中心公司】 已收入', '', '2020-09-20 13:51:31', '', '2020-09-20 13:51:31');
INSERT INTO `logistics_tbl` VALUES (1204, 'YT4793111423033', '2020-09-20 04:51:45', '【深圳转运中心】 已发出 下一站 【江门转运中心公司】', '', '2020-09-20 13:51:31', '', '2020-09-20 13:51:31');
INSERT INTO `logistics_tbl` VALUES (1205, 'YT4793111423033', '2020-09-20 10:13:04', '【江门转运中心公司】 已收入', '', '2020-09-20 13:51:31', '', '2020-09-20 13:51:31');
INSERT INTO `logistics_tbl` VALUES (1206, 'YT4793111423033', '2020-09-20 10:35:25', '【江门转运中心】 已发出 下一站 【广东省湛江市公司】', '', '2020-09-20 13:51:31', '', '2020-09-20 13:51:31');
INSERT INTO `logistics_tbl` VALUES (1207, 'JDVD00976222190', '2020-09-20 16:21:24', '京东快递 已收取快件', '', '2020-09-20 16:31:52', '', '2020-09-20 16:31:52');
INSERT INTO `logistics_tbl` VALUES (1208, 'JDVD00976222190', '2020-09-20 16:21:24', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-20 16:31:52', '', '2020-09-20 16:31:52');
INSERT INTO `logistics_tbl` VALUES (1209, 'JDVD00976222190', '2020-09-20 16:21:26', '您的快件已到达【广元苍溪营业部】', '', '2020-09-20 16:31:52', '', '2020-09-20 16:31:52');
INSERT INTO `logistics_tbl` VALUES (1210, 'YT4791786040000', '2020-09-19 20:56:32', '客户签收人: 磐石新城二区韵达快递超市 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18350034440，投诉电话：18059129309', '', '2020-09-20 20:00:01', '', '2020-09-20 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1211, 'YT4791895530363', '2020-09-20 11:41:38', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：13693433523，投诉电话：18521172098', '', '2020-09-20 20:00:04', '', '2020-09-20 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1212, 'YT4791911088503', '2020-09-20 09:14:20', '【四川省广元市苍溪县公司】 派件中  派件人: 屈大勇 电话 18683960483  如有疑问，请联系：0839-5880999', '', '2020-09-20 20:00:04', '', '2020-09-20 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1213, 'YT4791911088503', '2020-09-20 14:36:23', '客户签收人: 他人代收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18683960483，投诉电话：0839-5880999', '', '2020-09-20 20:00:04', '', '2020-09-20 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1214, 'YT4793111423033', '2020-09-20 18:50:51', '【广东省湛江市】 已发出 下一站 【广东省湛江市吴川市公司】', '', '2020-09-20 20:00:05', '', '2020-09-20 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1215, 'JDVD00976193380', '2020-09-20 16:21:20', '京东快递 已收取快件', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1216, 'JDVD00976193380', '2020-09-20 16:21:20', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1217, 'JDVD00976193380', '2020-09-20 16:21:23', '您的快件已到达【广元苍溪营业部】', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1218, 'JDVD00976193399', '2020-09-20 16:21:25', '京东快递 已收取快件', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1219, 'JDVD00976193399', '2020-09-20 16:21:25', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1220, 'JDVD00976193399', '2020-09-20 16:21:28', '您的快件已到达【广元苍溪营业部】', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1221, 'JDVD00976242970', '2020-09-20 16:21:28', '京东快递 已收取快件', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1222, 'JDVD00976242970', '2020-09-20 16:21:28', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1223, 'JDVD00976242970', '2020-09-20 16:21:30', '您的快件已到达【广元苍溪营业部】', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1224, 'JDVD00977312945', '2020-09-20 13:21:29', '揽收任务已分配给寇方正', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1225, 'JDVD00977312945', '2020-09-20 16:11:02', '京东快递 已收取快件', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1226, 'JDVD00977312945', '2020-09-20 16:11:02', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1227, 'JDVD00977312945', '2020-09-20 16:11:03', '您的快件已到达【广元苍溪营业部】', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1228, 'JDVD00977312945', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1229, 'JDVD00977312945', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-20 20:00:06', '', '2020-09-20 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1230, 'YT4801608577077', '2020-09-21 09:34:32', '添加快递单', 'admin', '2020-09-21 09:34:32', '', '2020-09-21 09:34:32');
INSERT INTO `logistics_tbl` VALUES (1231, 'YT4801608577077', '2020-09-21 08:05:17', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-21 11:13:07', '', '2020-09-21 11:13:07');
INSERT INTO `logistics_tbl` VALUES (1232, 'YT4801608577077', '2020-09-21 08:05:27', '【四川省直营市场部】 已发出 下一站 【成都转运中心公司】', '', '2020-09-21 11:13:07', '', '2020-09-21 11:13:07');
INSERT INTO `logistics_tbl` VALUES (1233, 'JDVD00976193380', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1234, 'JDVD00976193380', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1235, 'JDVD00976193380', '2020-09-21 00:17:17', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1236, 'JDVD00976193380', '2020-09-21 00:25:10', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1237, 'JDVD00976193380', '2020-09-21 01:37:24', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1238, 'JDVD00976193380', '2020-09-21 02:54:12', '您的快件已发车', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1239, 'JDVD00976193380', '2020-09-21 08:58:59', '您的快件已到达【重庆分拣中心】', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1240, 'JDVD00976193380', '2020-09-21 08:58:59', '重庆分拣中心分拣中心已收箱，箱号JDVD00976193380-1-1-', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1241, 'JDVD00976193380', '2020-09-21 09:18:32', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1242, 'JDVD00976193380', '2020-09-21 09:18:37', '您的快件由【重庆分拣中心】准备发往【重庆城南家园京东星配站】', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1243, 'JDVD00976193380', '2020-09-21 11:22:53', '您的快件已发车', '', '2020-09-21 11:58:05', '', '2020-09-21 11:58:05');
INSERT INTO `logistics_tbl` VALUES (1244, 'JDVD00976193399', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1245, 'JDVD00976193399', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1246, 'JDVD00976193399', '2020-09-21 00:12:35', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1247, 'JDVD00976193399', '2020-09-21 00:18:48', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1248, 'JDVD00976193399', '2020-09-21 01:05:55', '您的快件由【成都祥福分拣中心】准备发往【南充分拣场】', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1249, 'JDVD00976193399', '2020-09-21 01:25:11', '您的快件已发车', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1250, 'JDVD00976193399', '2020-09-21 04:48:11', '您的快件已到达【南充分拣场】', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1251, 'JDVD00976193399', '2020-09-21 04:49:20', '您的快件在【南充分拣场】分拣完成', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1252, 'JDVD00976193399', '2020-09-21 04:49:25', '您的快件由【南充分拣场】准备发往【南充仪陇金城营业部】', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1253, 'JDVD00976193399', '2020-09-21 07:16:30', '您的快件已发车', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1254, 'JDVD00976193399', '2020-09-21 08:43:15', '您的快件已到达【南充西充营业部】', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1255, 'JDVD00976193399', '2020-09-21 10:25:55', '您的快件在【南充西山营业部】收货完成', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1256, 'JDVD00976193399', '2020-09-21 10:25:56', '您的快件已到达【南充西山营业部】', '', '2020-09-21 11:58:15', '', '2020-09-21 11:58:15');
INSERT INTO `logistics_tbl` VALUES (1257, 'YT4793111423033', '2020-09-21 10:25:25', '【广东省湛江市吴川市公司】 已收入', '', '2020-09-21 15:06:04', '', '2020-09-21 15:06:04');
INSERT INTO `logistics_tbl` VALUES (1258, 'YT4793111423033', '2020-09-21 10:25:26', '【广东省湛江市吴川市公司】 派件中  派件人: 凌志浩 电话 19806862270  如有疑问，请联系：0759-5557729', '', '2020-09-21 15:06:04', '', '2020-09-21 15:06:04');
INSERT INTO `logistics_tbl` VALUES (1259, 'YT4793111423033', '2020-09-21 14:25:20', '客户签收人: 本人 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：19806862270，投诉电话：0759-5557729', '', '2020-09-21 15:06:04', '', '2020-09-21 15:06:04');
INSERT INTO `logistics_tbl` VALUES (1260, 'JDVD00976222190', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-21 15:09:37', '', '2020-09-21 15:09:37');
INSERT INTO `logistics_tbl` VALUES (1261, 'JDVD00976222190', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-21 15:09:37', '', '2020-09-21 15:09:37');
INSERT INTO `logistics_tbl` VALUES (1262, 'JDVD00976222190', '2020-09-21 00:12:55', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 15:09:37', '', '2020-09-21 15:09:37');
INSERT INTO `logistics_tbl` VALUES (1263, 'JDVD00976222190', '2020-09-21 00:18:22', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 15:09:37', '', '2020-09-21 15:09:37');
INSERT INTO `logistics_tbl` VALUES (1264, 'JDVD00976222190', '2020-09-21 00:18:27', '您的快件由【成都祥福分拣中心】准备发往【苏州昆山分拣中心】', '', '2020-09-21 15:09:37', '', '2020-09-21 15:09:37');
INSERT INTO `logistics_tbl` VALUES (1265, 'JDVD00976222190', '2020-09-21 01:55:13', '您的快件已发车', '', '2020-09-21 15:09:37', '', '2020-09-21 15:09:37');
INSERT INTO `logistics_tbl` VALUES (1266, 'JDVD00977312945', '2020-09-21 00:12:52', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1267, 'JDVD00977312945', '2020-09-21 00:18:21', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1268, 'JDVD00977312945', '2020-09-21 00:18:26', '您的快件由【成都祥福分拣中心】准备发往【成都青白江分拣中心】', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1269, 'JDVD00977312945', '2020-09-21 00:57:40', '您的快件已发车', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1270, 'JDVD00977312945', '2020-09-21 02:30:57', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1271, 'JDVD00977312945', '2020-09-21 05:29:41', '您的快件由【成都青白江分拣中心】准备发往【成都青龙场营业部】', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1272, 'JDVD00977312945', '2020-09-21 05:39:04', '您的快件已发车', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1273, 'JDVD00977312945', '2020-09-21 07:08:12', '您的快件在【成都青龙场营业部】收货完成', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1274, 'JDVD00977312945', '2020-09-21 07:08:13', '您的快件已到达【成都青龙场营业部】', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1275, 'JDVD00977312945', '2020-09-21 08:15:31', '您的快件正在派送中，请您准备签收（快递员：王思军，联系电话：18215687135）', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1276, 'JDVD00977312945', '2020-09-21 09:45:20', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-21 15:10:00', '', '2020-09-21 15:10:00');
INSERT INTO `logistics_tbl` VALUES (1277, 'JDVD00976185082', '2020-09-20 16:21:44', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1278, 'JDVD00976185082', '2020-09-20 16:21:44', '京东快递 已收取快件', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1279, 'JDVD00976185082', '2020-09-20 16:21:46', '您的快件已到达【广元苍溪营业部】', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1280, 'JDVD00976185082', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1281, 'JDVD00976185082', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1282, 'JDVD00976185082', '2020-09-21 00:20:08', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1283, 'JDVD00976185082', '2020-09-21 00:25:52', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1284, 'JDVD00976185082', '2020-09-21 00:25:57', '您的快件由【成都祥福分拣中心】准备发往【成都青白江分拣中心】', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1285, 'JDVD00976185082', '2020-09-21 01:04:18', '您的快件已发车', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1286, 'JDVD00976185082', '2020-09-21 01:45:33', '您的快件已到达【成都青白江分拣中心】', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1287, 'JDVD00976185082', '2020-09-21 03:16:56', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1288, 'JDVD00976185082', '2020-09-21 05:26:49', '您的快件由【成都青白江分拣中心】准备发往【成都经开营业部】', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1289, 'JDVD00976185082', '2020-09-21 05:43:37', '您的快件已发车', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1290, 'JDVD00976185082', '2020-09-21 08:29:30', '您的快件在【成都经开营业部】收货完成', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1291, 'JDVD00976185082', '2020-09-21 08:29:31', '您的快件已到达【成都经开营业部】', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1292, 'JDVD00976185082', '2020-09-21 08:49:08', '您的快件正在派送中，请您准备签收（快递员：邓云，联系电话：13688432407）', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1293, 'JDVD00976185082', '2020-09-21 10:34:01', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-21 15:10:20', '', '2020-09-21 15:10:20');
INSERT INTO `logistics_tbl` VALUES (1294, 'JDVD00976242970', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-21 15:10:52', '', '2020-09-21 15:10:52');
INSERT INTO `logistics_tbl` VALUES (1295, 'JDVD00976242970', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-21 15:10:52', '', '2020-09-21 15:10:52');
INSERT INTO `logistics_tbl` VALUES (1296, 'JDVD00976242970', '2020-09-21 00:14:29', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 15:10:52', '', '2020-09-21 15:10:52');
INSERT INTO `logistics_tbl` VALUES (1297, 'JDVD00976242970', '2020-09-21 00:20:52', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 15:10:52', '', '2020-09-21 15:10:52');
INSERT INTO `logistics_tbl` VALUES (1298, 'JDVD00976242970', '2020-09-21 00:20:57', '您的快件由【成都祥福分拣中心】准备发往【华中阳逻分拣中心】', '', '2020-09-21 15:10:52', '', '2020-09-21 15:10:52');
INSERT INTO `logistics_tbl` VALUES (1299, 'JDVD00976242970', '2020-09-21 02:25:14', '您的快件已发车', '', '2020-09-21 15:10:52', '', '2020-09-21 15:10:52');
INSERT INTO `logistics_tbl` VALUES (1300, 'JDVD00977342825', '2020-09-20 16:22:01', '京东快递 已收取快件', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1301, 'JDVD00977342825', '2020-09-20 16:22:01', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1302, 'JDVD00977342825', '2020-09-20 16:22:04', '您的快件已到达【广元苍溪营业部】', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1303, 'JDVD00977342825', '2020-09-20 18:45:21', '快递司机收箱', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1304, 'JDVD00977342825', '2020-09-20 18:45:25', '您的快件已发车', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1305, 'JDVD00977342825', '2020-09-21 00:11:38', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1306, 'JDVD00977342825', '2020-09-21 00:18:02', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1307, 'JDVD00977342825', '2020-09-21 00:18:07', '您的快件由【成都祥福分拣中心】准备发往【华中阳逻分拣中心】', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1308, 'JDVD00977342825', '2020-09-21 02:25:14', '您的快件已发车', '', '2020-09-21 15:11:12', '', '2020-09-21 15:11:12');
INSERT INTO `logistics_tbl` VALUES (1309, 'JDVD00976193380', '2020-09-21 12:36:43', '您的快件在【重庆城南家园京东星配站】收货完成', '', '2020-09-21 16:22:24', '', '2020-09-21 16:22:24');
INSERT INTO `logistics_tbl` VALUES (1310, 'JDVD00976193380', '2020-09-21 12:36:44', '您的快件已到达【重庆城南家园京东星配站】', '', '2020-09-21 16:22:24', '', '2020-09-21 16:22:24');
INSERT INTO `logistics_tbl` VALUES (1311, 'JDVD00976193380', '2020-09-21 13:56:18', '您的快件正在派送中，请您准备签收（快递员：陈羌飞，联系电话：18223563080）', '', '2020-09-21 16:22:24', '', '2020-09-21 16:22:24');
INSERT INTO `logistics_tbl` VALUES (1312, 'JDVD00976193380', '2020-09-21 15:45:02', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-21 16:22:24', '', '2020-09-21 16:22:24');
INSERT INTO `logistics_tbl` VALUES (1313, 'YT4803129557688', '2020-09-21 17:47:30', '添加快递单', 'admin', '2020-09-21 17:47:30', '', '2020-09-21 17:47:30');
INSERT INTO `logistics_tbl` VALUES (1314, 'YT4803127585304', '2020-09-21 17:47:55', '添加快递单', 'admin', '2020-09-21 17:47:55', '', '2020-09-21 17:47:55');
INSERT INTO `logistics_tbl` VALUES (1315, 'YT4791898277719', '2020-09-21 17:50:44', '客户签收人: 他人代签,代收点代收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：15931704117，投诉电话：0317-2215988', '', '2020-09-21 20:00:01', '', '2020-09-21 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1316, 'JDVD00976193399', '2020-09-21 18:27:30', '您的快件已到达【南充西山营业部】', '', '2020-09-21 20:00:01', '', '2020-09-21 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1317, 'JDVD00976193399', '2020-09-21 18:32:16', '快递司机收箱', '', '2020-09-21 20:00:01', '', '2020-09-21 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1318, 'JDVD00976193399', '2020-09-21 18:32:20', '您的快件已发车', '', '2020-09-21 20:00:01', '', '2020-09-21 20:00:01');
INSERT INTO `logistics_tbl` VALUES (1319, 'YT4803129557688', '2020-09-21 17:53:06', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-21 20:00:04', '', '2020-09-21 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1320, 'YT4803127585304', '2020-09-21 17:53:08', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-09-21 20:00:04', '', '2020-09-21 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1321, 'JDVD00983085779', '2020-09-22 11:11:50', '添加快递单', 'admin', '2020-09-22 11:11:50', '', '2020-09-22 11:11:50');
INSERT INTO `logistics_tbl` VALUES (1322, 'JDVD00976222190', '2020-09-22 08:19:23', '您的快件已到达【苏州昆山分拣中心】', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1323, 'JDVD00976222190', '2020-09-22 08:19:23', '苏州昆山分拣中心分拣中心已收箱，箱号JDVD00976222190-1-1-', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1324, 'JDVD00976222190', '2020-09-22 08:25:10', '您的快件在【苏州昆山分拣中心】分拣完成', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1325, 'JDVD00976222190', '2020-09-22 08:25:15', '您的快件由【苏州昆山分拣中心】准备发往【苏州淀山分拣中心】', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1326, 'JDVD00976222190', '2020-09-22 09:37:52', '您的快件已发车', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1327, 'JDVD00976222190', '2020-09-22 11:20:25', '您的快件在【苏州淀山分拣中心】分拣完成', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1328, 'JDVD00976222190', '2020-09-22 11:20:30', '您的快件由【苏州淀山分拣中心】准备发往【上海园新营业部】', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1329, 'JDVD00976222190', '2020-09-22 13:29:59', '您的快件已发车', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1330, 'JDVD00976222190', '2020-09-22 15:16:17', '您的快件已到达【上海园新营业部】', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1331, 'JDVD00976222190', '2020-09-22 15:31:57', '您的快件正在派送中，请您准备签收（快递员：王纪连，联系电话：15601681287）', '', '2020-09-22 15:37:21', '', '2020-09-22 15:37:21');
INSERT INTO `logistics_tbl` VALUES (1332, 'JDVD00983085779', '2020-09-22 11:06:09', '揽收任务已分配给寇方正', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1333, 'JDVD00983085779', '2020-09-22 16:07:52', '您的快件因【客户超时未准备好】操作再取，下次揽收时间2020-09-26 11:00-12:00', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1334, 'JDVD00983085779', '2020-09-22 18:20:24', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1335, 'JDVD00983085779', '2020-09-22 18:20:24', '京东快递 已收取快件', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1336, 'JDVD00983085779', '2020-09-22 18:20:26', '您的快件已到达【广元苍溪营业部】', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1337, 'JDVD00983085779', '2020-09-22 18:49:14', '快递司机收箱', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1338, 'JDVD00983085779', '2020-09-22 18:49:19', '您的快件已发车', '', '2020-09-22 22:31:06', '', '2020-09-22 22:31:06');
INSERT INTO `logistics_tbl` VALUES (1339, 'YT4803127585304', '2020-09-22 01:19:08', '【南充转运中心公司】 已收入', '', '2020-09-22 22:31:13', '', '2020-09-22 22:31:13');
INSERT INTO `logistics_tbl` VALUES (1340, 'YT4803127585304', '2020-09-22 01:47:20', '【南充转运中心】 已发出 下一站 【上海转运中心公司】', '', '2020-09-22 22:31:13', '', '2020-09-22 22:31:13');
INSERT INTO `logistics_tbl` VALUES (1341, 'YT4803129557688', '2020-09-21 23:24:32', '【南充转运中心公司】 已收入', '', '2020-09-22 22:31:13', '', '2020-09-22 22:31:13');
INSERT INTO `logistics_tbl` VALUES (1342, 'YT4803129557688', '2020-09-21 23:49:02', '【南充转运中心】 已发出 下一站 【上海转运中心公司】', '', '2020-09-22 22:31:13', '', '2020-09-22 22:31:13');
INSERT INTO `logistics_tbl` VALUES (1343, 'JDVD00976222190', '2020-09-22 16:24:48', '您的快件已由门卫代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-22 22:31:29', '', '2020-09-22 22:31:29');
INSERT INTO `logistics_tbl` VALUES (1344, 'YT4806894253200', '2020-09-23 08:11:47', '添加快递单', 'admin', '2020-09-23 08:11:47', '', '2020-09-23 08:11:47');
INSERT INTO `logistics_tbl` VALUES (1345, 'JDVD00983085779', '2020-09-23 02:01:08', '您的快件已到达【成都祥福分拣中心】', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1346, 'JDVD00983085779', '2020-09-23 02:01:41', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1347, 'JDVD00983085779', '2020-09-23 02:08:22', '您的快件由【成都祥福分拣中心】准备发往【重庆分拣中心】', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1348, 'JDVD00983085779', '2020-09-23 02:55:49', '您的快件已发车', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1349, 'JDVD00983085779', '2020-09-23 10:21:29', '重庆分拣中心分拣中心已收箱，箱号JDVD00983085779-1-1-', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1350, 'JDVD00983085779', '2020-09-23 10:21:29', '您的快件已到达【重庆分拣中心】', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1351, 'JDVD00983085779', '2020-09-23 10:40:54', '您的快件在【重庆分拣中心】分拣完成', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1352, 'JDVD00983085779', '2020-09-23 11:23:31', '您的快件由【重庆分拣中心】准备发往【重庆龙头寺营业部】', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1353, 'JDVD00983085779', '2020-09-23 11:24:16', '您的快件已发车', '', '2020-09-23 12:14:46', '', '2020-09-23 12:14:46');
INSERT INTO `logistics_tbl` VALUES (1354, 'YT4806894253200', '2020-09-23 09:03:33', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-23 15:50:16', '', '2020-09-23 15:50:16');
INSERT INTO `logistics_tbl` VALUES (1355, 'YT4806894253200', '2020-09-23 09:03:43', '【四川省直营市场部】 已发出 下一站 【昆明转运中心公司】', '', '2020-09-23 15:50:16', '', '2020-09-23 15:50:16');
INSERT INTO `logistics_tbl` VALUES (1356, 'JDVD00983085779', '2020-09-23 13:46:20', '您的快件在【重庆龙头寺营业部】收货完成', '', '2020-09-23 15:50:24', '', '2020-09-23 15:50:24');
INSERT INTO `logistics_tbl` VALUES (1357, 'JDVD00983085779', '2020-09-23 13:46:21', '您的快件已到达【重庆龙头寺营业部】', '', '2020-09-23 15:50:24', '', '2020-09-23 15:50:24');
INSERT INTO `logistics_tbl` VALUES (1358, 'JDVD00983085779', '2020-09-23 15:35:22', '您的快件正在派送中，请您准备签收（快递员：冉睿，联系电话：15696999917）', '', '2020-09-23 15:50:24', '', '2020-09-23 15:50:24');
INSERT INTO `logistics_tbl` VALUES (1359, 'YT4803127585304', '2020-09-23 11:14:28', '【浦东转运中心公司】 已收入', '', '2020-09-23 16:48:44', '', '2020-09-23 16:48:44');
INSERT INTO `logistics_tbl` VALUES (1360, 'YT4803127585304', '2020-09-23 11:42:30', '【浦东转运中心】 已发出 下一站 【上海市浦东机场公司】', '', '2020-09-23 16:48:44', '', '2020-09-23 16:48:44');
INSERT INTO `logistics_tbl` VALUES (1361, 'YT4803129557688', '2020-09-23 11:13:02', '【浦东转运中心公司】 已收入', '', '2020-09-23 16:48:44', '', '2020-09-23 16:48:44');
INSERT INTO `logistics_tbl` VALUES (1362, 'YT4803129557688', '2020-09-23 11:41:59', '【浦东转运中心】 已发出 下一站 【上海市浦东机场公司】', '', '2020-09-23 16:48:44', '', '2020-09-23 16:48:44');
INSERT INTO `logistics_tbl` VALUES (1363, 'YT4801608577077', '2020-09-23 21:02:47', '【昆明转运中心公司】 已收入', '', '2020-09-24 10:05:23', '', '2020-09-24 10:05:23');
INSERT INTO `logistics_tbl` VALUES (1364, 'YT4801608577077', '2020-09-23 21:32:32', '【昆明转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-09-24 10:05:23', '', '2020-09-24 10:05:23');
INSERT INTO `logistics_tbl` VALUES (1365, 'JDVD00983085779', '2020-09-23 16:53:55', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-24 11:41:26', '', '2020-09-24 11:41:26');
INSERT INTO `logistics_tbl` VALUES (1366, 'YT4803127585304', '2020-09-24 08:48:39', '【上海市浦东新区黄楼公司】 已收入', '', '2020-09-24 11:41:38', '', '2020-09-24 11:41:38');
INSERT INTO `logistics_tbl` VALUES (1367, 'YT4803127585304', '2020-09-24 09:22:04', '【上海市浦东新区黄楼公司】 派件中  派件人: 饶园丁 电话 17317214698 ', '', '2020-09-24 11:41:38', '', '2020-09-24 11:41:38');
INSERT INTO `logistics_tbl` VALUES (1368, 'YT4803129557688', '2020-09-24 08:50:34', '【上海市浦东新区黄楼公司】 已收入', '', '2020-09-24 11:41:38', '', '2020-09-24 11:41:38');
INSERT INTO `logistics_tbl` VALUES (1369, 'JDVD00977342825', '2020-09-21 21:37:10', '华中阳逻分拣中心分拣中心已收箱，箱号JDVD00977342825-1-1-', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1370, 'JDVD00977342825', '2020-09-21 21:37:10', '您的快件已到达【华中阳逻分拣中心】', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1371, 'JDVD00977342825', '2020-09-21 21:58:11', '您的快件在【华中阳逻分拣中心】分拣完成', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1372, 'JDVD00977342825', '2020-09-21 21:58:16', '您的快件由【华中阳逻分拣中心】准备发往【华中新洲分拣中心】', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1373, 'JDVD00977342825', '2020-09-21 22:36:16', '您的快件已发车', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1374, 'JDVD00977342825', '2020-09-21 23:22:29', '您的快件已到达【华中新洲分拣中心】', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1375, 'JDVD00977342825', '2020-09-21 23:22:29', '华中新洲分拣中心分拣中心已收箱，箱号JDVD00977342825-1-1-', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1376, 'JDVD00977342825', '2020-09-21 23:49:52', '您的快件在【华中新洲分拣中心】分拣完成', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1377, 'JDVD00977342825', '2020-09-21 23:49:57', '您的快件由【华中新洲分拣中心】准备发往【武汉龙城营业部】', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1378, 'JDVD00977342825', '2020-09-22 05:22:02', '您的快件已发车', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1379, 'JDVD00977342825', '2020-09-22 07:59:30', '您的快件在【武汉龙城营业部】收货完成', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1380, 'JDVD00977342825', '2020-09-22 07:59:31', '您的快件已到达【武汉龙城营业部】', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1381, 'JDVD00977342825', '2020-09-22 08:24:04', '您的快件正在派送中，请您准备签收（快递员：杨波，联系电话：18062010045）', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1382, 'JDVD00977342825', '2020-09-22 11:10:25', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-24 12:54:42', '', '2020-09-24 12:54:42');
INSERT INTO `logistics_tbl` VALUES (1383, 'YT4801608577077', '2020-09-24 12:41:31', '【重庆转运中心公司】 已收入', '', '2020-09-24 18:41:14', '', '2020-09-24 18:41:14');
INSERT INTO `logistics_tbl` VALUES (1384, 'YT4801608577077', '2020-09-24 13:04:14', '【重庆转运中心】 已发出 下一站 【重庆市万州区公司】', '', '2020-09-24 18:41:14', '', '2020-09-24 18:41:14');
INSERT INTO `logistics_tbl` VALUES (1385, 'YT4801608577077', '2020-09-24 18:06:29', '【重庆市万州区公司】 已收入', '', '2020-09-24 18:41:14', '', '2020-09-24 18:41:14');
INSERT INTO `logistics_tbl` VALUES (1386, 'YT4801608577077', '2020-09-24 18:20:45', '【重庆市万州区】 已发出 下一站 【重庆市云阳县公司】', '', '2020-09-24 18:41:14', '', '2020-09-24 18:41:14');
INSERT INTO `logistics_tbl` VALUES (1387, 'YT4812020656267', '2020-09-25 08:11:33', '添加快递单', 'admin', '2020-09-25 08:11:33', '', '2020-09-25 08:11:33');
INSERT INTO `logistics_tbl` VALUES (1388, 'YT4811974436620', '2020-09-25 08:11:59', '添加快递单', 'admin', '2020-09-25 08:11:59', '', '2020-09-25 08:11:59');
INSERT INTO `logistics_tbl` VALUES (1389, 'YT4801608577077', '2020-09-25 08:06:01', '【重庆市云阳县公司】 已收入', '', '2020-09-25 09:19:48', '', '2020-09-25 09:19:48');
INSERT INTO `logistics_tbl` VALUES (1390, 'YT4801608577077', '2020-09-25 08:37:52', '【重庆市云阳县公司】 派件中  派件人: 冉燃 电话 18521172394  如有疑问，请联系：023-55121886', '', '2020-09-25 09:19:48', '', '2020-09-25 09:19:48');
INSERT INTO `logistics_tbl` VALUES (1391, 'YT4803127585304', '2020-09-24 15:34:17', '圆通合作点【兔喜快递超市】快件已到达508弄73号104室驿站,联系电话17802118959', '', '2020-09-25 12:43:41', '', '2020-09-25 12:43:41');
INSERT INTO `logistics_tbl` VALUES (1392, 'YT4803129557688', '2020-09-24 11:45:39', '【上海市浦东新区黄楼公司】 派件中  派件人: 饶园丁 电话 17317214698 ', '', '2020-09-25 12:43:41', '', '2020-09-25 12:43:41');
INSERT INTO `logistics_tbl` VALUES (1393, 'YT4803129557688', '2020-09-24 15:34:37', '圆通合作点【兔喜快递超市】快件已到达508弄73号104室驿站,联系电话17802118959', '', '2020-09-25 12:43:41', '', '2020-09-25 12:43:41');
INSERT INTO `logistics_tbl` VALUES (1394, 'YT4801608577077', '2020-09-25 09:58:32', '客户签收人: 同事或家人 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521172394，投诉电话：023-55121886', '', '2020-09-25 12:43:45', '', '2020-09-25 12:43:45');
INSERT INTO `logistics_tbl` VALUES (1395, 'YT4812020656267', '2020-09-25 09:03:39', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-25 12:45:37', '', '2020-09-25 12:45:37');
INSERT INTO `logistics_tbl` VALUES (1396, 'YT4812020656267', '2020-09-25 09:03:49', '【四川省直营市场部】 已发出 下一站 【昆明转运中心公司】', '', '2020-09-25 12:45:37', '', '2020-09-25 12:45:37');
INSERT INTO `logistics_tbl` VALUES (1397, 'YT4812020656267', '2020-09-26 19:03:34', '【昆明转运中心公司】 已收入', '', '2020-09-27 14:51:49', '', '2020-09-27 14:51:49');
INSERT INTO `logistics_tbl` VALUES (1398, 'YT4812020656267', '2020-09-26 19:06:20', '【昆明转运中心】 已发出 下一站 【北京转运中心公司】', '', '2020-09-27 14:51:49', '', '2020-09-27 14:51:49');
INSERT INTO `logistics_tbl` VALUES (1399, 'YT4806894253200', '2020-09-25 21:20:26', '【成都转运中心公司】 已收入', '', '2020-09-27 17:47:02', '', '2020-09-27 17:47:02');
INSERT INTO `logistics_tbl` VALUES (1400, 'YT4806894253200', '2020-09-25 21:23:05', '【成都转运中心】 已发出 下一站 【华北昌平城配中心公司】', '', '2020-09-27 17:47:02', '', '2020-09-27 17:47:02');
INSERT INTO `logistics_tbl` VALUES (1401, 'YT4806894253200', '2020-09-27 11:56:24', '【华北昌平城配中心公司】 已收入', '', '2020-09-27 17:47:02', '', '2020-09-27 17:47:02');
INSERT INTO `logistics_tbl` VALUES (1402, 'YT4806894253200', '2020-09-27 11:57:55', '【华北昌平城配中心】 已发出 下一站 【北京市海淀区科贸公司】', '', '2020-09-27 17:47:02', '', '2020-09-27 17:47:02');
INSERT INTO `logistics_tbl` VALUES (1403, 'YT4806894253200', '2020-09-27 11:58:50', '【北京市海淀区科贸公司】 已收入', '', '2020-09-27 17:47:02', '', '2020-09-27 17:47:02');
INSERT INTO `logistics_tbl` VALUES (1404, 'YT4806894253200', '2020-09-27 15:52:48', '【北京市海淀区科贸公司】 派件中  派件人: 梁震坤 电话 18521850993 。 圆通快递小哥每天已测体温，请放心收寄快递 如有疑问，请联系：010-53807290', '', '2020-09-27 17:47:02', '', '2020-09-27 17:47:02');
INSERT INTO `logistics_tbl` VALUES (1405, 'YT4811974436620', '2020-09-25 09:03:43', '【四川省直营市场部公司】 已收件 取件人: 会理石榴仓 (19160279925)', '', '2020-09-27 23:29:35', '', '2020-09-27 23:29:35');
INSERT INTO `logistics_tbl` VALUES (1406, 'YT4811974436620', '2020-09-25 09:03:53', '【四川省直营市场部】 已发出 下一站 【昆明转运中心公司】', '', '2020-09-27 23:29:35', '', '2020-09-27 23:29:35');
INSERT INTO `logistics_tbl` VALUES (1407, 'YT4811974436620', '2020-09-26 17:30:06', '【昆明转运中心公司】 已收入', '', '2020-09-27 23:29:35', '', '2020-09-27 23:29:35');
INSERT INTO `logistics_tbl` VALUES (1408, 'YT4811974436620', '2020-09-26 18:20:06', '【昆明转运中心】 已发出', '', '2020-09-27 23:29:35', '', '2020-09-27 23:29:35');
INSERT INTO `logistics_tbl` VALUES (1409, 'YT4806894253200', '2020-09-27 20:12:17', '【北京市海淀区科贸公司】 失败签收录入 杨少婷 （010-60936479）', '', '2020-09-28 09:10:08', '', '2020-09-28 09:10:08');
INSERT INTO `logistics_tbl` VALUES (1410, 'YT4811974436620', '2020-09-28 21:32:20', '【成都转运中心公司】 已收入', '', '2020-09-28 22:10:20', '', '2020-09-28 22:10:20');
INSERT INTO `logistics_tbl` VALUES (1411, 'YT4811974436620', '2020-09-28 21:41:34', '【成都转运中心】 已发出 下一站 【自贡转运中心公司】', '', '2020-09-28 22:10:20', '', '2020-09-28 22:10:20');
INSERT INTO `logistics_tbl` VALUES (1412, 'JDVD00976193399', '2020-09-22 07:27:32', '您的快件已发车', '', '2020-09-28 22:45:50', '', '2020-09-28 22:45:50');
INSERT INTO `logistics_tbl` VALUES (1413, 'JDVD00976193399', '2020-09-22 11:35:13', '您的快件已到达【南充仪陇金城营业部】', '', '2020-09-28 22:45:50', '', '2020-09-28 22:45:50');
INSERT INTO `logistics_tbl` VALUES (1414, 'JDVD00976193399', '2020-09-22 11:59:52', '您的快件正在派送中，请您准备签收（快递员：龙杰，联系电话：18582385101）', '', '2020-09-28 22:45:50', '', '2020-09-28 22:45:50');
INSERT INTO `logistics_tbl` VALUES (1415, 'JDVD00976193399', '2020-09-22 17:28:09', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-28 22:45:50', '', '2020-09-28 22:45:50');
INSERT INTO `logistics_tbl` VALUES (1416, 'JDVD00976242970', '2020-09-21 21:30:29', '华中阳逻分拣中心分拣中心已收箱，箱号JDVD00976242970-1-1-', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1417, 'JDVD00976242970', '2020-09-21 21:30:29', '您的快件已到达【华中阳逻分拣中心】', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1418, 'JDVD00976242970', '2020-09-21 21:40:45', '您的快件在【华中阳逻分拣中心】分拣完成', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1419, 'JDVD00976242970', '2020-09-21 21:40:50', '您的快件由【华中阳逻分拣中心】准备发往【镇江句容分拣中心】', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1420, 'JDVD00976242970', '2020-09-22 01:53:15', '您的快件已发车', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1421, 'JDVD00976242970', '2020-09-22 11:09:52', '您的快件已到达【镇江句容分拣中心】', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1422, 'JDVD00976242970', '2020-09-22 11:14:08', '您的快件在【镇江句容分拣中心】分拣完成', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1423, 'JDVD00976242970', '2020-09-22 11:14:13', '您的快件由【镇江句容分拣中心】准备发往【南京高新营业部】', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1424, 'JDVD00976242970', '2020-09-22 11:57:09', '您的快件已发车', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1425, 'JDVD00976242970', '2020-09-22 15:15:04', '您的快件在【南京高新营业部】收货完成', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1426, 'JDVD00976242970', '2020-09-22 15:15:05', '您的快件已到达【南京高新营业部】', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1427, 'JDVD00976242970', '2020-09-22 15:44:23', '您的快件正在派送中，请您准备签收（快递员：王星桥，联系电话：19852837047）', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1428, 'JDVD00976242970', '2020-09-22 16:52:20', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1429, 'YT4803129557688', '2020-09-25 22:07:15', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：17317214698', '', '2020-09-28 22:45:51', '', '2020-09-28 22:45:51');
INSERT INTO `logistics_tbl` VALUES (1430, 'YT4803127585304', '2020-09-25 22:07:23', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：17317214698', '', '2020-09-28 22:45:52', '', '2020-09-28 22:45:52');
INSERT INTO `logistics_tbl` VALUES (1431, 'YT4812020656267', '2020-09-28 21:22:24', '【北京转运中心公司】 已收入', '', '2020-09-28 22:45:52', '', '2020-09-28 22:45:52');
INSERT INTO `logistics_tbl` VALUES (1432, 'YT4812020656267', '2020-09-28 21:26:05', '【北京转运中心】 已发出 下一站 【北京市朝阳区北苑公司】', '', '2020-09-28 22:45:52', '', '2020-09-28 22:45:52');
INSERT INTO `logistics_tbl` VALUES (1433, 'YT4812020656267', '2020-09-29 06:01:39', '【北京市朝阳区北苑公司】 已收入', '', '2020-09-29 08:09:48', '', '2020-09-29 08:09:48');
INSERT INTO `logistics_tbl` VALUES (1434, 'YT4812020656267', '2020-09-29 06:02:05', '【北京市朝阳区北苑公司】 派件中  派件人: 刘阳阳 电话 18521852490 。 圆通快递小哥每天已测体温，请放心收寄快递 如有疑问，请联系：010-53572996', '', '2020-09-29 08:09:48', '', '2020-09-29 08:09:48');
INSERT INTO `logistics_tbl` VALUES (1435, 'YT4812020656267', '2020-09-29 07:55:33', '快件已存放至佳兴园西门入口菜鸟智能柜【自提柜】，请及时取件。有问题请联系派件员18521852490', '', '2020-09-29 08:09:48', '', '2020-09-29 08:09:48');
INSERT INTO `logistics_tbl` VALUES (1436, '9880201780019', '2020-09-29 13:34:59', '添加快递单', 'admin', '2020-09-29 13:34:59', '', '2020-09-29 13:34:59');
INSERT INTO `logistics_tbl` VALUES (1437, '9880201780019', '2020-09-28 22:47:29', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:孙光琼,电话:13378374371', '', '2020-09-29 13:35:02', '', '2020-09-29 13:35:02');
INSERT INTO `logistics_tbl` VALUES (1438, '9880201780019', '2020-09-28 23:46:13', '[凉山彝族自治州]离开【四川会理县中心】,下一站【成都邮件处理中心】', '', '2020-09-29 13:35:02', '', '2020-09-29 13:35:02');
INSERT INTO `logistics_tbl` VALUES (1439, '9880201780019', '2020-09-28 23:52:09', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-09-29 13:35:02', '', '2020-09-29 13:35:02');
INSERT INTO `logistics_tbl` VALUES (1440, '9880201780019', '2020-09-28 23:54:28', '[凉山彝族自治州]到达【四川会理县中心】', '', '2020-09-29 13:35:02', '', '2020-09-29 13:35:02');
INSERT INTO `logistics_tbl` VALUES (1441, '9880201780019', '2020-09-29 11:43:05', '[成都市]到达【成都邮件处理中心】', '', '2020-09-29 13:35:02', '', '2020-09-29 13:35:02');
INSERT INTO `logistics_tbl` VALUES (1442, '9880201780019', '2020-09-29 13:01:47', '[成都市]离开【成都邮件处理中心】,下一站【成都双流区华阳营业部】', '', '2020-09-29 13:35:02', '', '2020-09-29 13:35:02');
INSERT INTO `logistics_tbl` VALUES (1443, 'YT4823368045243', '2020-09-29 16:21:57', '添加快递单', 'admin', '2020-09-29 16:21:57', '', '2020-09-29 16:21:57');
INSERT INTO `logistics_tbl` VALUES (1444, '9880209246003', '2020-09-30 08:10:31', '添加快递单', 'admin', '2020-09-30 08:10:31', '', '2020-09-30 08:10:31');
INSERT INTO `logistics_tbl` VALUES (1445, 'YT4812020656267', '2020-09-29 19:20:40', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521852490，投诉电话：010-53572996。疫情期间圆通每天对网点多次消毒，快递小哥每天测量体温，佩戴口罩', '', '2020-09-30 15:45:16', '', '2020-09-30 15:45:16');
INSERT INTO `logistics_tbl` VALUES (1446, '9880218742576', '2020-10-01 08:10:26', '添加快递单', 'admin', '2020-10-01 08:10:26', '', '2020-10-01 08:10:26');
INSERT INTO `logistics_tbl` VALUES (1447, '9880218751881', '2020-10-01 08:10:47', '添加快递单', 'admin', '2020-10-01 08:10:47', '', '2020-10-01 08:10:47');
INSERT INTO `logistics_tbl` VALUES (1448, '9880218737813', '2020-10-01 08:11:04', '添加快递单', 'admin', '2020-10-01 08:11:04', '', '2020-10-01 08:11:04');
INSERT INTO `logistics_tbl` VALUES (1449, '9880218728209', '2020-10-01 08:11:20', '添加快递单', 'admin', '2020-10-01 08:11:20', '', '2020-10-01 08:11:20');
INSERT INTO `logistics_tbl` VALUES (1450, '9880218728209', '2020-09-30 20:54:25', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:吴开贵,电话:17620176885', '', '2020-10-01 08:11:33', '', '2020-10-01 08:11:33');
INSERT INTO `logistics_tbl` VALUES (1451, '9880218728209', '2020-09-30 23:26:00', '[凉山彝族自治州]离开【四川会理县中心】,下一站【成都邮件处理中心】', '', '2020-10-01 08:11:33', '', '2020-10-01 08:11:33');
INSERT INTO `logistics_tbl` VALUES (1452, '9880218728209', '2020-09-30 23:55:38', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-01 08:11:33', '', '2020-10-01 08:11:33');
INSERT INTO `logistics_tbl` VALUES (1453, '9880218728209', '2020-10-01 04:05:11', '[凉山彝族自治州]到达【四川会理县中心】', '', '2020-10-01 08:11:33', '', '2020-10-01 08:11:33');
INSERT INTO `logistics_tbl` VALUES (1454, '9880218737813', '2020-09-30 20:54:34', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:吴开贵,电话:17620176885', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1455, '9880218737813', '2020-09-30 23:26:00', '[凉山彝族自治州]离开【四川会理县中心】,下一站【成都邮件处理中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1456, '9880218737813', '2020-09-30 23:55:38', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1457, '9880218737813', '2020-10-01 04:05:11', '[凉山彝族自治州]到达【四川会理县中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1458, '9880218742576', '2020-09-30 20:54:31', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:吴开贵,电话:17620176885', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1459, '9880218742576', '2020-09-30 23:26:00', '[凉山彝族自治州]离开【四川会理县中心】,下一站【成都邮件处理中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1460, '9880218742576', '2020-09-30 23:55:38', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1461, '9880218742576', '2020-10-01 04:05:11', '[凉山彝族自治州]到达【四川会理县中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1462, '9880218751881', '2020-09-30 20:54:36', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:吴开贵,电话:17620176885', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1463, '9880218751881', '2020-09-30 23:26:00', '[凉山彝族自治州]离开【四川会理县中心】,下一站【成都邮件处理中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1464, '9880218751881', '2020-09-30 23:55:38', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1465, '9880218751881', '2020-10-01 04:05:11', '[凉山彝族自治州]到达【四川会理县中心】', '', '2020-10-01 08:11:34', '', '2020-10-01 08:11:34');
INSERT INTO `logistics_tbl` VALUES (1466, 'YT4828059531673', '2020-10-02 11:33:12', '添加快递单', 'admin', '2020-10-02 11:33:12', '', '2020-10-02 11:33:12');
INSERT INTO `logistics_tbl` VALUES (1467, '9880218728209', '2020-10-02 11:39:48', '[成都市]到达【成都邮件处理中心】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1468, '9880218728209', '2020-10-02 12:44:36', '[成都市]离开【成都邮件处理中心】,下一站【公平营业部】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1469, '9880218737813', '2020-10-02 11:39:48', '[成都市]到达【成都邮件处理中心】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1470, '9880218742576', '2020-10-02 11:39:48', '[成都市]到达【成都邮件处理中心】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1471, '9880218742576', '2020-10-02 12:44:36', '[成都市]离开【成都邮件处理中心】,下一站【公平营业部】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1472, '9880218751881', '2020-10-02 11:39:48', '[成都市]到达【成都邮件处理中心】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1473, '9880218751881', '2020-10-02 12:44:36', '[成都市]离开【成都邮件处理中心】,下一站【公平营业部】', '', '2020-10-02 14:15:48', '', '2020-10-02 14:15:48');
INSERT INTO `logistics_tbl` VALUES (1474, '9880209246003', '2020-09-29 20:33:31', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:吴开贵,电话:17620176885', '', '2020-10-02 15:10:42', '', '2020-10-02 15:10:42');
INSERT INTO `logistics_tbl` VALUES (1475, '9880209246003', '2020-09-29 23:56:38', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-02 15:10:42', '', '2020-10-02 15:10:42');
INSERT INTO `logistics_tbl` VALUES (1476, '9880209246003', '2020-09-30 00:00:04', '[凉山彝族自治州]到达【四川会理县中心】', '', '2020-10-02 15:10:42', '', '2020-10-02 15:10:42');
INSERT INTO `logistics_tbl` VALUES (1477, '9880209246003', '2020-09-30 00:18:40', '[凉山彝族自治州]离开【四川会理县中心】,下一站【成都航站】', '', '2020-10-02 15:10:42', '', '2020-10-02 15:10:42');
INSERT INTO `logistics_tbl` VALUES (1478, '9880209246003', '2020-09-30 18:09:35', '[成都市]到达【成都航站】（经转）', '', '2020-10-02 15:10:42', '', '2020-10-02 15:10:42');
INSERT INTO `logistics_tbl` VALUES (1479, '9880209246003', '2020-09-30 20:56:36', '[成都市]离开【成都航站】,下一站【北京综合邮件处理中心】（经转）', '', '2020-10-02 15:10:42', '', '2020-10-02 15:10:42');
INSERT INTO `logistics_tbl` VALUES (1480, 'YT4823368045243', '2020-09-29 17:41:01', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1481, 'YT4823368045243', '2020-09-29 20:40:12', '【南充转运中心公司】 已收入', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1482, 'YT4823368045243', '2020-09-29 20:42:13', '【南充转运中心】 已发出 下一站 【重庆转运中心公司】', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1483, 'YT4823368045243', '2020-09-30 03:26:56', '【重庆转运中心公司】 已收入', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1484, 'YT4823368045243', '2020-09-30 04:16:56', '【重庆转运中心】 已发出', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1485, 'YT4823368045243', '2020-10-01 07:10:15', '【北京转运中心公司】 已收入', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1486, 'YT4823368045243', '2020-10-01 07:12:28', '【北京转运中心】 已发出 下一站 【华北昌平城配中心公司】', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1487, 'YT4823368045243', '2020-10-01 19:13:29', '【华北昌平城配中心公司】 已收入', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1488, 'YT4823368045243', '2020-10-01 19:14:43', '【华北昌平城配中心】 已发出 下一站 【北京市海淀区北师大公司】', '', '2020-10-02 15:10:43', '', '2020-10-02 15:10:43');
INSERT INTO `logistics_tbl` VALUES (1489, '9880218728209', '2020-10-02 14:15:16', '[成都市]到达【公平营业部】', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1490, '9880218728209', '2020-10-02 14:39:25', '[成都市]【公平营业部】安排投递,投递员:卢永志,电话:19827574748,揽投部电话:028-83284405', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1491, '9880218728209', '2020-10-02 15:07:03', '[成都市]已签收,他人代收：菜鸟驿站。,投递员:卢永志,电话:19827574748', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1492, '9880218742576', '2020-10-02 14:15:16', '[成都市]到达【公平营业部】', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1493, '9880218742576', '2020-10-02 14:39:25', '[成都市]【公平营业部】安排投递,投递员:卢永志,电话:19827574748,揽投部电话:028-83284405', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1494, '9880218742576', '2020-10-02 15:07:03', '[成都市]已签收,他人代收：菜鸟驿站。,投递员:卢永志,电话:19827574748', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1495, '9880218751881', '2020-10-02 14:15:16', '[成都市]到达【公平营业部】', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1496, '9880218751881', '2020-10-02 14:39:25', '[成都市]【公平营业部】安排投递,投递员:卢永志,电话:19827574748,揽投部电话:028-83284405', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1497, '9880218751881', '2020-10-02 15:07:03', '[成都市]已签收,他人代收：菜鸟驿站。,投递员:卢永志,电话:19827574748', '', '2020-10-02 15:35:22', '', '2020-10-02 15:35:22');
INSERT INTO `logistics_tbl` VALUES (1498, 'YT4828059531673', '2020-10-02 16:30:55', '【四川省直营市场部公司】 已收件 取件人: 罗小玲 (15583199578)', '', '2020-10-02 21:07:01', '', '2020-10-02 21:07:01');
INSERT INTO `logistics_tbl` VALUES (1499, 'YT4828059531673', '2020-10-02 19:54:49', '【南充转运中心公司】 已收入', '', '2020-10-02 21:07:01', '', '2020-10-02 21:07:01');
INSERT INTO `logistics_tbl` VALUES (1500, 'YT4828059531673', '2020-10-02 19:57:06', '【南充转运中心】 已发出 下一站 【成都转运中心公司】', '', '2020-10-02 21:07:01', '', '2020-10-02 21:07:01');
INSERT INTO `logistics_tbl` VALUES (1501, 'YT4828059531673', '2020-10-03 03:06:10', '【成都转运中心公司】 已收入', '', '2020-10-03 09:12:46', '', '2020-10-03 09:12:46');
INSERT INTO `logistics_tbl` VALUES (1502, 'YT4828059531673', '2020-10-03 03:16:37', '【成都转运中心】 已发出 下一站 【虎门转运中心公司】', '', '2020-10-03 09:12:46', '', '2020-10-03 09:12:46');
INSERT INTO `logistics_tbl` VALUES (1503, 'YT4828059531673', '2020-10-04 08:27:04', '【虎门转运中心公司】 已收入', '', '2020-10-04 12:25:22', '', '2020-10-04 12:25:22');
INSERT INTO `logistics_tbl` VALUES (1504, 'YT4828059531673', '2020-10-04 09:00:27', '【虎门转运中心】 已发出 下一站 【江门转运中心公司】', '', '2020-10-04 12:25:22', '', '2020-10-04 12:25:22');
INSERT INTO `logistics_tbl` VALUES (1505, 'YT4828059531673', '2020-10-04 12:24:16', '【江门转运中心公司】 已收入', '', '2020-10-04 12:25:22', '', '2020-10-04 12:25:22');
INSERT INTO `logistics_tbl` VALUES (1506, 'YT4828059531673', '2020-10-04 13:14:16', '【江门转运中心】 已发出', '', '2020-10-04 18:56:58', '', '2020-10-04 18:56:58');
INSERT INTO `logistics_tbl` VALUES (1507, 'YT4828059531673', '2020-10-05 03:45:15', '【佛山转运中心公司】 已收入', '', '2020-10-05 07:41:09', '', '2020-10-05 07:41:09');
INSERT INTO `logistics_tbl` VALUES (1508, 'YT4828059531673', '2020-10-05 03:55:12', '【佛山转运中心】 已发出 下一站 【广东省广州市南沙区自贸区公司】', '', '2020-10-05 07:41:09', '', '2020-10-05 07:41:09');
INSERT INTO `logistics_tbl` VALUES (1509, 'YT4828059531673', '2020-10-06 07:21:39', '【广东省广州市南沙区自贸区公司】 已收入', '', '2020-10-06 12:14:27', '', '2020-10-06 12:14:27');
INSERT INTO `logistics_tbl` VALUES (1510, 'YT4828059531673', '2020-10-06 08:35:24', '【广东省广州市南沙区自贸区】 已发出 下一站 【广东省广州市自贸区东井分部公司】', '', '2020-10-06 12:14:27', '', '2020-10-06 12:14:27');
INSERT INTO `logistics_tbl` VALUES (1511, 'YT4828059531673', '2020-10-06 11:27:16', '【广东省广州市自贸区东井分部公司】 派件中  派件人: 吴少龙 电话 13288689994 ', '', '2020-10-06 12:14:27', '', '2020-10-06 12:14:27');
INSERT INTO `logistics_tbl` VALUES (1512, 'YT4828059531673', '2020-10-06 15:34:45', '快件已暂存至广州南沙富力华庭五街3号店菜鸟驿站，如有疑问请联系13719045738', '', '2020-10-06 17:43:00', '', '2020-10-06 17:43:00');
INSERT INTO `logistics_tbl` VALUES (1513, 'YT4828059531673', '2020-10-06 22:26:38', '客户签收人: 已签收，签收人凭取货码签收。 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：13288689994', '', '2020-10-07 00:46:46', '', '2020-10-07 00:46:46');
INSERT INTO `logistics_tbl` VALUES (1514, '4309032943340', '2020-10-07 17:36:52', '添加快递单', 'admin', '2020-10-07 17:36:52', '', '2020-10-07 17:36:52');
INSERT INTO `logistics_tbl` VALUES (1515, '4309032935655', '2020-10-07 17:37:13', '添加快递单', 'admin', '2020-10-07 17:37:13', '', '2020-10-07 17:37:13');
INSERT INTO `logistics_tbl` VALUES (1516, '4309032960869', '2020-10-07 17:37:52', '添加快递单', 'admin', '2020-10-07 17:37:52', '', '2020-10-07 17:37:52');
INSERT INTO `logistics_tbl` VALUES (1517, '4309033023721', '2020-10-07 17:38:05', '添加快递单', 'admin', '2020-10-07 17:38:05', '', '2020-10-07 17:38:05');
INSERT INTO `logistics_tbl` VALUES (1518, '4309032908240', '2020-10-07 17:44:39', '添加快递单', 'admin', '2020-10-07 17:44:39', '', '2020-10-07 17:44:39');
INSERT INTO `logistics_tbl` VALUES (1519, '4309032908240', '2020-10-07 16:57:35', '【广元市】四川苍溪县公司 已揽收', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1520, '4309032908240', '2020-10-08 01:44:10', '【成都市】已到达 四川成都分拨中心', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1521, '4309032908240', '2020-10-08 02:26:39', '【成都市】已离开 四川成都分拨中心；发往 四川成都国际广场公司', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1522, '4309032908240', '2020-10-08 06:47:31', '【成都市】已到达 四川成都国际广场公司', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1523, '4309032908240', '2020-10-08 07:41:14', '【成都市】四川成都国际广场公司 快递员 何泽润华18200224476 正在为您派件【95114/95121/9501395546为韵达快递员外呼专属号码，请放心接听】', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1524, '4309032935655', '2020-10-07 16:57:19', '【广元市】四川苍溪县公司 已揽收', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1525, '4309032935655', '2020-10-08 01:44:07', '【成都市】已到达 四川成都分拨中心', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1526, '4309032935655', '2020-10-08 02:23:44', '【成都市】已离开 四川成都分拨中心；发往 四川成都国际广场公司', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1527, '4309032935655', '2020-10-08 06:35:03', '【成都市】已到达 四川成都国际广场公司', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1528, '4309032935655', '2020-10-08 07:46:39', '【成都市】四川成都国际广场公司 快递员 何泽润华18200224476 正在为您派件【95114/95121/9501395546为韵达快递员外呼专属号码，请放心接听】', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1529, '4309032943340', '2020-10-07 16:57:24', '【广元市】四川苍溪县公司 已揽收', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1530, '4309032943340', '2020-10-08 01:44:05', '【成都市】已到达 四川成都分拨中心', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1531, '4309032943340', '2020-10-08 02:02:15', '【成都市】已离开 四川成都分拨中心；发往 四川成都国际广场公司', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1532, '4309032943340', '2020-10-08 06:34:51', '【成都市】已到达 四川成都国际广场公司', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1533, '4309032943340', '2020-10-08 07:38:43', '【成都市】四川成都国际广场公司 快递员 何泽润华18200224476 正在为您派件【95114/95121/9501395546为韵达快递员外呼专属号码，请放心接听】', '', '2020-10-08 09:12:17', '', '2020-10-08 09:12:17');
INSERT INTO `logistics_tbl` VALUES (1534, '4309032960869', '2020-10-07 16:57:41', '【广元市】四川苍溪县公司 已揽收', '', '2020-10-08 09:12:43', '', '2020-10-08 09:12:43');
INSERT INTO `logistics_tbl` VALUES (1535, '4309032960869', '2020-10-08 00:48:06', '【成都市】已到达 四川成都分拨中心', '', '2020-10-08 09:12:43', '', '2020-10-08 09:12:43');
INSERT INTO `logistics_tbl` VALUES (1536, '4309032960869', '2020-10-08 00:50:24', '【成都市】已离开 四川成都分拨中心；发往 上海分拨中心', '', '2020-10-08 09:12:43', '', '2020-10-08 09:12:43');
INSERT INTO `logistics_tbl` VALUES (1537, '4309033023721', '2020-10-07 16:57:30', '【广元市】四川苍溪县公司 已揽收', '', '2020-10-08 09:12:50', '', '2020-10-08 09:12:50');
INSERT INTO `logistics_tbl` VALUES (1538, '4309033023721', '2020-10-08 00:48:27', '【成都市】已到达 四川成都分拨中心', '', '2020-10-08 09:12:50', '', '2020-10-08 09:12:50');
INSERT INTO `logistics_tbl` VALUES (1539, '4309033023721', '2020-10-08 00:50:44', '【成都市】已离开 四川成都分拨中心；发往 上海分拨中心', '', '2020-10-08 09:12:50', '', '2020-10-08 09:12:50');
INSERT INTO `logistics_tbl` VALUES (1540, 'JDVD01028720803', '2020-10-08 16:26:48', '添加快递单', 'admin', '2020-10-08 16:26:48', '', '2020-10-08 16:26:48');
INSERT INTO `logistics_tbl` VALUES (1541, 'JDVD01028714548', '2020-10-08 16:27:24', '添加快递单', 'admin', '2020-10-08 16:27:24', '', '2020-10-08 16:27:24');
INSERT INTO `logistics_tbl` VALUES (1542, 'JDVD01028714548', '2020-10-08 15:43:50', '揽收任务已分配给王沛', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1543, 'JDVD01028714548', '2020-10-08 15:47:11', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1544, 'JDVD01028714548', '2020-10-08 15:47:11', '京东快递 已收取快件', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1545, 'JDVD01028714548', '2020-10-08 15:47:13', '您的快件已到达【广元苍溪营业部】', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1546, 'JDVD01028720803', '2020-10-08 15:43:50', '揽收任务已分配给王沛', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1547, 'JDVD01028720803', '2020-10-08 15:47:09', '京东快递 已收取快件', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1548, 'JDVD01028720803', '2020-10-08 15:47:09', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1549, 'JDVD01028720803', '2020-10-08 15:47:11', '您的快件已到达【广元苍溪营业部】', '', '2020-10-08 16:27:44', '', '2020-10-08 16:27:44');
INSERT INTO `logistics_tbl` VALUES (1550, 'JDVD01028714548', '2020-10-08 17:21:23', '快递司机收箱', '', '2020-10-08 20:18:25', '', '2020-10-08 20:18:25');
INSERT INTO `logistics_tbl` VALUES (1551, 'JDVD01028714548', '2020-10-08 17:21:28', '您的快件已发车', '', '2020-10-08 20:18:25', '', '2020-10-08 20:18:25');
INSERT INTO `logistics_tbl` VALUES (1552, 'JDVD01028720803', '2020-10-08 17:21:23', '快递司机收箱', '', '2020-10-08 20:18:26', '', '2020-10-08 20:18:26');
INSERT INTO `logistics_tbl` VALUES (1553, 'JDVD01028720803', '2020-10-08 17:21:28', '您的快件已发车', '', '2020-10-08 20:18:26', '', '2020-10-08 20:18:26');
INSERT INTO `logistics_tbl` VALUES (1554, 'JDVD01028714548', '2020-10-08 20:53:28', '您的快件已到达【南充分拣场】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1555, 'JDVD01028714548', '2020-10-08 20:53:51', '您的快件在【南充分拣场】分拣完成', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1556, 'JDVD01028714548', '2020-10-08 20:53:56', '您的快件由【南充分拣场】准备发往【成都青白江分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1557, 'JDVD01028714548', '2020-10-08 21:57:48', '您的快件已发车', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1558, 'JDVD01028714548', '2020-10-09 01:52:22', '您的快件已到达【成都青白江分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1559, 'JDVD01028714548', '2020-10-09 01:52:22', '成都青白江分拣中心分拣中心已收箱，箱号JDVD01028714548-1-1-', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1560, 'JDVD01028714548', '2020-10-09 01:56:36', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1561, 'JDVD01028714548', '2020-10-09 01:56:41', '您的快件由【成都青白江分拣中心】准备发往【成都祥福分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1562, 'JDVD01028714548', '2020-10-09 02:08:12', '您的快件已发车', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1563, 'JDVD01028714548', '2020-10-09 02:30:08', '您的快件已到达【成都祥福分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1564, 'JDVD01028714548', '2020-10-09 04:14:26', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1565, 'JDVD01028714548', '2020-10-09 04:14:31', '您的快件由【成都祥福分拣中心】准备发往【廊坊固安分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1566, 'JDVD01028714548', '2020-10-09 05:55:59', '您的快件已发车', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1567, 'JDVD01028720803', '2020-10-08 21:01:20', '您的快件已到达【南充分拣场】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1568, 'JDVD01028720803', '2020-10-08 21:01:22', '您的快件在【南充分拣场】分拣完成', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1569, 'JDVD01028720803', '2020-10-08 21:01:27', '您的快件由【南充分拣场】准备发往【成都青白江分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1570, 'JDVD01028720803', '2020-10-08 21:57:48', '您的快件已发车', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1571, 'JDVD01028720803', '2020-10-09 01:55:51', '您的快件已到达【成都青白江分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1572, 'JDVD01028720803', '2020-10-09 01:55:51', '成都青白江分拣中心分拣中心已收箱，箱号JDVD01028720803-1-1-', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1573, 'JDVD01028720803', '2020-10-09 02:06:51', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1574, 'JDVD01028720803', '2020-10-09 03:40:17', '您的快件由【成都青白江分拣中心】准备发往【成都祥福分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1575, 'JDVD01028720803', '2020-10-09 03:51:38', '您的快件已发车', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1576, 'JDVD01028720803', '2020-10-09 04:25:41', '您的快件已到达【成都祥福分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1577, 'JDVD01028720803', '2020-10-09 04:29:27', '您的快件在【成都祥福分拣中心】分拣完成', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1578, 'JDVD01028720803', '2020-10-09 04:29:32', '您的快件由【成都祥福分拣中心】准备发往【廊坊固安分拣中心】', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1579, 'JDVD01028720803', '2020-10-09 05:55:59', '您的快件已发车', '', '2020-10-09 10:32:56', '', '2020-10-09 10:32:56');
INSERT INTO `logistics_tbl` VALUES (1580, 'JDVD01028714548', '2020-10-10 07:08:49', '您的快件已到达【廊坊固安分拣中心】', '', '2020-10-10 10:46:45', '', '2020-10-10 10:46:45');
INSERT INTO `logistics_tbl` VALUES (1581, 'JDVD01028714548', '2020-10-10 07:17:00', '您的快件在【廊坊固安分拣中心】分拣完成', '', '2020-10-10 10:46:45', '', '2020-10-10 10:46:45');
INSERT INTO `logistics_tbl` VALUES (1582, 'JDVD01028720803', '2020-10-10 07:06:19', '您的快件已到达【廊坊固安分拣中心】', '', '2020-10-10 10:46:45', '', '2020-10-10 10:46:45');
INSERT INTO `logistics_tbl` VALUES (1583, 'JDVD01028720803', '2020-10-10 07:09:53', '您的快件在【廊坊固安分拣中心】分拣完成', '', '2020-10-10 10:46:45', '', '2020-10-10 10:46:45');
INSERT INTO `logistics_tbl` VALUES (1584, '4309032908240', '2020-10-08 17:40:51', '【代收点】您的快件已送达 成都中海兰庭北一门店 保管，地址：快件已暂存至成都中海兰庭北一门店菜鸟驿站如有疑问请联系19183972983，请及时领取，如有疑问请电联快递员：何泽润华【18200224476】。', '', '2020-10-10 12:27:45', '', '2020-10-10 12:27:45');
INSERT INTO `logistics_tbl` VALUES (1585, '4309032935655', '2020-10-08 17:40:16', '【代收点】您的快件已送达 成都中海兰庭北一门店 保管，地址：快件已暂存至成都中海兰庭北一门店菜鸟驿站如有疑问请联系19183972983，请及时领取，如有疑问请电联快递员：何泽润华【18200224476】。', '', '2020-10-10 12:27:45', '', '2020-10-10 12:27:45');
INSERT INTO `logistics_tbl` VALUES (1586, '4309032943340', '2020-10-08 17:40:17', '【代收点】您的快件已送达 成都中海兰庭北一门店 保管，地址：快件已暂存至成都中海兰庭北一门店菜鸟驿站如有疑问请联系19183972983，请及时领取，如有疑问请电联快递员：何泽润华【18200224476】。', '', '2020-10-10 12:27:45', '', '2020-10-10 12:27:45');
INSERT INTO `logistics_tbl` VALUES (1587, '4309032960869', '2020-10-09 08:52:13', '【上海市】已到达 上海分拨中心', '', '2020-10-10 12:27:51', '', '2020-10-10 12:27:51');
INSERT INTO `logistics_tbl` VALUES (1588, '4309032960869', '2020-10-09 08:57:02', '【上海市】已离开 上海分拨中心；发往 上海杨浦区五角场公司', '', '2020-10-10 12:27:51', '', '2020-10-10 12:27:51');
INSERT INTO `logistics_tbl` VALUES (1589, '4309032960869', '2020-10-09 14:10:01', '【上海市】已到达 上海杨浦区五角场公司', '', '2020-10-10 12:27:51', '', '2020-10-10 12:27:51');
INSERT INTO `logistics_tbl` VALUES (1590, '4309032960869', '2020-10-09 18:15:53', '【上海市】上海杨浦区五角场公司 快递员 叶献中18221791695 正在为您派件【95114/95121/9501395546为韵达快递员外呼专属号码，请放心接听】', '', '2020-10-10 12:27:51', '', '2020-10-10 12:27:51');
INSERT INTO `logistics_tbl` VALUES (1591, '4309032960869', '2020-10-09 20:12:05', '【上海市】您的快件已签收,签收人：同事，如有疑问请电联快递员：叶献中【18221791695】。', '', '2020-10-10 12:27:51', '', '2020-10-10 12:27:51');
INSERT INTO `logistics_tbl` VALUES (1592, '4309033023721', '2020-10-09 08:52:11', '【上海市】已到达 上海分拨中心', '', '2020-10-10 12:27:57', '', '2020-10-10 12:27:57');
INSERT INTO `logistics_tbl` VALUES (1593, '4309033023721', '2020-10-09 09:00:31', '【上海市】已离开 上海分拨中心；发往 上海宝山区大场公司', '', '2020-10-10 12:27:57', '', '2020-10-10 12:27:57');
INSERT INTO `logistics_tbl` VALUES (1594, '4309033023721', '2020-10-09 16:25:05', '【上海市】上海宝山区大场公司云琪服务部 快递员 杨忠18721919856 正在为您派件【95114/95121/9501395546为韵达快递员外呼专属号码，请放心接听】', '', '2020-10-10 12:27:57', '', '2020-10-10 12:27:57');
INSERT INTO `logistics_tbl` VALUES (1595, '4309033023721', '2020-10-09 19:06:11', '【上海市】您的快件已被 工业园门卫岗亭保安室 代签收，地址：城市工业园区，如有疑问请电联快递员：杨忠【18721919856】。', '', '2020-10-10 12:27:57', '', '2020-10-10 12:27:57');
INSERT INTO `logistics_tbl` VALUES (1596, 'JDVD01028714548', '2020-10-10 11:07:40', '您的快件由【廊坊固安分拣中心】准备发往【北京大兴分拣中心】', '', '2020-10-10 13:53:52', '', '2020-10-10 13:53:52');
INSERT INTO `logistics_tbl` VALUES (1597, 'JDVD01028714548', '2020-10-10 12:11:54', '您的快件已发车', '', '2020-10-10 13:53:52', '', '2020-10-10 13:53:52');
INSERT INTO `logistics_tbl` VALUES (1598, 'JDVD01028714548', '2020-10-10 13:18:03', '您的快件已到达【北京大兴分拣中心】', '', '2020-10-10 13:53:52', '', '2020-10-10 13:53:52');
INSERT INTO `logistics_tbl` VALUES (1599, 'JDVD01028714548', '2020-10-10 13:41:21', '您的快件在【北京大兴分拣中心】分拣完成', '', '2020-10-10 13:53:52', '', '2020-10-10 13:53:52');
INSERT INTO `logistics_tbl` VALUES (1600, 'JDVD01028720803', '2020-10-10 11:07:40', '您的快件由【廊坊固安分拣中心】准备发往【北京大兴分拣中心】', '', '2020-10-10 13:53:53', '', '2020-10-10 13:53:53');
INSERT INTO `logistics_tbl` VALUES (1601, 'JDVD01028720803', '2020-10-10 12:11:54', '您的快件已发车', '', '2020-10-10 13:53:53', '', '2020-10-10 13:53:53');
INSERT INTO `logistics_tbl` VALUES (1602, 'JDVD01028720803', '2020-10-10 13:19:17', '您的快件已到达【北京大兴分拣中心】', '', '2020-10-10 13:53:53', '', '2020-10-10 13:53:53');
INSERT INTO `logistics_tbl` VALUES (1603, 'JDVD01028720803', '2020-10-10 13:19:58', '您的快件在【北京大兴分拣中心】分拣完成', '', '2020-10-10 13:53:53', '', '2020-10-10 13:53:53');
INSERT INTO `logistics_tbl` VALUES (1604, 'JDVD01028714548', '2020-10-10 14:52:57', '您的快件由【北京大兴分拣中心】准备发往【北京胡家垡营业部】', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1605, 'JDVD01028714548', '2020-10-10 14:59:27', '您的快件已发车', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1606, 'JDVD01028714548', '2020-10-10 15:59:25', '您的快件已到达【北京胡家垡营业部】', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1607, 'JDVD01028714548', '2020-10-10 16:17:28', '您的快件正在派送中，请您准备签收（快递员：向祖岭，联系电话：18601329580）', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1608, 'JDVD01028714548', '2020-10-10 16:59:48', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1609, 'JDVD01028720803', '2020-10-10 14:52:57', '您的快件由【北京大兴分拣中心】准备发往【北京胡家垡营业部】', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1610, 'JDVD01028720803', '2020-10-10 14:59:27', '您的快件已发车', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1611, 'JDVD01028720803', '2020-10-10 16:01:28', '您的快件已到达【北京胡家垡营业部】', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1612, 'JDVD01028720803', '2020-10-10 16:15:43', '您的快件正在派送中，请您准备签收（快递员：向祖岭，联系电话：18601329580）', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1613, 'JDVD01028720803', '2020-10-10 16:59:48', '您的快件已由本人签收，感谢您使用京东物流，期待再次为您服务', '', '2020-10-10 19:19:06', '', '2020-10-10 19:19:06');
INSERT INTO `logistics_tbl` VALUES (1614, '4309032908240', '2020-10-09 20:10:41', '【代收点】您的快件已签收,签收人在：成都中海兰庭北一门店(已签收签收人凭取货码签收) 领取', '', '2020-10-12 18:57:14', '', '2020-10-12 18:57:14');
INSERT INTO `logistics_tbl` VALUES (1615, '4309032943340', '2020-10-09 20:10:38', '【代收点】您的快件已签收,签收人在：成都中海兰庭北一门店(已签收签收人凭取货码签收) 领取', '', '2020-10-12 18:57:14', '', '2020-10-12 18:57:14');
INSERT INTO `logistics_tbl` VALUES (1616, '9880218737813', '2020-10-03 05:51:20', '[成都市]离开【成都邮件处理中心】,下一站【公平营业部】', '', '2020-10-13 00:00:36', '', '2020-10-13 00:00:36');
INSERT INTO `logistics_tbl` VALUES (1617, '9880218737813', '2020-10-03 07:30:02', '[成都市]到达【公平营业部】', '', '2020-10-13 00:00:36', '', '2020-10-13 00:00:36');
INSERT INTO `logistics_tbl` VALUES (1618, '9880218737813', '2020-10-03 08:59:21', '[成都市]【公平营业部】安排投递,投递员:卢永志,电话:19827574748,揽投部电话:028-83284405', '', '2020-10-13 00:00:36', '', '2020-10-13 00:00:36');
INSERT INTO `logistics_tbl` VALUES (1619, '9880218737813', '2020-10-03 09:43:38', '[成都市]已签收,他人代收：菜鸟驿站。,投递员:卢永志,电话:19827574748', '', '2020-10-13 00:00:36', '', '2020-10-13 00:00:36');
INSERT INTO `logistics_tbl` VALUES (1620, '9880336228325', '2020-10-14 10:06:39', '添加快递单', 'admin', '2020-10-14 10:06:39', '', '2020-10-14 10:06:39');
INSERT INTO `logistics_tbl` VALUES (1621, '9880336237785', '2020-10-14 10:06:48', '添加快递单', 'admin', '2020-10-14 10:06:48', '', '2020-10-14 10:06:48');
INSERT INTO `logistics_tbl` VALUES (1622, 'JDVD01047764919', '2020-10-14 15:03:51', '添加快递单', 'admin', '2020-10-14 15:03:51', '', '2020-10-14 15:03:51');
INSERT INTO `logistics_tbl` VALUES (1623, '9880336228325', '2020-10-13 18:39:32', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:储倩,电话:13330979458', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1624, '9880336228325', '2020-10-13 21:55:16', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1625, '9880336228325', '2020-10-13 22:07:46', '[凉山彝族自治州]离开【四川会理县中心】,下一站【凉山邮件处理中心】', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1626, '9880336228325', '2020-10-14 08:21:04', '[凉山彝族自治州]到达【凉山邮件处理中心】', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1627, '9880336237785', '2020-10-13 18:39:22', '[凉山彝族自治州]【会理大宗收寄处理组】已收寄,揽投员:储倩,电话:13330979458', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1628, '9880336237785', '2020-10-13 21:55:16', '[凉山彝族自治州]离开【会理大宗收寄处理组】,下一站【四川会理县中心】', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1629, '9880336237785', '2020-10-13 22:07:46', '[凉山彝族自治州]离开【四川会理县中心】,下一站【凉山邮件处理中心】', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1630, '9880336237785', '2020-10-14 08:21:04', '[凉山彝族自治州]到达【凉山邮件处理中心】', '', '2020-10-14 20:51:44', '', '2020-10-14 20:51:44');
INSERT INTO `logistics_tbl` VALUES (1631, '9880336228325', '2020-10-14 22:06:45', '[凉山彝族自治州]离开【凉山邮件处理中心】,下一站【成都航站】', '', '2020-10-15 09:11:16', '', '2020-10-15 09:11:16');
INSERT INTO `logistics_tbl` VALUES (1632, '9880336228325', '2020-10-15 05:29:01', '[成都市]到达【成都航站】（经转）', '', '2020-10-15 09:11:16', '', '2020-10-15 09:11:16');
INSERT INTO `logistics_tbl` VALUES (1633, '9880336237785', '2020-10-14 22:06:45', '[凉山彝族自治州]离开【凉山邮件处理中心】,下一站【成都航站】', '', '2020-10-15 09:11:16', '', '2020-10-15 09:11:16');
INSERT INTO `logistics_tbl` VALUES (1634, '9880336237785', '2020-10-15 05:29:01', '[成都市]到达【成都航站】（经转）', '', '2020-10-15 09:11:16', '', '2020-10-15 09:11:16');
INSERT INTO `logistics_tbl` VALUES (1635, 'JDVD01047764919', '2020-10-14 15:01:04', '揽收任务已分配给王沛', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1636, 'JDVD01047764919', '2020-10-14 17:13:37', '京东快递 已收取快件', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1637, 'JDVD01047764919', '2020-10-14 17:13:37', '您的快件已由【广元苍溪营业部】揽收完成', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1638, 'JDVD01047764919', '2020-10-14 17:13:38', '您的快件已到达【广元苍溪营业部】', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1639, 'JDVD01047764919', '2020-10-14 17:27:32', '快递司机收箱', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1640, 'JDVD01047764919', '2020-10-14 17:27:36', '您的快件已发车', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1641, 'JDVD01047764919', '2020-10-14 19:53:47', '您的快件已到达【南充分拣中心】', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1642, 'JDVD01047764919', '2020-10-14 19:53:58', '您的快件在【南充分拣中心】分拣完成', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1643, 'JDVD01047764919', '2020-10-14 19:54:03', '您的快件由【南充分拣中心】准备发往【成都青白江分拣中心】', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1644, 'JDVD01047764919', '2020-10-14 21:54:34', '您的快件已发车', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1645, 'JDVD01047764919', '2020-10-15 01:40:29', '您的快件已到达【成都青白江分拣中心】', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1646, 'JDVD01047764919', '2020-10-15 01:40:29', '成都青白江分拣中心分拣中心已收箱，箱号JDVD01047764919-1-1-', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1647, 'JDVD01047764919', '2020-10-15 01:46:04', '您的快件在【成都青白江分拣中心】分拣完成', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1648, 'JDVD01047764919', '2020-10-15 01:46:09', '您的快件由【成都青白江分拣中心】准备发往【昆明分拣中心】', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1649, 'JDVD01047764919', '2020-10-15 03:52:00', '您的快件已发车', '', '2020-10-15 17:58:55', '', '2020-10-15 17:58:55');
INSERT INTO `logistics_tbl` VALUES (1650, 'JDVD01047764919', '2020-10-15 21:26:17', '您的快件已到达【昆明分拣中心】', '', '2020-10-16 00:34:56', '', '2020-10-16 00:34:56');
INSERT INTO `logistics_tbl` VALUES (1651, 'JDVD01047764919', '2020-10-15 21:38:56', '您的快件在【昆明分拣中心】分拣完成', '', '2020-10-16 00:34:56', '', '2020-10-16 00:34:56');
INSERT INTO `logistics_tbl` VALUES (1652, 'JDVD01047764919', '2020-10-15 21:39:01', '您的快件由【昆明分拣中心】准备发往【昆明梁源营业部】', '', '2020-10-16 00:34:56', '', '2020-10-16 00:34:56');
INSERT INTO `logistics_tbl` VALUES (1653, 'JDVD01047764919', '2020-10-16 05:39:07', '您的快件已发车', '', '2020-10-16 22:41:55', '', '2020-10-16 22:41:55');
INSERT INTO `logistics_tbl` VALUES (1654, 'JDVD01047764919', '2020-10-16 07:48:10', '您的快件在【昆明梁源营业部】收货完成', '', '2020-10-16 22:41:55', '', '2020-10-16 22:41:55');
INSERT INTO `logistics_tbl` VALUES (1655, 'JDVD01047764919', '2020-10-16 07:48:11', '您的快件已到达【昆明梁源营业部】', '', '2020-10-16 22:41:55', '', '2020-10-16 22:41:55');
INSERT INTO `logistics_tbl` VALUES (1656, 'JDVD01047764919', '2020-10-16 08:18:25', '您的快件正在派送中，请您准备签收（快递员：张刚文，联系电话：19187555598）', '', '2020-10-16 22:41:55', '', '2020-10-16 22:41:55');
INSERT INTO `logistics_tbl` VALUES (1657, 'JDVD01047764919', '2020-10-16 14:05:35', '您的快件已由客户指定地点代收，感谢您使用京东物流，期待再次为您服务', '', '2020-10-16 22:41:55', '', '2020-10-16 22:41:55');
INSERT INTO `logistics_tbl` VALUES (1658, '9880336228325', '2020-10-15 21:30:04', '[成都市]离开【成都航站】,下一站【北京综合邮件处理中心】（经转）', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1659, '9880336228325', '2020-10-17 16:04:46', '[北京市]到达【北京综合邮件处理中心】', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1660, '9880336228325', '2020-10-18 05:21:05', '[北京市]离开【北京重件进口作业区】,下一站【塔院营业部】', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1661, '9880336228325', '2020-10-18 07:26:11', '[北京市]到达【塔院营业部】', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1662, '9880336228325', '2020-10-18 09:40:26', '[北京市]【塔院营业部】安排投递,投递员:薛国鹏,电话:18519361582,揽投部电话:010-62017030', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1663, '9880336228325', '2020-10-18 09:42:33', '[北京市]已签收,他人代收：门口,投递员:薛国鹏,电话:18519361582', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1664, '9880336237785', '2020-10-15 21:30:04', '[成都市]离开【成都航站】,下一站【北京综合邮件处理中心】（经转）', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1665, '9880336237785', '2020-10-17 16:04:46', '[北京市]到达【北京综合邮件处理中心】', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1666, '9880336237785', '2020-10-18 05:21:05', '[北京市]离开【北京重件进口作业区】,下一站【塔院营业部】', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1667, '9880336237785', '2020-10-18 07:26:11', '[北京市]到达【塔院营业部】', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1668, '9880336237785', '2020-10-18 08:06:23', '[北京市]【塔院营业部】安排投递,投递员:薛国鹏,电话:18519361582,揽投部电话:010-62017030', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1669, '9880336237785', '2020-10-18 09:42:33', '[北京市]已签收,他人代收：门口,投递员:薛国鹏,电话:18519361582', '', '2020-10-18 16:44:11', '', '2020-10-18 16:44:11');
INSERT INTO `logistics_tbl` VALUES (1670, 'YT4806894253200', '2020-09-30 16:44:59', '客户签收人: 张 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18521850993，投诉电话：010-53807290。疫情期间圆通每天对网点多次消毒，快递小哥每天测量体温，佩戴口罩', '', '2020-10-22 20:00:02', '', '2020-10-22 20:00:02');
INSERT INTO `logistics_tbl` VALUES (1671, 'YT4811974436620', '2020-09-29 03:00:07', '【自贡转运中心公司】 已收入', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1672, 'YT4811974436620', '2020-09-29 03:03:47', '【自贡转运中心】 已发出 下一站 【四川省泸州市公司】', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1673, 'YT4811974436620', '2020-09-29 13:27:32', '【四川省泸州市江阳区万象汇公司】 派件中  派件人: 杨先兵 电话 18181865511 ', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1674, 'YT4811974436620', '2020-09-29 21:14:54', '客户签收人: 本人签收 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18181865511', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1675, '9880201780019', '2020-09-29 14:13:31', '[成都市]到达【成都双流区华阳营业部】', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1676, '9880201780019', '2020-09-29 15:34:44', '[成都市]【成都双流区华阳营业部】安排投递,投递员:邱凯,电话:18180484929,揽投部电话:028-85769395', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1677, '9880201780019', '2020-09-29 16:42:31', '[成都市]邮件投递到成都宏达世纪锦城店菜鸟驿站,投递员:邱凯,电话:18180484929', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1678, '9880201780019', '2020-09-30 12:32:25', '[成都市]收件人已取走邮件', '', '2020-10-22 20:00:04', '', '2020-10-22 20:00:04');
INSERT INTO `logistics_tbl` VALUES (1679, 'YT4823368045243', '2020-10-03 10:41:48', '【北京市海淀区北师大北航分部公司】 派件中  派件人: 郭世君 电话 18513509329 ', '', '2020-10-22 20:00:05', '', '2020-10-22 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1680, 'YT4823368045243', '2020-10-03 16:49:39', '客户签收人: 张 已签收  感谢使用圆通速递，期待再次为您服务 如有疑问请联系：18513509329', '', '2020-10-22 20:00:05', '', '2020-10-22 20:00:05');
INSERT INTO `logistics_tbl` VALUES (1681, '9880209246003', '2020-10-02 22:48:08', '[北京市]到达【北京综合邮件处理中心】', '', '2020-10-22 20:00:06', '', '2020-10-22 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1682, '9880209246003', '2020-10-03 05:19:17', '[北京市]离开【北京重件进口作业区】,下一站【塔院营业部】', '', '2020-10-22 20:00:06', '', '2020-10-22 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1683, '9880209246003', '2020-10-03 08:02:31', '[北京市]到达【塔院营业部】', '', '2020-10-22 20:00:06', '', '2020-10-22 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1684, '9880209246003', '2020-10-03 08:24:01', '[北京市]【塔院营业部】安排投递,投递员:薛国鹏,电话:18519361582,揽投部电话:010-62017030', '', '2020-10-22 20:00:06', '', '2020-10-22 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1685, '9880209246003', '2020-10-03 09:51:02', '[北京市]已签收,他人代收：家,投递员:薛国鹏,电话:18519361582', '', '2020-10-22 20:00:06', '', '2020-10-22 20:00:06');
INSERT INTO `logistics_tbl` VALUES (1686, '776260237245573', '2020-11-26 14:37:25', '添加快递单', '13520697936', '2020-11-26 14:37:25', '', '2020-11-26 14:37:25');
INSERT INTO `logistics_tbl` VALUES (1687, '776260237245573', '2020-11-26 21:38:50', '【江西省赣州市场部三部】的收件员【崇仙脐橙(18679701992)】已收件', '', '2020-11-27 20:21:25', '', '2020-11-27 20:21:25');
INSERT INTO `logistics_tbl` VALUES (1688, '776260237245573', '2020-11-27 17:24:26', '快件由【江西省赣州市场部三部】发往【江西赣州转运中心】', '', '2020-11-27 20:21:25', '', '2020-11-27 20:21:25');
INSERT INTO `logistics_tbl` VALUES (1689, '776260237245573', '2020-11-27 20:19:40', '快件已到达【江西赣州转运中心】扫描员是【王焕新】', '', '2020-11-28 09:52:38', '', '2020-11-28 09:52:38');
INSERT INTO `logistics_tbl` VALUES (1690, '776260237245573', '2020-11-27 20:32:55', '快件由【江西赣州转运中心】发往【江西南昌转运中心】', '', '2020-11-28 09:52:38', '', '2020-11-28 09:52:38');
INSERT INTO `logistics_tbl` VALUES (1691, '776260237245573', '2020-11-28 21:57:48', '快件已到达【江西南昌转运中心】扫描员是【张润华】', '', '2020-11-29 11:42:42', '', '2020-11-29 11:42:42');
INSERT INTO `logistics_tbl` VALUES (1692, '776260237245573', '2020-11-28 22:08:39', '快件由【江西南昌转运中心】发往【北京转运中心】', '', '2020-11-29 11:42:42', '', '2020-11-29 11:42:42');
INSERT INTO `logistics_tbl` VALUES (1693, '776260237245573', '2020-11-29 19:56:34', '快件已到达【北京转运中心】扫描员是【王翠平】', '', '2020-11-30 15:04:23', '', '2020-11-30 15:04:23');
INSERT INTO `logistics_tbl` VALUES (1694, '776260237245573', '2020-11-29 20:08:18', '快件由【北京转运中心】发往【北京朝阳鸿博公司】', '', '2020-11-30 15:04:23', '', '2020-11-30 15:04:23');
INSERT INTO `logistics_tbl` VALUES (1695, '776260237245573', '2020-11-30 06:49:50', '快件已到达【北京朝阳鸿博公司】扫描员是【业李旭超】', '', '2020-11-30 15:04:23', '', '2020-11-30 15:04:23');
INSERT INTO `logistics_tbl` VALUES (1696, '776260237245573', '2020-11-30 07:49:18', '【北京朝阳鸿博公司】的派件员【业郝虎亮(17319390421)】正在为您派件，如有疑问请联系派件员，联系电话【17319390421】', '', '2020-11-30 15:04:23', '', '2020-11-30 15:04:23');
INSERT INTO `logistics_tbl` VALUES (1697, '776260237245573', '2020-11-30 08:54:47', '已签收(17319390421)，签收人是【家门口代收】,如有疑问请联系:17319390421,风里来，雨里去，汗也撒泪也流，申通小哥一刻不停留。不求服务惊天下，但求好评动我心，给个好评呗！！', '', '2020-11-30 15:04:23', '', '2020-11-30 15:04:23');

-- ----------------------------
-- Table structure for op_log_tbl
-- ----------------------------
DROP TABLE IF EXISTS `op_log_tbl`;
CREATE TABLE `op_log_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operator` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作人',
  `receive` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '接收方',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
  `topic` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作主题',
  `bundle` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '子主体',
  `pid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '细分主体',
  `ip_addr` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作IP地址',
  `num0` int(11) NULL DEFAULT NULL COMMENT '附加数字1',
  `num1` int(11) NULL DEFAULT NULL COMMENT '附加数字2',
  `num2` int(11) NULL DEFAULT NULL COMMENT '附加数字3',
  `attach0` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '附加字符串1',
  `attach1` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '附加字符串2',
  `attach2` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '附加字符串3',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 120 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of op_log_tbl
-- ----------------------------
INSERT INTO `op_log_tbl` VALUES (1, 'admin', '', '2020-09-02 23:38:40', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (2, 'admin', '', '2020-09-02 23:39:21', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (3, 'admin', '', '2020-09-03 23:07:09', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (4, 'admin', '', '2020-09-04 01:14:18', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (5, 'admin', '', '2020-09-04 01:19:27', '帐号系统', '登录', 'pc登录', '117.136.63.61', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (6, 'admin', '', '2020-09-04 01:19:42', '帐号系统', '登录', 'pc登录', '117.136.63.61', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (7, 'admin', '', '2020-09-07 00:56:18', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (8, 'admin', '', '2020-09-08 00:44:58', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (9, 'admin', '', '2020-09-08 00:45:12', '帐号系统', '登录', 'pc登录', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (10, 'admin', '', '2020-09-08 01:07:21', '产品系统', '修改', '会理县特级石榴', '116.237.23.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (11, 'admin', '', '2020-09-11 21:54:27', '帐号系统', '登录', 'pc登录', '118.113.18.12', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (12, '13795896867', '', '2020-09-12 01:11:05', '帐号系统', '登录', 'pc登录', '::1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (13, '13795896867', '', '2020-09-12 01:18:45', '帐号系统', '登录', 'pc登录', '::1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (14, '13795896867', '', '2020-09-12 01:19:05', '帐号系统', '登录', 'pc登录', '::1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (15, 'admin', '', '2020-09-12 01:19:28', '帐号系统', '登录', 'pc登录', '::1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (16, 'admin', '', '2020-09-12 01:20:40', '帐号系统', '登录', 'pc登录', '::1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (17, 'admin', '', '2020-09-12 15:48:00', '帐号系统', '登录', 'pc登录', '117.136.63.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (18, 'admin', '', '2020-09-14 09:59:58', '帐号系统', '登录', 'pc登录', '117.136.65.148', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (19, 'admin', '', '2020-09-15 16:36:45', '帐号系统', '登录', 'pc登录', '117.136.65.148', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (20, '13795896867', '', '2020-09-15 18:13:12', '帐号系统', '登录', 'pc登录', '58.247.131.189', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (21, 'admin', '', '2020-09-16 10:06:01', '帐号系统', '登录', 'pc登录', '223.104.212.96', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (22, 'admin', '', '2020-09-27 14:51:42', '帐号系统', '登录', 'pc登录', '223.104.9.220', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (23, '787940823', '', '2020-09-29 20:52:18', '帐号系统', '登录', 'pc登录', '223.104.212.231', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (24, 'admin', '', '2020-09-30 15:39:33', '帐号系统', '登录', 'pc登录', '58.33.96.221', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (25, 'admin', '', '2020-10-05 07:41:01', '帐号系统', '登录', 'pc登录', '223.104.215.144', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (26, 'admin', '', '2020-10-07 14:38:04', '帐号系统', '登录', 'pc登录', '117.176.130.61', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (27, 'admin', '', '2020-10-07 14:38:30', '帐号系统', '登录', 'pc登录', '117.176.130.61', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (28, 'admin', '', '2020-10-13 13:02:09', '帐号系统', '登录', 'pc登录', '223.104.24.109', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (29, 'admin', '', '2020-10-16 21:13:23', '帐号系统', '登录', 'pc登录', '117.136.0.189', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (30, 'admin', '', '2020-10-21 12:45:14', '帐号系统', '登录', 'pc登录', '58.247.131.187', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (31, 'admin', '', '2020-10-25 18:34:51', '帐号系统', '登录', 'pc登录', '112.65.48.212', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (32, 'admin', '', '2020-11-24 21:03:23', '帐号系统', '登录', 'pc登录', '124.79.51.121', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (33, 'admin', '', '2020-11-24 21:03:29', '产品系统', '修改', '信丰苹果脐橙', '124.79.51.121', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (34, '13520697936', '', '2020-11-25 19:21:13', '帐号系统', '登录', 'pc登录', '223.72.50.176', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (35, '13520697936', '', '2020-11-25 19:22:08', '帐号系统', '登录', 'pc登录', '223.72.50.176', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (36, 'admin', '', '2020-11-25 21:27:30', '帐号系统', '登录', 'pc登录', '124.79.51.121', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (37, 'admin', '', '2020-11-26 14:31:46', '帐号系统', '登录', 'pc登录', '223.104.218.188', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (38, 'admin', '', '2020-11-26 20:12:36', '帐号系统', '登录', 'pc登录', '211.161.200.251', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (39, 'admin', '', '2020-11-29 11:42:38', '帐号系统', '登录', 'pc登录', '124.79.51.121', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (40, 'admin', '', '2020-11-30 15:20:31', '帐号系统', '登录', 'pc登录', '218.89.136.197', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (41, 'admin', '', '2020-11-30 15:22:03', '帐号系统', '登录', 'pc登录', '218.89.136.197', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (42, '13520697936', '', '2020-12-06 12:31:33', '帐号系统', '登录', 'pc登录', '223.72.51.138', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (43, 'admin', '', '2020-12-09 21:22:21', '帐号系统', '登录', 'pc登录', '124.79.51.121', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (44, '13520697936', '', '2020-12-09 21:46:33', '帐号系统', '登录', 'pc登录', '223.72.51.49', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (45, '13520697936', '', '2020-12-09 21:47:08', '帐号系统', '登录', 'pc登录', '223.72.51.49', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (46, 'admin', '', '2020-12-10 16:47:50', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (47, 'admin', '', '2020-12-10 16:49:33', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (48, 'admin', '', '2020-12-10 16:58:52', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (49, 'admin', '', '2020-12-10 17:02:08', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (50, 'admin', '', '2020-12-10 17:02:59', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (51, 'admin', '', '2020-12-10 17:03:31', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (52, 'admin', '', '2020-12-10 17:04:00', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (53, 'admin', '', '2020-12-10 17:07:30', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (54, 'admin', '', '2020-12-10 17:10:27', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (55, 'admin', '', '2020-12-10 17:10:53', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (56, 'admin', '', '2020-12-10 17:13:26', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (57, 'admin', '', '2020-12-10 17:13:42', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (58, 'admin', '', '2020-12-10 17:15:16', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (59, 'admin', '', '2020-12-10 17:15:51', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (60, 'admin', '', '2020-12-10 17:16:10', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (61, 'admin', '', '2020-12-10 17:16:15', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (62, 'admin', '', '2020-12-10 17:17:23', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (63, 'admin', '', '2020-12-10 17:19:50', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (64, 'admin', '', '2020-12-10 17:20:19', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (65, 'admin', '', '2020-12-10 17:20:21', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (66, 'admin', '', '2020-12-10 17:21:03', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (67, 'admin', '', '2020-12-10 17:21:47', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (68, 'admin', '', '2020-12-10 17:22:51', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (69, 'admin', '', '2020-12-10 17:23:37', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (70, 'admin', '', '2020-12-10 17:28:03', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (71, 'admin', '', '2020-12-10 17:30:48', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (72, 'admin', '', '2020-12-10 17:31:55', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (73, 'admin', '', '2020-12-10 17:32:57', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (74, 'admin', '', '2020-12-10 17:35:21', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (75, 'admin', '', '2020-12-10 17:43:37', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (76, 'admin', '', '2020-12-10 17:45:30', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (77, 'admin', '', '2020-12-10 17:46:11', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (78, 'admin', '', '2020-12-10 17:47:26', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (79, 'admin', '', '2020-12-10 17:48:51', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (80, 'admin', '', '2020-12-10 20:55:27', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (81, 'admin', '', '2020-12-10 21:52:43', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (82, 'admin', '', '2020-12-29 17:19:04', '帐号系统', '登录', 'pc登录', '211.161.200.251', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (83, 'admin', '', '2020-12-29 17:35:37', '帐号系统', '登录', 'pc登录', '119.4.252.109', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (84, 'admin', '', '2020-12-29 17:36:12', '帐号系统', '登录', 'pc登录', '119.4.252.109', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (85, 'admin', '', '2020-12-30 11:43:34', '帐号系统', '登录', 'pc登录', '223.104.9.146', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (86, 'admin', '', '2020-12-30 13:38:19', '帐号系统', '登录', 'pc登录', '119.4.252.93', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (87, 'admin', '', '2020-12-30 13:55:51', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (88, 'admin', '', '2020-12-30 14:12:16', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (89, 'admin', '', '2020-12-30 14:15:55', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (90, 'admin', '', '2020-12-30 14:16:41', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (91, 'admin', '', '2020-12-30 14:19:50', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (92, 'admin', '', '2020-12-30 14:23:23', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (93, 'admin', '', '2020-12-30 14:24:22', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (94, 'admin', '', '2020-12-30 14:39:17', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (95, 'admin', '', '2020-12-30 14:52:28', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (96, 'admin', '', '2020-12-30 15:01:40', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (97, 'admin', '', '2020-12-30 15:05:16', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (98, 'admin', '', '2020-12-30 15:07:40', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (99, 'admin', '', '2020-12-30 15:09:50', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (100, 'admin', '', '2020-12-30 15:17:58', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (101, 'admin', '', '2020-12-30 17:42:38', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (102, 'admin', '', '2020-12-31 10:03:24', '帐号系统', '登录', 'pc登录', '101.206.166.176', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (103, 'admin', '', '2020-12-31 10:03:33', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (104, 'admin', '', '2020-12-31 10:56:36', '产品系统', '修改', '竹笋', '101.206.166.176', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (105, 'admin', '', '2021-01-13 10:45:35', '帐号系统', '登录', 'pc登录', '119.4.253.246', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (106, 'admin', '', '2021-01-13 10:45:52', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (107, 'admin', '', '2021-01-18 10:29:45', '帐号系统', '登录', 'pc登录', '101.206.168.32', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (108, 'admin', '', '2021-01-18 10:29:59', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (109, 'admin', '', '2021-01-21 15:51:12', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (110, 'admin', '', '2021-01-22 09:17:08', '帐号系统', '登录', 'pc登录', '101.206.166.124', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (111, 'admin', '', '2021-01-22 09:38:24', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (112, 'admin', '', '2021-01-22 18:55:05', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (113, 'admin', '', '2021-01-25 11:04:27', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (114, 'admin', '', '2021-01-27 10:25:33', '帐号系统', '登录', 'pc登录', '192.168.42.79', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (115, 'admin', '', '2021-02-20 09:54:22', '帐号系统', '登录', 'pc登录', '127.0.0.1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (116, 'admin', '', '2021-03-06 17:09:11', '帐号系统', '登录', 'pc登录', '211.161.200.251', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (117, 'admin', '', '2021-03-06 17:09:27', '帐号系统', '登录', 'pc登录', '::1', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (118, 'admin', '', '2021-06-22 17:43:38', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');
INSERT INTO `op_log_tbl` VALUES (119, 'admin', '', '2021-06-29 09:31:12', '帐号系统', '登录', 'pc登录', '192.168.43.240', 0, 0, 0, '', '', '');

-- ----------------------------
-- Table structure for post_tbl
-- ----------------------------
DROP TABLE IF EXISTS `post_tbl`;
CREATE TABLE `post_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递名称',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递id',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递logo地址',
  `exp_phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递电话',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '快递公司信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of post_tbl
-- ----------------------------
INSERT INTO `post_tbl` VALUES (1, '顺丰速运', 'SF', 'https://cdn.kuaidi100.com/images/all/56/shunfeng.png', '95338');
INSERT INTO `post_tbl` VALUES (2, '中通快递', 'ZTO', 'https://cdn.kuaidi100.com/images/all/56/zhongtong.png', '95311');
INSERT INTO `post_tbl` VALUES (3, '申通快递', 'STO', 'https://cdn.kuaidi100.com/images/all/56/shentong.png', '95543');
INSERT INTO `post_tbl` VALUES (4, '圆通速递', 'YTO', 'https://cdn.kuaidi100.com/images/all/56/yuantong.png', '95554');
INSERT INTO `post_tbl` VALUES (5, '韵达速递', 'YD', 'https://cdn.kuaidi100.com/images/all/56/yunda.png', '95546');
INSERT INTO `post_tbl` VALUES (6, '德邦快递', 'DBL', 'https://cdn.kuaidi100.com/images/all/56/debangkuaidi.png', '95353');
INSERT INTO `post_tbl` VALUES (7, '京东快递', 'JD', 'https://cdn.kuaidi100.com/images/all/56/jd.png', '950616');
INSERT INTO `post_tbl` VALUES (8, 'EMS', 'EMS', 'https://cdn.kuaidi100.com/images/all/56/ems.png', '11183');

-- ----------------------------
-- Table structure for product_ad_group_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_ad_group_tbl`;
CREATE TABLE `product_ad_group_tbl`  (
  `id` int(11) NOT NULL,
  `main_gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '主商品id\n',
  `sub_gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加商品id',
  `vaild` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否有效',
  `sort_id` int(11) NOT NULL COMMENT '排序(越大越前)',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `main_sub_gpid`(`main_gpid`, `sub_gpid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '首页精选列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_ad_group_tbl
-- ----------------------------
INSERT INTO `product_ad_group_tbl` VALUES (1, '1280145890979352576', '1303017018819088384', 0, 2, '', NULL, '', NULL);
INSERT INTO `product_ad_group_tbl` VALUES (2, '1280146196534398976', '1280145890979352576', 0, 2, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for product_ad_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_ad_tbl`;
CREATE TABLE `product_ad_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '跳转路径',
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '卡片展示图片',
  `sort_id` int(11) NOT NULL DEFAULT 0 COMMENT '排序(越大越前)',
  `type` int(11) NULL DEFAULT NULL COMMENT '类型(1:轮播图广告，2:类型广告，3:主销广告)',
  `attach` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `vaild` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否有效',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gpid`(`url`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '广告轮播图' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_ad_tbl
-- ----------------------------
INSERT INTO `product_ad_tbl` VALUES (4, '/pages/product/product?gpid=1280145890979352576', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/banner01.jpg', 4, 1, 'rgb(23,42,8)', 1, '', NULL, '', NULL, '', NULL);
INSERT INTO `product_ad_tbl` VALUES (5, '/pages/product/product?gpid=1280145890979352576', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/banner02.jpg', 4, 1, 'rgb(23,42,8)', 1, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `product_ad_tbl` VALUES (8, '/pages/product/list?type_id=2', '/static/temp/c3.png', 5, 2, '环球美食', 0, NULL, '2020-08-25 01:44:54', NULL, NULL, NULL, NULL);
INSERT INTO `product_ad_tbl` VALUES (9, '/pages/product/list?type_id=2', '/static/temp/c5.png', 4, 2, '个护美妆', 0, '', '2020-08-25 01:44:54', '', NULL, '', NULL);
INSERT INTO `product_ad_tbl` VALUES (10, '/pages/product/list?type_id=2', '/static/temp/c6.png', 3, 2, '营养保健', 0, '', '2020-08-25 01:44:54', '', NULL, '', NULL);
INSERT INTO `product_ad_tbl` VALUES (11, '/pages/product/list?type_id=2', '/static/temp/c7.png', 2, 2, '家居厨卫', 0, '', '2020-08-25 01:44:54', '', NULL, '', NULL);
INSERT INTO `product_ad_tbl` VALUES (12, '/pages/product/list?type_id=2', '/static/temp/c8.png', 1, 2, '速食生鲜', 0, '', '2020-08-25 01:44:54', '', NULL, '', NULL);
INSERT INTO `product_ad_tbl` VALUES (13, '/pages/coupon/getCoupon', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/coupon01.jpg', 1, 3, '', 1, '', '2020-08-25 01:44:54', '', NULL, '', NULL);

-- ----------------------------
-- Table structure for product_info_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_info_tbl`;
CREATE TABLE `product_info_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增',
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `qty` bigint(20) NOT NULL COMMENT '已购买数量(销量)',
  `img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '轮播图',
  `video` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频信息',
  `icon` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品图标',
  `service` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '商品服务',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `gpid`(`gpid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品附加属性' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_info_tbl
-- ----------------------------
INSERT INTO `product_info_tbl` VALUES (7, '1280145890979352576', 199, '[\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_banner01.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_banner02.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_banner03.jpg\"]', '', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_icon.jpg', '原产货源 · 假一赔十 ·', '', NULL, '', NULL);
INSERT INTO `product_info_tbl` VALUES (8, '1280146196534398976', 36, '[\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_banner01.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_banner02.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_banner03.jpg\"]', '', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_icon.jpg', '原产货源 · 假一赔十 ·', '', NULL, '', NULL);
INSERT INTO `product_info_tbl` VALUES (16, '1303017018819088384', 30, '[\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_banner01.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_banner02.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_banner03.jpg\",\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_banner04.jpg\"]', '', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_icon.jpg', '原产货源1· 假一赔十 ·', 'admin', '2020-09-08 01:07:21', 'admin', '2020-09-08 01:07:21');
INSERT INTO `product_info_tbl` VALUES (17, '1331221900822581248', 2, '[\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/jc/jc_banner01.jpg\"]', '', 'https://hospital.xxjwxc.cn/commcn/api/v1/file/image/jc/jc_icon.jpg', '12月20号前下单现摘直发.72小时内发货', 'admin', '2020-11-24 21:03:29', 'admin', '2020-11-24 21:03:29');
INSERT INTO `product_info_tbl` VALUES (18, '1344477525782302720', 0, '[\"http://localhost:8001/commcn/api/v1/file/image/1591532127P1H2.png\",\"http://localhost:8001/commcn/api/v1/file/image/1591532127P1H2.png\"]', '', 'http://localhost:8001/commcn/api/v1/file/image/1591532127P1H2.png', '', 'admin', '2020-12-31 10:56:36', 'admin', '2020-12-31 10:56:36');

-- ----------------------------
-- Table structure for product_sku_price_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_sku_price_tbl`;
CREATE TABLE `product_sku_price_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `sku_list` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '属性列表',
  `premium` bigint(20) NULL DEFAULT NULL COMMENT '商品单价',
  `dist_amount` bigint(20) NULL DEFAULT NULL COMMENT '分享收益',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `sku_list`(`sku_list`) USING BTREE,
  INDEX `gpid`(`gpid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品sku 属性溢价表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_sku_price_tbl
-- ----------------------------
INSERT INTO `product_sku_price_tbl` VALUES (1, '1280145890979352576', '8', 0, NULL, '', '2020-09-03 01:50:03', '', NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (2, '1280145890979352576', '9', 4300, NULL, NULL, '2020-09-03 01:50:03', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (3, '1280145890979352576', '10', 3300, NULL, '', '2020-09-03 01:52:00', '', NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (4, '1280145890979352576', '11', 12300, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (5, '1280146196534398976', '12', 0, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (6, '1280146196534398976', '13', 1000, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (7, '1303017018819088384', '14', 0, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (8, '1303017018819088384', '15', 990, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (9, '1303017018819088384', '16', 1500, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (10, '1303017018819088384', '17', 2990, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (11, '1303017018819088384', '18', 3990, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (12, '1303017018819088384', '19', 6200, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, '', NULL);
INSERT INTO `product_sku_price_tbl` VALUES (13, '1331221900822581248', '21', 6000, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, NULL, NULL);
INSERT INTO `product_sku_price_tbl` VALUES (14, '1331221900822581248', '20', 0, NULL, NULL, '2020-09-03 01:52:00', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for product_sku_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_sku_tbl`;
CREATE TABLE `product_sku_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `type_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签类型',
  `tag_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签名字',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '删除者',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `gpid_tag`(`gpid`, `tag_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品sku 属性' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_sku_tbl
-- ----------------------------
INSERT INTO `product_sku_tbl` VALUES (8, '1280145890979352576', '规格', '生态猕猴桃大果5斤装', '', NULL, '', NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (9, '1280145890979352576', '规格', '生态猕猴桃大果10斤装', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (10, '1280145890979352576', '规格', '精品礼盒装猕猴桃大果5斤装', '', NULL, '', NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (11, '1280145890979352576', '规格', '精品礼盒装猕猴桃大果10斤装', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (12, '1280146196534398976', '规格', '约8斤装', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (13, '1280146196534398976', '规格', '约10斤装', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (14, '1303017018819088384', '规格', '一级小乖乖突尼斯软籽石榴八斤装(净重6斤,10-12个)', '', NULL, '', NULL, '', '2020-09-22 11:20:25');
INSERT INTO `product_sku_tbl` VALUES (15, '1303017018819088384', '规格', '花皮果突尼斯软籽石榴十斤装(净重8斤,10-12个)', NULL, NULL, NULL, NULL, '', '2020-09-28 16:18:01');
INSERT INTO `product_sku_tbl` VALUES (16, '1303017018819088384', '规格', '一级突尼斯软籽石榴六枚装(净重5斤)', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (17, '1303017018819088384', '规格', '一级突尼斯软籽石榴十斤装(净重8斤,11-12个)', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (18, '1303017018819088384', '规格', '一级突尼斯软籽石榴十斤装(净重8斤,9-10个)', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (19, '1303017018819088384', '规格', '(土豪装)特级突尼斯软籽石榴十斤装(净重8斤,6-8个)', NULL, NULL, NULL, NULL, '', NULL);
INSERT INTO `product_sku_tbl` VALUES (20, '1331221900822581248', '规格', '10斤裝(部分地区包邮)', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `product_sku_tbl` VALUES (21, '1331221900822581248', '规格', '20斤裝(部分地区包邮)', NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for product_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_tbl`;
CREATE TABLE `product_tbl`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增',
  `gpid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品id',
  `gtid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品分类id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品名字',
  `price` bigint(20) NOT NULL COMMENT '商品价格(分)',
  `original_price` bigint(20) NOT NULL COMMENT '商品原始价格(分)',
  `dist_amount` bigint(20) NULL DEFAULT NULL COMMENT '分享收益',
  `platform_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商铺客户id',
  `qty` bigint(20) NOT NULL COMMENT '数量(库存)',
  `shipment_type_id` bigint(20) NOT NULL COMMENT '运输方式类型',
  `shipment_fee` bigint(20) NOT NULL COMMENT '运费',
  `context` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品详情',
  `volume_weight` bigint(20) NOT NULL DEFAULT 0 COMMENT '体积重',
  `actual_weight` bigint(20) NOT NULL DEFAULT 0 COMMENT '实际重',
  `source_region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '货源地名',
  `max_buy_qty` int(11) NOT NULL COMMENT '最大购买数量',
  `is_select` tinyint(1) NULL DEFAULT NULL COMMENT '是否推荐',
  `vaild` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否有效',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name_id`(`name`, `platform_id`) USING BTREE COMMENT '商品名索引',
  UNIQUE INDEX `gpid`(`gpid`) USING BTREE COMMENT '商品索引',
  CONSTRAINT `product_tbl_ibfk_1` FOREIGN KEY (`gpid`) REFERENCES `product_info_tbl` (`gpid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_tbl
-- ----------------------------
INSERT INTO `product_tbl` VALUES (3, '1280145890979352576', '水果', '苍溪红心猕猴桃', 7500, 24000, NULL, '13795896867', 1000, 0, 0, '[{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_ctx01.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_ctx02.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_ctx03.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_ctx04.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/mht/mht_ctx05.jpg\"},{\"type\":3,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/video/mht_01.mp4\"}]', 0, 0, '', 0, NULL, 1, '', NULL, '', NULL);
INSERT INTO `product_tbl` VALUES (4, '1280146196534398976', '水果', '原生态苍溪雪梨', 5800, 6800, NULL, '13795896867', 1000, 0, 0, '[{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_ctx01.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_ctx02.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_ctx03.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/xl/xl_ctx04.jpg\"}]', 0, 0, '', 0, NULL, 1, '', NULL, '', NULL);
INSERT INTO `product_tbl` VALUES (12, '1303017018819088384', '水果', '会理县突尼斯软籽石榴', 5000, 13800, NULL, '787940823', 1000, 0, 0, '[{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx01.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx02.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx03.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx04.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx05.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx06.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx07.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx08.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/sl/sl_ctx09.jpg\"}]', 0, 0, '', 0, NULL, 0, 'admin', '2020-09-08 01:07:21', 'admin', '2020-09-08 01:07:21');
INSERT INTO `product_tbl` VALUES (13, '1331221900822581248', '水果', '赣南脐橙', 7800, 18800, NULL, '13520697936', 1000, 0, 0, '[{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/jc/jc_ctx01.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/jc/jc_ctx02.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/jc/jc_ctx03.jpg\"},{\"type\":2,\"context\":\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/jc/jc_ctx04.jpg\"}]', 0, 0, '', 0, NULL, 1, 'admin', '2020-11-24 21:03:29', 'admin', '2020-11-24 21:03:29');

-- ----------------------------
-- Table structure for product_type_tbl
-- ----------------------------
DROP TABLE IF EXISTS `product_type_tbl`;
CREATE TABLE `product_type_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gtid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品类型id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `gtid`(`gtid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品类型' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product_type_tbl
-- ----------------------------
INSERT INTO `product_type_tbl` VALUES (2, '水果');
INSERT INTO `product_type_tbl` VALUES (3, '蔬菜');

-- ----------------------------
-- Table structure for proposal_tbl
-- ----------------------------
DROP TABLE IF EXISTS `proposal_tbl`;
CREATE TABLE `proposal_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '用户id',
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账单id',
  `type` int(11) NOT NULL DEFAULT 0 COMMENT '操作：-1：取消，1：申请售后，',
  `contact` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '联系方式',
  `remak` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `imgs` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图片',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `bill_id`(`bill_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '建议与反馈' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of proposal_tbl
-- ----------------------------
INSERT INTO `proposal_tbl` VALUES (1, 'oXebs4t0GxU93i319WkKEVxB5AyU', '', 3, '，，', '', '[\"https://hospital.xxjwxc.cn/commcn/api/v1/file/image/15981976002sKS.jpg\"]', '', '2020-08-23 23:46:40', '', '2020-08-23 23:46:40');
INSERT INTO `proposal_tbl` VALUES (2, 'oXebs4kdAodjFZvrDldisECoFfA4', '', 3, '开业大吉 祝福大麦～', '', '[]', '', '2020-09-04 14:07:26', '', '2020-09-04 14:07:26');
INSERT INTO `proposal_tbl` VALUES (3, 'oXebs4mlJGzhbgcchTbYuE1YVMjg', '', 3, 'CVVoMy8V', 'Wbh5nbI2', '[]', '', '2020-09-10 18:40:32', '', '2020-09-10 18:40:32');

-- ----------------------------
-- Table structure for shipment_tbl
-- ----------------------------
DROP TABLE IF EXISTS `shipment_tbl`;
CREATE TABLE `shipment_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bill_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账单id',
  `shipment_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递单号',
  `post_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '快递id',
  `status` int(11) NULL DEFAULT NULL COMMENT '快递状态  2在途中 3已签收 4 问题件',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `bill_shipment`(`bill_id`, `shipment_id`) USING BTREE,
  INDEX `post_key`(`post_key`) USING BTREE,
  CONSTRAINT `shipment_tbl_ibfk_1` FOREIGN KEY (`post_key`) REFERENCES `post_tbl` (`code`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 169 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户快递单信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of shipment_tbl
-- ----------------------------
INSERT INTO `shipment_tbl` VALUES (14, '20200905184201084404396720976982', 'YT4765645114700', 'YTO', 3, '', '2020-09-06 08:13:25', '', '2020-09-09 13:01:26');
INSERT INTO `shipment_tbl` VALUES (15, '20200905100426684334582162751618', 'YT4766119173574', 'YTO', 3, '', '2020-09-06 10:48:20', '', '2020-09-09 13:01:26');
INSERT INTO `shipment_tbl` VALUES (16, '20200905101006867950594888872266', 'JDVD00934676229', 'JD', 3, '', '2020-09-06 11:30:23', '', '2020-09-07 16:26:14');
INSERT INTO `shipment_tbl` VALUES (17, '20200905101006867950594888872266', 'JDVD00934692194', 'JD', 3, '', '2020-09-06 11:31:17', '', '2020-09-08 20:35:15');
INSERT INTO `shipment_tbl` VALUES (19, '20200905101031873249134724151780', 'JDVD00934679583', 'JD', 3, '', '2020-09-06 11:36:03', '', '2020-09-07 16:26:34');
INSERT INTO `shipment_tbl` VALUES (20, '20200905101031873249134724151780', 'JDVD00934685489', 'JD', 3, '', '2020-09-06 11:36:44', '', '2020-09-07 22:01:42');
INSERT INTO `shipment_tbl` VALUES (21, '20200905101110670537826772928058', 'JDVD00934659318', 'JD', 3, '', '2020-09-06 11:38:58', '', '2020-09-07 16:26:40');
INSERT INTO `shipment_tbl` VALUES (22, '20200905101110670537826772928058', 'JDVD00934683078', 'JD', 3, '', '2020-09-06 11:39:44', '', '2020-09-09 13:01:27');
INSERT INTO `shipment_tbl` VALUES (23, '20200904172854405140828894451431', 'JDVD00936221520', 'JD', 3, '', '2020-09-06 12:25:17', '', '2020-09-08 13:59:55');
INSERT INTO `shipment_tbl` VALUES (24, '20200904185814730172010744654212', 'JDVD00936221539', 'JD', 3, '', '2020-09-06 12:26:34', '', '2020-09-08 00:05:05');
INSERT INTO `shipment_tbl` VALUES (25, '20200905114645150932699740537643', 'JDVD00936215570', 'JD', 3, '', '2020-09-06 12:27:42', '', '2020-09-08 13:59:55');
INSERT INTO `shipment_tbl` VALUES (26, '20200904150227320207650506656575', 'JDVD00936201478', 'JD', 3, '', '2020-09-06 12:35:33', '', '2020-09-08 13:59:56');
INSERT INTO `shipment_tbl` VALUES (27, '20200904170709728039091360431790', 'JDVD00936221512', 'JD', 3, '', '2020-09-06 12:37:26', '', '2020-09-08 22:11:56');
INSERT INTO `shipment_tbl` VALUES (28, '20200904141800543481256192445779', 'JDVD00936215561', 'JD', 3, '', '2020-09-06 12:41:13', '', '2020-09-09 18:38:25');
INSERT INTO `shipment_tbl` VALUES (29, '20200904141914753054214958634643', 'JDVD00936213420', 'JD', 3, '', '2020-09-06 12:42:33', '', '2020-09-08 22:11:58');
INSERT INTO `shipment_tbl` VALUES (30, '20200906151849829902741555933210', 'YT4767034499495', 'YTO', 3, '', '2020-09-06 16:49:02', '', '2020-09-09 13:01:27');
INSERT INTO `shipment_tbl` VALUES (31, '20200906143007980723177113769080', 'YT4767033731540', 'YTO', 3, '', '2020-09-06 16:49:18', '', '2020-09-09 13:01:28');
INSERT INTO `shipment_tbl` VALUES (32, '20200906142851502720104214641259', 'YT4767033732744', 'YTO', 3, '', '2020-09-06 16:49:43', '', '2020-09-09 13:01:14');
INSERT INTO `shipment_tbl` VALUES (33, '20200906142917360342872133924722', 'YT4767030836799', 'YTO', 3, '', '2020-09-06 16:49:54', '', '2020-09-09 13:01:28');
INSERT INTO `shipment_tbl` VALUES (34, '20200904203011408428773933556708', '773054585661279', 'STO', 3, '', '2020-09-06 17:53:11', '', '2020-09-08 19:50:13');
INSERT INTO `shipment_tbl` VALUES (40, '20200904141816901816912762088270', '773054585245568', 'STO', 3, '', '2020-09-06 18:39:28', '', '2020-09-09 13:01:29');
INSERT INTO `shipment_tbl` VALUES (41, '20200905195839992093507738258324', '773054585660092', 'STO', 3, '', '2020-09-06 23:29:01', '', '2020-09-08 11:20:23');
INSERT INTO `shipment_tbl` VALUES (42, '20200905195839992093507738258324', '773054585660134', 'STO', 3, '', '2020-09-06 23:29:28', '', '2020-09-08 11:20:24');
INSERT INTO `shipment_tbl` VALUES (43, '20200905195839992093507738258324', '773054585733964', 'STO', 3, '', '2020-09-06 23:30:07', '', '2020-09-08 11:20:24');
INSERT INTO `shipment_tbl` VALUES (44, '20200905195839992093507738258324', '773054583991096', 'STO', 3, '', '2020-09-06 23:30:27', '', '2020-09-08 11:20:23');
INSERT INTO `shipment_tbl` VALUES (45, '20200905195520295672341079648272', '773054585812874', 'STO', 3, '', '2020-09-06 23:31:58', '', '2020-09-08 12:51:20');
INSERT INTO `shipment_tbl` VALUES (46, '20200905195520295672341079648272', '773054585660663', 'STO', 3, '', '2020-09-06 23:32:17', '', '2020-09-08 12:51:19');
INSERT INTO `shipment_tbl` VALUES (47, '20200905195520295672341079648272', '773054585734185', 'STO', 3, '', '2020-09-06 23:32:33', '', '2020-09-08 12:51:19');
INSERT INTO `shipment_tbl` VALUES (48, '20200905195520295672341079648272', '773054585245386', 'STO', 3, '', '2020-09-06 23:33:31', '', '2020-09-08 12:51:19');
INSERT INTO `shipment_tbl` VALUES (49, '20200905220758283601615677103152', '773054585582825', 'STO', 3, '', '2020-09-06 23:33:58', '', '2020-09-09 22:34:16');
INSERT INTO `shipment_tbl` VALUES (50, '20200905230754819424429232943536', '773054583991235', 'STO', 3, '', '2020-09-06 23:35:12', '', '2020-09-08 13:59:57');
INSERT INTO `shipment_tbl` VALUES (52, '20200904142253922366777669901009', '773054585583004', 'STO', 3, '', '2020-09-06 23:37:05', '', '2020-09-08 13:59:57');
INSERT INTO `shipment_tbl` VALUES (53, '20200904152858231509030325604741', '773054585734702', 'STO', 3, '', '2020-09-06 23:38:01', '', '2020-09-07 22:02:01');
INSERT INTO `shipment_tbl` VALUES (54, '20200905101212164920377959660604', '773054585734878', 'STO', 3, '', '2020-09-06 23:40:17', '', '2020-09-08 09:25:10');
INSERT INTO `shipment_tbl` VALUES (55, '20200905101140560649495877421640', '773054585734933', 'STO', 3, '', '2020-09-06 23:40:49', '', '2020-09-07 22:01:18');
INSERT INTO `shipment_tbl` VALUES (56, '20200905100916721302892902213519', '773054585510778', 'STO', 3, '', '2020-09-06 23:41:35', '', '2020-09-08 13:59:57');
INSERT INTO `shipment_tbl` VALUES (57, '20200905100916721302892902213519', '773054585661840', 'STO', 3, '', '2020-09-06 23:42:22', '', '2020-09-08 13:59:58');
INSERT INTO `shipment_tbl` VALUES (58, '20200905101140560649495877421640', '773054585813745', 'STO', 3, '', '2020-09-06 23:42:54', '', '2020-09-07 22:01:18');
INSERT INTO `shipment_tbl` VALUES (59, '20200905101140560649495877421640', '773054584544493', 'STO', 3, '', '2020-09-06 23:43:09', '', '2020-09-07 22:01:17');
INSERT INTO `shipment_tbl` VALUES (60, '20200905101140560649495877421640', '773054583991345', 'STO', 3, '', '2020-09-06 23:43:21', '', '2020-09-07 22:01:17');
INSERT INTO `shipment_tbl` VALUES (61, '20200906172640442707334024777544', 'YT4768219922650', 'YTO', 3, '', '2020-09-07 11:02:08', '', '2020-09-12 00:23:26');
INSERT INTO `shipment_tbl` VALUES (62, '20200906205834271379132055472660', 'YT4768216385233', 'YTO', 3, '', '2020-09-07 11:03:09', '', '2020-09-09 22:34:17');
INSERT INTO `shipment_tbl` VALUES (63, '20200906205834271379132055472660', 'YT4768216384701', 'YTO', 3, '', '2020-09-07 11:04:22', '', '2020-09-09 22:34:18');
INSERT INTO `shipment_tbl` VALUES (64, '20200906211742769675900115363704', 'YT4768219923622', 'YTO', 3, '', '2020-09-07 11:04:38', '', '2020-09-09 22:34:18');
INSERT INTO `shipment_tbl` VALUES (65, '20200907113041606114995807914455', 'YT4768493934859', 'YTO', 3, '', '2020-09-07 12:58:45', '', '2020-09-08 13:59:58');
INSERT INTO `shipment_tbl` VALUES (66, '20200907113041606114995807914455', 'YT4768496031904', 'YTO', 3, '', '2020-09-07 12:59:07', '', '2020-09-08 13:59:58');
INSERT INTO `shipment_tbl` VALUES (67, '20200907112957847817507786300510', 'YT4768496704095', 'YTO', 3, '', '2020-09-07 12:59:27', '', '2020-09-08 13:59:58');
INSERT INTO `shipment_tbl` VALUES (68, '20200907112957847817507786300510', 'YT4768497494319', 'YTO', 3, '', '2020-09-07 12:59:37', '', '2020-09-08 13:59:58');
INSERT INTO `shipment_tbl` VALUES (69, '20200905172354473834479740994785', '773054739511845', 'STO', 3, '', '2020-09-07 15:05:40', '', '2020-09-08 19:49:55');
INSERT INTO `shipment_tbl` VALUES (70, '20200906172734374023330061786458', '773054741020685', 'STO', 3, '', '2020-09-07 15:06:05', '', '2020-09-08 19:48:54');
INSERT INTO `shipment_tbl` VALUES (71, '20200907112957847817507786300510', '773054740940829', 'STO', 3, '', '2020-09-07 15:19:48', '', '2020-09-08 13:59:59');
INSERT INTO `shipment_tbl` VALUES (72, '20200908001844369571285664356405', 'YT4770457089144', 'YTO', 3, '', '2020-09-08 11:38:35', '', '2020-09-12 19:49:12');
INSERT INTO `shipment_tbl` VALUES (73, '20200908001844369571285664356405', 'YT4770460243366', 'YTO', 3, '', '2020-09-08 11:38:45', '', '2020-09-10 12:09:39');
INSERT INTO `shipment_tbl` VALUES (74, '20200907200833307196699851825808', '773054906506161', 'STO', 3, '', '2020-09-08 15:56:26', '', '2020-09-13 19:35:47');
INSERT INTO `shipment_tbl` VALUES (75, '20200907200432080454900544836795', '773054904868479', 'STO', 3, '', '2020-09-08 15:56:39', '', '2020-09-12 00:47:41');
INSERT INTO `shipment_tbl` VALUES (76, '20200907200833307196699851825808', '773054904635511', 'STO', 3, '', '2020-09-08 15:56:57', '', '2020-09-14 16:10:25');
INSERT INTO `shipment_tbl` VALUES (77, '20200907155031329533846776004093', '773054906506161', 'STO', 3, '', '2020-09-08 15:57:18', '', '2020-09-13 19:35:47');
INSERT INTO `shipment_tbl` VALUES (78, '20200905101110670537826772928058', '773054906075645', 'STO', 3, '', '2020-09-08 15:58:09', '', '2020-09-12 00:47:42');
INSERT INTO `shipment_tbl` VALUES (79, '20200905101110670537826772928058', '773054906427998', 'STO', 3, '', '2020-09-08 15:58:15', '', '2020-09-12 00:47:42');
INSERT INTO `shipment_tbl` VALUES (80, '20200908213407864360175015892977', 'YT4773416999820', 'YTO', 3, '', '2020-09-09 13:47:07', '', '2020-09-12 00:47:42');
INSERT INTO `shipment_tbl` VALUES (81, '20200908213407864360175015892977', 'YT4773415789541', 'YTO', 3, '', '2020-09-09 13:47:18', '', '2020-09-12 00:47:42');
INSERT INTO `shipment_tbl` VALUES (82, '20200908213407864360175015892977', 'YT4773419880126', 'YTO', 3, '', '2020-09-09 13:47:32', '', '2020-09-12 00:47:42');
INSERT INTO `shipment_tbl` VALUES (83, '20200908213407864360175015892977', 'YT4773417016273', 'YTO', 3, '', '2020-09-09 13:47:40', '', '2020-09-12 00:47:42');
INSERT INTO `shipment_tbl` VALUES (84, '20200908213407864360175015892977', 'YT4773419154635', 'YTO', 3, '', '2020-09-09 13:47:54', '', '2020-09-12 00:47:43');
INSERT INTO `shipment_tbl` VALUES (85, '20200908213407864360175015892977', 'YT4773417656935', 'YTO', 3, '', '2020-09-09 13:48:01', '', '2020-09-12 00:47:43');
INSERT INTO `shipment_tbl` VALUES (86, '20200908213407864360175015892977', 'YT4773418439870', 'YTO', 3, '', '2020-09-09 13:48:12', '', '2020-09-12 00:47:43');
INSERT INTO `shipment_tbl` VALUES (87, '20200908213407864360175015892977', 'YT4773419910786', 'YTO', 3, '', '2020-09-09 13:48:17', '', '2020-09-12 00:47:43');
INSERT INTO `shipment_tbl` VALUES (88, '20200908213407864360175015892977', 'YT4773419902199', 'YTO', 3, '', '2020-09-09 13:48:23', '', '2020-09-12 00:47:43');
INSERT INTO `shipment_tbl` VALUES (89, '20200908213407864360175015892977', 'YT4773418461886', 'YTO', 3, '', '2020-09-09 13:48:35', '', '2020-09-14 22:04:15');
INSERT INTO `shipment_tbl` VALUES (90, '20200908213407864360175015892977', 'YT4773418460044', 'YTO', 3, '', '2020-09-09 13:48:40', '', '2020-09-12 00:47:43');
INSERT INTO `shipment_tbl` VALUES (91, '20200908213407864360175015892977', 'YT4773415806150', 'YTO', 3, '', '2020-09-09 13:48:52', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (92, '20200908213407864360175015892977', 'YT4773415801755', 'YTO', 3, '', '2020-09-09 13:48:57', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (93, '20200908213407864360175015892977', 'YT4773416451220', 'YTO', 3, '', '2020-09-09 13:49:03', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (94, '20200908213407864360175015892977', 'YT4773419169365', 'YTO', 3, '', '2020-09-09 13:49:08', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (95, '20200908213407864360175015892977', 'YT4773419890573', 'YTO', 3, '', '2020-09-09 13:49:17', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (96, '20200908213407864360175015892977', 'YT4773419172800', 'YTO', 3, '', '2020-09-09 13:49:22', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (97, '20200908213407864360175015892977', 'YT4773417677734', 'YTO', 3, '', '2020-09-09 13:49:27', '', '2020-09-12 00:47:44');
INSERT INTO `shipment_tbl` VALUES (98, '20200908213407864360175015892977', 'YT4773417665418', 'YTO', 3, '', '2020-09-09 13:49:31', '', '2020-09-12 00:47:45');
INSERT INTO `shipment_tbl` VALUES (99, '20200908213407864360175015892977', 'YT4773419884144', 'YTO', 3, '', '2020-09-09 13:49:36', '', '2020-09-12 00:47:45');
INSERT INTO `shipment_tbl` VALUES (100, '20200909132911945714226744764476', 'YT4774201651604', 'YTO', 3, '', '2020-09-09 16:31:26', '', '2020-09-14 20:00:01');
INSERT INTO `shipment_tbl` VALUES (101, '20200909224510473591859264498515', 'YT4775134367550', 'YTO', 3, '', '2020-09-10 09:01:00', '', '2020-09-14 20:00:01');
INSERT INTO `shipment_tbl` VALUES (102, '20200909175523231976514692794715', 'YT4775131545884', 'YTO', 3, '', '2020-09-10 09:14:45', '', '2020-09-12 20:00:05');
INSERT INTO `shipment_tbl` VALUES (103, '20200909181944566133514501774934', 'YT4775132833220', 'YTO', 3, '', '2020-09-10 09:15:03', '', '2020-09-12 19:49:02');
INSERT INTO `shipment_tbl` VALUES (104, '20200909174129539543023550749456', 'YT4775132837766', 'YTO', 3, '', '2020-09-10 09:15:23', '', '2020-09-11 13:11:32');
INSERT INTO `shipment_tbl` VALUES (105, '20200909180219226293407553151149', 'YT4775135853366', 'YTO', 3, '', '2020-09-10 09:15:46', '', '2020-09-11 23:05:35');
INSERT INTO `shipment_tbl` VALUES (106, '20200909182243435271729424888372', 'YT4775135848765', 'YTO', 3, '', '2020-09-10 09:16:20', '', '2020-09-11 23:05:59');
INSERT INTO `shipment_tbl` VALUES (107, '20200909173357807763164552479653', 'YT4775134358570', 'YTO', 3, '', '2020-09-10 09:17:54', '', '2020-09-11 13:11:00');
INSERT INTO `shipment_tbl` VALUES (108, '20200909173953701539348896530383', 'YT4775135096174', 'YTO', 3, '', '2020-09-10 09:18:08', '', '2020-09-11 18:34:49');
INSERT INTO `shipment_tbl` VALUES (109, '20200909173555738268632785367647', 'YT4775135098761', 'YTO', 3, '', '2020-09-10 09:18:16', '', '2020-09-13 18:26:55');
INSERT INTO `shipment_tbl` VALUES (110, '20200911133940601871451794926551', 'YT4779927240766', 'YTO', 3, '', '2020-09-11 21:54:42', '', '2020-09-13 19:34:50');
INSERT INTO `shipment_tbl` VALUES (111, '20200910225201640326016963945994', 'YT4779924412828 ', 'YTO', 3, '', '2020-09-11 21:55:46', '', '2020-09-13 19:35:15');
INSERT INTO `shipment_tbl` VALUES (112, '20200911113523228859309119284489', 'YT4779925047925', 'YTO', 3, '', '2020-09-11 21:56:07', '', '2020-09-13 19:35:07');
INSERT INTO `shipment_tbl` VALUES (113, '20200913224618924871225599168767', 'YT4785137386427', 'YTO', 3, '', '2020-09-14 11:17:12', '', '2020-09-15 11:51:06');
INSERT INTO `shipment_tbl` VALUES (114, '20200914153209705330988504385799', 'YT4788197734601', 'YTO', 3, '', '2020-09-15 16:37:28', '', '2020-09-19 20:00:00');
INSERT INTO `shipment_tbl` VALUES (115, '20200914153209705330988504385799', 'YT4788195739310', 'YTO', 3, '', '2020-09-15 16:37:42', '', '2020-09-19 20:00:01');
INSERT INTO `shipment_tbl` VALUES (116, '20200914152842722331661049895077', 'YT4788197145362', 'YTO', 3, '', '2020-09-15 16:41:02', '', '2020-09-17 20:00:03');
INSERT INTO `shipment_tbl` VALUES (117, '20200915133509274259700616093749', 'YT4788893009381', 'YTO', 3, '', '2020-09-15 16:41:22', '', '2020-09-16 17:39:10');
INSERT INTO `shipment_tbl` VALUES (123, '20200916170648321670796277731896', 'YT4791786040000', 'YTO', 3, '', '2020-09-16 19:15:55', '', '2020-09-20 20:00:01');
INSERT INTO `shipment_tbl` VALUES (124, '20200916095823206092291933165418', 'YT4791007099990', 'YTO', 2, '', '2020-09-16 19:16:23', '', '2020-09-16 19:16:23');
INSERT INTO `shipment_tbl` VALUES (125, '20200916170607973462602149328237', 'YT4791894898331', 'YTO', 3, '', '2020-09-17 07:41:22', '', '2020-09-20 09:10:28');
INSERT INTO `shipment_tbl` VALUES (126, '20200916101331330390174948627859', 'YT4791895530363', 'YTO', 3, '', '2020-09-17 07:53:24', '', '2020-09-20 20:00:04');
INSERT INTO `shipment_tbl` VALUES (127, '20200916100118906678032272912298', 'YT4791879822002', 'YTO', 3, '', '2020-09-17 07:54:23', '', '2020-09-19 20:00:06');
INSERT INTO `shipment_tbl` VALUES (128, '20200916095846708669332593966577', 'YT4791898277719', 'YTO', 3, '', '2020-09-17 07:57:02', '', '2020-09-21 20:00:01');
INSERT INTO `shipment_tbl` VALUES (129, '20200916102843054887933764793234', 'YT4791911088503', 'YTO', 3, '', '2020-09-17 08:02:40', '', '2020-09-20 20:00:04');
INSERT INTO `shipment_tbl` VALUES (130, '20200917000138344371884833639098', 'YT4793111423033', 'YTO', 3, '', '2020-09-17 14:09:02', '', '2020-09-21 15:06:04');
INSERT INTO `shipment_tbl` VALUES (131, '20200916220459204895794235096025', 'YT4793107685946', 'YTO', 3, '', '2020-09-17 14:09:19', '', '2020-09-19 20:00:06');
INSERT INTO `shipment_tbl` VALUES (133, '20200918010640642627935054150104', ' YT4795520163866', 'YTO', 2, '', '2020-09-18 21:37:13', '', '2020-09-18 21:37:13');
INSERT INTO `shipment_tbl` VALUES (134, '20200918010640642627935054150104', 'YT4795518645450', 'YTO', 2, '', '2020-09-18 21:37:35', '', '2020-09-18 21:37:35');
INSERT INTO `shipment_tbl` VALUES (135, '20200919004125750220078134736612', 'JDVD00976222190', 'JD', 3, '', '2020-09-20 09:03:44', '', '2020-09-22 22:31:29');
INSERT INTO `shipment_tbl` VALUES (136, '20200918212222804517077714055863', 'JDVD00976193380', 'JD', 3, '', '2020-09-20 09:05:44', '', '2020-09-21 16:22:24');
INSERT INTO `shipment_tbl` VALUES (137, '20200918211743948841920866386370', 'JDVD00976193399', 'JD', 3, '', '2020-09-20 09:06:07', '', '2020-09-28 22:45:50');
INSERT INTO `shipment_tbl` VALUES (138, '20200918211256738001240093951339', 'JDVD00976185082', 'JD', 3, '', '2020-09-20 09:06:42', '', '2020-09-21 15:10:20');
INSERT INTO `shipment_tbl` VALUES (139, '20200918205121505111823459706356', 'JDVD00976242970', 'JD', 3, '', '2020-09-20 09:07:58', '', '2020-09-28 22:45:51');
INSERT INTO `shipment_tbl` VALUES (140, '20200918211256738001240093951339', 'JDVD00973573239', 'JD', 3, '', '2020-09-20 09:43:46', '', '2020-09-20 10:54:46');
INSERT INTO `shipment_tbl` VALUES (141, '20200918214539667943671446001690', 'JDVD00977312945', 'JD', 3, '', '2020-09-20 13:32:10', '', '2020-09-21 15:10:00');
INSERT INTO `shipment_tbl` VALUES (142, '20200919215051807129549924894234', 'JDVD00977342825', 'JD', 3, '', '2020-09-20 13:32:33', '', '2020-09-24 12:54:42');
INSERT INTO `shipment_tbl` VALUES (143, '20200920111346347682582153851631', 'YT4801608577077', 'YTO', 3, '', '2020-09-21 09:34:32', '', '2020-09-25 12:43:45');
INSERT INTO `shipment_tbl` VALUES (144, '20200920225337966405965162970066', 'YT4803129557688', 'YTO', 3, '', '2020-09-21 17:47:30', '', '2020-09-28 22:45:51');
INSERT INTO `shipment_tbl` VALUES (145, '20200920225337966405965162970066', 'YT4803127585304', 'YTO', 3, '', '2020-09-21 17:47:55', '', '2020-09-28 22:45:52');
INSERT INTO `shipment_tbl` VALUES (146, '20200921225715668107825830084149', 'JDVD00983085779', 'JD', 3, '', '2020-09-22 11:11:50', '', '2020-09-24 11:41:26');
INSERT INTO `shipment_tbl` VALUES (147, '20200921111652624127792659218763', 'YT4806894253200', 'YTO', 3, '', '2020-09-23 08:11:47', '', '2020-10-22 20:00:02');
INSERT INTO `shipment_tbl` VALUES (148, '20200923144058552818290275559056', 'YT4812020656267', 'YTO', 3, '', '2020-09-25 08:11:33', '', '2020-09-30 15:45:16');
INSERT INTO `shipment_tbl` VALUES (149, '20200921150602088390042186894130', 'YT4811974436620', 'YTO', 3, '', '2020-09-25 08:11:59', '', '2020-10-22 20:00:04');
INSERT INTO `shipment_tbl` VALUES (150, '20200928161448855787209481488398', '9880201780019', 'EMS', 3, '', '2020-09-29 13:34:59', '', '2020-10-22 20:00:04');
INSERT INTO `shipment_tbl` VALUES (151, '20200929152056700145188914387520', 'YT4823368045243', 'YTO', 3, '', '2020-09-29 16:21:57', '', '2020-10-22 20:00:05');
INSERT INTO `shipment_tbl` VALUES (152, '20200929152056700145188914387520', '9880209246003', 'EMS', 3, '', '2020-09-30 08:10:31', '', '2020-10-22 20:00:06');
INSERT INTO `shipment_tbl` VALUES (153, '20200929234323924609999790176394', '9880218742576', 'EMS', 3, '', '2020-10-01 08:10:26', '', '2020-10-02 15:35:22');
INSERT INTO `shipment_tbl` VALUES (154, '20200929234323924609999790176394', '9880218751881', 'EMS', 3, '', '2020-10-01 08:10:47', '', '2020-10-02 15:35:22');
INSERT INTO `shipment_tbl` VALUES (155, '20200929234323924609999790176394', '9880218737813', 'EMS', 3, '', '2020-10-01 08:11:04', '', '2020-10-13 00:00:36');
INSERT INTO `shipment_tbl` VALUES (156, '20200929234323924609999790176394', '9880218728209', 'EMS', 3, '', '2020-10-01 08:11:20', '', '2020-10-02 15:35:22');
INSERT INTO `shipment_tbl` VALUES (157, '20200930153824681834078381261434', 'YT4828059531673', 'YTO', 3, '', '2020-10-02 11:33:12', '', '2020-10-07 00:46:46');
INSERT INTO `shipment_tbl` VALUES (158, '20201007131507381968807761569225', '4309032943340', 'YD', 3, '', '2020-10-07 17:36:52', '', '2020-10-12 18:57:14');
INSERT INTO `shipment_tbl` VALUES (159, '20201007131507381968807761569225', '4309032935655', 'YD', 3, '', '2020-10-07 17:37:13', '', '2020-10-10 12:27:45');
INSERT INTO `shipment_tbl` VALUES (160, '20201007131026472772406962514736', '4309032960869', 'YD', 3, '', '2020-10-07 17:37:52', '', '2020-10-10 12:27:51');
INSERT INTO `shipment_tbl` VALUES (161, '20201007130625914345795382843099', '4309033023721', 'YD', 3, '', '2020-10-07 17:38:05', '', '2020-10-10 12:27:57');
INSERT INTO `shipment_tbl` VALUES (162, '20201007131507381968807761569225', '4309032908240', 'YD', 3, '', '2020-10-07 17:44:39', '', '2020-10-12 18:57:14');
INSERT INTO `shipment_tbl` VALUES (163, '20201008123403155119933925257989', 'JDVD01028720803', 'JD', 3, '', '2020-10-08 16:26:48', '', '2020-10-10 19:19:06');
INSERT INTO `shipment_tbl` VALUES (164, '20201008123403155119933925257989', 'JDVD01028714548', 'JD', 3, '', '2020-10-08 16:27:24', '', '2020-10-10 19:19:06');
INSERT INTO `shipment_tbl` VALUES (165, '20201013130024637638374055013401', '9880336228325', 'EMS', 3, '', '2020-10-14 10:06:39', '', '2020-10-18 16:44:11');
INSERT INTO `shipment_tbl` VALUES (166, '20201013130024637638374055013401', '9880336237785', 'EMS', 3, '', '2020-10-14 10:06:48', '', '2020-10-18 16:44:11');
INSERT INTO `shipment_tbl` VALUES (167, '20201014100202889773388574814316', 'JDVD01047764919', 'JD', 3, '', '2020-10-14 15:03:51', '', '2020-10-16 22:41:55');
INSERT INTO `shipment_tbl` VALUES (168, '20201125181640452848153660757904', '776260237245573', 'STO', 3, '', '2020-11-26 14:37:25', '', '2020-11-30 15:04:23');

-- ----------------------------
-- Table structure for user_link_log_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_link_log_tbl`;
CREATE TABLE `user_link_log_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '微信用户唯一标识符',
  `link_open_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '关联的微信用户唯一标识符',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信用户信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_link_log_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for user_link_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_link_tbl`;
CREATE TABLE `user_link_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '被邀请人',
  `parent_user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邀请人',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信用户信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_link_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for user_tbl
-- ----------------------------
DROP TABLE IF EXISTS `user_tbl`;
CREATE TABLE `user_tbl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '微信用户唯一标识符',
  `vip` tinyint(1) NULL DEFAULT NULL COMMENT '是否vip',
  `dist_vip` tinyint(1) NULL DEFAULT NULL COMMENT '是否分销vip',
  `balance` bigint(20) NULL DEFAULT NULL COMMENT '用户余额',
  `points` bigint(20) NULL DEFAULT NULL COMMENT '用户积分',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信用户信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_tbl
-- ----------------------------

-- ----------------------------
-- Table structure for wx_userinfo
-- ----------------------------
DROP TABLE IF EXISTS `wx_userinfo`;
CREATE TABLE `wx_userinfo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '微信用户唯一标识符\n\n微信用户唯一标识符',
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `avatar_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。',
  `gender` int(11) NULL DEFAULT NULL COMMENT '用户的性别，值为1时是男性，值为2时是女性，值为0时是未知',
  `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户所在城市',
  `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户所在省份',
  `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户所在国家',
  `language` varbinary(16) NULL DEFAULT NULL COMMENT '用户的语言，简体中文为zh_CN',
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户绑定的手机号',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '微信用户信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of wx_userinfo
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
