package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) DelRoleList(ctx context.Context, req *v1.DelRoleListReq) (res *v1.DelRoleListRes, err error) {
	if _, err := dao.Role.Ctx(ctx).Where("id = ?", req.ID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete role with ID=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete role with ID=%d: %w", req.ID, err)
	}
	return res, nil
}
