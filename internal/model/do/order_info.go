// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure of table order_info for DAO operations like Where/Data.
type OrderInfo struct {
	g.Meta     `orm:"table:order_info, do:true"`
	Id         any         // 订单主键
	UserId     any         // 用户ID
	GoodsId    any         // 商品ID
	GoodsName  any         // 冗余的商品名称（防商品改名）
	GoodsCount any         // 购买数量
	GoodsPrice any         // 购买时的单价
	Status     any         // 订单状态：0新建未支付，1已支付，2已发货，3已收货，4已退款，5已关闭
	CreateTime *gtime.Time // 订单创建时间
	PayTime    *gtime.Time // 支付时间
}
