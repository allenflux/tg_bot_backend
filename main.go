package main

import (
	_ "tg_bot_backend/internal/packed"

	//_ "tg_bot_backend/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"tg_bot_backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
