

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Weixin]
- [Waiting to write...]

--------------------

### Oauth

#### 简要描述：

- [微信授权获取登录信息]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/weixin.oauth

#### 请求方式：

- post

#### 请求参数:

- ` OauthReq ` : 微信授权请求信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`jscode` | `是`|string|微信端jscode   |


#### 请求示例:
```
{
     "jscode": ""
}
```

#### 返回参数说明:

- ` OauthResp ` : 微信授权返回信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`session_id` | 否|string|用户sessionid 用于当前交互使用   |
|`open_id` | 否|string|openid   |
|`overdue_time` | 否|int64|逾期时间点(时间戳)   |


#### 返回示例:
	
```
{
     "open_id": "",
     "overdue_time": 0,
     "session_id": ""
}
```

#### 备注:

- 微信授权获取登录信息

--------------------

### RefundPay

#### 简要描述：

- [退款]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/weixin.refund_pay

#### 请求方式：

- post

#### 请求参数:

- ` RefundPayReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`billId` | 否|string|账单id   |
|`refundFee` | 否|int64|退款金额   |


#### 请求示例:
```
{
     "billId": "",
     "refundFee": 0
}
```

#### 返回参数说明:

- ` RefundPayResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|bool|是否退款成功   |


#### 返回示例:
	
```
{
     "status": false
}
```

#### 备注:

- 退款

--------------------

### UpdateUserInfo

#### 简要描述：

- [更新用户信息]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/weixin.update_user_info

#### 请求方式：

- post

#### 请求参数:

- ` WxUserinfo ` : 用户信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`nickName` | 否|string|用户昵称   |
|`avatarUrl` | 否|string|用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。   |
|`gender` | 否|int|用户的性别，值为1时是男性，值为2时是女性，值为0时是未知   |
|`city` | 否|string|用户所在城市   |
|`province` | 否|string|用户所在省份   |
|`country` | 否|string|用户所在国家   |
|`language` | 否|string|用户的语言，简体中文为zh_CN   |


#### 请求示例:
```
{
     "avatarUrl": "",
     "city": "",
     "country": "",
     "gender": 0,
     "language": "",
     "nickName": "",
     "province": ""
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

- 更新用户信息

--------------------

### GetQrcode

#### 简要描述：

- [获取微信二维码]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/weixin.get_qrcode

#### 请求方式：

- post

#### 请求参数:

- ` GetQrcodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`page` | 否|string|小程序页面头部   |
|`scene` | 否|string|附带参数   |


#### 请求示例:
```
{
     "page": "",
     "scene": ""
}
```

#### 返回参数说明:

- ` GetQrcodeResp ` : 获取二维码

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`url` | 否|string|二维码图片地址   |


#### 返回示例:
	
```
{
     "url": ""
}
```

#### 备注:

- 获取微信二维码
	

--------------------
--------------------

#### 自定义类型:


