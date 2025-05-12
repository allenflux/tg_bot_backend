package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {
	dbQuery := dao.Users.Ctx(ctx).
		Where("account = ?", req.Account)

	var users []entity.Users
	var totalCount int
	if err := dbQuery.ScanAndCount(&users, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count users: %v", err)
		return nil, fmt.Errorf("failed to fetch users list: %w", err)
	}

	if totalCount != 0 {
		g.Log().Errorf(ctx, "Existing accounts: %v", err)
		return nil, fmt.Errorf(" Existing accounts: %w", err)
	}

	user := &entity.Users{
		Account:    req.Account,
		Name:       req.Name,
		Role:       consts.UserPermissionMap[req.Role],
		Status:     consts.UserAvailable,
		CreateTime: gtime.Now(),
		Password:   req.Password,
	}
	_, err = dao.Users.Ctx(ctx).
		Data(user).
		InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to insert new users '%s': %v", req.Account, err)
		return nil, fmt.Errorf("failed to insert new user: %w", err)
	}
	return res, err
}
