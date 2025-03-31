// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CentralControl is the golang structure of table central_control for DAO operations like Where/Data.
type CentralControl struct {
	g.Meta            `orm:"table:central_control, do:true"`
	Id                interface{} //
	Name              interface{} //
	Domain            interface{} //
	NumberOfCustomers interface{} //
	NumberOfBusiness  interface{} //
	Note              interface{} //
	Status            interface{} //
	SecretKey         interface{} //
}
