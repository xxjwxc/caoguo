

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Coupon]
- [Waiting to write...]

--------------------

### GetAdminCouponInfo

#### 简要描述：

- [admin 获取优惠券列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/coupon.get_admin_coupon_info

#### 请求方式：

- post

#### 请求参数:

- ` Empty ` : 空

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetAdminCouponInfo ` : admin 优惠券相关信息(获取,更新或者添加)

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`couponList` | 否|[]`caoguo.AdminCouponInfo`|优惠券列表   |


#### 返回示例:
	
```
{
     "couponList": [
          {
               "couponType": 0,
               "denom": 0,
               "describle": "",
               "gpid": "",
               "greatMoney": 0,
               "groupName": "",
               "id": 0,
               "title": "",
               "vaild": false,
               "validity": 0
          }
     ]
}
```

#### 备注:

- admin 获取优惠券列表

--------------------

### GetMyCoupon

#### 简要描述：

- [获取我的优惠券]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/coupon.get_my_coupon

#### 请求方式：

- post

#### 请求参数:

- ` GetMyCouponReq ` : 获取我的优惠券

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetMyCouponResp ` : 获取我的优惠券返回

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`items` | 否|[]`caoguo.CouponGroupList`|优惠券列表   |


#### 返回示例:
	
```
{
     "items": [
          {
               "items": [
                    {
                         "couponType": 0,
                         "denom": 0,
                         "describle": "",
                         "gpid": "",
                         "greatMoney": "",
                         "groupName": "",
                         "id": 0,
                         "status": 0,
                         "title": "",
                         "type": 0,
                         "vaild": false,
                         "validity": ""
                    }
               ],
               "key": ""
          }
     ]
}
```

#### 备注:

- 获取我的优惠券

--------------------

### GetPromotionCoupon

#### 简要描述：

- [获取促销优惠券]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/coupon.get_promotion_coupon

#### 请求方式：

- post

#### 请求参数:

- ` GetPromotionCouponReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetPromotionCouponResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`items` | 否|[]`caoguo.CouponGroupList`|优惠券列表   |
|`totalFee` | 否|int64|总费用   |
|`ids` | 否|[]int64|   |


#### 返回示例:
	
```
{
     "ids": [
          0
     ],
     "items": [
          {
               "items": [
                    {
                         "couponType": 0,
                         "denom": 0,
                         "describle": "",
                         "gpid": "",
                         "greatMoney": "",
                         "groupName": "",
                         "id": 0,
                         "status": 0,
                         "title": "",
                         "type": 0,
                         "vaild": false,
                         "validity": ""
                    }
               ],
               "key": ""
          }
     ],
     "totalFee": 0
}
```

#### 备注:

- 获取促销优惠券

--------------------

### GoGetCoupon

#### 简要描述：

- [一键领取优惠券]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/coupon.go_get_coupon

#### 请求方式：

- post

#### 请求参数:

- ` GoGetCouponReq ` : 领取优惠券

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`ids` | 否|[]int64|   |


#### 请求示例:
```
{
     "ids": [
          0
     ]
}
```

#### 返回参数说明:

- ` bool ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 一键领取优惠券

--------------------

### UpsetAdminCouponInfo

#### 简要描述：

- [admin 更新/添加 优惠券列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/coupon.upset_admin_coupon_info

#### 请求方式：

- post

#### 请求参数:

- ` GetAdminCouponInfo ` : admin 优惠券相关信息(获取,更新或者添加)

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`couponList` | 否|[]`caoguo.AdminCouponInfo`|优惠券列表   |


#### 请求示例:
```
{
     "couponList": [
          {
               "couponType": 0,
               "denom": 0,
               "describle": "",
               "gpid": "",
               "greatMoney": 0,
               "groupName": "",
               "id": 0,
               "title": "",
               "vaild": false,
               "validity": 0
          }
     ]
}
```

#### 返回参数说明:

- ` bool ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- admin 更新/添加 优惠券列表
	

--------------------
--------------------

#### 自定义类型:

#### ` caoguo `


- ` AdminCouponInfo ` : 优惠券返回信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|coupon id   |
|`groupName` | 否|string|分组名   |
|`title` | 否|string|优惠券名字   |
|`validity` | 否|int32|有效截止日期   |
|`gpid` | 否|string|商品优惠券商品id   |
|`denom` | 否|int64|面额   |
|`couponType` | 否|int32|优惠券类型(1:全场，2:单个商品)   |
|`greatMoney` | 否|int32|满多少可用   |
|`describle` | 否|string|优惠券描述   |
|`vaild` | 否|bool|是否有效   |



- ` CouponInfo ` : 优惠券信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|coupon id   |
|`groupName` | 否|string|分组名   |
|`title` | 否|string|优惠券名字   |
|`validity` | 否|string|有效截止日期   |
|`gpid` | 否|string|商品优惠券商品id   |
|`denom` | 否|int64|面额   |
|`couponType` | 否|int32|优惠券类型(1:全场，2:单个商品)   |
|`greatMoney` | 否|string|满多少可用   |
|`describle` | 否|string|优惠券描述   |
|`status` | 否|int32|0：默认，1:未使用(有效),2:已使用,-1:已过期   |
|`type` | 否|int32|优惠券类型，1:促销优惠券，2：用户已领取的优惠券   |
|`vaild` | 否|bool|是否有效   |



- ` CouponGroupList ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`key` | 否|string|groupname   |
|`items` | 否|[]`caoguo.CouponInfo`|优惠券列表   |




