package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) EditUser(ctx context.Context, req *v1.EditUserReq) (res *v1.EditUserRes, err error) {
	if _, err := dao.Users.Ctx(ctx).
		Data(g.Map{
			"account":  req.Account,
			"password": req.Password,
			"name":     req.Name,
			"role":     req.Role,
		}).
		Where("id = ?", req.ID).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update user with ID=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Log success and prepare response
	g.Log().Infof(ctx, "Successfully updated user with ID=%d", req.ID)
	return res, nil
}
