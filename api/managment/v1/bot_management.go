package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/model"
)

type GetBotBindListReq struct {
	g.Meta `path:"/bot/bind" tags:"bot" method:"get" summary:"查看绑定关系列表"`
	model.PageReq
}

type BotBindMap struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	NumberOfGroups int    `json:"number_of_groups"`
	NumberOfUsers  int    `json:"number_of_users"`
	GroupIDList    []int  `json:"group_id_list"`
	NumberOfBots   int    `json:"number_of_bots"`
	NumberOfRole   int    `json:"number_of_role"`
}

type GetBotBindListRes struct {
	Data []BotBindMap `json:"data"`
	commonApi.ListRes
}

type AddBotReq struct {
	g.Meta `path:"/bot/bind" tags:"bot" method:"post" summary:"添加机器人"`
	Name   string `json:"name" v:"required"`
	Alias  string `json:"alias" v:"required"`
}

type AddBotRes struct {
}
