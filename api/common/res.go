package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/consts"
)

// EmptyRes 不响应任何数据
type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

// ListRes 列表公共返回
type ListRes struct {
	CurrentPage      int `json:"current_page" dc:"当前页" `
	Total            int `json:"total" dc:"总数"`
	UnconditionalNum int `json:"unconditional_num" dc:"无条件查询的总数"`
}

type CommonRes struct {
	Status  consts.EnumResponseStatus `json:"status" dc:"响应状态0成功2警告" default:"0"`
	Message string                    `json:"message" dc:"响应消息"`
}
