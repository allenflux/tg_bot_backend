package consts

import "github.com/gogf/gf/v2/os/gtime"

type EnumsUserStatus int

const (
	EnumUserStatusEnable  EnumsUserStatus = 2
	EnumUserStatusDisable EnumsUserStatus = 1
)

const (
	BotStatus = iota
	BotStatusAvailable
	BotStatusUnAvailable
)

const (
	UserStatus = iota
	UserAvailable
	UserUnAvailable
)

const (
	CentralControlStatus = iota
	CentralControlStatusAvailable
	CentralControlStatusUnAvailable
)

const (
	GroupType = iota
	GroupTypeForCustomer
	GroupTypeForBusiness
)

const (
	GreetingStatus = iota
	GreetingStatusAvailable
	GreetingStatusUnAvailable
)

var UserPermissionMap = map[string]int{
	"管理员": 1,
}

func (e EnumsUserStatus) IsValid() bool {
	return e >= EnumUserStatusDisable && e <= EnumUserStatusEnable
}

const (
	PageSize = 10 //分页长度
)

var StartTime = gtime.New()

type EnumResponseStatus int

const (
	EnumResponseSuccess EnumResponseStatus = 0 //成功
	EnumResponseWarning EnumResponseStatus = 2 //警告
)

const HKDToRMB = 0.93

const TaskFilePath = "./resource/file"
