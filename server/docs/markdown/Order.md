

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Order]
- [Waiting to write...]

--------------------

### Pay

#### 简要描述：

- [支付]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.pay

#### 请求方式：

- post

#### 请求参数:

- ` PayReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |


#### 请求示例:
```
{
     "billId": ""
}
```

#### 返回参数说明:

- ` PayResp ` : 下单

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|bool|状态 1:成功，-1:失败   |
|`billId` | 否|string|账单id   |
|`info` | 否|map (string,string)|支付相关信息   |


#### 返回示例:
	
```
{
     "billId": "",
     "info": "",
     "status": false
}
```

#### 备注:

- 支付

--------------------

### DealSystem

#### 简要描述：

- [意见与反馈]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.deal_system

#### 请求方式：

- post

#### 请求参数:

- ` DealOrderReq ` : 处理订单

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |
|`type` | 否|int32|操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈   |
|`contact` | 否|string|联系方式   |
|`remak` | 否|string|备注   |
|`imgs` | 否|[]string|售后图片   |


#### 请求示例:
```
{
     "billId": "",
     "contact": "",
     "imgs": [
          ""
     ],
     "remak": "",
     "type": 0
}
```

#### 返回参数说明:

- ` Empty ` : 空

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 意见与反馈

--------------------

### GetBillInfoByAdmin

#### 简要描述：

- [获取订单信息]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.get_bill_info_by_admin

#### 请求方式：

- post

#### 请求参数:

- ` GetMyOrdersReq ` : 获取订单列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|int32|支付状态(-1:已退款,1:待支付,2:已支付,3:已取消)   |
|`pageNumber` | 否|int32|当前页数   |
|`search` | 否|string|搜索内容   |


#### 请求示例:
```
{
     "pageNumber": 0,
     "search": "",
     "status": 0
}
```

#### 返回参数说明:

- ` GetMyOrdersResp ` : 获取订单列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`totalPages` | 否|int32|总页数   |
|`pageNumber` | 否|int32|当前页数   |
|`items` | 否|[]`caoguo.BillOrderInfo`|账单列表   |


#### 返回示例:
	
```
{
     "items": [
          {
               "addr": {
                    "address": "",
                    "addressName": "",
                    "area": "",
                    "default": false,
                    "id": 0,
                    "mobile": "",
                    "name": ""
               },
               "billId": "",
               "checkStatus": 0,
               "couponTitle": "",
               "items": [
                    {
                         "distAmount": 0,
                         "gpid": "",
                         "icon": "",
                         "name": "",
                         "number": 0,
                         "price": 0,
                         "skuList": "",
                         "skuVal": ""
                    }
               ],
               "number": 0,
               "shipmentInfo": [
                    {
                         "shipmentId": "",
                         "shipmentName": ""
                    }
               ],
               "stateTip": "",
               "stateTipColor": "",
               "status": 0,
               "time": "",
               "totalFee": 0
          }
     ],
     "pageNumber": 0,
     "totalPages": 0
}
```

#### 备注:

- 获取订单信息

--------------------

### GetShipmentPost

#### 简要描述：

- [获取快递列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.get_shipment_post

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

- ` GetShipmentPostResp ` : 获取快递列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`items` | 否|[]`caoguo.PostInfo`|   |


#### 返回示例:
	
```
{
     "items": [
          {
               "expPhone": "",
               "icon": "",
               "id": "",
               "name": ""
          }
     ]
}
```

#### 备注:

- 获取快递列表

--------------------

### GetAfterMsg

#### 简要描述：

- [获取回复信息]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.get_after_msg

#### 请求方式：

- post

#### 请求参数:

- ` AfterMsgReq ` : 售后回复请求

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |


#### 请求示例:
```
{
     "billId": ""
}
```

#### 返回参数说明:

- ` AfterMsgResp ` : 售后回复请求

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |
|`items` | 否|[]`caoguo.AfterMsg`|搜后回复信息   |


#### 返回示例:
	
```
{
     "billId": "",
     "items": [
          {
               "contact": "",
               "imgs": [
                    ""
               ],
               "remak": "",
               "time": ""
          }
     ]
}
```

#### 备注:

- 获取回复信息

--------------------

### GetMyOrders

#### 简要描述：

- [获取所有订单]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.get_my_orders

#### 请求方式：

- post

#### 请求参数:

- ` GetMyOrdersReq ` : 获取订单列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|int32|支付状态(-1:已退款,1:待支付,2:已支付,3:已取消)   |
|`pageNumber` | 否|int32|当前页数   |
|`search` | 否|string|搜索内容   |


#### 请求示例:
```
{
     "pageNumber": 0,
     "search": "",
     "status": 0
}
```

#### 返回参数说明:

- ` GetMyOrdersResp ` : 获取订单列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`totalPages` | 否|int32|总页数   |
|`pageNumber` | 否|int32|当前页数   |
|`items` | 否|[]`caoguo.BillOrderInfo`|账单列表   |


#### 返回示例:
	
```
{
     "items": [
          {
               "addr": {
                    "address": "",
                    "addressName": "",
                    "area": "",
                    "default": false,
                    "id": 0,
                    "mobile": "",
                    "name": ""
               },
               "billId": "",
               "checkStatus": 0,
               "couponTitle": "",
               "items": [
                    {
                         "distAmount": 0,
                         "gpid": "",
                         "icon": "",
                         "name": "",
                         "number": 0,
                         "price": 0,
                         "skuList": "",
                         "skuVal": ""
                    }
               ],
               "number": 0,
               "shipmentInfo": [
                    {
                         "shipmentId": "",
                         "shipmentName": ""
                    }
               ],
               "stateTip": "",
               "stateTipColor": "",
               "status": 0,
               "time": "",
               "totalFee": 0
          }
     ],
     "pageNumber": 0,
     "totalPages": 0
}
```

#### 备注:

- 获取所有订单

--------------------

### GetMyOrdersConfig

#### 简要描述：

- [获取订单配置]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.get_my_orders_config

#### 请求方式：

- post

#### 请求参数:

- ` GetMyOrdersConfigReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`isAdmin` | 否|bool|是否用户端   |


