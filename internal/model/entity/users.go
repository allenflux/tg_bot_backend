// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id         int         `json:"id"         orm:"id"          description:""` //
	Account    string      `json:"account"    orm:"account"     description:""` //
	Name       string      `json:"name"       orm:"name"        description:""` //
	Role       int         `json:"role"       orm:"role"        description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	Password   string      `json:"password"   orm:"password"    description:""` //
}
