// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SeckillOrder is the golang structure of table seckill_order for DAO operations like Where/Data.
type SeckillOrder struct {
	g.Meta  `orm:"table:seckill_order, do:true"`
	Id      any // 秒杀订单凭证ID
	UserId  any // 用户ID
	OrderId any // 关联的真实订单ID
	GoodsId any // 秒杀商品ID
}
