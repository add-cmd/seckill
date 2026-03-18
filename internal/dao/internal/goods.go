// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsDao is the data access object for the table goods.
type GoodsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GoodsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GoodsColumns defines and stores column names for the table goods.
type GoodsColumns struct {
	Id          string // 商品ID
	GoodsName   string // 商品名称
	GoodsTitle  string // 商品标题
	GoodsImg    string // 商品图片URL
	GoodsDetail string // 商品详情
	GoodsPrice  string // 商品日常标价
	GoodsStock  string // 商品日常库存
	Status      string // 状态：1上架，0下架
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// goodsColumns holds the columns for the table goods.
var goodsColumns = GoodsColumns{
	Id:          "id",
	GoodsName:   "goods_name",
	GoodsTitle:  "goods_title",
	GoodsImg:    "goods_img",
	GoodsDetail: "goods_detail",
	GoodsPrice:  "goods_price",
	GoodsStock:  "goods_stock",
	Status:      "status",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewGoodsDao creates and returns a new DAO object for table data access.
func NewGoodsDao(handlers ...gdb.ModelHandler) *GoodsDao {
	return &GoodsDao{
		group:    "default",
		table:    "goods",
		columns:  goodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GoodsDao) Columns() GoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GoodsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
