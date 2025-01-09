package entity

const (
	ErrLoadRadixTree          = 10001 // 加载路由树出错
	ErrLoadRegexRules         = 10002 // 加载正则路由信息出错
	ErrGetRouteConfFromNode   = 10003 // 从radixTreeNode解析路由信息出错
	ErrFindRouteConfFromRegex = 10004 //从正则信息获取获取路由信息出错
	ErrMatchRoute             = 10005 // 未匹配到路由
	ErrContentType            = 10006 // 非法的 content type
	ErrReadHTTPReqBody        = 10007 // 读取 http.req.Body 出错
	ErrTransHTTPRsp           = 10008 // trpc.rsp 转成 http.rsp 出错
	ErrVerifAuthFail          = 10009 // 服务端鉴权失败
)
