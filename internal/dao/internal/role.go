// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleDao is the data access object for the table role.
type RoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RoleColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RoleColumns defines and stores column names for the table role.
type RoleColumns struct {
	Id       string //
	Name     string //
	BotId    string //
	Cmd      string //
	UserSize string //
	BotName  string //
}

// roleColumns holds the columns for the table role.
var roleColumns = RoleColumns{
	Id:       "id",
	Name:     "name",
	BotId:    "bot_id",
	Cmd:      "cmd",
	UserSize: "user_size",
	BotName:  "bot_name",
}

// NewRoleDao creates and returns a new DAO object for table data access.
func NewRoleDao(handlers ...gdb.ModelHandler) *RoleDao {
	return &RoleDao{
		group:    "default",
		table:    "role",
		columns:  roleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoleDao) Columns() RoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
