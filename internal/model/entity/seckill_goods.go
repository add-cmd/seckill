// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillGoods is the golang structure for table seckill_goods.
type SeckillGoods struct {
	Id           int64       `json:"id"           orm:"id"            description:"秒杀商品配置ID"`      // 秒杀商品配置ID
	GoodsId      int64       `json:"goodsId"      orm:"goods_id"      description:"关联的基础商品ID"`     // 关联的基础商品ID
	SeckillPrice float64     `json:"seckillPrice" orm:"seckill_price" description:"秒杀价"`           // 秒杀价
	SeckillStock int         `json:"seckillStock" orm:"seckill_stock" description:"秒杀专属库存（极其重要）"`  // 秒杀专属库存（极其重要）
	StartTime    *gtime.Time `json:"startTime"    orm:"start_time"    description:"秒杀开始时间"`        // 秒杀开始时间
	EndTime      *gtime.Time `json:"endTime"      orm:"end_time"      description:"秒杀结束时间"`        // 秒杀结束时间
	Version      int         `json:"version"      orm:"version"       description:"并发版本号（预留给乐观锁）"` // 并发版本号（预留给乐观锁）
}
