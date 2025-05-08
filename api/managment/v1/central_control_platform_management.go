package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/model"
)

type CentralControlListReq struct {
	g.Meta `path:"/central-control" tags:"central-control" method:"get" summary:"查看中控管理平台列表"`
	model.PageReq
	KeyWordSearch string `json:"key_word_search" dc:"关键字查询"`
}

type CentralControlMap struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Domain            string `json:"domain"`
	SecretKey         string `json:"secret_key"`
	NumberOfCustomers int    `json:"number_of_customers"`
	NumberOfBusiness  int    `json:"number_of_business"`
	Note              string `json:"note"`
	Status            int    `json:"status" dc:"1-可用 2-禁用"`
}

type CentralControlListRes struct {
	Data []CentralControlMap `json:"data"`
	commonApi.ListRes
}

type UpdateCentralControlReq struct {
	g.Meta    `path:"/central-control" tags:"central-control" method:"post" summary:"添加中控平台"`
	Name      string `json:"name" v:"required"`
	Domain    string `json:"domain" v:"required"`
	SecretKey string `json:"secret_key" v:"required"`
	Note      string `json:"note" `
}

type UpdateCentralControlRes struct {
}

type PutCentralControlReq struct {
	g.Meta `path:"/central-control" tags:"central-control" method:"put" summary:"编辑中控平台"`
	ID     int    `json:"id" v:"required"`
	Name   string `json:"name" v:"required"`
	Note   string `json:"note" `
}

type PutCentralControlRes struct {
}

type DelCentralControlReq struct {
	g.Meta `path:"/central-control" tags:"central-control" method:"delete" summary:"删除中控平台"`
	ID     int `json:"id" v:"required"`
}

type DelCentralControlRes struct {
}

type CentralControlGroupListResData struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type CentralControlGroupListReq struct {
	g.Meta `path:"/central-control/customer/group" tags:"central-control" method:"get" summary:"查看中控管理平台列表中的客户群信息"`
	ID     int `json:"id" v:"required"`
}

type CentralControlGroupListRes struct {
	Data []CentralControlGroupListResData `json:"data"`
}

type CentralControlGroupList2BusinessReq struct {
	g.Meta `path:"/central-control/business/group" tags:"central-control" method:"get" summary:"查看中控管理平台列表中的渠道群信息"`
	ID     int `json:"id" v:"required"`
}

type CentralControlGroupList2BusinessRes struct {
	Data []CentralControlGroupListResData `json:"data"`
}

type CentralControlStatusSwitchReq struct {
	g.Meta `path:"/central-control/status/switch" tags:"central-control" method:"get" summary:"更新中控状态平台"`
	ID     int `json:"id" v:"required"`
}

type CentralControlStatusSwitchRes struct{}
