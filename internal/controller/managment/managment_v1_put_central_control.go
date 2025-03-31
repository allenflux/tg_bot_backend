package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) PutCentralControl(ctx context.Context, req *v1.PutCentralControlReq) (res *v1.PutCentralControlRes, err error) {
	if _, err := dao.CentralControl.Ctx(ctx).
		Data(g.Map{
			"name": req.Name,
			"note": req.Note,
		}).
		Where("id = ?", req.ID).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update CentralControl with ID=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to update CentralControl: %w", err)
	}

	// Log success and prepare response
	g.Log().Infof(ctx, "Successfully updated CentralControl with ID=%d", req.ID)
	return res, nil
}
