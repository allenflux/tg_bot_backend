package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetBindTgUsers(ctx context.Context, req *v1.GetBindTgUsersReq) (res *v1.GetBindTgUsersRes, err error) {
	dbQuery := dao.TgUsers.Ctx(ctx).
		Where("role_id = ?", req.RoleId)

	var users []entity.TgUsers
	var totalCount int
	if err := dbQuery.ScanAndCount(&users, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count TgUsers list: %v", err)
		return nil, fmt.Errorf("failed to fetch TgUsers list: %w", err)
	}

	// Prepare response
	res = &v1.GetBindTgUsersRes{
		Data: make([]v1.BindTgUsersData, len(users)),
	}

	// Map data to response structure
	for i, user := range users {
		res.Data[i] = v1.BindTgUsersData{
			Name:      user.TgName,
			TgAccount: user.TgAccount,
			ID:        user.Id,
		}
	}
	return res, nil
}
