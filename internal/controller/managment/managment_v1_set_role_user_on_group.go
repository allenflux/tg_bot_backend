package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) SetRoleUserOnGroup(ctx context.Context, req *v1.SetRoleUserOnGroupReq) (res *v1.SetRoleUserOnGroupRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// Store all affected role IDs to update user count later
		affectedRoleIDs := map[int]struct{}{req.RoleID: {}}

		for _, groupID := range req.GroupID {
			var user entity.TgUsers

			// Fetch user by group ID and account
			err = dao.TgUsers.Ctx(ctx).TX(tx).
				Where("group_id = ? AND tg_account = ?", groupID, req.TgUser).
				Scan(&user)
			if err != nil {
				g.Log().Errorf(ctx, "Failed to query TgUsers for group_id [%d], user [%s]: %v", groupID, req.TgUser, err)
				return fmt.Errorf("failed to query TgUsers: %w", err)
			}

			// Track the original role ID for later role update
			if user.RoleId != 0 {
				affectedRoleIDs[user.RoleId] = struct{}{}
			}

			// Update user's role
			_, err = dao.TgUsers.Ctx(ctx).TX(tx).
				Data(g.Map{"role_id": req.RoleID}).
				Where("group_id = ? AND tg_account = ?", groupID, req.TgUser).
				Update()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to update role for user [%s] in group [%d]: %v", req.TgUser, groupID, err)
				return fmt.Errorf("failed to update TgUsers: %w", err)
			}
		}

		// Update role user count for all affected roles
		for roleID := range affectedRoleIDs {
			userCount, err := dao.TgUsers.Ctx(ctx).TX(tx).
				Where("role_id = ?", roleID).
				Count()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to count users for role_id [%d]: %v", roleID, err)
				return fmt.Errorf("failed to count users for role: %w", err)
			}

			_, err = dao.Role.Ctx(ctx).TX(tx).
				Data(g.Map{"user_size": userCount}).
				Where("id = ?", roleID).
				Update()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to update user_size for role_id [%d]: %v", roleID, err)
				return fmt.Errorf("failed to update role user size: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v1.SetRoleUserOnGroupRes{}, nil
}
