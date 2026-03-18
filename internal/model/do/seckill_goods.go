// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillGoods is the golang structure of table seckill_goods for DAO operations like Where/Data.
type SeckillGoods struct {
	g.Meta       `orm:"table:seckill_goods, do:true"`
	Id           any         // 秒杀商品配置ID
	GoodsId      any         // 关联的基础商品ID
	SeckillPrice any         // 秒杀价
	SeckillStock any         // 秒杀专属库存（极其重要）
	StartTime    *gtime.Time // 秒杀开始时间
	EndTime      *gtime.Time // 秒杀结束时间
	Version      any         // 并发版本号（预留给乐观锁）
}
