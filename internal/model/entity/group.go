// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Group is the golang structure for table group.
type Group struct {
	Id               int    `json:"id"               orm:"id"                 description:""` //
	Name             string `json:"name"             orm:"name"               description:""` //
	CentralControlId int    `json:"centralControlId" orm:"central_control_id" description:""` //
	TgLink           string `json:"tgLink"           orm:"tg_link"            description:""` //
	Type             int    `json:"type"             orm:"type"               description:""` //
	Size             int    `json:"size"             orm:"size"               description:""` //
	BotSize          int    `json:"botSize"          orm:"bot_size"           description:""` //
	RoleSize         int    `json:"roleSize"         orm:"role_size"          description:""` //
	GroupChatId      string `json:"groupChatId"      orm:"group_chat_id"      description:""` //
}
