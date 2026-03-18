// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure for table goods.
type Goods struct {
	Id          int64       `json:"id"          orm:"id"           description:"商品ID"`       // 商品ID
	GoodsName   string      `json:"goodsName"   orm:"goods_name"   description:"商品名称"`       // 商品名称
	GoodsTitle  string      `json:"goodsTitle"  orm:"goods_title"  description:"商品标题"`       // 商品标题
	GoodsImg    string      `json:"goodsImg"    orm:"goods_img"    description:"商品图片URL"`    // 商品图片URL
	GoodsDetail string      `json:"goodsDetail" orm:"goods_detail" description:"商品详情"`       // 商品详情
	GoodsPrice  float64     `json:"goodsPrice"  orm:"goods_price"  description:"商品日常标价"`     // 商品日常标价
	GoodsStock  int         `json:"goodsStock"  orm:"goods_stock"  description:"商品日常库存"`     // 商品日常库存
	Status      int         `json:"status"      orm:"status"       description:"状态：1上架，0下架"` // 状态：1上架，0下架
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`       // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:"更新时间"`       // 更新时间
}
