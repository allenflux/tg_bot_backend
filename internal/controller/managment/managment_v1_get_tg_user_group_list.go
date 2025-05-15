package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

type TreeGroupData struct {
	Value    string               `json:"value"`
	Children []v1.TgUserGroupData `json:"children"`
}

func (c *ControllerV1) GetTgUserGroupList(ctx context.Context, req *v1.GetTgUserGroupListReq) (res *v1.GetTgUserGroupListRes, err error) {
	// Get Tg Users Group ID list
	dbQuery := dao.TgUsers.Ctx(ctx).
		Where("tg_account = ?", req.TgAccount)

	var tgUsers []entity.TgUsers
	var totalCount int
	if err := dbQuery.ScanAndCount(&tgUsers, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count TgUsers: %v", err)
		return nil, fmt.Errorf("failed to fetch TgUsers list: %w", err)
	}

	if totalCount == 0 {
		return nil, fmt.Errorf("invalid count TgUsers: %v", totalCount)
	}

	// Get Group Info
	groupIds := make([]int, 0)
	for _, tgUser := range tgUsers {
		groupIds = append(groupIds, tgUser.GroupId)
	}

	dbQuery = dao.Group.Ctx(ctx).
		WhereIn("id ", groupIds)

	var groups []entity.Group
	totalCount = 0
	if err := dbQuery.ScanAndCount(&groups, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count Group: %v", err)
		return nil, fmt.Errorf("failed to fetch Group list: %w", err)
	}

	if totalCount == 0 {
		return nil, fmt.Errorf("invalid count Group: %v", totalCount)
	}

	// Get Platform Info
	platformIds := make([]int, 0)
	for _, group := range groups {
		platformIds = append(platformIds, group.CentralControlId)
	}

	dbQuery = dao.CentralControl.Ctx(ctx).
		WhereIn("id ", platformIds)

	var platforms []entity.CentralControl
	totalCount = 0
	if err := dbQuery.ScanAndCount(&platforms, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count platforms: %v", err)
		return nil, fmt.Errorf("failed to fetch platforms list: %w", err)
	}

	if totalCount == 0 {
		return nil, fmt.Errorf("invalid count platforms: %v", totalCount)
	}

	groupMap := make(map[int][]v1.TgUserGroupData, len(groups))
	for _, group := range groups {
		groupMap[group.CentralControlId] = append(groupMap[group.CentralControlId], v1.TgUserGroupData{
			Name:       group.Name,
			ID:         group.Id,
			PlatformID: group.CentralControlId,
			Value:      group.Name,
		})
	}

	resultMap := make(map[string][]v1.TgUserGroupData, len(platforms))
	for _, platform := range platforms {
		resultMap[platform.Name] = groupMap[platform.Id]
	}

	treeData := make([]TreeGroupData, 0, len(resultMap))
	for key, children := range resultMap {
		treeData = append(treeData, TreeGroupData{
			Value:    key,
			Children: children,
		})
	}

	res = &v1.GetTgUserGroupListRes{
		Data:     resultMap,
		Resource: treeData,
	}

	return res, nil
}
