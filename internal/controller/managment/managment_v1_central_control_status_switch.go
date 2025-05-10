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

func (c *ControllerV1) CentralControlStatusSwitch(ctx context.Context, req *v1.CentralControlStatusSwitchReq) (res *v1.CentralControlStatusSwitchRes, err error) {
	dbQuery := dao.CentralControl.Ctx(ctx).
		Where("id = ?", req.ID)

	var centrals []entity.CentralControl
	var totalCount int
	if err := dbQuery.ScanAndCount(&centrals, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count CentralControlList: %v", err)
		return nil, fmt.Errorf("failed to fetch CentralControl list: %w", err)
	}

	if totalCount == 0 {
		g.Log().Errorf(ctx, "Invalid id: %v", err)
		return nil, fmt.Errorf(" Invalid id : %w", err)
	}

	status := consts.CentralControlStatus

	switch centrals[0].Status {
	case consts.CentralControlStatusAvailable:
		status = consts.CentralControlStatusUnAvailable
	case consts.CentralControlStatusUnAvailable:
		status = consts.CentralControlStatusAvailable
	default:
		panic("unhandled default case")
	}

	_, err = dao.CentralControl.Ctx(ctx).
		Data(g.Map{
			"status": status,
		}).
		Where("id = ?", req.ID).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to update new CentralControl '%s': %v", req.ID, err)
		return nil, fmt.Errorf("failed to update new CentralControl: %w", err)
	}
	return res, err
}
