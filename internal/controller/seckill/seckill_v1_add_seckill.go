package seckill

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"seckill/api/seckill/v1"
)

func (c *ControllerV1) AddSeckill(ctx context.Context, req *v1.AddSeckillReq) (res *v1.AddSeckillRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