#### 请求示例:
```
{
     "isAdmin": false
}
```

#### 返回参数说明:

- ` GetMyOrdersConfigResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`navList` | 否|[]`caoguo.NavList`|类型列表   |


#### 返回示例:
	
```
{
     "navList": [
          {
               "loadingType": "",
               "state": 0,
               "text": ""
          }
     ]
}
```

#### 备注:

- 获取订单配置

--------------------

### GetOrdertrackInfo

#### 简要描述：

- [获取物流轨迹信息]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.get_ordertrack_info

#### 请求方式：

- post

#### 请求参数:

- ` GetOrdertrackInfoReq ` : 获取订单物流信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|订单号   |


#### 请求示例:
```
{
     "billId": ""
}
```

#### 返回参数说明:

- ` GetOrdertrackInfoResp ` : 获取订单物流信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|订单id   |
|`deliveryStatus` | 否|int32|快递状态 1已签收 2配送中   |
|`postName` | 否|string|快递名称   |
|`logo` | 否|string|快递logo   |
|`expPhone` | 否|string|快递电话   |
|`postNo` | 否|string|快递单号   |
|`addr` | 否|string|收货地址   |
|`list` | 否|[]`caoguo.Logistics`|物流信息   |


#### 返回示例:
	
```
{
     "addr": "",
     "billId": "",
     "deliveryStatus": 0,
     "expPhone": "",
     "list": [
          {
               "context": "",
               "time": "",
               "timeArr": [
                    ""
               ]
          }
     ],
     "logo": "",
     "postName": "",
     "postNo": ""
}
```

#### 备注:

- 获取物流轨迹信息

--------------------

### PlaceOrder

#### 简要描述：

- [下单]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.place_order

#### 请求方式：

- post

#### 请求参数:

- ` PlaceOrderReq ` : 下单

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`buyType` | 否|int32|购买类型:1,来自直接购买，2:来自购物车   |
|`ids` | 否|[]int64|购买列表   |
|`addrId` | 否|int64|地址id   |
|`couponId` | 否|int64|优惠券id   |
|`desc` | 否|string|备注   |


#### 请求示例:
```
{
     "addrId": 0,
     "buyType": 0,
     "couponId": 0,
     "desc": "",
     "ids": [
          0
     ]
}
```

#### 返回参数说明:

- ` PayResp ` : 下单

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|bool|状态 1:成功，-1:失败   |
|`billId` | 否|string|账单id   |
|`info` | 否|map (string,string)|支付相关信息   |


#### 返回示例:
	
```
{
     "billId": "",
     "info": "",
     "status": false
}
```

#### 备注:

- 下单

--------------------

### AddBillShipment

#### 简要描述：

- [订单添加运单号]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.add_bill_shipment

#### 请求方式：

- post

#### 请求参数:

- ` AddBillShipmentReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|订单号   |
|`shipmentId` | 否|string|快递单号   |
|`postKey` | 否|string|快递id   |


#### 请求示例:
```
{
     "billId": "",
     "postKey": "",
     "shipmentId": ""
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

- 订单添加运单号

--------------------

### CheckPay

#### 简要描述：

- [支付成功check一次]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.check_pay

#### 请求方式：

- post

#### 请求参数:

- ` PayReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |


#### 请求示例:
```
{
     "billId": ""
}
```

#### 返回参数说明:

- ` PayResp ` : 下单

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|bool|状态 1:成功，-1:失败   |
|`billId` | 否|string|账单id   |
|`info` | 否|map (string,string)|支付相关信息   |


#### 返回示例:
	
```
{
     "billId": "",
     "info": "",
     "status": false
}
```

#### 备注:

- 支付成功check一次

--------------------

### DealOrder

#### 简要描述：

- [账单/订单处理]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.deal_order

#### 请求方式：

- post

#### 请求参数:

- ` DealOrderReq ` : 处理订单

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |
|`type` | 否|int32|操作：-1：取消，1：申请售后，2:删除订单,3:意见反馈   |
|`contact` | 否|string|联系方式   |
|`remak` | 否|string|备注   |
|`imgs` | 否|[]string|售后图片   |


