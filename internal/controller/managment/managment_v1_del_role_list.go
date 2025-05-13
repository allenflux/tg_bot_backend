package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) DelRoleList(ctx context.Context, req *v1.DelRoleListReq) (res *v1.DelRoleListRes, err error) {
	dbQuery := dao.Role.Ctx(ctx).
		Where("id = ?", req.ID)
	var roles []entity.Role
	var totalCount int
	if err = dbQuery.ScanAndCount(&roles, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and roles Group List: %v", err)
		return nil, fmt.Errorf("failed to fetch roles list: %w", err)
	}

	if totalCount == 0 {
		return &v1.DelRoleListRes{}, nil
	}

	if roles[0].UserSize != 0 {
		g.Log().Errorf(ctx, "The Role has been bound, so it cannot be deleted: %v", err)
		return nil, fmt.Errorf("the Role has been bound, so it cannot be deleted: %w", err)
	}

	if _, err := dao.Role.Ctx(ctx).Where("id = ?", req.ID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete role with ID=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete role with ID=%d: %w", req.ID, err)
	}
	return res, nil
}
