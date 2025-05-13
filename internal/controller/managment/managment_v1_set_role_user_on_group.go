package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) SetRoleUserOnGroup(ctx context.Context, req *v1.SetRoleUserOnGroupReq) (res *v1.SetRoleUserOnGroupRes, err error) {
	for _, value := range req.GroupID {
		_, err = dao.TgUsers.Ctx(ctx).
			Data(g.Map{
				"role_id": req.RoleID,
			}).
			Where("group_id = ?", value).
			Where("tg_account = ?", req.TgUser).
			Update()
		if err != nil {
			g.Log().Errorf(ctx, "Failed to update  TgUsers '%s': %v", req.TgUser, err)
			return nil, fmt.Errorf("failed to update  TgUsers: %w", err)
		}
	}

	// Update Role Table
	roleUserCount, err := dao.TgUsers.Ctx(ctx).
		Where("role_id = ?", req.RoleID).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to count  TgUsers '%s': %v", req.TgUser, err)
		return nil, fmt.Errorf("failed to count  TgUsers: %w", err)
	}

	_, err = dao.Role.Ctx(ctx).
		Data(g.Map{
			"user_size": roleUserCount,
		}).
		Where("id = ?", req.RoleID).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to update  Role '%s': %v", req.TgUser, err)
		return nil, fmt.Errorf("failed to update  Role: %w", err)
	}
	return res, nil
}
