package main

import (
	_ "tg_bot_backend/internal/packed"

	_ "tg_bot_backend/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"tg_bot_backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
