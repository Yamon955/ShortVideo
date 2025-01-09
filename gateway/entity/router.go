// Package entity 公共函数、结构体、常量
package entity

import "regexp"

// RouteConf 路由配置
type RouteConf struct {
	Path          string         `json:"path"`
	IsRegex       int            `json:"is_regex"`
	Service       string         `json:"service"`
	Method        string         `json:"method"`
	Target        string         `json:"target"`
	Timeout       string         `json:"timeout"`
	IsAllowNoAuth bool           `json:"is_allow_no_auth"`
	Protocol      string         `json:"protocol"`
	Regexp        *regexp.Regexp `json:"-"`
	CalleeMethod  string         `json:"-"`
}
