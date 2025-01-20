package uploader

import (
	"context"
)

// MediaFileUploadReq 媒体文件上传所需内容
type MediaFileUploadReq struct {
	UID  uint64
	VID  uint64
	Data []byte
	Size int64
	Tags []string
}

// MediaFileUploader 文件上传器
type MediaFileUploader interface {
	// VideoUpload 视频上传
	VideoUpload(ctx context.Context, req *MediaFileUploadReq) (int64, error)
}

// NewMediaFileUploader 创建上传器对象
func NewMediaFileUploader() MediaFileUploader {
	return &uploaderImpl{}
}
