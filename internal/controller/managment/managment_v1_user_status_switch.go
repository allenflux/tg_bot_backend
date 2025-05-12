package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) UserStatusSwitch(ctx context.Context, req *v1.UserStatusSwitchReq) (res *v1.UserStatusSwitchRes, err error) {
	dbQuery := dao.Users.Ctx(ctx).
		Where("id = ?", req.ID)

	var users []entity.Users
	var totalCount int
	if err := dbQuery.ScanAndCount(&users, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count users: %v", err)
		return nil, fmt.Errorf("failed to fetch users list: %w", err)
	}

	if totalCount == 0 {
		g.Log().Errorf(ctx, "Invalid id: %v", err)
		return nil, fmt.Errorf(" Invalid id : %w", err)
	}

	status := consts.UserStatus

	switch users[0].Status {
	case consts.UserAvailable:
		status = consts.UserUnAvailable
	case consts.UserUnAvailable:
		status = consts.UserAvailable
	default:
		panic("unhandled default case")
	}

	_, err = dao.Users.Ctx(ctx).
		Data(g.Map{
			"status": status,
		}).
		Where("id = ?", req.ID).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to update  Users '%s': %v", req.ID, err)
		return nil, fmt.Errorf("failed to update  Users: %w", err)
	}
	return res, err
}
