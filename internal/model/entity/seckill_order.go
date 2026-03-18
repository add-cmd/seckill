// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SeckillOrder is the golang structure for table seckill_order.
type SeckillOrder struct {
	Id      int64 `json:"id"      orm:"id"       description:"秒杀订单凭证ID"`  // 秒杀订单凭证ID
	UserId  int64 `json:"userId"  orm:"user_id"  description:"用户ID"`      // 用户ID
	OrderId int64 `json:"orderId" orm:"order_id" description:"关联的真实订单ID"` // 关联的真实订单ID
	GoodsId int64 `json:"goodsId" orm:"goods_id" description:"秒杀商品ID"`    // 秒杀商品ID
}
