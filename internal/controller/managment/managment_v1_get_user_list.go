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

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	dbQuery := dao.Users.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id desc")

	if req.KeyWordSearch != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+req.KeyWordSearch+"%")
	}

	var users []entity.Users
	var totalCount int
	if err := dbQuery.ScanAndCount(&users, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count UsersList: %v", err)
		return nil, fmt.Errorf("failed to fetch users list: %w", err)
	}

	// Prepare response
	res = &v1.GetUserListRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]v1.UserMap, len(users)),
	}

	// Map data to response structure
	for i, user := range users {
		res.Data[i] = v1.UserMap{
			ID:         user.Id,
			Account:    user.Account,
			Name:       user.Name,
			Role:       user.Role,
			Status:     user.Status,
			CreateTime: user.CreateTime,
		}
	}

	return res, nil
}
