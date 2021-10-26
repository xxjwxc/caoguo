

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Login]
- [Waiting to write...]

--------------------

### Login

#### 简要描述：

- [登录]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/login/login

#### 请求方式：

- POST

#### 请求参数:

- ` LoginReq ` : 登陆请求

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`username` | `是`|string|用户名   |
|`password` | `是`|string|密码(用户输入的md5值,大写)   |


#### 请求示例:
```
{
     "password": "",
     "username": ""
}
```

#### 返回参数说明:


#### 返回示例:
	
```
{}
```

#### 备注:

- 登录

--------------------

### RefreshToken

#### 简要描述：

- [刷新token管理员]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/login.refresh_token

#### 请求方式：

- post

#### 请求参数:

- ` RefreshReq ` : reflashToken

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`refresh_token` | 否|string|token   |


#### 请求示例:
```
{
     "refresh_token": ""
}
```

#### 返回参数说明:


#### 返回示例:
	
```
{}
```

#### 备注:

- 刷新token管理员
	

--------------------
--------------------

#### 自定义类型:


