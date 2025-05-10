package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) UpdateBot(ctx context.Context, req *v1.UpdateBotReq) (res *v1.UpdateBotRes, err error) {

	dbQuery := dao.Bot.Ctx(ctx).
		Where("id = ?", req.ID)

	var bots []entity.Bot
	var totalCount int
	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count bots: %v", err)
		return nil, fmt.Errorf("failed to fetch bots : %w", err)
	}

	if totalCount == 0 {
		g.Log().Errorf(ctx, "Invalid id: %v", err)
		return nil, fmt.Errorf(" Invalid id : %w", err)
	}
	botInfo := bots[0]
	fillBotInfo(req, &botInfo)

	_, err = dao.CentralControl.Ctx(ctx).
		Data(g.Map{
			"name":      req.Name,
			"account":   req.Account,
			"greeting":  req.Greeting,
			"bot_token": req.BotToken,
		}).
		Where("id = ?", req.ID).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to update new CentralControl '%s': %v", req.ID, err)
		return nil, fmt.Errorf("failed to update new CentralControl: %w", err)
	}
	return res, err
}

func fillEmpty(target *string, source string) {
	if *target == "" {
		*target = source
	}
}

func fillBotInfo(req *v1.UpdateBotReq, botInfo *entity.Bot) {
	fillEmpty(&req.BotToken, botInfo.BotToken)
	fillEmpty(&req.Account, botInfo.Account)
	fillEmpty(&req.Greeting, botInfo.Greeting)
	fillEmpty(&req.Name, botInfo.Name)
}
