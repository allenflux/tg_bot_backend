package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetRoleBotList(ctx context.Context, req *v1.GetRoleBotListReq) (res *v1.GetRoleBotListRes, err error) {
	dbQuery := dao.Bot.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id desc")

	var bots []entity.Bot
	var totalCount int
	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count bots list: %v", err)
		return nil, fmt.Errorf("failed to fetch bots list: %w", err)
	}

	// Prepare response
	res = &v1.GetRoleBotListRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]v1.RoleBotData, len(bots)),
	}

	// Map data to response structure
	for i, bot := range bots {
		res.Data[i] = v1.RoleBotData{
			ID:   bot.Id,
			Name: bot.Name,
		}
	}
	return res, nil
}
