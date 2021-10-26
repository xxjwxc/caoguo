package product

// PContext 商品内容字段
type PContext struct {
	Type    int    `json:"type"`    // 类型，1:文字，2:图片，3：视频
	Context string `json:"context"` // 内容
}

// AddReq ...
type AddReq struct {
	Name     string     `json:"name" binding:"required"`  // 商品名
	GType    string     `json:"type" binding:"required"`  // 商品类型
	Price    int64      `json:"price" binding:"required"` // 商品价格(分)
	Qty      int64      `json:"qty"`                      // 商品数量
	Contexts []PContext `json:"contexts"`                 // 商品详情
	Image    []string   `json:"imgs"`                     // 轮播图(图片，或者视频都可以)
	Icon     string     `json:"icon"`                     // 商品图标
	// MaxBuyQty int        `json:"max_buy_qty"` // 最大购买数量
}

// SkuInfo sku 信息
type SkuInfo struct {
	TypeName string `json:"type_name" binding:"required"` // 标签类型
	TagName  string `json:"tag_name" binding:"required"`  // 标签名字
	Premium  int64  `json:"premium" binding:"required"`   // 溢价值
}

// AddSkuReq sku
type AddSkuReq struct {
	Gpid string    `json:"gpid"  binding:"required"` // 商品gpid
	Skus []SkuInfo `json:"skus" binding:"required"`  // 商品sku信息
}

// UpdateSelectReq 更新推荐
type UpdateSelectReq struct {
	Gpid     string `json:"gpid"  binding:"required"`    // 商品gpid
	IsSelect bool   `json:"isSelect" binding:"required"` // 是否推荐
}

type ReqGetContext struct {
}

const (
	_ProductPageLen = 10 // 一次最大请求
)
