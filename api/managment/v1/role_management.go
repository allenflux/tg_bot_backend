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
	Bot          string `json:"bot" dc:"关联机器人"`
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
}

type RoleBotData struct {
	Name string `json:"name" dc:"机器人名称"`
	ID   int    `json:"id" dc:"机器人id"`
}
type GetRoleBotListRes struct {
	Data []RoleBotData `json:"data"`
}
