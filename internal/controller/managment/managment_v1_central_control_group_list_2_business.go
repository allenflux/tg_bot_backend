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

func (c *ControllerV1) CentralControlGroupList2Business(ctx context.Context, req *v1.CentralControlGroupList2BusinessReq) (res *v1.CentralControlGroupList2BusinessRes, err error) {
	centralControlId := req.ID

	var groups []entity.Group
	var totalCount int
	if err := dao.Group.Ctx(ctx).Where("central_control_id = ?", centralControlId).Where("type = ?", consts.GroupTypeForBusiness).ScanAndCount(&groups, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count GroupsList: %v", err)
		return nil, fmt.Errorf("failed to fetch group list: %w", err)
	}

	// Prepare response
	res = &v1.CentralControlGroupList2BusinessRes{
		Data: make([]v1.CentralControlGroupListResData, len(groups)),
	}

	// Map data to response structure
	for i, group := range groups {
		res.Data[i] = v1.CentralControlGroupListResData{
			Name: group.Name,
			Link: group.TgLink,
		}
	}
	return res, nil
}
