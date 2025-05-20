package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/model"
)

type GetGroupListReq struct {
	g.Meta `path:"/group" tags:"group" method:"get" summary:"查看群组列表"`
	model.PageReq
	KeyWordSearch string `json:"key_word_search" dc:"关键字查询"`
	GroupType     int    `json:"group_type" default:"0" dc:"0 -不筛选 1 - 客户群 2- 渠道群"`
	PlatformID    int    `json:"platform_id" default:"0" dc:"0 -不筛选 如果需要筛选会有单独返回所有平台id和对应名称的接口提供"`
}

type GroupMap struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	PlatformName         string `json:"platform_name"`
	TGLink               string `json:"tg_link"`
	GroupType            int    `json:"group_type"  dc:"1 - 客户群 2- 渠道群"`
	GroupSize            int    `json:"group_size"`
	PlatformIdAndGroupId string `json:"platform_id_and_group_id"`
	AssociatedRobot      int    `json:"associated_robot"`
	AssociatedRole       int    `json:"associated_role"`
}

type GetGroupListRes struct {
	Data []GroupMap `json:"data"`
	commonApi.ListRes
}

type GetGroupMemberDetailReq struct {
	g.Meta `path:"/group/members" tags:"group" method:"get" summary:"获取群组人数"`
	ID     int `json:"id" dc:"Group ID"`
}
type GroupMemberDetail struct {
	TgId   string `json:"tg_id"`
	TgName string `json:"tg_name"`
}
type GetGroupMemberDetailRes struct {
	Data []GroupMemberDetail `json:"data"`
}
