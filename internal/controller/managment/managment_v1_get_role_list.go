package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error) {
	dbQuery := dao.Role.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id desc")

	var roles []entity.Role
	var totalCount int
	if err = dbQuery.ScanAndCount(&roles, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and roles Group List: %v", err)
		return nil, fmt.Errorf("failed to fetch roles list: %w", err)
	}

	// Prepare response
	res = &v1.GetRoleListRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]v1.RoleMap, len(roles)),
	}

	// Map data to response structure
	for i, v := range roles {
		res.Data[i] = v1.RoleMap{
			ID:           v.Id,
			Name:         v.Name,
			Bot:          v.BotName,
			Cmd:          v.Cmd,
			NumberOfBind: v.UserSize,
		}
	}
	return res, err
}
