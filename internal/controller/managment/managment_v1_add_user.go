package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

const (
	UserStatus = 0
	UserAvailable
	UserUnAvailable
)

func (c *ControllerV1) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {
	user := &entity.Users{
		Account:    req.Account,
		Name:       req.Name,
		Role:       req.Role,
		Status:     UserAvailable,
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
