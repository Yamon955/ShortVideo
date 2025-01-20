package downloader

import (
	"context"
	"mime/multipart"
)

// MediaFile 媒体文件
type MediaFile struct {
	Data []byte
	Size int64
}

// MediaFileDownloader 文件下载器
type MediaFileDownloader interface {
	// PieceDownloadVideoFile 分片下载
	PieceDownloadVideoFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*MediaFile, error)
	// DownloadVideoFile 普通下载
	DownloadVideoFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*MediaFile, error)
}

// NewMediaFileDownloader 创建下载器对象
func NewMediaFileDownloader() MediaFileDownloader {
	return &downloaderImpl{}
}
