package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"tg_bot_backend/internal/bot"
	_ "tg_bot_backend/internal/packed"

	"tg_bot_backend/internal/cmd"
)

func init() {
	g.Log().SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)
}

func main() {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")
	ctx := gctx.GetInitCtx()
	//g.Redis().Del(ctx, "-4957481659")
	bot.InitBotApiChanFromMysql(ctx, bot.AwesomeBotApiChan)
	go bot.MakeBotApiClientPipLine(ctx, bot.AwesomeBotApiChan)
	go bot.MakeTgGroupPipLine(ctx, bot.AwesomeGroupChan)
	cmd.Main.Run(ctx)
}
