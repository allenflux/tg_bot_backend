// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TgUsers is the golang structure of table tg_users for DAO operations like Where/Data.
type TgUsers struct {
	g.Meta    `orm:"table:tg_users, do:true"`
	Id        interface{} //
	TgAccount interface{} //
	GroupId   interface{} //
	TgName    interface{} //
	RoleId    interface{} //
	FirstName interface{} //
	LastName  interface{} //
	IsBot     interface{} //
	Phone     interface{} //
	TgId      interface{} //
}
