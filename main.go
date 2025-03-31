package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	_ "tg_bot_backend/internal/packed"

	//_ "tg_bot_backend/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"tg_bot_backend/internal/cmd"
)

func main() {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")
	cmd.Main.Run(gctx.GetInitCtx())
}
