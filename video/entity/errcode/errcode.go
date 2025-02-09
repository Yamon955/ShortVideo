package errcode

const (
	ErrPublishFailed   = 30001 // 上传失败
	ErrReadConfig      = 30002 // 读取配置出错
	ErrUnmarshalConfig = 30003 // 解析配置出错
	ErrNotVideoType    = 30004 // 非视频文件类型
	ErrInsertVideo     = 30005 // 视频信息写入失败
	ErrInvalidStartID  = 30006 // 非法的偏移量
	ErrDBOperation     = 30007 // 数据库操作失败
)
