package oplog

import (
	"time"

	"github.com/xxjwxc/public/tools"
)

const (
	contentCount int = 10
)

//
type SearchParam struct {
	Page_num   int        `json:"page_num"`             //第几页
	Text       string     `json:"text"`                 //条件
	Start_time tools.Time `json:"start_time,omitempty"` //
	End_time   tools.Time `json:"end_time,omitempty"`   //
}

//
type OpLogInfo struct {
	Operate_id  string    `json:"operate_id"`  //操作ID
	Operator    string    `json:"operator"`    //操作人
	Topic       string    `json:"topic"`       //操作topic
	Bundle      string    `json:"bundle"`      //操作bundle
	Create_time time.Time `json:"create_time"` //操作时间
}

//
type MineCenter struct {
	Type   string              `json:"type"`   //订单、点赞、收藏和全部
	Count  []int               `json:"count"`  //
	Recent *Wx_user_log_view   `json:"recent"` //最新的
	List   []*Wx_user_log_view `json:"list"`   //
}

//
type GetCenterParam struct {
	SessionId string `json:"sessionId"`
	Type      int    `json:"type"` //类型 1：收藏 2：点赞 3：订单 0：全部
}

//
type Wx_user_log_view struct {
	Id           int       `gorm:"primary_key" json:"-"`
	Open_id      string    `json:"open_id"`                 //
	Goods_id     string    `json:"goods_id"`                //
	Goods_name   string    `json:"goods_name"`              //
	Type_name    string    `json:"type_name"`               //商品类型
	Pic_url      []string  `gorm:"-" json:"pic_url"`        //商品图片
	Pic_url_str  string    `gorm:"column:pic_url" json:"-"` //商品图片
	Order_id     string    `json:"order_id"`                //
	Type         int       `json:"type"`                    //操作类型（1：收藏 2：点赞 3：订单）
	Create_time  time.Time `json:"create_time"`             //
	Code         string    `json:"code"`                    //商品码
	Default_url  string    `json:"default_url"`             //默认icon（文字、音视频url）
	Content_type string    `json:"content_type"`            //发布内容类型
}
