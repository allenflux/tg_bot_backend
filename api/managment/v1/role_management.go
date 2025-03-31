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
	Name   string `json:"name" v:"required"`
	BotID  int    `json:"bot_id" v:"required"`
	Cmd    []int  `json:"cmd" dc:"可调用指令" v:"required"`
}

type AddRoleListRes struct {
}
