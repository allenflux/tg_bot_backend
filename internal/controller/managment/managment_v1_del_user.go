package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) DelUser(ctx context.Context, req *v1.DelUserReq) (res *v1.DelUserRes, err error) {
	if _, err := dao.Users.Ctx(ctx).Where("id = ?", req.ID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete user with UserID=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete user with ProjectID=%d: %w", req.ID, err)
	}
	return res, nil
}
