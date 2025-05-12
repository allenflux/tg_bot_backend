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

func (c *ControllerV1) SwitchBotGreetingStatus(ctx context.Context, req *v1.SwitchBotGreetingStatusReq) (res *v1.SwitchBotGreetingStatusRes, err error) {
	dbQuery := dao.Bot.Ctx(ctx).
		Where("id = ?", req.ID)

	var bots []entity.Bot
	var totalCount int
	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count bots: %v", err)
		return nil, fmt.Errorf("failed to fetch bots list: %w", err)
	}

	if totalCount == 0 {
		g.Log().Errorf(ctx, "Invalid id: %v", err)
		return nil, fmt.Errorf(" Invalid id : %w", err)
	}

	status := consts.GreetingStatus

	switch bots[0].Status {
	case consts.GreetingStatusAvailable:
		status = consts.GreetingStatusUnAvailable
	case consts.GreetingStatusUnAvailable:
		status = consts.GreetingStatusAvailable
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
		g.Log().Errorf(ctx, "Failed to update bots '%s': %v", req.ID, err)
		return nil, fmt.Errorf("failed to update bots: %w", err)
	}
	return res, err
}
