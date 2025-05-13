// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Role is the golang structure for table role.
type Role struct {
	Id       int    `json:"id"       orm:"id"        description:""` //
	Name     string `json:"name"     orm:"name"      description:""` //
	BotId    int    `json:"botId"    orm:"bot_id"    description:""` //
	Cmd      string `json:"cmd"      orm:"cmd"       description:""` //
	UserSize int    `json:"userSize" orm:"user_size" description:""` //
	BotName  string `json:"botName"  orm:"bot_name"  description:""` //
}
