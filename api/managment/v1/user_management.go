package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/model"
)

type GetUserListReq struct {
	g.Meta `path:"/user" tags:"user" method:"get" summary:"获取账号列表"`
	model.PageReq
	KeyWordSearch string `json:"key_word_search" dc:"关键字查询"`
}

type UserMap struct {
	ID         int         `json:"id"`
	Account    string      `json:"account"`
	Name       string      `json:"name"`
	Role       int         `json:"role"`
	Status     int         `json:"status" dc:"1-禁用 2-可用"`
	CreateTime *gtime.Time `json:"create_time"`
}

type GetUserListRes struct {
	Data []UserMap `json:"data"`
	commonApi.ListRes
}

type AddUserReq struct {
	g.Meta   `path:"/user" tags:"user" method:"post" summary:"新增账号"`
	Account  string `json:"account" v:"required"`
	Name     string `json:"name" v:"required"`
	Password string `json:"password" v:"required"`
	Role     int    `json:"role" v:"required"`
}

type AddUserRes struct {
}

type EditUserReq struct {
	g.Meta   `path:"/user" tags:"user" method:"put" summary:"编辑账号"`
	ID       int    `json:"id" v:"required"`
	Account  string `json:"account" v:"required"`
	Name     string `json:"name" v:"required"`
	Password string `json:"password" `
	Role     string `json:"role" v:"required"`
}

type EditUserRes struct {
}

type DelUserReq struct {
	g.Meta `path:"/user" tags:"user" method:"delete" summary:"删除账号"`
	ID     int `json:"id" v:"required"`
}

type DelUserRes struct {
}
