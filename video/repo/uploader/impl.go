package uploader

import (
	"context"
)

// uploaderImpl 上传器实现
type uploaderImpl struct {
}

// VideoUpload 视频上传
func (u *uploaderImpl) VideoUpload(ctx context.Context, req *MediaFileUploadReq) (int64, error) {

	return 0, nil
}
