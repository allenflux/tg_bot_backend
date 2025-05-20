package managment

import (
	"context"
	"tg_bot_backend/internal/consts"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetBotCmdList(ctx context.Context, req *v1.GetBotCmdListReq) (res *v1.GetBotCmdListRes, err error) {
	newBotCmd := func(name string) v1.BotCmdData {
		return v1.BotCmdData{
			Name:     name,
			Describe: consts.BotCmdDesc[name],
		}
	}

	res = &v1.GetBotCmdListRes{
		Data: []v1.BotCmdData{
			newBotCmd(consts.BotCmdBind),
			//newBotCmd(consts.BotCmdTopUp),
			newBotCmd(consts.BotCmdUnbind),
		},
	}
	return res, nil
}
