// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	Id         int64       `json:"id"         orm:"id"          description:"订单主键"`                                 // 订单主键
	UserId     int64       `json:"userId"     orm:"user_id"     description:"用户ID"`                                 // 用户ID
	GoodsId    int64       `json:"goodsId"    orm:"goods_id"    description:"商品ID"`                                 // 商品ID
	GoodsName  string      `json:"goodsName"  orm:"goods_name"  description:"冗余的商品名称（防商品改名）"`                       // 冗余的商品名称（防商品改名）
	GoodsCount int         `json:"goodsCount" orm:"goods_count" description:"购买数量"`                                 // 购买数量
	GoodsPrice float64     `json:"goodsPrice" orm:"goods_price" description:"购买时的单价"`                               // 购买时的单价
	Status     int         `json:"status"     orm:"status"      description:"订单状态：0新建未支付，1已支付，2已发货，3已收货，4已退款，5已关闭"` // 订单状态：0新建未支付，1已支付，2已发货，3已收货，4已退款，5已关闭
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"订单创建时间"`                               // 订单创建时间
	PayTime    *gtime.Time `json:"payTime"    orm:"pay_time"    description:"支付时间"`                                 // 支付时间
}
