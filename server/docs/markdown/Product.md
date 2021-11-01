

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Product]
- [Waiting to write...]

--------------------

### ReplaceProduct

#### 简要描述：

- [添加一个产品(可能更新)]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.replace_product

#### 请求方式：

- post

#### 请求参数:

- ` AddReq ` : ...

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`name` | `是`|string|商品名   |
|`type` | `是`|string|商品类型   |
|`price` | `是`|int64|商品价格(分)   |
|`qty` | 否|int64|商品数量   |
|`contexts` | 否|[]`product.PContext`|商品详情   |
|`imgs` | 否|[]string|轮播图(图片，或者视频都可以)   |
|`icon` | 否|string|商品图标   |


#### 请求示例:
```
{
     "contexts": [
          {
               "context": "",
               "type": 0
          }
     ],
     "icon": "",
     "imgs": [
          ""
     ],
     "name": "",
     "price": 0,
     "qty": 0,
     "type": ""
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

- 添加一个产品(可能更新)

--------------------

### ReplaceSku

#### 简要描述：

- [添加一个产品Sku(可能更新)]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.replace_sku

#### 请求方式：

- post

#### 请求参数:

- ` AddSkuReq ` : sku

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | `是`|string|商品gpid   |
|`skus` | `是`|[]`product.SkuInfo`|商品sku信息   |


#### 请求示例:
```
{
     "gpid": "",
     "skus": [
          {
               "premium": 0,
               "tag_name": "",
               "type_name": ""
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

- 添加一个产品Sku(可能更新)

--------------------

### UpdateSelect

#### 简要描述：

- [修改商品推荐信息]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.update_select

#### 请求方式：

- post

#### 请求参数:

- ` UpdateSelectReq ` : 更新推荐

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | `是`|string|商品gpid   |
|`isSelect` | `是`|bool|是否推荐   |


#### 请求示例:
```
{
     "gpid": "",
     "isSelect": false
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

- 修改商品推荐信息

--------------------

### AddToBuyTmpCart

#### 简要描述：

- [添加到直接购买虚拟购物车里面]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.add_to_buy_tmp_cart

#### 请求方式：

- post

#### 请求参数:

- ` AddToBuyTmpCartReq ` : 直接购买添加到购物车里

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品id   |
|`number` | 否|int32|数量(目前默认为1)   |
|`skuList` | 否|[]int32|sku信息   |


#### 请求示例:
```
{
     "gpid": "",
     "number": 0,
     "skuList": [
          0
     ]
}
```

#### 返回参数说明:

- ` AddToBuyTmpCartResp ` : 添加到虚拟购物车时返回id

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|   |


#### 返回示例:
	
```
{
     "id": 0
}
```

#### 备注:

- 添加到直接购买虚拟购物车里面

--------------------

### GetFavorite

#### 简要描述：

- [获取收藏列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_favorite

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

- ` GetFavoriteResp ` : 获取收藏列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`items` | 否|[]`caoguo.SampleProductInfo`|收藏列表   |


#### 返回示例:
	
```
{
     "items": [
          {
               "gpid": "",
               "icon": "",
               "img": [
                    ""
               ],
               "name": "",
               "originalPrice": 0,
               "percent": 0,
               "price": 0,
               "sales": 0
          }
     ]
}
```

#### 备注:

- 获取收藏列表

--------------------

### GetProductList

#### 简要描述：

- [admin 获取商品列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_product_list

#### 请求方式：

- post

#### 请求参数:

- ` GetProductListReq ` : admin 获取商品列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid (为空则获取列表)   |
|`index` | 否|int32|取第几页(默认0)   |


#### 请求示例:
```
{
     "gpid": "",
     "index": 0
}
```

#### 返回参数说明:

- ` GetProductListResp ` : admin 获取商品列表返回值

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`productList` | 否|[]`caoguo.Product`|商品列表   |
|`page` | 否|`common.Page`|总共   |


#### 返回示例:
	
```
{
     "page": {
          "index": 0,
          "limit": 0,
          "size": 0
     },
     "productList": [
          {
               "PlatformID": "",
               "couponList": [
                    {
                         "couponType": 0,
                         "denom": 0,
                         "describle": "",
                         "gpid": "",
                         "greatMoney": 0,
                         "groupName": "",
                         "id": 0,
                         "status": 0,
                         "title": "",
                         "type": 0,
                         "vaild": false,
                         "validity": 0
                    }
               ],
               "createdAt": "",
               "distAmount": 0,
               "gpid": "",
               "gtid": "",
               "icon": "",
               "img": [
                    ""
               ],
               "imgDetail": [
                    ""
               ],
               "isSelect": false,
               "name": "",
               "originalPrice": 0,
               "price": 0,
               "qty": 0,
               "sales": 0,
               "service": "",
               "shareTit": "",
               "shipmentFee": 0,
               "skuPrice": [
                    {
                         "distAmount": 0,
                         "id": 0,
                         "premium": 0,
                         "skuList": ""
                    }
               ],
               "skus": [
                    {
                         "id": 0,
                         "tagName": "",
                         "typeName": ""
                    }
               ],
               "updatedAt": "",
               "vaild": false,
               "videoDetail": [
                    ""
               ]
          }
     ]
}
```

#### 备注:

- admin 获取商品列表

--------------------

### GetAdminAdInfo

#### 简要描述：

- [admin 获取首页信息(轮播图广告，类型广告，主销广告,精选列表)]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_admin_ad_info

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

- ` GetAdminAdInfo ` : Resp admin 获取相关广告信息(可更新或者添加)

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`adList` | 否|[]`caoguo.AdminAdInfo`|广告相关信息   |
|`groupList` | 否|[]`caoguo.AdminAdGroupInfo`|首页精选列表   |


#### 返回示例:
	
```
{
     "adList": [
          {
               "attach": "",
               "id": 0,
               "img": "",
               "sortId": 0,
               "type": 0,
               "url": "",
               "vaild": false
          }
     ],
     "groupList": [
          {
               "id": 0,
               "mainGpid": "",
               "sortId": 0,
               "subGpid": "",
               "vaild": false
          }
     ]
}
```

#### 备注:

- admin 获取首页信息(轮播图广告，类型广告，主销广告,精选列表)

--------------------

### GetCartList

#### 简要描述：

- [获取购物车内容]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_cart_list

#### 请求方式：

- post

#### 请求参数:

- ` GetCartListReq ` : 获取购物车

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetCartListResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`items` | 否|[]`caoguo.CartInfo`|购物车列表   |


#### 返回示例:
	
```
{
     "items": [
          {
               "icon": "",
               "id": 0,
               "name": "",
               "number": 0,
               "originalPrice": 0,
               "price": 0,
               "skuVal": "",
               "stock": 0
          }
     ]
}
```

#### 备注:

- 获取购物车内容

--------------------

### GetProductByType

#### 简要描述：

- [通过商品类型获取商品列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_product_by_type

#### 请求方式：

- post

#### 请求参数:

- ` GetProductByTypeReq ` : 通过类型获取商品列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`typeId` | 否|string|类型   |
|`pageNumber` | 否|int32|当前页数   |


#### 请求示例:
```
{
     "pageNumber": 0,
     "typeId": ""
}
```

#### 返回参数说明:

- ` GetProductByTypeResp ` : 通过类型获取商品列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`typeId` | 否|string|类型   |
|`pageNumber` | 否|int32|当前页数   |
|`items` | 否|[]`caoguo.SampleProductInfo`|当前列表   |


#### 返回示例:
	
```
{
     "items": [
          {
               "gpid": "",
               "icon": "",
               "img": [
                    ""
               ],
               "name": "",
               "originalPrice": 0,
               "percent": 0,
               "price": 0,
               "sales": 0
          }
     ],
     "pageNumber": 0,
     "typeId": ""
}
```

#### 备注:

- 通过商品类型获取商品列表

--------------------

### UpsetAdminAdInfo

#### 简要描述：

- [admin 更新首页信息(轮播图广告，类型广告，主销广告,精选列表)(id为0,添加，id&gt;0 删除)]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.upset_admin_ad_info

#### 请求方式：

- post

#### 请求参数:

- ` GetAdminAdInfo ` : Resp admin 获取相关广告信息(可更新或者添加)

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`adList` | 否|[]`caoguo.AdminAdInfo`|广告相关信息   |
|`groupList` | 否|[]`caoguo.AdminAdGroupInfo`|首页精选列表   |


#### 请求示例:
```
{
     "adList": [
          {
               "attach": "",
               "id": 0,
               "img": "",
               "sortId": 0,
               "type": 0,
               "url": "",
               "vaild": false
          }
     ],
     "groupList": [
          {
               "id": 0,
               "mainGpid": "",
               "sortId": 0,
               "subGpid": "",
               "vaild": false
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

- admin 更新首页信息(轮播图广告，类型广告，主销广告,精选列表)(id为0,添加，id&gt;0 删除)

--------------------

### UpsetSkuPrice

#### 简要描述：

- [更新/设置产品sku价格]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.upset_sku_price

#### 请求方式：

- post

#### 请求参数:

- ` UpsetSkuPriceReq ` : admin 更新或设置商品价格

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|用于更新单价及分享收益信息,0:表示没有sku属性时   |
|`gpid` | 否|string|商品gpid,(更新单个sku价格，gpid可以为空)   |
|`premium` | 否|int64|商品单价   |
|`distAmount` | 否|int64|分享收益   |
|`skuList` | 否|[]int32|属性列表   |
|`tag` | 否|int32|0:添加，1:更新，-1：删除(只允许删除sku price)   |


#### 请求示例:
```
{
     "distAmount": 0,
     "gpid": "",
     "id": 0,
     "premium": 0,
     "skuList": [
          0
     ],
     "tag": 0
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

- 更新/设置产品sku价格

--------------------

### AddAddress

#### 简要描述：

- [添加或者修改地址]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.add_address

#### 请求方式：

- post

#### 请求参数:

- ` AddAddressReq ` : 添加或者修改地址

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`addr` | 否|`caoguo.AddressInfo`|地址   |
|`op` | 否|int32|1:添加，2:修改，3:删除   |


#### 请求示例:
```
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
     "op": 0
}
```

#### 返回参数说明:

- ` int64 ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 添加或者修改地址

--------------------

### ChangeCart

#### 简要描述：

- [改变]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.change_cart

#### 请求方式：

- post

#### 请求参数:

- ` ChangeCardReq ` : 修改数量 or 删除购物车

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|id   |
|`number` | 否|int64|商品数量 (0:删除)   |


#### 请求示例:
```
{
     "id": 0,
     "number": 0
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

- 改变

--------------------

### GetAddress

#### 简要描述：

- [获取地址]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_address

#### 请求方式：

- post

#### 请求参数:

- ` GetAddressReq ` : 获取地址

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`isDefault` | 否|bool|1:获取默认地址   |


#### 请求示例:
```
{
     "isDefault": false
}
```

#### 返回参数说明:

- ` GetAddressResp ` : 返回地址信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`addrs` | 否|[]`caoguo.AddressInfo`|地址   |


#### 返回示例:
	
```
{
     "addrs": [
          {
               "address": "",
               "addressName": "",
               "area": "",
               "default": false,
               "id": 0,
               "mobile": "",
               "name": ""
          }
     ]
}
```

#### 备注:

- 获取地址

--------------------

### AddToCart

#### 简要描述：

- [加入购物车]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.add_to_cart

#### 请求方式：

- post

#### 请求参数:

- ` AddCartReq ` : 添加购物车

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品id   |
|`number` | 否|int32|数量   |
|`skuList` | 否|[]int32|sku信息   |


#### 请求示例:
```
{
     "gpid": "",
     "number": 0,
     "skuList": [
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

- 加入购物车

--------------------

### GetProductDetails

#### 简要描述：

- [获取订单详情]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_product_details

#### 请求方式：

- post

#### 请求参数:

- ` GetProductDetailsReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |


#### 请求示例:
```
{
     "gpid": ""
}
```

#### 返回参数说明:

- ` GetProductDetailsRsp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`info` | 否|`caoguo.ProductInfo`|商品信息   |


#### 返回示例:
	
```
{
     "info": {
          "bzList": [
               ""
          ],
          "conList": [
               ""
          ],
          "couponName": "",
          "couponPct": 0,
          "gpid": "",
          "icon": "",
          "img": [
               ""
          ],
          "imgDetail": [
               ""
          ],
          "isFavorite": false,
          "name": "",
          "originalPrice": 0,
          "price": 0,
          "richText": "",
          "sales": 0,
          "shareTit": "",
          "sku": {
               "gpid": "",
               "groups": [
                    {
                         "Items": [
                              {
                                   "id": 0,
                                   "tagName": "",
                                   "typeName": ""
                              }
                         ],
                         "typeName": ""
                    }
               ]
          },
          "skuPrice": {
               "gpid": "",
               "list": [
                    {
                         "distAmount": 0,
                         "id": 0,
                         "premium": 0,
                         "skuList": ""
                    }
               ]
          },
          "stock": "",
          "vdDetail": [
               ""
          ],
          "views": ""
     }
}
```

#### 备注:

- 获取订单详情

--------------------

### Favorite

#### 简要描述：

- [收藏 or 取消收藏]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.favorite

#### 请求方式：

- post

#### 请求参数:

- ` FavoriteReq ` : 收藏&amp;amp;取消收藏

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品id   |
|`isFavorite` | 否|bool|是否收藏   |


#### 请求示例:
```
{
     "gpid": "",
     "isFavorite": false
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

- 收藏 or 取消收藏

--------------------

### GetAdInfo

#### 简要描述：

- [获取广告主信息列表(可能更新)]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_ad_info

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

- ` GetAdResp ` : 获取ad广告内容

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`rotationList` | 否|[]`caoguo.AdInfo`|广告轮播图信息   |
|`typeList` | 否|[]`caoguo.AdInfo`|广告页类型信息   |
|`masterInfo` | 否|`caoguo.AdInfo`|主销广告   |


#### 返回示例:
	
```
{
     "masterInfo": {
          "attach": "",
          "img": "",
          "sortId": 0,
          "url": ""
     },
     "rotationList": [
          {
               "attach": "",
               "img": "",
               "sortId": 0,
               "url": ""
          }
     ],
     "typeList": [
          {
               "attach": "",
               "img": "",
               "sortId": 0,
               "url": ""
          }
     ]
}
```

#### 备注:

- 获取广告主信息列表(可能更新)

--------------------

### GetBoutiqueGroup

#### 简要描述：

- [获取团购列表]

#### 请求URL:

- http://localhost:9001/commcn/api/v1/product.get_boutique_group

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

- ` BoutiqueGroupResp ` : 团购列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`groups` | 否|[]`caoguo.BoutiqueGroupInfo`|团购值   |
|`likes` | 否|[]`caoguo.SampleProductInfo`|推荐列表   |


#### 返回示例:
	
```
{
     "groups": [
          {
               "mainProduct": {
                    "gpid": "",
                    "icon": "",
                    "img": [
                         ""
                    ],
                    "name": "",
                    "originalPrice": 0,
                    "percent": 0,
                    "price": 0,
                    "sales": 0
               },
               "subProduct": {
                    "gpid": "",
                    "icon": "",
                    "img": [
                         ""
                    ],
                    "name": "",
                    "originalPrice": 0,
                    "percent": 0,
                    "price": 0,
                    "sales": 0
               }
          }
     ],
     "likes": [
          {
               "gpid": "",
               "icon": "",
               "img": [
                    ""
               ],
               "name": "",
               "originalPrice": 0,
               "percent": 0,
               "price": 0,
               "sales": 0
          }
     ]
}
```

#### 备注:

- 获取团购列表
	

--------------------
--------------------

#### 自定义类型:

#### ` caoguo `


- ` SampleProductInfo ` : 简单商品信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`name` | 否|string|商品名   |
|`price` | 否|int64|商品价格   |
|`originalPrice` | 否|int64|商品原始价格   |
|`sales` | 否|int64|商品销量   |
|`img` | 否|[]string|商品轮播图   |
|`icon` | 否|string|商品图标   |
|`percent` | 否|int64|团购进度条   |



- ` SkuInfo ` : sku信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|id   |
|`typeName` | 否|string|类型名字   |
|`tagName` | 否|string|标签名称   |



- ` SkuPriceGroup ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|id(用于更新单价及分享收益信息)   |
|`skuList` | 否|string|属性列表(逗号小到大排序)   |
|`premium` | 否|int64|商品单价   |
|`distAmount` | 否|int64|分享收益   |



- ` Coupon ` : 优惠券信息

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
|`status` | 否|int32|0：默认，1:未使用(有效),2:已使用,-1:已过期   |
|`type` | 否|int32|优惠券类型，1:促销优惠券，2：用户已领取的优惠券   |
|`vaild` | 否|bool|是否有效   |



- ` Product ` : 商品信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`name` | 否|string|商品名   |
|`gtid` | 否|string|商品分类   |
|`price` | 否|int64|商品价格(分)   |
|`originalPrice` | 否|int64|商品原始价格(分)   |
|`distAmount` | 否|int64|分享收益（-----todo:更新）   |
|`PlatformID` | 否|string|商铺客户id   |
|`qty` | 否|int64|数量(库存)   |
|`shipmentFee` | 否|int64|运费(分)(------todo:更新)   |
|`imgDetail` | 否|[]string|图文详情   |
|`videoDetail` | 否|[]string|视频详情   |
|`vaild` | 否|bool|是否有效   |
|`updatedAt` | 否|string|更新时间   |
|`createdAt` | 否|string|创建时间   |
|`sales` | 否|int64|已购买数量(销量)// 商品附加信息   |
|`img` | 否|[]string|商品轮播图   |
|`icon` | 否|string|商品图标   |
|`service` | 否|string|商品服务   |
|`shareTit` | 否|string|分享提示   |
|`isSelect` | 否|bool|是否精选(推荐)   |
|`skus` | 否|[]`caoguo.SkuInfo`|商品属性   |
|`skuPrice` | 否|[]`caoguo.SkuPriceGroup`|商品价格明细   |
|`couponList` | 否|[]`caoguo.Coupon`|商品优惠价   |



- ` AdminAdInfo ` : 广告页信息信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|id(id&amp;gt;0:修改，id=0:添加)   |
|`url` | 否|string|页面跳转 (页内跳转)   |
|`img` | 否|string|展示图片   |
|`sortId` | 否|int32|排序(越大越前)   |
|`type` | 否|int32|类型(1:轮播图广告，2:类型广告，3:主销广告)   |
|`attach` | 否|string|附加信息,type 对应关系(1:rgb(23,42,8),2:类型名字,3:空)   |
|`vaild` | 否|bool|是否有效   |



- ` AdminAdGroupInfo ` : 首页精选列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|id(id&amp;gt;0:修改，id=0:添加)   |
|`mainGpid` | 否|string|主商品id   |
|`subGpid` | 否|string|附加商品id   |
|`sortId` | 否|int32|排序(越大越前)   |
|`vaild` | 否|bool|是否有效   |



- ` CartInfo ` : 购物车信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`id` | 否|int64|id   |
|`icon` | 否|string|icon   |
|`skuVal` | 否|string|sku信息   |
|`stock` | 否|int64|库存   |
|`name` | 否|string|商品名字   |
|`price` | 否|int64|商品价格   |
|`number` | 否|int64|商品数量   |
|`originalPrice` | 否|int64|商品原始价格   |



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



- ` SkuGroup ` : sku grup信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`typeName` | 否|string|类型名字   |
|`Items` | 否|[]`caoguo.SkuInfo`|sku信息列表   |



- ` ProducktSkuInfo ` : 产品列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`groups` | 否|[]`caoguo.SkuGroup`|sku信息列表   |



- ` ProducktSkuPriceInfo ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`list` | 否|[]`caoguo.SkuPriceGroup`|sku信息列表   |



- ` ProductInfo ` : 商品信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`gpid` | 否|string|商品gpid   |
|`img` | 否|[]string|商品轮播图   |
|`name` | 否|string|商品名   |
|`price` | 否|int64|商品价格   |
|`originalPrice` | 否|int64|商品原始价格   |
|`couponPct` | 否|int32|折扣占比   |
|`sales` | 否|int64|商品销量   |
|`stock` | 否|string|库存   |
|`views` | 否|string|浏览量   |
|`shareTit` | 否|string|分享提示   |
|`sku` | 否|`caoguo.ProducktSkuInfo`|商品分类   |
|`skuPrice` | 否|`caoguo.ProducktSkuPriceInfo`|商品溢价明细   |
|`couponName` | 否|string|优惠券   |
|`conList` | 否|[]string|促销活动信息   |
|`bzList` | 否|[]string|服务   |
|`imgDetail` | 否|[]string|图文详情   |
|`vdDetail` | 否|[]string|视频详情   |
|`richText` | 否|string|富文本信息   |
|`icon` | 否|string|商品图标   |
|`isFavorite` | 否|bool|是否收藏   |



- ` AdInfo ` : 广告页信息信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`url` | 否|string|页面跳转   |
|`img` | 否|string|卡片展示图片   |
|`sortId` | 否|int32|排序(越大越前)   |
|`attach` | 否|string|附加   |



- ` BoutiqueGroupInfo ` : 精品商品团购列表

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`mainProduct` | 否|`caoguo.SampleProductInfo`|主推荐   |
|`subProduct` | 否|`caoguo.SampleProductInfo`|附加推荐   |



#### ` common `


- ` Page ` : 分页

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`size` | 否|int32|每页尺寸   |
|`index` | 否|int32|第几页   |
|`limit` | 否|int32|总页数   |



#### ` product `


- ` PContext ` : 商品内容字段

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`type` | 否|int|类型，1:文字，2:图片，3：视频   |
|`context` | 否|string|内容   |



- ` SkuInfo ` : sku 信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`type_name` | `是`|string|标签类型   |
|`tag_name` | `是`|string|标签名字   |
|`premium` | `是`|int64|溢价值   |




