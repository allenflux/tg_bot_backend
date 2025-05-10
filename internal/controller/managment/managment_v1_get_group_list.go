package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	commonApi "tg_bot_backend/api/common"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetGroupList(ctx context.Context, req *v1.GetGroupListReq) (res *v1.GetGroupListRes, err error) {
	dbQuery := dao.Group.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id desc")

	if req.KeyWordSearch != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+req.KeyWordSearch+"%")
	}

	if req.GroupType != consts.GroupType {
		dbQuery = dbQuery.Where("type = ?", req.GroupType)
	}

	if req.PlatformID != 0 {
		dbQuery = dbQuery.Where("central_control_id = ?", req.PlatformID)
	}

	var groups []entity.Group
	var totalCount int
	if err := dbQuery.ScanAndCount(&groups, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count Group List: %v", err)
		return nil, fmt.Errorf("failed to fetch Group list: %w", err)
	}

	// Prepare response
	res = &v1.GetGroupListRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]v1.GroupMap, len(groups)),
	}

	// Map data to response structure
	for i, v := range groups {
		platform, err := GetPlatform(v.CentralControlId, ctx)
		if err != nil {
			g.Log().Errorf(ctx, "Failed to query and count platformName : %v", err)
			return nil, fmt.Errorf("failed to fetch platformName : %w", err)
		}
		res.Data[i] = v1.GroupMap{
			ID:                   v.Id,
			Name:                 v.Name,
			PlatformName:         platform.Name,
			TGLink:               v.TgLink,
			GroupType:            v.Type,
			GroupSize:            v.Size,
			PlatformIdAndGroupId: fmt.Sprintf("%d+%d", v.CentralControlId, v.Id),
			AssociatedRobot:      v.BotSize,
			AssociatedRole:       v.RoleSize,
		}
	}
	return res, err
}

func GetPlatform(id int, ctx context.Context) (res *entity.CentralControl, err error) {
	platform := new(entity.CentralControl)
	err = dao.CentralControl.Ctx(ctx).Where("id", id).Scan(platform)
	return platform, err
}
