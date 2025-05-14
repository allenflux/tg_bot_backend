package managment

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/frame/g"
	bot2 "tg_bot_backend/internal/bot"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) AddBot(ctx context.Context, req *v1.AddBotReq) (res *v1.AddBotRes, err error) {
	myBot, err := tgbotapi.NewBotAPI(req.BotToken)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to init bot: %v", err)
		return nil, fmt.Errorf(" Failed to init bot: %w", err)
	}
	myBot.Debug = true
	bot2.AwesomeBotApiChan <- myBot

	bot := &entity.Bot{
		Account:        req.Account,
		Name:           req.Name,
		Greeting:       req.Greeting,
		GreetingStatus: consts.GreetingStatusAvailable,
		Status:         consts.BotStatusAvailable,
		Photo:          "",
		BotToken:       req.BotToken,
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
