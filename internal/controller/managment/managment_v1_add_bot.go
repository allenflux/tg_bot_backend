package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

const (
	BotStatus = iota
	BotStatusAvailable
	BotStatusUnAvailable
)

const (
	GreetingStatus = iota
	GreetingStatusAvailable
	GreetingStatusUnAvailable
)

func (c *ControllerV1) AddBot(ctx context.Context, req *v1.AddBotReq) (res *v1.AddBotRes, err error) {
	bot := &entity.Bot{
		Account:        req.Account,
		Name:           req.Name,
		Greeting:       req.Greeting,
		GreetingStatus: GreetingStatusAvailable,
		Status:         BotStatusAvailable,
		Photo:          "",
	}
	_, err = dao.Bot.Ctx(ctx).
		Data(bot).
		InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to insert new Bot '%s': %v", req.Name, err)
		return nil, fmt.Errorf("failed to insert new Bot: %w", err)
	}
	return res, err
}
