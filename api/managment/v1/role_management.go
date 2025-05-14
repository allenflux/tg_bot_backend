package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/model"
)

type GetRoleListReq struct {
	g.Meta `path:"/role" tags:"role" method:"get" summary:"查看角色列表"`
	model.PageReq
}

type RoleMap struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Bot          string `json:"bot" dc:"关联机器人名字"`
	Cmd          string `json:"cmd" dc:"可调用指令"`
	NumberOfBind int    `json:"number_of_bind"`
}

type GetRoleListRes struct {
	Data []RoleMap `json:"data"`
	commonApi.ListRes
}

type DelRoleListReq struct {
	g.Meta `path:"/role" tags:"role" method:"delete" summary:"删除角色数据"`
	ID     int `json:"id" v:"required"`
}

type DelRoleListRes struct {
}

type AddRoleListReq struct {
	g.Meta `path:"/role" tags:"role" method:"post" summary:"新增角色数据"`
	Name   string   `json:"name" v:"required" dc:"角色名称"`
	BotID  int      `json:"bot_id" v:"required"`
	Cmd    []string `json:"cmd" dc:"可调用指令,从/role/cmd" v:"required"`
}

type AddRoleListRes struct {
}

type GetBotCmdListReq struct {
	g.Meta `path:"/role/cmd" tags:"role" method:"get" summary:"获取新增角色时使用的命令"`
}

type BotCmdData struct {
	Name     string `json:"name" dc:"只提交这个字段的数组就好"`
	Describe string `json:"describe" dc:"<UNK>"`
}
type GetBotCmdListRes struct {
	Data []BotCmdData `json:"data"`
}

type GetRoleBotListReq struct {
	g.Meta `path:"/role/bot" tags:"role" method:"get" summary:"获取新增角色时使用的机器人列表"`
	model.PageReq
}

type RoleBotData struct {
	Name string `json:"name" dc:"机器人名称"`
	ID   int    `json:"id" dc:"机器人id"`
}
type GetRoleBotListRes struct {
	Data []RoleBotData `json:"data"`
	commonApi.ListRes
}

type SetRoleUserOnGroupReq struct {
	g.Meta  `path:"/role/bind" tags:"role" method:"post" summary:"一键绑定"`
	TgUser  string `json:"tg_user" v:"required" dc:"tg user ID《搜索到的用户的tg账号》"`
	GroupID []int  `json:"group_id" v:"required" dc:"group ID"`
	RoleID  int    `json:"role_id" v:"required" dc:"role ID"`
}

type SetRoleUserOnGroupRes struct{}

type SearchTgUserAccountReq struct {
	g.Meta `path:"/role/tg/account" tags:"role" method:"get" summary:"搜索tg用户信息"`
	model.PageReq
	KeyWordSearch string `json:"key_word_search" dc:"关键字查询"`
}

type SearchTgUserAccountData struct {
	Name          string `json:"name" dc:"用户名"`
	TgUserAccount string `json:"tg_user_account"  dc:"tg user account<提交时候的参数>"`
}
type SearchTgUserAccountRes struct {
	Data []SearchTgUserAccountData `json:"data"`
	commonApi.ListRes
}

type GetTgUserGroupListReq struct {
	g.Meta    `path:"/role/tg/group" tags:"role" method:"get" summary:"tg用户信息group信息"`
	TgAccount string `p:"account" dc:"<UNK>" v:"required"`
}

type TgUserGroupData struct {
	Name       string `json:"name" dc:"<UNK>"`
	ID         int    `json:"id" dc:"<最终提交的group数组的入参>id"`
	PlatformID int    `json:"platform_id" dc:"<<UNK>ID>"`
}

type GetTgUserGroupListRes struct {
	Data map[string][]TgUserGroupData `json:"data"`
}

type GetBindTgUsersReq struct {
	g.Meta `path:"/role/tg/bind" tags:"role" method:"get" summary:"获取已绑定的用户"`
	RoleId int `json:"role_id" v:"required" dc:"role ID"`
}
type BindTgUsersData struct {
	Name      string `json:"name" dc:"<UNK>"`
	TgAccount string `json:"account" dc:"<UNK>"`
	ID        int    `json:"id" dc:"<UNK>id"`
}

type GetBindTgUsersRes struct {
	Data []BindTgUsersData `json:"data"`
}
