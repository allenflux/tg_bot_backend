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

func (c *ControllerV1) SearchTgUserAccount(ctx context.Context, req *v1.SearchTgUserAccountReq) (res *v1.SearchTgUserAccountRes, err error) {
	dbQuery := dao.TgUsers.Ctx(ctx).
		//Page(req.PageNum, req.PageSize).
		Order("id desc").
		Where("tg_account = ?", req.KeyWordSearch)

	//if req.KeyWordSearch != "" {
	//	dbQuery = dbQuery.Where("tg_account LIKE ?", "%"+req.KeyWordSearch+"%")
	//}

	var tgUsers []entity.TgUsers
	var totalCount int
	if err := dbQuery.ScanAndCount(&tgUsers, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count TgUsers: %v", err)
		return nil, fmt.Errorf("failed to fetch TgUsers list: %w", err)
	}

	// Prepare response
	res = &v1.SearchTgUserAccountRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]v1.SearchTgUserAccountData, len(tgUsers)),
	}

	// Map data to response structure
	for i, v := range tgUsers {
		res.Data[i] = v1.SearchTgUserAccountData{
			Name:          v.TgName,
			TgUserAccount: v.TgAccount,
		}
	}

	return res, nil
}
