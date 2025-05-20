package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetPlatformList(ctx context.Context, req *v1.GetPlatformListReq) (res *v1.GetPlatformListRes, err error) {
	dbQuery := dao.CentralControl.Ctx(ctx).
		//Page(req.PageNum, req.PageSize).
		Order("id desc")

	//if req.KeyWordSearch != "" {
	//	dbQuery = dbQuery.Where("name LIKE ?", "%"+req.KeyWordSearch+"%")
	//}

	var centrals []entity.CentralControl
	var totalCount int
	if err := dbQuery.ScanAndCount(&centrals, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count CentralControlList: %v", err)
		return nil, fmt.Errorf("failed to fetch CentralControl list: %w", err)
	}

	// Prepare response
	res = &v1.GetPlatformListRes{
		Data: make([]v1.GetPlatformData, len(centrals)),
	}

	// Map data to response structure
	for i, central := range centrals {
		res.Data[i] = v1.GetPlatformData{
			PlatformId:   central.Id,
			PlatformName: central.Name,
		}
	}

	return res, nil
}
