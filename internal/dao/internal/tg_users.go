// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgUsersDao is the data access object for the table tg_users.
type TgUsersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TgUsersColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TgUsersColumns defines and stores column names for the table tg_users.
type TgUsersColumns struct {
	Id        string //
	TgAccount string //
	GroupId   string //
	TgName    string //
	RoleId    string //
}

// tgUsersColumns holds the columns for the table tg_users.
var tgUsersColumns = TgUsersColumns{
	Id:        "id",
	TgAccount: "tg_account",
	GroupId:   "group_id",
	TgName:    "tg_name",
	RoleId:    "role_id",
}

// NewTgUsersDao creates and returns a new DAO object for table data access.
func NewTgUsersDao(handlers ...gdb.ModelHandler) *TgUsersDao {
	return &TgUsersDao{
		group:    "default",
		table:    "tg_users",
		columns:  tgUsersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TgUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TgUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TgUsersDao) Columns() TgUsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TgUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TgUsersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TgUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
