// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package managment

import (
	"context"

	"tg_bot_backend/api/managment/v1"
)

type IManagmentV1 interface {
	AddBot(ctx context.Context, req *v1.AddBotReq) (res *v1.AddBotRes, err error)
	DeleteBot(ctx context.Context, req *v1.DeleteBotReq) (res *v1.DeleteBotRes, err error)
	UpdateBot(ctx context.Context, req *v1.UpdateBotReq) (res *v1.UpdateBotRes, err error)
	GetBotList(ctx context.Context, req *v1.GetBotListReq) (res *v1.GetBotListRes, err error)
	SwitchBotStatus(ctx context.Context, req *v1.SwitchBotStatusReq) (res *v1.SwitchBotStatusRes, err error)
	SwitchBotGreetingStatus(ctx context.Context, req *v1.SwitchBotGreetingStatusReq) (res *v1.SwitchBotGreetingStatusRes, err error)
	CentralControlList(ctx context.Context, req *v1.CentralControlListReq) (res *v1.CentralControlListRes, err error)
	UpdateCentralControl(ctx context.Context, req *v1.UpdateCentralControlReq) (res *v1.UpdateCentralControlRes, err error)
	PutCentralControl(ctx context.Context, req *v1.PutCentralControlReq) (res *v1.PutCentralControlRes, err error)
	DelCentralControl(ctx context.Context, req *v1.DelCentralControlReq) (res *v1.DelCentralControlRes, err error)
	CentralControlGroupList(ctx context.Context, req *v1.CentralControlGroupListReq) (res *v1.CentralControlGroupListRes, err error)
	CentralControlGroupList2Business(ctx context.Context, req *v1.CentralControlGroupList2BusinessReq) (res *v1.CentralControlGroupList2BusinessRes, err error)
	CentralControlStatusSwitch(ctx context.Context, req *v1.CentralControlStatusSwitchReq) (res *v1.CentralControlStatusSwitchRes, err error)
	GetPlatformList(ctx context.Context, req *v1.GetPlatformListReq) (res *v1.GetPlatformListRes, err error)
	GetGroupList(ctx context.Context, req *v1.GetGroupListReq) (res *v1.GetGroupListRes, err error)
	GetGroupMemberDetail(ctx context.Context, req *v1.GetGroupMemberDetailReq) (res *v1.GetGroupMemberDetailRes, err error)
	GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error)
	DelRoleList(ctx context.Context, req *v1.DelRoleListReq) (res *v1.DelRoleListRes, err error)
	AddRoleList(ctx context.Context, req *v1.AddRoleListReq) (res *v1.AddRoleListRes, err error)
	GetBotCmdList(ctx context.Context, req *v1.GetBotCmdListReq) (res *v1.GetBotCmdListRes, err error)
	GetRoleBotList(ctx context.Context, req *v1.GetRoleBotListReq) (res *v1.GetRoleBotListRes, err error)
	SetRoleUserOnGroup(ctx context.Context, req *v1.SetRoleUserOnGroupReq) (res *v1.SetRoleUserOnGroupRes, err error)
	SearchTgUserAccount(ctx context.Context, req *v1.SearchTgUserAccountReq) (res *v1.SearchTgUserAccountRes, err error)
	GetTgUserGroupList(ctx context.Context, req *v1.GetTgUserGroupListReq) (res *v1.GetTgUserGroupListRes, err error)
	GetBindTgUsers(ctx context.Context, req *v1.GetBindTgUsersReq) (res *v1.GetBindTgUsersRes, err error)
	GetRoute(ctx context.Context, req *v1.GetRouteReq) (res *v1.GetRouteRes, err error)
	GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error)
	AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error)
	EditUser(ctx context.Context, req *v1.EditUserReq) (res *v1.EditUserRes, err error)
	DelUser(ctx context.Context, req *v1.DelUserReq) (res *v1.DelUserRes, err error)
	LoginUser(ctx context.Context, req *v1.LoginUserReq) (res *v1.LoginUserRes, err error)
	UserStatusSwitch(ctx context.Context, req *v1.UserStatusSwitchReq) (res *v1.UserStatusSwitchRes, err error)
	UserPermission(ctx context.Context, req *v1.UserPermissionReq) (res *v1.UserPermissionRes, err error)
}
