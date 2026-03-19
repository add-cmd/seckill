package seckill

import (
	"context"
	"errors"
	"fmt"
	"seckill/api/seckill/v1"
	"seckill/internal/dao"
	"seckill/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// 定义神级 Lua 脚本：极其严谨的原子操作
const luaScript = `
local stockKey = KEYS[1]
local userKey = KEYS[2]
local userId = ARGV[1]

-- 1. 判断是否重复购买
if redis.call('sismember', userKey, userId) == 1 then
    return 2 
end

-- 2. 判断库存是否充足
local stock = tonumber(redis.call('get', stockKey))
if stock == nil or stock <= 0 then
    return 1 
end

-- 3. 扣库存并记录用户
redis.call('decr', stockKey)
redis.call('sadd', userKey, userId)
return 0
`

func DoSeckill(ctx context.Context, req *v1.DoSeckillReq) (orderId int64, err error) {
	// 定义 Redis 的 Key
	stockKey := fmt.Sprintf("seckill:stock:%d", req.GoodsId)
	userKey := fmt.Sprintf("seckill:users:%d", req.GoodsId)

	// ==========================================
	// ⚡️ 第一阶段：Redis 内存层拦截 (10 万并发在这里被挡下)
	// ==========================================
	result, err := g.Redis().Eval(ctx, luaScript, 2, []string{stockKey, userKey}, []any{req.UserId})
	if err != nil {
		g.Log().Error(ctx, "Redis 执行 Lua 失败", err)
		return 0, errors.New("服务器开小差了，请稍后再试")
	}

	code := result.Int()
	if code == 1 {
		return 0, errors.New("手慢了，商品已被抢光")
	}
	if code == 2 {
		return 0, errors.New("您已经抢购过了，请勿重复下单")
	}

	// ==========================================
	// 💾 第二阶段：MySQL 落盘阶段 (只有极少数幸运儿能走到这里)
	// 此时并发量已经被严格限制在了商品库存数以内 (例如 100 个)，MySQL 毫无压力
	// ==========================================

	// 在真实的工业级 V3.0 中，走到这里通常是直接向 MQ 发送一条消息就返回给前端了。
	// 目前 V2.0 我们先同步写库，因为现在的并发量 MySQL 已经完全吃得消了。

	err = dao.SeckillGoods.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 获取商品信息 (主要为了拿到价格生成订单)
		var goods *entity.SeckillGoods
		tx.Model(dao.SeckillGoods.Table()).Where(dao.SeckillGoods.Columns().GoodsId, req.GoodsId).Scan(&goods)

		// 2. 扣减 MySQL 的真实库存 (兜底)
		_, err = tx.Model(dao.SeckillGoods.Table()).
			Where(dao.SeckillGoods.Columns().GoodsId, req.GoodsId).
			Decrement(dao.SeckillGoods.Columns().SeckillStock, 1)
		if err != nil {
			return err
		}

		// 3. 写入防刷凭证 (双重保险)
		seckillOrder := &entity.SeckillOrder{
			UserId:  req.UserId,
			GoodsId: req.GoodsId,
		}
		_, err = tx.Model(dao.SeckillOrder.Table()).Insert(seckillOrder)
		if err != nil {
			return err
		}

		// 4. 生成真实订单
		orderInfo := &entity.OrderInfo{
			UserId:     req.UserId,
			GoodsId:    req.GoodsId,
			GoodsCount: 1,
			GoodsPrice: goods.SeckillPrice,
			Status:     0,
		}
		insertRes, err := tx.Model(dao.OrderInfo.Table()).Insert(orderInfo)
		if err != nil {
			return err
		}

		orderId, _ = insertRes.LastInsertId()
		return nil
	})

	return orderId, err
}
