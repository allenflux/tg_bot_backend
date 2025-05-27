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

const (
	PlatformPathLogin        = "/api/login_admin"
	PlatformVerifyCustomerId = "/api/customer/query"
	PlatformVerifyBusinessId = "/api/admin/channel/detail/"
	PlatformAddCustomerFind  = "/api/customer/finance/add"
)

var UserPermissionMap = map[string]int{
	"管理员": 1,
}

var PermissionUserMap = map[int]string{
	1: "管理员",
}

const (
	BotCmdTopUp  = "topup"
	BotCmdBind   = "bind"
	BotCmdUnbind = "unbind"
)

var BotCmdDesc = map[string]string{
	BotCmdTopUp:  "加款",
	BotCmdBind:   "绑定",
	BotCmdUnbind: "解绑",
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
