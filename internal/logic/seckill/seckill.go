package seckill

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"

	"seckill/api/seckill/v1"
	"seckill/internal/dao"
	"seckill/internal/model/entity"
)

func DoSeckill(ctx context.Context, req *v1.DoSeckillReq) (orderId int64, err error) {
	//开起数据库事务：保证扣库存和建订单同时成功或同时失败
	err = dao.SeckillGoods.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		//1.查询秒杀商品信息
		var goods *entity.SeckillGoods
		err := tx.Model(dao.SeckillGoods.Table()).Where(dao.SeckillGoods.Columns().GoodsId, req.GoodsId).Scan(&goods)
		if err != nil || goods == nil {
			return errors.New("秒杀商品不存在")
		}

		//2.校验活动时间
		now := gtime.Now() // 获取当前时间，返回的也是 *gtime.Time 类型
		if goods.StartTime != nil && now.Before(goods.StartTime) {
			return errors.New("秒杀还未开始")
		}
		if goods.EndTime != nil && now.After(goods.EndTime) {
			return errors.New("秒杀已经结束")
		}

		//3.校验库存
		if goods.SeckillStock <= 0 {
			return errors.New("商品已售空")
		}
		//4.尝试扣减库存
		//只要库存>0 才允许扣减，防止并发下扣成负数
		result, err := tx.Model(dao.SeckillGoods.Table()).
			Where(dao.SeckillGoods.Columns().GoodsId, req.GoodsId).
			Where("seckill_stock > 0").
			Decrement(dao.SeckillGoods.Columns().SeckillStock, 1)
		if err != nil {
			return err
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("手慢了，商品已被抢光")
		}
		// 5.写入秒杀订单凭证
		//我们在数据库建了联合唯一索引 uk_user_goods,如果同一个用户并发发了两次请求，第二次走到这里会直接报 Duplicate entry 错误，事务回滚！
		seckillOrder := &entity.SeckillOrder{
			UserId:  req.UserId,
			GoodsId: req.GoodsId,
		}
		_, err = tx.Model(dao.SeckillOrder.Table()).Insert(seckillOrder)
		if err != nil {
			return errors.New("您已经抢购过了，请勿重复下单")
		}

		//6.创建真正的业务订单
		orderInfo := &entity.OrderInfo{
			UserId:     req.UserId,
			GoodsId:    req.GoodsId,
			GoodsPrice: goods.SeckillPrice,
			Status:     0, //新建未支付
		}
		insertRes, err := tx.Model(dao.OrderInfo.Table()).Insert(orderInfo)
		if err != nil {
			return err
		}
		orderId, err = insertRes.LastInsertId()
		return nil
	})
	return orderId, err
}
