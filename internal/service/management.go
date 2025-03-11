// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "tg_bot_backend/api/managment/v1"
)

type (
	IManagement interface {
		AddUser(ctx context.Context, req *v1.AddUserReq) (*v1.AddUserRes, error)
	}
)

var (
	localManagement IManagement
)

func Management() IManagement {
	if localManagement == nil {
		panic("implement not found for interface IManagement, forgot register?")
	}
	return localManagement
}

func RegisterManagement(i IManagement) {
	localManagement = i
}
