package managment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
