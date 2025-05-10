package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) DeleteBot(ctx context.Context, req *v1.DeleteBotReq) (res *v1.DeleteBotRes, err error) {
	if _, err := dao.Bot.Ctx(ctx).Where("id = ?", req.ID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete Bot with Bot=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete Bot with Bot=%d: %w", req.ID, err)
	}
	return res, nil
}
