// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SeckillGoodsDao is the data access object for the table seckill_goods.
type SeckillGoodsDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SeckillGoodsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SeckillGoodsColumns defines and stores column names for the table seckill_goods.
type SeckillGoodsColumns struct {
	Id           string // 秒杀商品配置ID
	GoodsId      string // 关联的基础商品ID
	SeckillPrice string // 秒杀价
	SeckillStock string // 秒杀专属库存（极其重要）
	StartTime    string // 秒杀开始时间
	EndTime      string // 秒杀结束时间
	Version      string // 并发版本号（预留给乐观锁）
}

// seckillGoodsColumns holds the columns for the table seckill_goods.
var seckillGoodsColumns = SeckillGoodsColumns{
	Id:           "id",
	GoodsId:      "goods_id",
	SeckillPrice: "seckill_price",
	SeckillStock: "seckill_stock",
	StartTime:    "start_time",
	EndTime:      "end_time",
	Version:      "version",
}

// NewSeckillGoodsDao creates and returns a new DAO object for table data access.
func NewSeckillGoodsDao(handlers ...gdb.ModelHandler) *SeckillGoodsDao {
	return &SeckillGoodsDao{
		group:    "default",
		table:    "seckill_goods",
		columns:  seckillGoodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SeckillGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SeckillGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SeckillGoodsDao) Columns() SeckillGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SeckillGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SeckillGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SeckillGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
