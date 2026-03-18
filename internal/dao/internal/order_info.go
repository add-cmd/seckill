// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderInfoDao is the data access object for the table order_info.
type OrderInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OrderInfoColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OrderInfoColumns defines and stores column names for the table order_info.
type OrderInfoColumns struct {
	Id         string // 订单主键
	UserId     string // 用户ID
	GoodsId    string // 商品ID
	GoodsName  string // 冗余的商品名称（防商品改名）
	GoodsCount string // 购买数量
	GoodsPrice string // 购买时的单价
	Status     string // 订单状态：0新建未支付，1已支付，2已发货，3已收货，4已退款，5已关闭
	CreateTime string // 订单创建时间
	PayTime    string // 支付时间
}

// orderInfoColumns holds the columns for the table order_info.
var orderInfoColumns = OrderInfoColumns{
	Id:         "id",
	UserId:     "user_id",
	GoodsId:    "goods_id",
	GoodsName:  "goods_name",
	GoodsCount: "goods_count",
	GoodsPrice: "goods_price",
	Status:     "status",
	CreateTime: "create_time",
	PayTime:    "pay_time",
}

// NewOrderInfoDao creates and returns a new DAO object for table data access.
func NewOrderInfoDao(handlers ...gdb.ModelHandler) *OrderInfoDao {
	return &OrderInfoDao{
		group:    "default",
		table:    "order_info",
		columns:  orderInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrderInfoDao) Columns() OrderInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrderInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
