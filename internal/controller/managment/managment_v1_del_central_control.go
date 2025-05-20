package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) DelCentralControl(ctx context.Context, req *v1.DelCentralControlReq) (res *v1.DelCentralControlRes, err error) {
	if req.ID == 0 {
		return nil, fmt.Errorf("invalid Id %d", req.ID)
	}
	if _, err := dao.CentralControl.Ctx(ctx).Where("id = ?", req.ID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete CentralControl with CentralControlID=%d: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete CentralControl with CentralControlID=%d: %w", req.ID, err)
	}
	return res, nil
}
