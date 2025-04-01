package v1

import "github.com/gogf/gf/v2/frame/g"

type Response struct {
	Success bool    `json:"success"`
	Data    []Route `json:"data"`
}

type Route struct {
	Path      string  `json:"path"`
	Name      string  `json:"name,omitempty"`
	Component string  `json:"component,omitempty"`
	Meta      Meta    `json:"meta"`
	Children  []Route `json:"children,omitempty"`
}

type Meta struct {
	Icon       string   `json:"icon,omitempty"`
	Title      string   `json:"title"`
	Rank       int      `json:"rank,omitempty"`
	Roles      []string `json:"roles,omitempty"`
	ShowParent bool     `json:"showParent,omitempty"`
}

type GetRouteReq struct {
	g.Meta `path:"/route" tags:"route" method:"get" summary:"返回路由权限信息"`
}

type GetRouteRes struct {
	Data []Route `json:"data"`
}
