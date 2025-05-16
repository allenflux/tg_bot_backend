// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TgUsers is the golang structure for table tg_users.
type TgUsers struct {
	Id        int    `json:"id"        orm:"id"         description:""` //
	TgAccount string `json:"tgAccount" orm:"tg_account" description:""` //
	GroupId   int    `json:"groupId"   orm:"group_id"   description:""` //
	TgName    string `json:"tgName"    orm:"tg_name"    description:""` //
	RoleId    int    `json:"roleId"    orm:"role_id"    description:""` //
	FirstName string `json:"firstName" orm:"first_name" description:""` //
	LastName  string `json:"lastName"  orm:"last_name"  description:""` //
	IsBot     int    `json:"isBot"     orm:"is_bot"     description:""` //
	Phone     string `json:"phone"     orm:"phone"      description:""` //
	TgId      string `json:"tgId"      orm:"tg_id"      description:""` //
}
