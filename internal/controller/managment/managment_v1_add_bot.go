package managment

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	bot2 "tg_bot_backend/internal/bot"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) AddBot(ctx context.Context, req *v1.AddBotReq) (res *v1.AddBotRes, err error) {
	if ok, err := dao.Bot.Ctx(ctx).Where("bot_token = ?", req.BotToken).Exist(); err != nil {
		g.Log().Errorf(ctx, "Failed to check if [%s] exists: %v", req.BotToken, err)
		return nil, fmt.Errorf(" Failed to check if [%s] exists: %v", req.BotToken, err)
	} else if ok {
		g.Log().Errorf(ctx, "Bot [%s] already exists", req.BotToken)
		return nil, fmt.Errorf(" Bot [%s] already exists", req.BotToken)
	}

	myBot, err := tgbotapi.NewBotAPI(req.BotToken)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to init bot: %v", err)
		return nil, fmt.Errorf(" Failed to init bot: %w", err)
	}
	myBot.Debug = true
	bot2.AwesomeBotApiChan <- myBot

	bot := &entity.Bot{
		Account:        strconv.FormatInt(myBot.Self.ID, 10),
		Name:           myBot.Self.UserName,
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
		g.Log().Errorf(ctx, "Failed to insert new Bot '%s': %v", req.BotToken, err)
		return nil, fmt.Errorf("failed to insert new Bot: %w", err)
	}
	return res, err
}
