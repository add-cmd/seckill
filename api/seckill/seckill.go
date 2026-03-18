// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package seckill

import (
	"context"

	"seckill/api/seckill/v1"
)

type ISeckillV1 interface {
	AddSeckill(ctx context.Context, req *v1.AddSeckillReq) (res *v1.AddSeckillRes, err error)
	DoSeckill(ctx context.Context, req *v1.DoSeckillReq) (res *v1.DoSeckillRes, err error)
}
