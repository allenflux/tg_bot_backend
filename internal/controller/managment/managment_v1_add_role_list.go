package managment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) AddRoleList(ctx context.Context, req *v1.AddRoleListReq) (res *v1.AddRoleListRes, err error) {

	dbQuery := dao.Role.Ctx(ctx).
		Where("name = ?", req.Name)

	var users []entity.Users
	var totalCount int
	if err := dbQuery.ScanAndCount(&users, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "query db error roles : %v", err)
		return nil, fmt.Errorf(" query db error roles: %w", err)
	}

	if totalCount != 0 {
		g.Log().Errorf(ctx, "Existing roles : %v", err)
		return nil, fmt.Errorf(" Existing roles: %w", err)
	}

	data, err := json.Marshal(req.Cmd)
	if err != nil {
		g.Log().Errorf(ctx, "JSON marshal failed: '%s': %v", req.Name, err)
		return nil, fmt.Errorf("JSON marshal failed:: %w", err)
	}
	jsonStr := string(data)

	dbQuery = dao.Bot.Ctx(ctx).
		Where("id = ?", req.BotID)

	var bots []entity.Bot
	totalCount = 0
	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "query db error bot : %v", err)
		return nil, fmt.Errorf(" query db error bot: %w", err)
	}

	if totalCount == 0 {
		g.Log().Errorf(ctx, "Not Existing bot : %v", err)
		return nil, fmt.Errorf("not Existing bot: %w", err)
	}

	role := &entity.Role{
		Name:     req.Name,
		BotId:    req.BotID,
		BotName:  bots[0].Name,
		Cmd:      jsonStr,
		UserSize: 0,
	}
	_, err = dao.Role.Ctx(ctx).
		Data(role).
		InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to insert new Role '%s': %v", req.Name, err)
		return nil, fmt.Errorf("failed to insert new Role: %w", err)
	}
	return res, err
}
