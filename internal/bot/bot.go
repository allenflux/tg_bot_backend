package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"
)

var AwesomeBotApiChan = make(chan *tgbotapi.BotAPI, 100)

func InitBotApiChanFromMysql(ctx context.Context, payload chan<- *tgbotapi.BotAPI) {
	dbQuery := dao.Bot.Ctx(ctx).
		Order("id desc")
	var bots []entity.Bot
	var totalCount int
	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count bots list: %v", err)
	}
	for _, bot := range bots {
		myBot, err := tgbotapi.NewBotAPI(bot.BotToken)
		if err != nil {
			g.Log().Errorf(ctx, "Failed to init bot: %v", err)
			continue
		}
		myBot.Debug = true
		payload <- myBot
	}
}

func MakeBotApiClientPipLine(ctx context.Context, payload <-chan *tgbotapi.BotAPI) {
	for {
		select {
		case bot, ok := <-payload:
			if !ok {
				g.Log().Error(ctx, "payload channel closed")
				return
			}
			go Program(ctx, bot)
		case <-ctx.Done():
			g.Log().Info(ctx, "BotApiClientPipLine closed")
			return
		}
	}
}

func Program(ctx context.Context, bot *tgbotapi.BotAPI) {
	g.Log().Infof(ctx, "Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
