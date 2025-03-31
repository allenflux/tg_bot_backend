package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) AddRoleList(ctx context.Context, req *v1.AddRoleListReq) (res *v1.AddRoleListRes, err error) {
	role := &entity.Role{
		Name:     req.Name,
		BotId:    req.BotID,
		Cmd:      "req.Cmd",
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
