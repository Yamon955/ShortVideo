package entity

const (
	ErrLoadRadixTree          = 10001 // 加载路由树出错
	ErrLoadRegexRules         = 10002 // 加载正则路由信息出错
	ErrGetRouteConfFromNode   = 10003 // 从radixTreeNode解析路由信息出错
	ErrFindRouteConfFromRegex = 10004 //从正则信息获取获取路由信息出错
	ErrMatchRoute             = 10005 // 未匹配到路由
)
