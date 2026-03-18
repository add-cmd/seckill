package seckill

import (
	"context"
	"seckill/internal/logic/seckill"

	"seckill/api/seckill/v1"
)

func (c *ControllerV1) DoSeckill(ctx context.Context, req *v1.DoSeckillReq) (res *v1.DoSeckillRes, err error) {
	// 调用业务逻辑层
	orderId, err := seckill.DoSeckill(ctx, req)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.DoSeckillRes{
		OrderId: orderId,
	}, nil
}
