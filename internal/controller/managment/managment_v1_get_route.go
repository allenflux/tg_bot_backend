package managment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"

	"tg_bot_backend/api/managment/v1"
)

func (c *ControllerV1) GetRoute(ctx context.Context, req *v1.GetRouteReq) (res *v1.GetRouteRes, err error) {
	jsonData := []byte(hardCodeResponse) // 你的JSON数据
	var response v1.Response
	if err = json.Unmarshal(jsonData, &response); err != nil {
		g.Log().Errorf(ctx, "Failed to Unmarshal jsonData: %v", err)
		return nil, fmt.Errorf("failed to Unmarshal jsonData: %w", err)
	}
	return &v1.GetRouteRes{
		Data: response.Data,
	}, nil
}

var hardCodeResponse string = `
{
    "success": true,
    "data": [
        {
            "path": "/system",
            "meta": {
                "icon": "ri:settings-3-line",
                "title": "menus.pureSysManagement",
                "rank": 14
            },
            "children": [
                {
                    "path": "/system/user/index",
                    "name": "SystemUser",
                    "meta": {
                        "icon": "ri:admin-line",
                        "title": "menus.pureUser",
                        "roles": [
                            "admin"
                        ]
                    }
                }
            ]
        },
        {
            "path": "/roles",
            "meta": {
                "icon": "ep:user-filled",
                "title": "角色权限设置",
                "rank": 2
            },
            "children": [
                {
                    "path": "/roles/permission/index",
                    "name": "rolePermissionSettings",
                    "meta": {
                        "icon": "ep:baseball",
                        "title": "角色权限设置",
                        "roles": [
                            "admin"
                        ],
                        "showParent": true
                    }
                }
            ]
        },
        {
            "path": "/groups",
            "meta": {
                "icon": "ep:user-filled",
                "title": "群组管理",
                "rank": 2
            },
            "children": [
                {
                    "path": "/groups/management/index",
                    "name": "groupManagement",
                    "meta": {
                        "icon": "ep:baseball",
                        "title": "群组管理",
                        "roles": [
                            "admin",
                            "common"
                        ],
                        "showParent": true
                    }
                }
            ]
        },
		{
			"path": "/robot",
			"meta": {
				"icon": "ep:user-filled",
				"title": "机器人管理",
				"rank": 2
			},
			"children": [
				{
					"path": "/robot/base/index",
					"name": "baseSettings",
					"meta": {
						"icon": "ep:baseball",
						"title": "基础管理",
						"roles": [
							"admin",
							"common"
						],
						"showParent": true
					}
				},
			]
		},
        {
            "path": "/controls",
            "meta": {
                "icon": "ri:bookmark-2-line",
                "title": "中控平台",
                "rank": 1
            },
            "children": [
                {
                    "path": "/controls/management/index",
                    "name": "management",
                    "meta": {
                        "title": "中控平台管理",
                        "roles": [
                            "admin"
                        ],
                        "showParent": true
                    }
                }
            ]
        }
    ]
}

`
