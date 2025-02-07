package errcode

const (
	ErrPublishFailed   = 30001 // 上传失败
	ErrReadConfig      = 30002 // 读取配置出错
	ErrUnmarshalConfig = 30003 // 解析配置出错
	ErrNotVideoType    = 30004 // 非视频文件类型
	ErrInsertVideo     = 30005 // 视频信息写入失败
)
