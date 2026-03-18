// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure of table goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta      `orm:"table:goods, do:true"`
	Id          any         // 商品ID
	GoodsName   any         // 商品名称
	GoodsTitle  any         // 商品标题
	GoodsImg    any         // 商品图片URL
	GoodsDetail any         // 商品详情
	GoodsPrice  any         // 商品日常标价
	GoodsStock  any         // 商品日常库存
	Status      any         // 状态：1上架，0下架
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
