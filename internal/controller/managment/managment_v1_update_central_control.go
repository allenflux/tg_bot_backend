package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) UpdateCentralControl(ctx context.Context, req *v1.UpdateCentralControlReq) (res *v1.UpdateCentralControlRes, err error) {
	data := &entity.CentralControl{
		Name:              req.Name,
		Domain:            req.Domain,
		NumberOfCustomers: 0,
		NumberOfBusiness:  0,
		Note:              req.Note,
		Status:            consts.CentralControlStatusAvailable,
		SecretKey:         req.SecretKey,
	}
	_, err = dao.CentralControl.Ctx(ctx).
		Data(data).
		InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to insert new CentralControl '%s': %v", req.Name, err)
		return nil, fmt.Errorf("failed to insert new CentralControl: %w", err)
	}
	return res, err
}
