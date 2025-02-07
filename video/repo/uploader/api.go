package uploader

import (
	"context"
	"mime/multipart"

	"github.com/Yamon955/ShortVideo/video/repo/minio"
)

// MediaFileUploadReq 媒体文件上传所需内容
type MediaFileUploadReq struct {
	UID  uint64
	File multipart.File
	Tags []string
}

// MediaFileUploader 文件上传器
type MediaFileUploader interface {
	// VideoUpload 视频上传
	VideoUpload(ctx context.Context, req *MediaFileUploadReq) (uint64, error)
}

// NewMediaFileUploader 创建上传器对象
func NewMediaFileUploader() MediaFileUploader {
	return &uploaderImpl{
		MinIOClient: minio.GetClient(),
	}
}
