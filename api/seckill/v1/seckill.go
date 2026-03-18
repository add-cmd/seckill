package v1

import "github.com/gogf/gf/v2/frame/g"

// 1. 后台：添加秒杀商品配置 Req
type AddSeckillReq struct {
	// g.Meta 定义了路由路径、请求方法和 Swagger 文档信息，极其方便
	g.Meta       `path:"/admin/seckill/add" method:"post" tags:"后台管理" summary:"添加秒杀配置"`
	GoodsId      int64   `json:"goodsId" v:"required#商品ID不能为空"`
	SeckillPrice float64 `json:"seckillPrice" v:"required#秒杀价不能为空"`
	SeckillStock int     `json:"seckillStock" v:"required|min:1#秒杀库存不能小于1"`
	StartTime    string  `json:"startTime" v:"required#开始时间不能为空"`
	EndTime      string  `json:"endTime" v:"required#结束时间不能为空"`
}

// 添加成功后不需要返回特别的数据，留空即可
type AddSeckillRes struct{}

// 2. 前台：用户点击抢购 Req
type DoSeckillReq struct {
	g.Meta  `path:"/app/seckill/do" method:"post" tags:"前台用户" summary:"执行秒杀抢购"`
	GoodsId int64 `json:"goodsId" v:"required#商品ID不能为空"`
	// V1 版本为了方便测试，我们让前端直接传 userId。
	// 在未来的工业级标准中，这里应该是无参数的，userId 会通过网关层的 JWT Token 鉴权后拦截注入到 Context 中。
	UserId int64 `json:"userId" v:"required#用户ID不能为空"`
}

// 抢购成功后，需要返回生成的订单号给前端
type DoSeckillRes struct {
	OrderId int64 `json:"orderId" dc:"创建成功的订单ID"`
}
