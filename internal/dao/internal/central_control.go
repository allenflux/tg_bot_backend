// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CentralControlDao is the data access object for the table central_control.
type CentralControlDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  CentralControlColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// CentralControlColumns defines and stores column names for the table central_control.
type CentralControlColumns struct {
	Id                string //
	Name              string //
	Domain            string //
	NumberOfCustomers string //
	NumberOfBusiness  string //
	Note              string //
	Status            string //
	SecretKey         string //
	ApiUsername       string //
}

// centralControlColumns holds the columns for the table central_control.
var centralControlColumns = CentralControlColumns{
	Id:                "id",
	Name:              "name",
	Domain:            "domain",
	NumberOfCustomers: "number_of_customers",
	NumberOfBusiness:  "number_of_business",
	Note:              "note",
	Status:            "status",
	SecretKey:         "secret_key",
	ApiUsername:       "api_username",
}

// NewCentralControlDao creates and returns a new DAO object for table data access.
func NewCentralControlDao(handlers ...gdb.ModelHandler) *CentralControlDao {
	return &CentralControlDao{
		group:    "default",
		table:    "central_control",
		columns:  centralControlColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CentralControlDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CentralControlDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CentralControlDao) Columns() CentralControlColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CentralControlDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CentralControlDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CentralControlDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
