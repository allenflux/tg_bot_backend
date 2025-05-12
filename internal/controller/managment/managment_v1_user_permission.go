package managment

import (
	"context"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) UserPermission(ctx context.Context, req *v1.UserPermissionReq) (res *v1.UserPermissionRes, err error) {
	res = &v1.UserPermissionRes{
		Data: make([]v1.UserPermissionData, 0),
	}
	data := v1.UserPermissionData{
		Name:     "管理员",
		Describe: "管理员角色",
	}
	res.Data = append(res.Data, data)
	return res, nil
}
