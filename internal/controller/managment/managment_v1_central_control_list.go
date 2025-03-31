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

func (c *ControllerV1) CentralControlList(ctx context.Context, req *v1.CentralControlListReq) (res *v1.CentralControlListRes, err error) {
	dbQuery := dao.CentralControl.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id desc")

	if req.KeyWordSearch != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+req.KeyWordSearch+"%")
	}

	var centrals []entity.CentralControl
	var totalCount int
	if err := dbQuery.ScanAndCount(&centrals, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count CentralControlList: %v", err)
		return nil, fmt.Errorf("failed to fetch CentralControl list: %w", err)
	}

	// Prepare response
	res = &v1.CentralControlListRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]v1.CentralControlMap, len(centrals)),
	}

	// Map data to response structure
	for i, central := range centrals {
		res.Data[i] = v1.CentralControlMap{
			ID:                central.Id,
			Name:              central.Name,
			Domain:            central.Domain,
			SecretKey:         central.SecretKey,
			NumberOfCustomers: central.NumberOfCustomers,
			NumberOfBusiness:  central.NumberOfBusiness,
			Note:              central.Note,
			Status:            central.Status,
		}
	}

	return res, nil
}
