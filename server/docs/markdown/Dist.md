

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Dist]
- [Waiting to write...]

--------------------

### GetDistList

#### 简要描述：

- [获取可分销列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/dist.get_dist_list

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

- ` GetDistListResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`info` | 否|`caoguo.DistDoc`|详情   |
|`items` | 否|[]`caoguo.DistProductInfo`|列表   |


#### 返回示例:
	
```
{
     "info": {
          "detail": "",
          "imgs": [
               ""
          ],
          "isVip": false
     },
     "items": [
          {
               "gpid": "",
               "icon": "",
               "name": "",
               "originalPrice": 0,
               "price": 0,
               "shareAmounts": [
                    ""
               ]
          }
     ]
}
```

#### 备注:

- 获取可分销列表

--------------------

### RequestDist

#### 简要描述：

- [申请分销]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/dist.request_dist

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

- ` ResultResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`result` | 否|bool|   |
|`msg` | 否|string|   |


#### 返回示例:
	
```
{
     "msg": "",
     "result": false
}
```

#### 备注:

- 申请分销
	

--------------------
--------------------

#### 自定义类型:

#### ` caoguo `


- ` DistDoc ` : 分销说明详细信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`isVip` | 否|bool|是否vip   |
|`detail` | 否|string|文字详情   |
|`imgs` | 否|[]string|图片详情   |



- ` DistProductInfo ` : 分销商品信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`name` | 否|string|商品名   |
|`price` | 否|int64|商品价格   |
|`originalPrice` | 否|int64|商品原始价格   |
|`icon` | 否|string|商品图标   |
|`shareAmounts` | 否|[]string|分享受益金   |




