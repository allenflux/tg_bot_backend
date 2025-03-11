package management

import (
	"context"
	v1 "tg_bot_backend/api/managment/v1"
)

const (
	UserStatus = 0
	UserAvailable
	UserUnAvailable
)

func (s *sManagement) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {

	return res, err
}

func (s *sManagement) GetUserList(ctx context.Context, req *v1.GetUserListReq) (*v1.GetUserListRes, error) {

	return nil, nil
}

func (s *sManagement) EditUser(ctx context.Context, req *v1.EditUserReq) (*v1.EditUserRes, error) {
	return nil, nil
}

func (s *sManagement) DelUser(ctx context.Context, req *v1.DelUserReq) (*v1.DelUserRes, error) {
	return nil, nil
}
