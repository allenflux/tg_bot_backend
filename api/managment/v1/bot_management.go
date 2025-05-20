package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/model"
)

//type GetBotBindListReq struct {
//	g.Meta `path:"/bot/bind" tags:"bot" method:"get" summary:"查看绑定关系列表"`
//	model.PageReq
//}

//type BotBindMap struct {
//	ID             int    `json:"id"`
//	Account        string `json:"account"`
//	Name           string `json:"name"`
//	NumberOfGroups int    `json:"number_of_groups"`
//	NumberOfUsers  int    `json:"number_of_users"`
//	GroupIDList    []int  `json:"group_id_list"`
//	NumberOfBots   int    `json:"number_of_bots"`
//	NumberOfRole   int    `json:"number_of_role"`
//}

//type GetBotBindListRes struct {
//	Data []BotBindMap `json:"data"`
//	commonApi.ListRes
//}

type AddBotReq struct {
	g.Meta `path:"/bot" tags:"bot" method:"post" summary:"添加机器人"`
	//Name     string `json:"name" v:"required" dc:"机器人名称"`
	//Account  string `json:"account" v:"required" dc:"机器人账号"`
	Greeting string `json:"greeting" dc:"欢迎语"`
	BotToken string `json:"bot_token" v:"required" dc:"机器人token"`
}

type AddBotRes struct {
}

type DeleteBotReq struct {
	g.Meta `path:"/bot" tags:"bot" method:"delete" summary:"删除机器人"`
	ID     int `json:"id" v:"required" dc:"<机器人>ID"`
}
type DeleteBotRes struct{}

type UpdateBotReq struct {
	g.Meta `path:"/bot" tags:"bot" method:"put" summary:"更新机器人"`
	ID     int `json:"id" v:"required" dc:"<机器人>ID"`
	//Name     string `json:"name"  dc:"机器人名称"`
	//Account  string `json:"account"  dc:"机器人账号"`
	Greeting string `json:"greeting" dc:"欢迎语"`
	//BotToken string `json:"bot_token" dc:"机器人token"`
}
type UpdateBotRes struct{}

type GetBot struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Account        string `json:"account"`
	Greeting       string `json:"greeting"`
	BotToken       string `json:"bot_token"`
	GreetingStatus int    `json:"greetingStatus"`
	Status         int    `json:"status"`
	Photo          string
}
type GetBotListReq struct {
	g.Meta `path:"/bot" tags:"bot" method:"get" summary:"获取机器人列表"`
	model.PageReq
	KeyWordSearch string `json:"key_word_search" dc:"关键字查询"`
}
type GetBotListRes struct {
	Data []GetBot `json:"data"`
	commonApi.ListRes
}

type SwitchBotStatusReq struct {
	g.Meta `path:"/bot/status" tags:"bot" method:"get" summary:"切换机器人状态"`
	ID     int `json:"id" v:"required" dc:"<机器人>ID"`
}

type SwitchBotStatusRes struct{}

type SwitchBotGreetingStatusReq struct {
	g.Meta `path:"/bot/greeting/status" tags:"bot" method:"get" summary:"切换机器人欢迎语状态"`
	ID     int `json:"id" v:"required" dc:"<机器人>ID"`
}

type SwitchBotGreetingStatusRes struct{}