#### 请求示例:
```
{
     "billId": "",
     "contact": "",
     "imgs": [
          ""
     ],
     "remak": "",
     "type": 0
}
```

#### 返回参数说明:

- ` Empty ` : 空

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 账单/订单处理

--------------------

### ReckonFee

#### 简要描述：

- [费用计算]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/order.reckon_fee

#### 请求方式：

- post

#### 请求参数:

- ` ReckonFeeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`buyType` | 否|int32|购买类型:1,来自直接购买，2:来自购物车   |
|`ids` | 否|[]int64|购买列表   |
|`addrId` | 否|int64|地址id   |
|`couponId` | 否|int64|优惠券id   |


#### 请求示例:
```
{
     "addrId": 0,
     "buyType": 0,
     "couponId": 0,
     "ids": [
          0
     ]
}
```

#### 返回参数说明:

- ` ReckonFeeResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`orderInfo` | 否|[]`caoguo.OrderProductInfo`|订单信息   |
|`feeDetails` | 否|[]`caoguo.FeeDetail`|费用明细   |
|`couponList` | 否|[]`caoguo.CouponInfo`|优惠券列表   |
|`couponDetail` | 否|string|优惠券内容   |
|`couponId` | 否|int64|优惠券id   |
|`promotionDetail` | 否|string|促销内容   |
|`vendorName` | 否|string|供应商名称   |
|`vendorImg` | 否|string|供应商图标   |
|`totalFee` | 否|int64|总金额   |
|`shipmentFee` | 否|int64|运输费   |
|`couponFee` | 否|int64|优惠金额   |


#### 返回示例:
	
```
{
     "couponDetail": "",
     "couponFee": 0,
     "couponId": 0,
     "couponList": [
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
     "feeDetails": [
          {
               "color": "",
               "key": "",
               "value": ""
          }
     ],
     "orderInfo": [
          {
               "distAmount": 0,
               "gpid": "",
               "icon": "",
               "name": "",
               "number": 0,
               "price": 0,
               "skuList": "",
               "skuVal": ""
          }
     ],
     "promotionDetail": "",
     "shipmentFee": 0,
     "totalFee": 0,
     "vendorImg": "",
     "vendorName": ""
}
```

#### 备注:

- 费用计算
	

--------------------
--------------------

#### 自定义类型:

#### ` caoguo `


- ` ShipmentInfo ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`shipmentId` | 否|string|快递单号   |
|`shipmentName` | 否|string|快递名   |



- ` OrderProductInfo ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`name` | 否|string|商品名   |
|`price` | 否|int64|商品价格   |
|`icon` | 否|string|商品图标   |
|`skuVal` | 否|string|sku描述   |
|`skuList` | 否|string|sku信息   |
|`number` | 否|int64|购买数量   |
|`distAmount` | 否|int64|分销金额   |



- ` AddressInfo ` : 地址信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|   |
|`name` | 否|string|名字   |
|`mobile` | 否|string|手机号   |
|`addressName` | 否|string|地址名称   |
|`address` | 否|string|详细地址   |
|`area` | 否|string|单元门牌号   |
|`default` | 否|bool|是否默认   |



- ` BillOrderInfo ` : 订单信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|int32|支付状态(-1:已退款,1:待支付,2:已支付,3:已取消)   |
|`billId` | 否|string|账单id   |
|`totalFee` | 否|int64|总费用   |
|`time` | 否|string|时间   |
|`stateTip` | 否|string|状态描述   |
|`stateTipColor` | 否|string|状态颜色   |
|`number` | 否|int64|数量   |
|`checkStatus` | 否|int64|点击状态   |
|`couponTitle` | 否|string|优惠券描述   |
|`shipmentInfo` | 否|[]`caoguo.ShipmentInfo`|订单快递单号   |
|`items` | 否|[]`caoguo.OrderProductInfo`|订单信息   |
|`addr` | 否|`caoguo.AddressInfo`|配送地址   |



- ` PostInfo ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`name` | 否|string|快递名称   |
|`id` | 否|string|快递id   |
|`icon` | 否|string|快递logo地址   |
|`expPhone` | 否|string|快递电话   |



- ` AfterMsg ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`contact` | 否|string|联系方式   |
|`remak` | 否|string|备注   |
|`time` | 否|string|   |
|`imgs` | 否|[]string|售后图片   |



- ` NavList ` : 类型列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`state` | 否|int32|支付状态(0:全部,-1:已退款,1:待支付,2:已支付,3:已取消)   |
|`text` | 否|string|内容   |
|`loadingType` | 否|string|&amp;#39;more&amp;#39;,   |



- ` Logistics ` : 物流信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`time` | 否|string|时间   |
|`timeArr` | 否|[]string|时间分割([&amp;#39;2020-04-12&amp;#39;, &amp;#39;13:00:57&amp;#39;])   |
|`context` | 否|string|描述信息   |



- ` FeeDetail ` : 费用明细

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`key` | 否|string|名字   |
|`value` | 否|string|值   |
|`color` | 否|string|颜色   |



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




