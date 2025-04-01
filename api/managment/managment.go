// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package managment

import (
	"context"

	"tg_bot_backend/api/managment/v1"
)

type IManagmentV1 interface {
	GetBotBindList(ctx context.Context, req *v1.GetBotBindListReq) (res *v1.GetBotBindListRes, err error)
	AddBot(ctx context.Context, req *v1.AddBotReq) (res *v1.AddBotRes, err error)
	CentralControlList(ctx context.Context, req *v1.CentralControlListReq) (res *v1.CentralControlListRes, err error)
	UpdateCentralControl(ctx context.Context, req *v1.UpdateCentralControlReq) (res *v1.UpdateCentralControlRes, err error)
	PutCentralControl(ctx context.Context, req *v1.PutCentralControlReq) (res *v1.PutCentralControlRes, err error)
	DelCentralControl(ctx context.Context, req *v1.DelCentralControlReq) (res *v1.DelCentralControlRes, err error)
	GetGroupList(ctx context.Context, req *v1.GetGroupListReq) (res *v1.GetGroupListRes, err error)
	GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error)
	DelRoleList(ctx context.Context, req *v1.DelRoleListReq) (res *v1.DelRoleListRes, err error)
	AddRoleList(ctx context.Context, req *v1.AddRoleListReq) (res *v1.AddRoleListRes, err error)
	GetRoute(ctx context.Context, req *v1.GetRouteReq) (res *v1.GetRouteRes, err error)
	GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error)
	AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error)
	EditUser(ctx context.Context, req *v1.EditUserReq) (res *v1.EditUserRes, err error)
	DelUser(ctx context.Context, req *v1.DelUserReq) (res *v1.DelUserRes, err error)
	LoginUser(ctx context.Context, req *v1.LoginUserReq) (res *v1.LoginUserRes, err error)
}
