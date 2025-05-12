package managment

//
//func (c *ControllerV1) GetBotBindList(ctx context.Context, req *v1.GetBotBindListReq) (res *v1.GetBotBindListRes, err error) {
//	dbQuery := dao.Bot.Ctx(ctx).
//		Page(req.PageNum, req.PageSize).
//		Order("id desc")
//
//	var bots []entity.Bot
//	var totalCount int
//	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
//		g.Log().Errorf(ctx, "Failed to query and count bots List: %v", err)
//		return nil, fmt.Errorf("failed to fetch bots list: %w", err)
//	}
//
//	// Prepare response
//	res = &v1.GetBotBindListRes{
//		ListRes: commonApi.ListRes{Total: totalCount},
//		Data:    make([]v1.BotBindMap, len(bots)),
//	}
//
//	// Map data to response structure
//	for i, v := range bots {
//		res.Data[i] = v1.BotBindMap{
//			ID:             v.Id,
//			Name:           v.Name,
//			Account:        v.Account,
//			NumberOfGroups: 0,
//			NumberOfUsers:  0,
//			GroupIDList:    nil,
//			NumberOfBots:   0,
//			NumberOfRole:   0,
//		}
//	}
//
//	return res, nil
//}
