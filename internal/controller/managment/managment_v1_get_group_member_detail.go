package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetGroupMemberDetail(ctx context.Context, req *v1.GetGroupMemberDetailReq) (res *v1.GetGroupMemberDetailRes, err error) {
	var tgUsers []entity.TgUsers
	var totalCount int
	if err = dao.TgUsers.Ctx(ctx).Where("group_id = ?", req.ID).
		ScanAndCount(&tgUsers, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, fmt.Errorf("%w", err)
	}

	res = &v1.GetGroupMemberDetailRes{
		Data: make([]v1.GroupMemberDetail, totalCount),
	}

	for k, v := range tgUsers {
		res.Data[k] = v1.GroupMemberDetail{
			TgId:   v.TgId,
			TgName: v.TgName,
		}
	}

	return res, nil
}
